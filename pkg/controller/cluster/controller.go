/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	awsplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/aws_plan/utils"
	azureplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/azure_plan/utils"
	clusterutils "github.com/frodenas/pks-k8s-api/pkg/controller/cluster/utils"
	clusternsxtutils "github.com/frodenas/pks-k8s-api/pkg/controller/cluster_nsxt/utils"
	gcpplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/gcp_plan/utils"
	kubernetesprofileutils "github.com/frodenas/pks-k8s-api/pkg/controller/kubernetes_profile/utils"
	odbplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/odb_plan/utils"
	vsphereplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/vsphere_plan/utils"
	"github.com/frodenas/pks-k8s-api/pkg/provisioner"
	provisionertypes "github.com/frodenas/pks-k8s-api/pkg/provisioner/types"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	// ClusterFinalizer is set on Reconcile callback.
	ClusterFinalizer = "cluster_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.cluster")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "validation_count",
			Help:      "Total number of cluster validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "validation_errors_count",
			Help:      "Total number of cluster validation errors",
		},
		[]string{"namespace", "name"},
	)

	creationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "creation_count",
			Help:      "Total number of cluster creation invocations",
		},
		[]string{"namespace", "name", "provisioner"},
	)

	creationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "creation_errors_count",
			Help:      "Total number of cluster creation errors",
		},
		[]string{"namespace", "name", "provisioner"},
	)

	creationTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "cluster",
			Name:      "creation_time_seconds",
			Help:      "Time in seconds spent creating a cluster",
		},
		[]string{"namespace", "name", "provisioner"},
	)

	deletionsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "deletion_count",
			Help:      "Total number of cluster deletion invocations",
		},
		[]string{"namespace", "name", "provisioner"},
	)

	deletionErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster",
			Name:      "deletion_errors_count",
			Help:      "Total number of cluster deletion errors",
		},
		[]string{"namespace", "name", "provisioner"},
	)

	deletionTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "cluster",
			Name:      "deletion_time_seconds",
			Help:      "Time in seconds spent deleting a cluster",
		},
		[]string{"namespace", "name", "provisioner"},
	)
)

func init() {
	metrics.Registry.MustRegister(validationsCounter)
	metrics.Registry.MustRegister(validationErrorsCounter)
	metrics.Registry.MustRegister(creationsCounter)
	metrics.Registry.MustRegister(creationErrorsCounter)
	metrics.Registry.MustRegister(creationTime)
	metrics.Registry.MustRegister(deletionsCounter)
	metrics.Registry.MustRegister(deletionErrorsCounter)
	metrics.Registry.MustRegister(deletionTime)
}

// Add creates a new Cluster Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileCluster{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("cluster-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("cluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Clusters
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.Cluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileCluster{}

// ReconcileCluster reconciles a Cluster object
type ReconcileCluster struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a Cluster object and makes changes based on the state read
// and what is in the Cluster.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=awsplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=awsplans/status,verbs=get;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=azureplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=azureplans/status,verbs=get;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=gcpplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=gcpplans/status,verbs=get;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans/status,verbs=get;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=odbplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=odbplans/status,verbs=get;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=kubernetesprofiles,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=kubernetesprofiles/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.Cluster{}
	if err := r.Get(context.TODO(), request.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	// Check for deletion.
	if !instance.DeletionTimestamp.IsZero() {
		return r.delete(instance)
	}

	return r.reconcile(instance)
}

func (r *ReconcileCluster) reconcile(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling Cluster `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(ClusterFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to Cluster `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, ClusterFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a ClusterValidated condition if absent.
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to Cluster `%s/%s`", instance.Namespace, instance.Name))
		condition := clusterutils.NewClusterCondition(
			pksv1alpha1.ClusterValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"Cluster has not yet been validated",
			"",
		)
		clusterutils.SetClusterCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If Cluster has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	// Check if Cluster requires using NSX-T resources.
	nsxt, err := clusternsxtutils.IsUsingNSXT(r.Client, instance)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error checking if Cluster `%s/%s` needs NSX-T resources: %v", instance.Namespace, instance.Name, err)
	}
	if nsxt {
		// Add a NSXTProvisioned condition if absent.
		if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterNSXTProvisioned); c == nil {
			log.Info(fmt.Sprintf("Adding `NSXTProvisioned` condition to Cluster `%s/%s`", instance.Namespace, instance.Name))
			condition := clusterutils.NewClusterCondition(
				pksv1alpha1.ClusterNSXTProvisioned,
				corev1.ConditionFalse,
				"ProvisioningPending",
				"NSX-T resources have not been yet provisioned",
				"",
			)
			clusterutils.SetClusterCondition(&instance.Status, *condition)
			return reconcile.Result{}, r.Status().Update(context.Background(), instance)
		}
	}

	// If any condition is not true, trigger the reconciliation again.
	if !clusterutils.AreAllClusterConditionsTrue(instance.Status) {
		log.Info(fmt.Sprintf("There are conditions for Cluster `%s/%s` that are not true", instance.Namespace, instance.Name))
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Trigger the creation if it has not yet been triggered.
	if instance.Status.LastOperation == nil ||
		instance.Status.LastOperation.Type != pksv1alpha1.ClusterLastOperationTypeCreate ||
		instance.Status.LastOperation.State == pksv1alpha1.ClusterLastOperationStateFailed {
		result, err := r.createCluster(instance)
		if err != nil {
			log.Error(err, fmt.Sprintf("error creating Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "CreationError", err.Error())
		}
		return result, err
	}

	// If creation is still in progress, check the status.
	if instance.Status.LastOperation.State != pksv1alpha1.ClusterLastOperationStateSucceeded {
		result, err := r.lastClusterOperation(instance)
		if err != nil {
			log.Error(err, fmt.Sprintf("error getting last operation for Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "CreationError", err.Error())
			return reconcile.Result{}, err
		}
		return result, err
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileCluster) validate(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating Cluster `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition if needed.
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterValidated); c == nil || c.Status != corev1.ConditionFalse {
		log.Info(fmt.Sprintf("Updating `Validated` condition to `False` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		condition := clusterutils.NewClusterCondition(
			pksv1alpha1.ClusterValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"Cluster has not yet been validated",
			"",
		)
		clusterutils.SetClusterCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Validate the spec details.
	if err := r.validateSpec(instance); err != nil {
		log.Error(err, fmt.Sprintf("Error validating Cluster `%s/%s`", instance.Namespace, instance.Name))
		validationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "ValidationError", err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, err
	}

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for Cluster `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "Cluster has been successfully validated")
	condition := clusterutils.NewClusterCondition(
		pksv1alpha1.ClusterValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"Cluster has been validated",
		"",
	)
	clusterutils.SetClusterCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileCluster) delete(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	// Trigger the deletion if it has not yet been triggered.
	if instance.Status.LastOperation != nil &&
		(instance.Status.LastOperation.Type != pksv1alpha1.ClusterLastOperationTypeDelete ||
			instance.Status.LastOperation.State == pksv1alpha1.ClusterLastOperationStateFailed) {
		result, err := r.deleteCluster(instance)
		if err != nil {
			log.Error(err, fmt.Sprintf("Error deleting Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "DeletionError", err.Error())
		}
		return result, err
	}

	// If deletion is still in progress, check the status.
	if instance.Status.LastOperation != nil && instance.Status.LastOperation.State != pksv1alpha1.ClusterLastOperationStateSucceeded {
		result, err := r.lastClusterOperation(instance)
		if err != nil {
			log.Error(err, fmt.Sprintf("error getting last operation for Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "DeletionError", err.Error())
			return reconcile.Result{}, err
		}
		return result, err
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(ClusterFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from Cluster `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(ClusterFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileCluster) createCluster(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Creating Cluster `%s/%s`", instance.Namespace, instance.Name))
	creationsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()

	clusterProvisioner, err := r.getClusterProvisioner(instance)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error getting provisioner credentials: %v", err)
	}

	provisionerLastOperation, err := clusterProvisioner.CreateCluster(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	provisionerLastOperationRaw, err := json.Marshal(provisionerLastOperation)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error marshalling provisioner last operation: %v", err)
	}

	instance.Status.LastOperation = &pksv1alpha1.ClusterLastOperation{
		Description:     "creating cluster",
		StartTime:       metav1.NewTime(time.Now()),
		LastUpdated:     metav1.NewTime(time.Now()),
		Type:            pksv1alpha1.ClusterLastOperationTypeCreate,
		ProvisionerData: string(provisionerLastOperationRaw),
	}
	switch provisionerLastOperation.State {
	case provisionertypes.ProvisionerOperationStateSucceeded:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Succeed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		creationTime.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Observe(time.Now().Sub(instance.Status.LastOperation.StartTime.Time).Seconds())
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateSucceeded
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	case provisionertypes.ProvisionerOperationStateInProgress:
		log.Info(fmt.Sprintf("Updating LastOperation status to `InProgress` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateInProgress
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	default:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Failed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		creationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "CreationError", provisionerLastOperation.Description)
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateFailed
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	}
}

func (r *ReconcileCluster) deleteCluster(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting Cluster `%s/%s`", instance.Namespace, instance.Name))
	deletionsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()

	clusterProvisioner, err := r.getClusterProvisioner(instance)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error getting provisioner: %v", err)
	}

	provisionerLastOperation, err := clusterProvisioner.DeleteCluster(instance)
	if err != nil {
		return reconcile.Result{}, err
	}

	provisionerLastOperationRaw, err := json.Marshal(provisionerLastOperation)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error marshalling provisioner last operation: %v", err)
	}

	instance.Status.LastOperation = &pksv1alpha1.ClusterLastOperation{
		Description:     "deleting cluster",
		StartTime:       metav1.NewTime(time.Now()),
		LastUpdated:     metav1.NewTime(time.Now()),
		Type:            pksv1alpha1.ClusterLastOperationTypeDelete,
		ProvisionerData: string(provisionerLastOperationRaw),
	}

	switch provisionerLastOperation.State {
	case provisionertypes.ProvisionerOperationStateSucceeded:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Succeed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		creationTime.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Observe(time.Now().Sub(instance.Status.LastOperation.StartTime.Time).Seconds())
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateSucceeded
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	case provisionertypes.ProvisionerOperationStateInProgress:
		log.Info(fmt.Sprintf("Updating LastOperation status to `InProgress` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateInProgress
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	default:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Failed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		deletionErrorsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "DeletionError", provisionerLastOperation.Description)
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateFailed
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	}
}

func (r *ReconcileCluster) lastClusterOperation(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	clusterProvisioner, err := r.getClusterProvisioner(instance)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error getting provisioner: %v", err)
	}

	oldProvisionerLastOperation := &provisionertypes.ProvisionerLastOperation{}
	if err := json.Unmarshal([]byte(instance.Status.LastOperation.ProvisionerData), oldProvisionerLastOperation); err != nil {
		return reconcile.Result{}, fmt.Errorf("error unmarshalling provisioner last operation: %v", err)
	}

	provisionerLastOperation, err := clusterProvisioner.LastOperation(instance, *oldProvisionerLastOperation)
	if err != nil {
		return reconcile.Result{}, err
	}

	instance.Status.LastOperation.LastUpdated = metav1.NewTime(time.Now())
	provisionerLastOperationRaw, err := json.Marshal(provisionerLastOperation)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("error marshalling provisioner last operation: %v", err)
	}
	instance.Status.LastOperation.ProvisionerData = string(provisionerLastOperationRaw)

	switch provisionerLastOperation.State {
	case provisionertypes.ProvisionerOperationStateSucceeded:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Succeed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		if instance.Status.LastOperation.Type == pksv1alpha1.ClusterLastOperationTypeCreate {
			creationTime.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Observe(time.Now().Sub(instance.Status.LastOperation.StartTime.Time).Seconds())
		} else {
			deletionTime.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Observe(time.Now().Sub(instance.Status.LastOperation.StartTime.Time).Seconds())
		}
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateSucceeded
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	case provisionertypes.ProvisionerOperationStateInProgress:
		log.Info(fmt.Sprintf("Updating LastOperation status to `InProgress` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateInProgress
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	default:
		log.Info(fmt.Sprintf("Updating LastOperation status to `Failed` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		if instance.Status.LastOperation.Type == pksv1alpha1.ClusterLastOperationTypeCreate {
			creationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()
		} else {
			deletionErrorsCounter.WithLabelValues(instance.Namespace, instance.Name, instance.Spec.ProvisionerSpec.Type).Inc()
		}
		r.recorder.Event(instance, corev1.EventTypeWarning, "LastOperationError", provisionerLastOperation.Description)
		instance.Status.LastOperation.State = pksv1alpha1.ClusterLastOperationStateFailed
		return reconcile.Result{Requeue: true, RequeueAfter: 10 * time.Second}, r.Status().Update(context.Background(), instance)
	}
}

func (r *ReconcileCluster) validateSpec(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating Spec for Cluster `%s/%s`", instance.Namespace, instance.Name))

	// Validate the Provisioner specification.
	if err := r.validateProvisionerSpec(instance); err != nil {
		return fmt.Errorf("error validating ProvisionerSpec: %v", err)
	}

	// Validate the Plan reference.
	if err := r.validatePlanRef(instance); err != nil {
		return fmt.Errorf("error validating PlanRef: %v", err)
	}

	// Validate the Kubernetes Profile reference.
	if err := r.validateKubernetesProfileRef(instance); err != nil {
		return fmt.Errorf("error validating KubernetesProfileRef: %v", err)
	}

	return nil
}

func (r *ReconcileCluster) validateProvisionerSpec(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating ProvisionerSpec for Cluster `%s/%s`", instance.Namespace, instance.Name))

	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.ProvisionerSpec.CredentialsSecretRef.Name, Namespace: instance.Namespace}, credentialsSecret); err != nil {
		return err
	}

	return nil
}

func (r *ReconcileCluster) validatePlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating PlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	switch instance.Spec.PlanRef.Kind {
	case pksv1alpha1.AWSPlanKind:
		if err := r.validateAWSPlanRef(instance); err != nil {
			return fmt.Errorf("error validating AWSPlan: %v", err)
		}
	case pksv1alpha1.AzurePlanKind:
		if err := r.validateAzurePlanRef(instance); err != nil {
			return fmt.Errorf("error validating AzurePlan: %v", err)
		}
	case pksv1alpha1.GCPPlanKind:
		if err := r.validateGCPPlanRef(instance); err != nil {
			return fmt.Errorf("error validating GCPPlan: %v", err)
		}
	case pksv1alpha1.VSpherePlanKind:
		if err := r.validateVSpherePlanRef(instance); err != nil {
			return fmt.Errorf("error validating VSpherePlan`: %v", err)
		}
	case pksv1alpha1.ODBPlanKind:
		if err := r.validateODBPlanRef(instance); err != nil {
			return fmt.Errorf("error validating ODBPlan`: %v", err)
		}
	default:
		return fmt.Errorf("plan Kind `%s` not supported", instance.Spec.PlanRef.Kind)
	}

	return nil
}

func (r *ReconcileCluster) validateAWSPlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating AWSPlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	awsPlan := &pksv1alpha1.AWSPlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, awsPlan); err != nil {
		return err
	}

	if !awsplanutils.AreAllAWSPlanConditionsTrue(awsPlan.Status) {
		return fmt.Errorf("plan is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) validateAzurePlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating AzurePlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	azurePlan := &pksv1alpha1.AzurePlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, azurePlan); err != nil {
		return err
	}

	if !azureplanutils.AreAllAzurePlanConditionsTrue(azurePlan.Status) {
		return fmt.Errorf("plan is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) validateGCPPlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating GCPPlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	gcpPlan := &pksv1alpha1.GCPPlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, gcpPlan); err != nil {
		return err
	}

	if !gcpplanutils.AreAllGCPPlanConditionsTrue(gcpPlan.Status) {
		return fmt.Errorf("plan is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) validateVSpherePlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating VSpherePlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return err
	}

	if !vsphereplanutils.AreAllVSpherePlanConditionsTrue(vSpherePlan.Status) {
		return fmt.Errorf("plan is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) validateODBPlanRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating ODBPPlanRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	odbPlan := &pksv1alpha1.ODBPlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, odbPlan); err != nil {
		return err
	}

	if !odbplanutils.AreAllODBPlanConditionsTrue(odbPlan.Status) {
		return fmt.Errorf("plan is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) validateKubernetesProfileRef(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Validating KubernetesProfileRef for Cluster `%s/%s`", instance.Namespace, instance.Name))

	kubernetesProfile := &pksv1alpha1.KubernetesProfile{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.KubernetesProfileRef.Name, Namespace: instance.Spec.KubernetesProfileRef.Namespace}, kubernetesProfile); err != nil {
		return err
	}

	if !kubernetesprofileutils.AreAllKubernetesProfileConditionsTrue(kubernetesProfile.Status) {
		return fmt.Errorf("kubernetes profile is not yet ready")
	}

	return nil
}

func (r *ReconcileCluster) getClusterProvisioner(instance *pksv1alpha1.Cluster) (provisioner.Provisioner, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.ProvisionerSpec.CredentialsSecretRef.Name, Namespace: instance.Namespace}, credentialsSecret); err != nil {
		return nil, fmt.Errorf("error getting provisioner credentials secret: %v", err)
	}

	clusterProvisioner, err := provisioner.New(instance.Spec.ProvisionerSpec.Type, r.Client, credentialsSecret)
	if err != nil {
		return nil, fmt.Errorf("error creating cluster provisioner: %v", err)
	}

	return clusterProvisioner, nil
}

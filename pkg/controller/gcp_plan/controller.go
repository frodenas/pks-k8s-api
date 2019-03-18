/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcpplan

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/gcp"
	gcpplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/gcp_plan/utils"
	"github.com/frodenas/pks-k8s-api/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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
	// GCPPlanFinalizer is set on Reconcile callback.
	GCPPlanFinalizer = "gcpplan_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.gcpplan")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "gcpplan",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "gcpplan",
			Name:      "validation_errors_count",
			Help:      "Total number of validation errors",
		},
		[]string{"namespace", "name"},
	)
)

func init() {
	metrics.Registry.MustRegister(validationsCounter)
	metrics.Registry.MustRegister(validationErrorsCounter)
}

// Add creates a new GCP Plan and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileGCPPlan{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("gcpplan-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("gcpplan-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to GCP Plan
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.GCPPlan{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileGCPPlan{}

// ReconcileGCPPlan reconciles a GCP Plan object
type ReconcileGCPPlan struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the GCP plan for a GCP Plan object and makes changes
// based on the state read and what is in the GCPPlan.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=gcpplans,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=gcpplans/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileGCPPlan) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.GCPPlan{}
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

func (r *ReconcileGCPPlan) reconcile(instance *pksv1alpha1.GCPPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(GCPPlanFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to GCP Plan `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, GCPPlanFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a GCPPlanValidated condition if absent.
	if c := gcpplanutils.GetGCPPlanCondition(instance.Status, pksv1alpha1.GCPPlanValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to GCP Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := gcpplanutils.NewGCPPlanCondition(
			pksv1alpha1.GCPPlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"GCP Plan has not yet been validated",
		)
		gcpplanutils.SetGCPPlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If GCP Plan has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileGCPPlan) validate(instance *pksv1alpha1.GCPPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating GCP Plan `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition if needed.
	if c := gcpplanutils.GetGCPPlanCondition(instance.Status, pksv1alpha1.GCPPlanValidated); c == nil || c.Status != corev1.ConditionFalse {
		log.Info(fmt.Sprintf("Updating `Validated` condition to `False` for GCP Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := gcpplanutils.NewGCPPlanCondition(
			pksv1alpha1.GCPPlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"GCP Plan has not yet been validated",
		)
		gcpplanutils.SetGCPPlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Validate the spec details.
	if err := r.validateSpec(instance); err != nil {
		log.Error(err, fmt.Sprintf("Error validating GCP Plan `%s/%s`", instance.Namespace, instance.Name))
		validationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "ValidationError", err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for GCP Plan `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "GCP Plan has been successfully validated")
	condition := gcpplanutils.NewGCPPlanCondition(
		pksv1alpha1.GCPPlanValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"GCP Plan has been validated",
	)
	gcpplanutils.SetGCPPlanCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileGCPPlan) delete(instance *pksv1alpha1.GCPPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("GCP Plan `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(GCPPlanFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from GCP Plan `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(GCPPlanFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileGCPPlan) listAssociatedClusters(instance *pksv1alpha1.GCPPlan) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.GCPPlanKind {
			if cluster.Spec.PlanRef.Namespace == instance.Namespace && cluster.Spec.PlanRef.Name == instance.Name {
				associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
			}
		}
	}

	return associatedClusters, nil
}

func (r *ReconcileGCPPlan) validateSpec(instance *pksv1alpha1.GCPPlan) error {
	log.Info(fmt.Sprintf("Validating Spec for GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	gcpProject, gcpJSONKey, err := r.getProviderCredentials(instance.Spec.ProviderSpec.CredentialsSecretRef)
	if err != nil {
		return fmt.Errorf("error getting provider credentials: %v", err)
	}

	// Build a GCP Client
	gcpClient, err := gcp.NewClient(gcpProject, gcpJSONKey)
	if err != nil {
		return err
	}

	// Validate the GCP Plan Compute specification.
	if err := r.validateComputeSpec(instance, gcpClient); err != nil {
		return fmt.Errorf("error validating ComputeProfile: %v", err)
	}

	// Validate the GCP Plan Network specification.
	if err := r.validateNetworkSpec(instance, gcpClient); err != nil {
		return fmt.Errorf("error validating NetworkProfile: %v", err)
	}

	// Validate the GCP Plan Storage specification.
	if err := r.validateStorageSpec(instance, gcpClient); err != nil {
		return fmt.Errorf("error validating StorageProfile: %v", err)
	}

	return nil
}

func (r *ReconcileGCPPlan) validateComputeSpec(instance *pksv1alpha1.GCPPlan, gcpClient gcp.Client) error {
	log.Info(fmt.Sprintf("Validating Compute Spec for GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	for _, zone := range instance.Spec.ComputeSpec.MastersSpec.Zones {
		if err := r.validateZoneSpec(instance.Spec.ProviderSpec.Region, zone, gcpClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone, err)
		}
	}

	for _, zone := range instance.Spec.ComputeSpec.WorkersSpec.Zones {
		if err := r.validateZoneSpec(instance.Spec.ProviderSpec.Region, zone, gcpClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone, err)
		}
	}

	return nil
}

func (r *ReconcileGCPPlan) validateZoneSpec(region string, zone string, gcpClient gcp.Client) error {
	_, err := gcpClient.GetZone(region, zone)
	if err != nil {
		return fmt.Errorf("error validating GCP Zone `%s/%s`: %v", region, zone, err)
	}

	return nil
}

func (r *ReconcileGCPPlan) validateNetworkSpec(instance *pksv1alpha1.GCPPlan, gcpClient gcp.Client) error {
	log.Info(fmt.Sprintf("Validating Network Spec for GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	_, err := gcpClient.GetNetwork(instance.Spec.NetworkSpec.Name)
	if err != nil {
		return fmt.Errorf("error validating GCP Network `%s`: %v", instance.Spec.NetworkSpec.Name, err)
	}

	return nil
}

func (r *ReconcileGCPPlan) validateStorageSpec(instance *pksv1alpha1.GCPPlan, gcpClient gcp.Client) error {
	log.Info(fmt.Sprintf("Validating Storage Spec for GCP Plan `%s/%s`", instance.Namespace, instance.Name))

	mastersZones := instance.Spec.ComputeSpec.MastersSpec.Zones
	for _, disk := range instance.Spec.StorageSpec.MastersSpec.Disks {
		for _, zone := range mastersZones {
			if err := r.validateDiskSpec(zone, disk.Type, gcpClient); err != nil {
				return fmt.Errorf("error validating Disk Type `%s`: %v", disk.Type, err)
			}
		}
	}

	workersZones := instance.Spec.ComputeSpec.MastersSpec.Zones
	for _, disk := range instance.Spec.StorageSpec.WorkersSpec.Disks {
		for _, zone := range workersZones {
			if err := r.validateDiskSpec(zone, disk.Type, gcpClient); err != nil {
				return fmt.Errorf("error validating Disk Type `%s`: %v", disk.Type, err)
			}
		}
	}

	return nil
}

func (r *ReconcileGCPPlan) validateDiskSpec(zone string, diskType string, gcpClient gcp.Client) error {
	_, err := gcpClient.GetDiskType(zone, diskType)
	if err != nil {
		return fmt.Errorf("error validating GCP Disk Type `%s` in Zone `%s`: %v", zone, diskType, err)
	}

	return nil
}

func (r *ReconcileGCPPlan) getProviderCredentials(secretRef corev1.SecretReference) (string, string, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: secretRef.Name, Namespace: secretRef.Namespace}, credentialsSecret); err != nil {
		return "", "", err
	}

	gcpProject := utils.GetSecretString(credentialsSecret, pksv1alpha1.GCPProviderCredentialsProjectKey)
	if gcpProject == "" {
		return "", "", fmt.Errorf("GCP Project is blank")
	}

	gcpJSONKey := utils.GetSecretString(credentialsSecret, pksv1alpha1.GCPProviderCredentialsJSONKeyKey)
	if gcpJSONKey == "" {
		return "", "", fmt.Errorf("GCP JSON Key is blank")
	}

	return gcpProject, gcpJSONKey, nil
}

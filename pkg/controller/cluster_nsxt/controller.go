/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package clusternsxt

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/nsxt"
	clusterutils "github.com/frodenas/pks-k8s-api/pkg/controller/cluster/utils"
	"github.com/frodenas/pks-k8s-api/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/networkmanager"
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
	// ClusterNSXTFinalizer is set on Reconcile callback.
	ClusterNSXTFinalizer = "cluster_nsxt_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.cluster.nsxt")

	creationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster_nsxt",
			Name:      "creation_count",
			Help:      "Total number of NSX-T resource creation invocations",
		},
		[]string{"namespace", "name"},
	)

	creationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster_nsxt",
			Name:      "creation_errors_count",
			Help:      "Total number of NSX-T resource creation errors",
		},
		[]string{"namespace", "name"},
	)

	creationTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "cluster_nsxt",
			Name:      "creation_time_seconds",
			Help:      "Time in seconds spent creating NSX-T resources",
		},
		[]string{"namespace", "name"},
	)

	deletionsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster_nsxt",
			Name:      "deletion_count",
			Help:      "Total number of NSX-T resource deletion invocations",
		},
		[]string{"namespace", "name"},
	)

	deletionErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "cluster_nsxt",
			Name:      "deletion_errors_count",
			Help:      "Total number of NSX-T resource deletion errors",
		},
		[]string{"namespace", "name"},
	)

	deletionTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: "cluster_nsxt",
			Name:      "deletion_time_seconds",
			Help:      "Time in seconds spent deleting NSX-T resources",
		},
		[]string{"namespace", "name"},
	)
)

func init() {
	metrics.Registry.MustRegister(creationsCounter)
	metrics.Registry.MustRegister(creationErrorsCounter)
	metrics.Registry.MustRegister(creationTime)
	metrics.Registry.MustRegister(deletionsCounter)
	metrics.Registry.MustRegister(deletionErrorsCounter)
	metrics.Registry.MustRegister(deletionTime)
}

// Add creates a new Cluster NSX-T Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileClusterNSXT{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("cluster-nsxt-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("cluster-nsxt-controller", mgr, controller.Options{Reconciler: r})
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

var _ reconcile.Reconciler = &ReconcileClusterNSXT{}

// ReconcileClusterNSXT reconciles a Cluster NSX-T object
type ReconcileClusterNSXT struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a Cluster object and makes changes based on the state read
// and what is in the Cluster.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileClusterNSXT) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.Cluster{}
	if err := r.Get(context.TODO(), request.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	// Return if the cluster is not using NSX-T
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterNSXTProvisioned); c == nil {
		return reconcile.Result{}, nil
	}

	// Check for deletion.
	if !instance.DeletionTimestamp.IsZero() {
		return r.delete(instance)
	}

	return r.reconcile(instance)
}

func (r *ReconcileClusterNSXT) reconcile(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling Cluster `%s/%s` for NSX-T", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(ClusterNSXTFinalizer) {
		log.Info(fmt.Sprintf("Adding NSX-T finalizer to Cluster `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, ClusterNSXTFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Return if the cluster has not yet been validated.
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterValidated); c == nil || c.Status != corev1.ConditionTrue {
		log.Info(fmt.Sprintf("Cluster `%s/%s` has not yet been validated", instance.Namespace, instance.Name))
		return reconcile.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	// Trigger the creation of NSX-T resources if needed.
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterNSXTProvisioned); c.Status != corev1.ConditionTrue {
		networkInfo, err := r.createNSXTResources(instance)
		if err != nil {
			log.Error(err, fmt.Sprintf("Error provisioning NSX-T resources for Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "NSXTProvisioningError", err.Error())
			// Delete NSX-T resources without checking the result.
			_ = r.deleteNSXTResources(instance)
			return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
		}

		// Update the NSX-T Provisioned condition.
		log.Info(fmt.Sprintf("Updating `NSXTProvisioned` condition to `True` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		r.recorder.Event(instance, corev1.EventTypeNormal, "NSXTProvisioning", "NSX-T resources have been successfully provisioned")
		condition := clusterutils.NewClusterCondition(
			pksv1alpha1.ClusterNSXTProvisioned,
			corev1.ConditionTrue,
			"ProvisioningSuccessful",
			"NSX-T resources have been successfully provisioned",
			networkInfo,
		)
		clusterutils.SetClusterCondition(&instance.Status, *condition)
		if err := r.Status().Update(context.Background(), instance); err != nil {
			log.Info(fmt.Sprintf("Error updating status for Cluster `%s/%s`, deleting NSX-T resources", instance.Namespace, instance.Name))
			// Delete NSX-T resources without checking the result.
			_ = r.deleteNSXTResources(instance)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileClusterNSXT) delete(instance *pksv1alpha1.Cluster) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting Cluster `%s/%s` for NSX-T", instance.Namespace, instance.Name))

	// Return if cluster has not yet been deprovisioned.
	if instance.Status.LastOperation != nil &&
		(instance.Status.LastOperation.Type != pksv1alpha1.ClusterLastOperationTypeDelete ||
			(instance.Status.LastOperation.Type == pksv1alpha1.ClusterLastOperationTypeDelete &&
				instance.Status.LastOperation.State != pksv1alpha1.ClusterLastOperationStateSucceeded)) {
		log.Info(fmt.Sprintf("Cluster `%s/%s` last operation has not yet finished", instance.Namespace, instance.Name))
		return reconcile.Result{Requeue: true, RequeueAfter: 5 * time.Second}, nil
	}

	// Trigger the deletion of NSX-T resources if needed.
	if c := clusterutils.GetClusterCondition(instance.Status, pksv1alpha1.ClusterNSXTProvisioned); c.Status != corev1.ConditionFalse {
		if err := r.deleteNSXTResources(instance); err != nil {
			log.Error(err, fmt.Sprintf("Error deprovisioning NSX-T resources for Cluster `%s/%s`", instance.Namespace, instance.Name))
			r.recorder.Event(instance, corev1.EventTypeWarning, "NSXTDeprovisioningError", err.Error())
			return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
		}

		// Update the NSX-T Provisioned condition.
		log.Info(fmt.Sprintf("Updating `NSXTProvisioned` condition to `False` for Cluster `%s/%s`", instance.Namespace, instance.Name))
		r.recorder.Event(instance, corev1.EventTypeNormal, "NSXTProvisioning", "NSX-T resources have been successfully deprovisioned")
		condition := clusterutils.NewClusterCondition(
			pksv1alpha1.ClusterNSXTProvisioned,
			corev1.ConditionFalse,
			"DeprovisioningSuccessful",
			"NSX-T resources have been successfully deprovisioned",
			"",
		)
		clusterutils.SetClusterCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(ClusterNSXTFinalizer) {
		log.Info(fmt.Sprintf("Removing NSX-T finalizer from Cluster `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(ClusterNSXTFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileClusterNSXT) createNSXTResources(instance *pksv1alpha1.Cluster) (string, error) {
	log.Info(fmt.Sprintf("Creating NSX-T resources for Cluster `%s/%s`", instance.Namespace, instance.Name))
	creationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Retrieve the associated vSphere Plan.
	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return "", err
	}

	// Create a NSX-T Client.
	nsxtClient, err := r.createNSXTClient(instance)
	if err != nil {
		return "", err
	}

	// Create a NSX-T Network Manager.
	nsxtSpec := &networkmanager.NSXTSpec{
		T0RouterID:       vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.T0RouterID,
		IPBlockID:        vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.IPBlockIDs[0],
		FloatingIPPoolID: vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.FloatingIPPoolIDs[0],
		NatMode:          vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.NatMode,
	}
	networkManager, err := nsxtClient.NewNetworkManager(nsxtSpec)
	if err != nil {
		return "", err
	}

	// Create the NSX-T resources.
	creationStartTS := time.Now()
	defer func() {
		creationTime.WithLabelValues(instance.Namespace, instance.Name).Observe(time.Now().Sub(creationStartTS).Seconds())
	}()
	nsxClusterSpec := &networkmanager.NSXTClusterSpec{
		T0RouterID:          vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.T0RouterID,
		IPBlockIDs:          vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.IPBlockIDs,
		LbFloatingIPPoolIDs: vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.FloatingIPPoolIDs,
		NatMode:             &vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.NatMode,
		WithLB:              true,
		LBSize:              strings.ToUpper(vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.LBSize),
	}
	networkInfo, err := nsxtClient.CreateClusterNetwork(networkManager, clusterutils.ClusterName(instance.Namespace, instance.Name), nsxClusterSpec)
	if err != nil {
		creationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		return "", err
	}

	networkInfoRaw, err := json.Marshal(networkInfo)
	if err != nil {
		return "", err
	}

	return string(networkInfoRaw), nil
}

func (r *ReconcileClusterNSXT) deleteNSXTResources(instance *pksv1alpha1.Cluster) error {
	log.Info(fmt.Sprintf("Deleting NSX-T resources for Cluster `%s/%s`", instance.Namespace, instance.Name))
	deletionsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Retrieve the associated vSphere Plan.
	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return err
	}

	// Create a NSX-T Client.
	nsxtClient, err := r.createNSXTClient(instance)
	if err != nil {
		return err
	}

	// Create a NSX-T Network Manager.
	nsxtSpec := &networkmanager.NSXTSpec{
		T0RouterID:       vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.T0RouterID,
		IPBlockID:        vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.IPBlockIDs[0],
		FloatingIPPoolID: vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.FloatingIPPoolIDs[0],
		NatMode:          vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.NatMode,
	}
	networkManager, err := nsxtClient.NewNetworkManager(nsxtSpec)
	if err != nil {
		return err
	}

	// Delete the NSX-T resources.
	deletionStartTS := time.Now()
	defer func() {
		deletionTime.WithLabelValues(instance.Namespace, instance.Name).Observe(time.Now().Sub(deletionStartTS).Seconds())
	}()
	if err = nsxtClient.DeleteClusterNetwork(networkManager, clusterutils.ClusterName(instance.Namespace, instance.Name)); err != nil {
		deletionErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		return err
	}

	return nil
}

func (r *ReconcileClusterNSXT) createNSXTClient(instance *pksv1alpha1.Cluster) (nsxt.Client, error) {
	// Retrieve the associated vSphere Plan.
	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return nil, err
	}

	// Create a NSX-T client.
	nsxtURL, nsxtUsername, nsxtPassword, err := r.getNSXTCredentials(vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.CredentialsSecretRef)
	if err != nil {
		return nil, fmt.Errorf("error getting NSX-T credentials: %v", err)
	}

	nsxtClient, err := nsxt.NewClient(nsxtURL, nsxtUsername, nsxtPassword, vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec.Insecure)
	if err != nil {
		return nil, err
	}

	return nsxtClient, nil
}

func (r *ReconcileClusterNSXT) getNSXTCredentials(secretRef corev1.SecretReference) (string, string, string, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: secretRef.Name, Namespace: secretRef.Namespace}, credentialsSecret); err != nil {
		return "", "", "", err
	}

	nsxtURL := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereNSXTCredentialsNSXTURLKey)
	if nsxtURL == "" {
		return "", "", "", fmt.Errorf("NSX-T Manager URL is blank")
	}

	nsxtUsername := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereNSXTCredentialsNSXTUsernameKey)
	if nsxtUsername == "" {
		return "", "", "", fmt.Errorf("NSX-T Manager Username is blank")
	}

	nsxtPassword := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereNSXTCredentialsNSXTPasswordKey)
	if nsxtPassword == "" {
		return "", "", "", fmt.Errorf("NSX-T Manager Password is blank")
	}

	return nsxtURL, nsxtUsername, nsxtPassword, nil
}

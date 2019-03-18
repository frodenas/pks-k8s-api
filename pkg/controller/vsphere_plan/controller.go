/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vsphereplan

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/nsxt"
	"github.com/frodenas/pks-k8s-api/pkg/client/vcenter"
	vsphereplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/vsphere_plan/utils"
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
	// VSpherePlanFinalizer is set on Reconcile callback.
	VSpherePlanFinalizer = "vsphereplan_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.vsphereplan")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "vsphereplan",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "vsphereplan",
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

// Add creates a new vSphere Plan and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileVSpherePlan{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("vsphereplan-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("vsphereplan-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to vSphere Plan
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.VSpherePlan{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileVSpherePlan{}

// ReconcileVSpherePlan reconciles a vSphere Plan object
type ReconcileVSpherePlan struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the vSphere plan for a vSphere Plan object and makes changes
// based on the state read and what is in the VSpherePlan.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=vsphereplans/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileVSpherePlan) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.VSpherePlan{}
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

func (r *ReconcileVSpherePlan) reconcile(instance *pksv1alpha1.VSpherePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(VSpherePlanFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, VSpherePlanFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a VSpherePlanValidated condition if absent.
	if c := vsphereplanutils.GetVSpherePlanCondition(instance.Status, pksv1alpha1.VSpherePlanValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := vsphereplanutils.NewVSpherePlanCondition(
			pksv1alpha1.VSpherePlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"VSphere Plan has not yet been validated",
		)
		vsphereplanutils.SetVSpherePlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If vSphere Plan has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileVSpherePlan) validate(instance *pksv1alpha1.VSpherePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition if needed.
	if c := vsphereplanutils.GetVSpherePlanCondition(instance.Status, pksv1alpha1.VSpherePlanValidated); c == nil || c.Status != corev1.ConditionFalse {
		log.Info(fmt.Sprintf("Updating `Validated` condition to `False` for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := vsphereplanutils.NewVSpherePlanCondition(
			pksv1alpha1.VSpherePlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"vSphere Plan has not yet been validated",
		)
		vsphereplanutils.SetVSpherePlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Validate the spec details.
	if err := r.validateSpec(instance); err != nil {
		log.Error(err, fmt.Sprintf("Error validating vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
		validationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "ValidationError", err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Update the validation condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "vSphere Plan has been successfully validated")
	condition := vsphereplanutils.NewVSpherePlanCondition(
		pksv1alpha1.VSpherePlanValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"VSphere Plan has been successfully validated",
	)
	vsphereplanutils.SetVSpherePlanCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileVSpherePlan) delete(instance *pksv1alpha1.VSpherePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("vSphere Plan `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(VSpherePlanFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from vSphere Plan `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(VSpherePlanFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileVSpherePlan) listAssociatedClusters(instance *pksv1alpha1.VSpherePlan) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.VSpherePlanKind {
			if cluster.Spec.PlanRef.Namespace == instance.Namespace && cluster.Spec.PlanRef.Name == instance.Name {
				associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
			}
		}
	}

	return associatedClusters, nil
}

func (r *ReconcileVSpherePlan) validateSpec(instance *pksv1alpha1.VSpherePlan) error {
	log.Info(fmt.Sprintf("Validating Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	vcURL, vcUsername, vcPassword, err := r.getProviderCredentials(instance.Spec.ProviderSpec.CredentialsSecretRef)
	if err != nil {
		return fmt.Errorf("error getting provider credentials: %v", err)
	}

	// Build a vCenter Client
	vcClient, err := vcenter.NewClient(
		context.TODO(),
		vcURL,
		vcUsername,
		vcPassword,
		instance.Spec.ProviderSpec.Insecure,
	)
	if err != nil {
		return err
	}

	if !vcClient.IsVC() {
		return fmt.Errorf("unable to connect to vCenter")
	}

	// Validate the vSphere Plan Compute specification.
	if err := r.validateComputeSpec(instance, vcClient); err != nil {
		return fmt.Errorf("error validating ComputeProfile: %v", err)
	}

	// Validate the vSphere Plan Network specification.
	if err := r.validateNetworkSpec(instance, vcClient); err != nil {
		return fmt.Errorf("error validating NetworkProfile: %v", err)
	}

	// Validate the vSphere Plan Storage specification.
	if err := r.validateStorageSpec(instance, vcClient); err != nil {
		return fmt.Errorf("error validating StorageProfile: %v", err)
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateComputeSpec(instance *pksv1alpha1.VSpherePlan, vcClient vcenter.Client) error {
	log.Info(fmt.Sprintf("Validating Compute Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	if _, err := vcClient.GetFolder(instance.Spec.ComputeSpec.MastersSpec.VMFolder); err != nil {
		return fmt.Errorf("error validating vCenter Folder `%s`: %v", instance.Spec.ComputeSpec.MastersSpec.VMFolder, err)

	}

	for _, zone := range instance.Spec.ComputeSpec.MastersSpec.Zones {
		if err := r.validateZoneSpec(zone, vcClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone.Name, err)
		}
	}

	if _, err := vcClient.GetFolder(instance.Spec.ComputeSpec.WorkersSpec.VMFolder); err != nil {
		return fmt.Errorf("error validating vCenter Folder `%s`: %v", instance.Spec.ComputeSpec.WorkersSpec.VMFolder, err)

	}

	for _, zone := range instance.Spec.ComputeSpec.WorkersSpec.Zones {
		if err := r.validateZoneSpec(zone, vcClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone.Name, err)
		}
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateZoneSpec(zone pksv1alpha1.VSphereZoneSpec, vcClient vcenter.Client) error {
	_, err := vcClient.GetDataCenter(zone.Datacenter)
	if err != nil {
		return fmt.Errorf("error validating vCenter Datacenter `%s`: %v", zone.Datacenter, err)
	}

	_, err = vcClient.GetComputeResource(zone.Datacenter, zone.Cluster)
	if err != nil {
		return fmt.Errorf("error validating vCenter Cluster `%s`: %v", zone.Cluster, err)
	}

	if zone.ResourcePool != "" {
		_, err := vcClient.GetResourcePool(zone.Datacenter, zone.ResourcePool)
		if err != nil {
			return fmt.Errorf("error validating vCenter Resource Pool `%s`: %v", zone.ResourcePool, err)
		}
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateNetworkSpec(instance *pksv1alpha1.VSpherePlan, vcClient vcenter.Client) error {
	log.Info(fmt.Sprintf("Validating Network Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	if instance.Spec.NetworkSpec.DVSNetworkSpec != nil {
		if err := r.validateVSphereNetworkSpec(instance, vcClient); err != nil {
			return fmt.Errorf("error validating vSphere Network: %v", err)
		}
	}

	if instance.Spec.NetworkSpec.NSXTNetworkSpec != nil {
		if err := r.validateNSXTNetworkSpec(instance, vcClient); err != nil {
			return fmt.Errorf("error validating NSX-T Network: %v", err)
		}
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateVSphereNetworkSpec(instance *pksv1alpha1.VSpherePlan, vcClient vcenter.Client) error {
	log.Info(fmt.Sprintf("Validating vSphere Network Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	_, err := vcClient.GetNetwork(instance.Spec.NetworkSpec.DVSNetworkSpec.Name)
	if err != nil {
		return fmt.Errorf("error validating vSphere Network name `%s`: %v", instance.Spec.NetworkSpec.DVSNetworkSpec.Name, err)
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateNSXTNetworkSpec(instance *pksv1alpha1.VSpherePlan, vcClient vcenter.Client) error {
	log.Info(fmt.Sprintf("Validating NSX-T Network Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	nsxtURL, nsxtUsername, nsxtPassword, err := r.getNSXTCredentials(instance.Spec.NetworkSpec.NSXTNetworkSpec.CredentialsSecretRef)
	if err != nil {
		return fmt.Errorf("error getting NSX-T credentials: %v", err)
	}

	nsxtClient, err := nsxt.NewClient(nsxtURL, nsxtUsername, nsxtPassword, instance.Spec.NetworkSpec.NSXTNetworkSpec.Insecure)
	if err != nil {
		return err
	}

	if _, err := nsxtClient.GetLogicalRouter(instance.Spec.NetworkSpec.NSXTNetworkSpec.T0RouterID); err != nil {
		return fmt.Errorf("error validating NSX-T T0 Router ID `%s`: %v", instance.Spec.NetworkSpec.NSXTNetworkSpec.T0RouterID, err)
	}

	for _, ipBlockID := range instance.Spec.NetworkSpec.NSXTNetworkSpec.IPBlockIDs {
		if _, err := nsxtClient.GetIPBlock(ipBlockID); err != nil {
			return fmt.Errorf("error validating NSX-T IP Block ID `%s`: %v", ipBlockID, err)
		}
	}

	for _, ipPoolID := range instance.Spec.NetworkSpec.NSXTNetworkSpec.FloatingIPPoolIDs {
		if _, err := nsxtClient.GetIPPool(ipPoolID); err != nil {
			return fmt.Errorf("error validating NSX-T Floating IP Pool ID `%s`: %v", ipPoolID, err)
		}
	}

	return nil
}

func (r *ReconcileVSpherePlan) validateStorageSpec(instance *pksv1alpha1.VSpherePlan, vcClient vcenter.Client) error {
	log.Info(fmt.Sprintf("Validating Storage Spec for vSphere Plan `%s/%s`", instance.Namespace, instance.Name))

	if _, err := vcClient.GetDatastore(instance.Spec.StorageSpec.MastersSpec.Datastore); err != nil {
		return fmt.Errorf("error validating vCenter Datastore `%s`: %v", instance.Spec.StorageSpec.MastersSpec.Datastore, err)

	}

	if _, err := vcClient.GetDatastore(instance.Spec.StorageSpec.WorkersSpec.Datastore); err != nil {
		return fmt.Errorf("error validating vCenter Datastore `%s`: %v", instance.Spec.StorageSpec.WorkersSpec.Datastore, err)

	}

	return nil
}

func (r *ReconcileVSpherePlan) getProviderCredentials(secretRef corev1.SecretReference) (string, string, string, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: secretRef.Name, Namespace: secretRef.Namespace}, credentialsSecret); err != nil {
		return "", "", "", err
	}

	vCenterURL := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereProviderCredentialsVCenterURLKey)
	if vCenterURL == "" {
		return "", "", "", fmt.Errorf("vCenter URL is blank")
	}

	vCenterUsername := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereProviderCredentialsVCenterUsernameKey)
	if vCenterUsername == "" {
		return "", "", "", fmt.Errorf("vCenter Username is blank")
	}

	vCenterPassword := utils.GetSecretString(credentialsSecret, pksv1alpha1.VSphereProviderCredentialsVCenterPasswordKey)
	if vCenterPassword == "" {
		return "", "", "", fmt.Errorf("vCenter Password is blank")
	}

	return vCenterURL, vCenterUsername, vCenterPassword, nil
}

func (r *ReconcileVSpherePlan) getNSXTCredentials(secretRef corev1.SecretReference) (string, string, string, error) {
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

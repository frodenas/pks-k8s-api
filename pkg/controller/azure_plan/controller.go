/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azureplan

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/azure"
	azureplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/azure_plan/utils"
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
	// AzurePlanFinalizer is set on Reconcile callback.
	AzurePlanFinalizer = "azureplan_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.azuresplan")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "azureplan",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "azureplan",
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

// Add creates a new Azure Plan and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAzurePlan{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("azureplan-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("azureplan-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Azure Plan
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.AzurePlan{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileAzurePlan{}

// ReconcileAzurePlan reconciles a Azure Plan object
type ReconcileAzurePlan struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the Azure plan for a Azure Plan object and makes changes
// based on the state read and what is in the AzurePlan.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=azureplans,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=azureplans/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileAzurePlan) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.AzurePlan{}
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

func (r *ReconcileAzurePlan) reconcile(instance *pksv1alpha1.AzurePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(AzurePlanFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to Azure Plan `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, AzurePlanFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a AzurePlanValidated condition if absent.
	if c := azureplanutils.GetAzurePlanCondition(instance.Status, pksv1alpha1.AzurePlanValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to Azure Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := azureplanutils.NewAzurePlanCondition(
			pksv1alpha1.AzurePlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"Azure Plan has not yet been validated",
		)
		azureplanutils.SetAzurePlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If Azure Plan has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileAzurePlan) validate(instance *pksv1alpha1.AzurePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating Azure Plan `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition if needed.
	if c := azureplanutils.GetAzurePlanCondition(instance.Status, pksv1alpha1.AzurePlanValidated); c == nil || c.Status != corev1.ConditionFalse {
		log.Info(fmt.Sprintf("Updating `Validated` condition to `False` for Azure Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := azureplanutils.NewAzurePlanCondition(
			pksv1alpha1.AzurePlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"Azure Plan has not yet been validated",
		)
		azureplanutils.SetAzurePlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Validate the spec details.
	if err := r.validateSpec(instance); err != nil {
		log.Error(err, fmt.Sprintf("Error validating Azure Plan `%s/%s`", instance.Namespace, instance.Name))
		validationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "ValidationError", err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for Azure Plan `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "Azure Plan has been successfully validated")
	condition := azureplanutils.NewAzurePlanCondition(
		pksv1alpha1.AzurePlanValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"AzurePlan has been validated",
	)
	azureplanutils.SetAzurePlanCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileAzurePlan) delete(instance *pksv1alpha1.AzurePlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("Azure Plan `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(AzurePlanFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from Azure Plan `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(AzurePlanFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileAzurePlan) listAssociatedClusters(instance *pksv1alpha1.AzurePlan) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.AzurePlanKind {
			if cluster.Spec.PlanRef.Namespace == instance.Namespace && cluster.Spec.PlanRef.Name == instance.Name {
				associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
			}
		}
	}

	return associatedClusters, nil
}

func (r *ReconcileAzurePlan) validateSpec(instance *pksv1alpha1.AzurePlan) error {
	log.Info(fmt.Sprintf("Validating Spec for Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	azureSubscriptionID, azureTenantID, azureClientID, azureClientSecret, err := r.getProviderCredentials(instance.Spec.ProviderSpec.CredentialsSecretRef)
	if err != nil {
		return fmt.Errorf("error getting provider credentials: %v", err)
	}

	// Build a Azure Client
	azureClient, err := azure.NewClient(
		context.TODO(),
		instance.Spec.ProviderSpec.Environment,
		instance.Spec.ProviderSpec.Location,
		instance.Spec.ProviderSpec.ResourceGroup,
		azureSubscriptionID,
		azureTenantID,
		azureClientID,
		azureClientSecret,
	)
	if err != nil {
		return err
	}

	// Validate the Azure Plan Compute specification.
	if err := r.validateComputeSpec(instance, azureClient); err != nil {
		return fmt.Errorf("error validating ComputeProfile: %v", err)
	}

	// Validate the Azure Plan Network specification.
	if err := r.validateNetworkSpec(instance, azureClient); err != nil {
		return fmt.Errorf("error validating NetworkProfile: %v", err)
	}

	// Validate the Azure Plan Storage specification.
	if err := r.validateStorageSpec(instance, azureClient); err != nil {
		return fmt.Errorf("error validating StorageProfile: %v", err)
	}

	return nil
}

func (r *ReconcileAzurePlan) validateComputeSpec(instance *pksv1alpha1.AzurePlan, azureClient azure.Client) error {
	log.Info(fmt.Sprintf("Validating Compute Spec for Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	// noop

	return nil
}

func (r *ReconcileAzurePlan) validateNetworkSpec(instance *pksv1alpha1.AzurePlan, azureClient azure.Client) error {
	log.Info(fmt.Sprintf("Validating Network Spec for Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	if _, err := azureClient.GetVnet(instance.Spec.NetworkSpec.Vnet); err != nil {
		return fmt.Errorf("error validating Azure Vnet `%s`: %v", instance.Spec.NetworkSpec.Vnet, err)
	}

	if _, err := azureClient.GetSubnet(instance.Spec.NetworkSpec.Vnet, instance.Spec.NetworkSpec.Subnet); err != nil {
		return fmt.Errorf("error validating Azure Subnet `%s`: %v", instance.Spec.NetworkSpec.Subnet, err)
	}

	return nil
}

func (r *ReconcileAzurePlan) validateStorageSpec(instance *pksv1alpha1.AzurePlan, azureClient azure.Client) error {
	log.Info(fmt.Sprintf("Validating Storage Spec for Azure Plan `%s/%s`", instance.Namespace, instance.Name))

	// noop

	return nil
}

func (r *ReconcileAzurePlan) getProviderCredentials(secretRef corev1.SecretReference) (string, string, string, string, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: secretRef.Name, Namespace: secretRef.Namespace}, credentialsSecret); err != nil {
		return "", "", "", "", err
	}

	azureSubscriptionID := utils.GetSecretString(credentialsSecret, pksv1alpha1.AzureProviderCredentialsSubscriptionIDKey)
	if azureSubscriptionID == "" {
		return "", "", "", "", fmt.Errorf("Azure Subscription ID is blank")
	}

	azureTenantID := utils.GetSecretString(credentialsSecret, pksv1alpha1.AzureProviderCredentialsTenantIDKey)
	if azureTenantID == "" {
		return "", "", "", "", fmt.Errorf("Azure Tenant ID is blank")
	}

	azureClientID := utils.GetSecretString(credentialsSecret, pksv1alpha1.AzureProviderCredentialsClientIDKey)
	if azureClientID == "" {
		return "", "", "", "", fmt.Errorf("Azure Client ID is blank")
	}

	azureClientSecret := utils.GetSecretString(credentialsSecret, pksv1alpha1.AzureProviderCredentialsClientSecretyKey)
	if azureClientSecret == "" {
		return "", "", "", "", fmt.Errorf("Azure Client Secret is blank")
	}

	return azureSubscriptionID, azureTenantID, azureClientID, azureClientSecret, nil
}

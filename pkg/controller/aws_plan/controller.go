/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package awsplan

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/aws"
	awsplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/aws_plan/utils"
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
	// AWSPlanFinalizer is set on Reconcile callback.
	AWSPlanFinalizer = "awsplan_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.awsplan")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "awsplan",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "awsplan",
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

// Add creates a new AWS Plan and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAWSPlan{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("awsplan-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("awsplan-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to AWS Plan
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.AWSPlan{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileAWSPlan{}

// ReconcileAWSPlan reconciles a AWS Plan object
type ReconcileAWSPlan struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the AWS plan for a AWS Plan object and makes changes
// based on the state read and what is in the AWSPlan.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=awsplans,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=awsplans/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileAWSPlan) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.AWSPlan{}
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

func (r *ReconcileAWSPlan) reconcile(instance *pksv1alpha1.AWSPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(AWSPlanFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to AWS Plan `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, AWSPlanFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a AWSPlanValidated condition if absent.
	if c := awsplanutils.GetAWSPlanCondition(instance.Status, pksv1alpha1.AWSPlanValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to AWS Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := awsplanutils.NewAWSPlanCondition(
			pksv1alpha1.AWSPlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"AWS Plan has not yet been validated",
		)
		awsplanutils.SetAWSPlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If AWS Plan has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileAWSPlan) validate(instance *pksv1alpha1.AWSPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating AWS Plan `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition if needed.
	if c := awsplanutils.GetAWSPlanCondition(instance.Status, pksv1alpha1.AWSPlanValidated); c == nil || c.Status != corev1.ConditionFalse {
		log.Info(fmt.Sprintf("Updating `Validated` condition to `False` for AWS Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := awsplanutils.NewAWSPlanCondition(
			pksv1alpha1.AWSPlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"AWS Plan has not yet been validated",
		)
		awsplanutils.SetAWSPlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// Validate the spec details.
	if err := r.validateSpec(instance); err != nil {
		log.Error(err, fmt.Sprintf("Error validating AWS Plan `%s/%s`", instance.Namespace, instance.Name))
		validationErrorsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()
		r.recorder.Event(instance, corev1.EventTypeWarning, "ValidationError", err.Error())
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for AWS Plan `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "AWS Plan has been successfully validated")
	condition := awsplanutils.NewAWSPlanCondition(
		pksv1alpha1.AWSPlanValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"AWS Plan has been validated",
	)
	awsplanutils.SetAWSPlanCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileAWSPlan) delete(instance *pksv1alpha1.AWSPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("AWS Plan `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(AWSPlanFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from AWS Plan `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(AWSPlanFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileAWSPlan) listAssociatedClusters(instance *pksv1alpha1.AWSPlan) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.AWSPlanKind {
			if cluster.Spec.PlanRef.Namespace == instance.Namespace && cluster.Spec.PlanRef.Name == instance.Name {
				associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
			}
		}
	}

	return associatedClusters, nil
}

func (r *ReconcileAWSPlan) validateSpec(instance *pksv1alpha1.AWSPlan) error {
	log.Info(fmt.Sprintf("Validating Spec for AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	awsAccessKey, awsSecretAccessKey, err := r.getProviderCredentials(instance.Spec.ProviderSpec.CredentialsSecretRef)
	if err != nil {
		return fmt.Errorf("error getting provider credentials: %v", err)
	}

	// Build a AWS Client
	awsClient, err := aws.NewClient(awsAccessKey, awsSecretAccessKey, instance.Spec.ProviderSpec.Region)
	if err != nil {
		return err
	}

	// Validate the AWS Plan Compute specification.
	if err := r.validateComputeSpec(instance, awsClient); err != nil {
		return fmt.Errorf("error validating ComputeProfile: %v", err)
	}

	// Validate the AWS Plan Network specification.
	if err := r.validateNetworkSpec(instance, awsClient); err != nil {
		return fmt.Errorf("error validating NetworkProfile: %v", err)
	}

	// Validate the AWS Plan Storage specification.
	if err := r.validateStorageSpec(instance, awsClient); err != nil {
		return fmt.Errorf("error validating StorageProfile: %v", err)
	}

	return nil
}

func (r *ReconcileAWSPlan) validateComputeSpec(instance *pksv1alpha1.AWSPlan, awsClient aws.Client) error {
	log.Info(fmt.Sprintf("Validating Compute Spec for AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	for _, zone := range instance.Spec.ComputeSpec.MastersSpec.Zones {
		if err := r.validateZoneSpec(zone, awsClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone, err)
		}
	}

	for _, zone := range instance.Spec.ComputeSpec.WorkersSpec.Zones {
		if err := r.validateZoneSpec(zone, awsClient); err != nil {
			return fmt.Errorf("error validating Zone `%s`: %v", zone, err)
		}
	}

	return nil
}

func (r *ReconcileAWSPlan) validateZoneSpec(zone string, awsClient aws.Client) error {
	_, err := awsClient.GetAvailabilityZone(zone)
	if err != nil {
		return fmt.Errorf("error validating AWS Zone `%s`: %v", zone, err)
	}

	return nil
}

func (r *ReconcileAWSPlan) validateNetworkSpec(instance *pksv1alpha1.AWSPlan, awsClient aws.Client) error {
	log.Info(fmt.Sprintf("Validating Network Spec for AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	_, err := awsClient.GetVPC(instance.Spec.NetworkSpec.VpcID)
	if err != nil {
		return fmt.Errorf("error validating AWS VPC ID `%s`: %v", instance.Spec.NetworkSpec.VpcID, err)
	}

	return nil
}

func (r *ReconcileAWSPlan) validateStorageSpec(instance *pksv1alpha1.AWSPlan, awsClient aws.Client) error {
	log.Info(fmt.Sprintf("Validating Storage Spec for AWS Plan `%s/%s`", instance.Namespace, instance.Name))

	// noop

	return nil
}

func (r *ReconcileAWSPlan) getProviderCredentials(secretRef corev1.SecretReference) (string, string, error) {
	credentialsSecret := &corev1.Secret{}
	if err := r.Get(context.TODO(), apitypes.NamespacedName{Name: secretRef.Name, Namespace: secretRef.Namespace}, credentialsSecret); err != nil {
		return "", "", err
	}

	awsAccessKey := utils.GetSecretString(credentialsSecret, pksv1alpha1.AWSProviderCredentialsAccessKeyKey)
	if awsAccessKey == "" {
		return "", "", fmt.Errorf("AWS Access Key is blank")
	}

	awsSecretAccessKey := utils.GetSecretString(credentialsSecret, pksv1alpha1.AWSProviderCredentialsSecretAccessKeyKey)
	if awsSecretAccessKey == "" {
		return "", "", fmt.Errorf("AWS Secret Access Key is blank")
	}

	return awsAccessKey, awsSecretAccessKey, nil
}

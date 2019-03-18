/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package kubernetesprofile

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	kubernetesprofileutils "github.com/frodenas/pks-k8s-api/pkg/controller/kubernetes_profile/utils"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
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
	// KubernetesProfileFinalizer is set on Reconcile callback.
	KubernetesProfileFinalizer = "kubernetesprofile_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.kubernetesprofile")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "kubernetesprofile",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "kubernetesprofile",
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

// Add creates a new Kubernetes Profile Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKubernetesProfile{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("kubernetesprofile-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("kubernetesprofile-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Kubernetes Profiles
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.KubernetesProfile{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileKubernetesProfile{}

// ReconcileKubernetesProfile reconciles a Kubernetes Profile object
type ReconcileKubernetesProfile struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the kubernetes profile for a Kubernetes Profile object and makes changes
// based on the state read and what is in the KubernetesProfile.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=kubernetesprofiles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=kubernetesprofiles/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileKubernetesProfile) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.KubernetesProfile{}
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

func (r *ReconcileKubernetesProfile) reconcile(instance *pksv1alpha1.KubernetesProfile) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(KubernetesProfileFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, KubernetesProfileFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a KubernetesProfileValidated condition if absent.
	if c := kubernetesprofileutils.GetKubernetesProfileCondition(instance.Status, pksv1alpha1.KubernetesProfileValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))
		condition := kubernetesprofileutils.NewKubernetesProfileCondition(
			pksv1alpha1.KubernetesProfileValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"Kubernetes Profile has not yet been validated",
		)
		kubernetesprofileutils.SetKubernetesProfileCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If Kubernetes Profile has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileKubernetesProfile) validate(instance *pksv1alpha1.KubernetesProfile) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "Kubernetes Profile has been successfully validated")
	condition := kubernetesprofileutils.NewKubernetesProfileCondition(
		pksv1alpha1.KubernetesProfileValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"Kubernetes Profile has been validated",
	)
	kubernetesprofileutils.SetKubernetesProfileCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileKubernetesProfile) delete(instance *pksv1alpha1.KubernetesProfile) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("Kubernetes Profile `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(KubernetesProfileFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(KubernetesProfileFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileKubernetesProfile) listAssociatedClusters(instance *pksv1alpha1.KubernetesProfile) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with Kubernetes Profile `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.KubernetesProfileRef.Namespace == instance.Namespace && cluster.Spec.KubernetesProfileRef.Name == instance.Name {
			associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
		}
	}

	return associatedClusters, nil
}

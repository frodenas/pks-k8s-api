/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package odbplan

import (
	"context"
	"fmt"
	"strings"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	odbplanutils "github.com/frodenas/pks-k8s-api/pkg/controller/odb_plan/utils"
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
	// ODBPlanFinalizer is set on Reconcile callback.
	ODBPlanFinalizer = "odbplan_controller.pks.vcna.io"
)

var (
	log = logf.Log.WithName("controller.odbplan")

	validationsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "odbplan",
			Name:      "validation_count",
			Help:      "Total number of validations",
		},
		[]string{"namespace", "name"},
	)

	validationErrorsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "odbplan",
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

// Add creates a new ODB Plan and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileODBPlan{
		Client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("odbplan-controller"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("odbplan-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to ODB Plan
	err = c.Watch(&source.Kind{Type: &pksv1alpha1.ODBPlan{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileODBPlan{}

// ReconcileODBPlan reconciles a ODB Plan object
type ReconcileODBPlan struct {
	client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the ODB plan for a ODB Plan object and makes changes
// based on the state read and what is in the ODBPlan.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=pks.vcna.io,resources=odbplans,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pks.vcna.io,resources=odbplans/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters,verbs=get;list;watch;
// +kubebuilder:rbac:groups=pks.vcna.io,resources=clusters/status,verbs=get;
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch
func (r *ReconcileODBPlan) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &pksv1alpha1.ODBPlan{}
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

func (r *ReconcileODBPlan) reconcile(instance *pksv1alpha1.ODBPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Reconciling ODB Plan `%s/%s`", instance.Namespace, instance.Name))

	// Add a finalizer if absent.
	finalizers := sets.NewString(instance.Finalizers...)
	if !finalizers.Has(ODBPlanFinalizer) {
		log.Info(fmt.Sprintf("Adding finalizer to ODB Plan `%s/%s`", instance.Namespace, instance.Name))
		instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, ODBPlanFinalizer)
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	// Add a ODBPlanValidated condition if absent.
	if c := odbplanutils.GetODBPlanCondition(instance.Status, pksv1alpha1.ODBPlanValidated); c == nil {
		log.Info(fmt.Sprintf("Adding `Validated` condition to ODB Plan `%s/%s`", instance.Namespace, instance.Name))
		condition := odbplanutils.NewODBPlanCondition(
			pksv1alpha1.ODBPlanValidated,
			corev1.ConditionFalse,
			"ValidationPending",
			"ODB Plan has not yet been validated",
		)
		odbplanutils.SetODBPlanCondition(&instance.Status, *condition)
		return reconcile.Result{}, r.Status().Update(context.Background(), instance)
	}

	// If ODB Plan has changes, validate it.
	if instance.Status.ObservedGeneration != instance.ObjectMeta.Generation {
		return r.validate(instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileODBPlan) validate(instance *pksv1alpha1.ODBPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Validating ODB Plan `%s/%s`", instance.Namespace, instance.Name))
	validationsCounter.WithLabelValues(instance.Namespace, instance.Name).Inc()

	// Update the validated condition.
	log.Info(fmt.Sprintf("Updating `Validated` condition to `True` for ODB Plan `%s/%s`", instance.Namespace, instance.Name))
	r.recorder.Event(instance, corev1.EventTypeNormal, "Validation", "ODB Plan has been successfully validated")
	condition := odbplanutils.NewODBPlanCondition(
		pksv1alpha1.ODBPlanValidated,
		corev1.ConditionTrue,
		"ValidationSuccessful",
		"ODB Plan has been validated",
	)
	odbplanutils.SetODBPlanCondition(&instance.Status, *condition)

	// Update the observed generation.
	instance.Status.ObservedGeneration = instance.ObjectMeta.Generation
	return reconcile.Result{}, r.Status().Update(context.Background(), instance)
}

func (r *ReconcileODBPlan) delete(instance *pksv1alpha1.ODBPlan) (reconcile.Result, error) {
	log.Info(fmt.Sprintf("Deleting ODB Plan `%s/%s`", instance.Namespace, instance.Name))

	// Check if there are clusters referencing the object.
	clusters, err := r.listAssociatedClusters(instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(clusters) > 0 {
		msg := fmt.Sprintf("ODB Plan `%s/%s` cannot be delete because it is still in use by Cluster(s): %s", instance.Namespace, instance.Name, strings.Join(clusters, ","))
		log.Info(msg)
		r.recorder.Event(instance, corev1.EventTypeWarning, "InUse", msg)
		return reconcile.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
	}

	// Remove the finalizer if present.
	finalizers := sets.NewString(instance.Finalizers...)
	if finalizers.Has(ODBPlanFinalizer) {
		log.Info(fmt.Sprintf("Removing finalizer from ODB Plan `%s/%s`", instance.Namespace, instance.Name))
		finalizers.Delete(ODBPlanFinalizer)
		instance.Finalizers = finalizers.UnsortedList()
		return reconcile.Result{}, r.Update(context.Background(), instance)
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileODBPlan) listAssociatedClusters(instance *pksv1alpha1.ODBPlan) ([]string, error) {
	log.Info(fmt.Sprintf("Listing Clusters associated with ODB Plan `%s/%s`", instance.Namespace, instance.Name))

	var associatedClusters []string
	clusters := &pksv1alpha1.ClusterList{}
	if err := r.List(context.TODO(), &client.ListOptions{}, clusters); err != nil {
		return nil, err
	}

	for _, cluster := range clusters.Items {
		if cluster.Spec.PlanRef.Kind == pksv1alpha1.ODBPlanKind {
			if cluster.Spec.PlanRef.Namespace == instance.Namespace && cluster.Spec.PlanRef.Name == instance.Name {
				associatedClusters = append(associatedClusters, fmt.Sprintf("%s/%s", cluster.Namespace, cluster.Name))
			}
		}
	}

	return associatedClusters, nil
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package cluster

import (
	"testing"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-cluster-name", Namespace: "fake-cluster-namespace"}}
var clusterKey = types.NamespacedName{Name: "fake-cluster-name", Namespace: "fake-cluster-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-cluster-name",
			Namespace: "fake-cluster-namespace",
		},
		Spec: pksv1alpha1.ClusterSpec{
			ExternalHostname:  "fake-cluster-external-hostname",
			Description:       "fake-cluster-description",
			NumWorkerReplicas: 32,
			ProvisionerSpec: pksv1alpha1.ProvisionerSpec{
				Type: "DUMMY",
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-provisioner-credentials-secret-name",
					Namespace: "fake-provisioner-credentials-secret-namespace",
				},
			},
			PlanRef: corev1.ObjectReference{
				Kind:      "fake-plan-kind",
				Name:      "fake-plan-name",
				Namespace: "fake-plan-namespace",
			},
			KubernetesProfileRef: corev1.ObjectReference{
				Name:      "fake-kubernetes-profile-name",
				Namespace: "fake-kubernetes-profile-namespace",
			},
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Cluster object and expect the Reconcile and Deployment to be created
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	cluster := &pksv1alpha1.Cluster{}
	g.Eventually(func() error { return c.Get(context.TODO(), clusterKey, cluster) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), clusterKey, cluster)
		return cluster.Finalizers
	}, timeout).Should(gomega.ContainElement(ClusterFinalizer))
}

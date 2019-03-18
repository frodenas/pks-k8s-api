/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package kubernetesprofile

import (
	"testing"
	"time"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-kubernetes-profile-name", Namespace: "fake-kubernetes-profile-namespace"}}
var kubernetesProfileKey = types.NamespacedName{Name: "fake-kubernetes-profile-name", Namespace: "fake-kubernetes-profile-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.KubernetesProfile{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-kubernetes-profile-name",
			Namespace: "fake-kubernetes-profile-namespace",
		},
		Spec: pksv1alpha1.KubernetesProfileSpec{
			Description: "fake-kubernetes-profile-description",
			Versions: pksv1alpha1.KubernetesVersionsSpec{
				Master: "fake-kubernetes-profile-version-master",
				Worker: "fake-kubernetes-profile-version-worker",
			},
			NetworkSpec: pksv1alpha1.KubernetesNetworkSpec{
				ServiceDomain:      "fake-kubernetes-profile-service-domain",
				ServicesCIDRBlocks: []string{"20.0.0.0/24"},
				PodsCIDRBlocks:     []string{"10.0.0.0/24"},
			},
		},
	}

	// Setup the Manager and Controller. Wrap the Controller Reconcile function so it writes each request to a
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

	// Create the Kubernetes Profile object
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	kubernetesProfile := &pksv1alpha1.KubernetesProfile{}
	g.Eventually(func() error { return c.Get(context.TODO(), kubernetesProfileKey, kubernetesProfile) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), kubernetesProfileKey, kubernetesProfile)
		return kubernetesProfile.Finalizers
	}, timeout).Should(gomega.ContainElement(KubernetesProfileFinalizer))
}

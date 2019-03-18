/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcpplan

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

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-gcp-plan-name", Namespace: "fake-gcp-plan-namespace"}}
var gcpPlanKey = types.NamespacedName{Name: "fake-gcp-plan-name", Namespace: "fake-gcp-plan-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.GCPPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-gcp-plan-name",
			Namespace: "fake-gcp-plan-namespace",
		},
		Spec: pksv1alpha1.GCPPlanSpec{
			Description: "fake-gcp-plan-description",
			ProviderSpec: pksv1alpha1.GCPProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-gcp-provider-credentials-secret-name",
					Namespace: "fake-gcp-provider-credentials-secret-namespace",
				},
				Region: "fake-gcp-region",
			},
			ComputeSpec: pksv1alpha1.GCPComputeSpec{
				MastersSpec: pksv1alpha1.GCPComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-gcp-zone-1"},
				},
				WorkersSpec: pksv1alpha1.GCPComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-gcp-zone-1"},
				},
			},
			NetworkSpec: pksv1alpha1.GCPNetworkSpec{
				Name: "fake-gcp-network-name",
				DNS:  []string{"fake-gcp-dns-server"},
			},
			StorageSpec: pksv1alpha1.GCPStorageSpec{
				MastersSpec: pksv1alpha1.GCPStorageMastersSpec{
					Disks: []pksv1alpha1.GCPDiskSpec{
						pksv1alpha1.GCPDiskSpec{
							SizeGb: 1,
							Type:   "fake-gcp-disk-type-1",
						},
					},
				},
				WorkersSpec: pksv1alpha1.GCPStorageWorkersSpec{
					Disks: []pksv1alpha1.GCPDiskSpec{
						pksv1alpha1.GCPDiskSpec{
							SizeGb: 1,
							Type:   "fake-gcp-disk-type-1",
						},
					},
				},
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

	// Create the GCP Plan object
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	gcpPlan := &pksv1alpha1.GCPPlan{}
	g.Eventually(func() error { return c.Get(context.TODO(), gcpPlanKey, gcpPlan) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), gcpPlanKey, gcpPlan)
		return gcpPlan.Finalizers
	}, timeout).Should(gomega.ContainElement(GCPPlanFinalizer))
}

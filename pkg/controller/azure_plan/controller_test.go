/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azureplan

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

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-azure-plan-name", Namespace: "fake-azure-plan-namespace"}}
var azurePlanKey = types.NamespacedName{Name: "fake-azure-plan-name", Namespace: "fake-azure-plan-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.AzurePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-azure-plan-name",
			Namespace: "fake-azure-plan-namespace",
		},
		Spec: pksv1alpha1.AzurePlanSpec{
			Description: "fake-azure-plan-description",
			ProviderSpec: pksv1alpha1.AzureProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-azure-provider-credentials-secret-name",
					Namespace: "fake-azure-provider-credentials-secret-namespace",
				},
				Location:      "fake-azure-location",
				ResourceGroup: "fake-azure-resource-group",
			},
			ComputeSpec: pksv1alpha1.AzureComputeSpec{
				MastersSpec: pksv1alpha1.AzureComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
				},
				WorkersSpec: pksv1alpha1.AzureComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
				},
			},
			NetworkSpec: pksv1alpha1.AzureNetworkSpec{
				Vnet:   "fake-azure-vnet",
				Subnet: "fake-azure-subnet",
				DNS:    []string{"fake-azure-dns-server"},
			},
			StorageSpec: pksv1alpha1.AzureStorageSpec{
				MastersSpec: pksv1alpha1.AzureStorageMastersSpec{
					Disks: []pksv1alpha1.AzureDiskSpec{
						pksv1alpha1.AzureDiskSpec{
							SizeGb:             1,
							StorageAccountType: "Standard_LRS",
							Caching:            "None",
						},
					},
				},
				WorkersSpec: pksv1alpha1.AzureStorageWorkersSpec{
					Disks: []pksv1alpha1.AzureDiskSpec{
						pksv1alpha1.AzureDiskSpec{
							SizeGb:             1,
							StorageAccountType: "Premium_LRS",
							Caching:            "ReadOnly",
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

	// Create the Azure Plan object
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	azurePlan := &pksv1alpha1.AzurePlan{}
	g.Eventually(func() error { return c.Get(context.TODO(), azurePlanKey, azurePlan) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), azurePlanKey, azurePlan)
		return azurePlan.Finalizers
	}, timeout).Should(gomega.ContainElement(AzurePlanFinalizer))
}

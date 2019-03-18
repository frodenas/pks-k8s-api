/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageAzurePlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-azure-plan-name",
		Namespace: "fake-azure-plan-namespace",
	}
	created := &AzurePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-azure-plan-name",
			Namespace: "fake-azure-plan-namespace",
		},
		Spec: AzurePlanSpec{
			Description: "fake-azure-plan-description",
			ProviderSpec: AzureProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-azure-provider-credentials-secret-name",
					Namespace: "fake-azure-provider-credentials-secret-namespace",
				},
				Location:      "fake-azure-location",
				ResourceGroup: "fake-azure-resource-group",
			},
			ComputeSpec: AzureComputeSpec{
				MastersSpec: AzureComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
				},
				WorkersSpec: AzureComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
				},
			},
			NetworkSpec: AzureNetworkSpec{
				Vnet:   "fake-azure-vnet",
				Subnet: "fake-azure-subnet",
				DNS:    []string{"fake-azure-dns-server"},
			},
			StorageSpec: AzureStorageSpec{
				MastersSpec: AzureStorageMastersSpec{
					Disks: []AzureDiskSpec{
						AzureDiskSpec{
							SizeGb:             1,
							StorageAccountType: "Standard_LRS",
							Caching:            "None",
						},
					},
				},
				WorkersSpec: AzureStorageWorkersSpec{
					Disks: []AzureDiskSpec{
						AzureDiskSpec{
							SizeGb:             1,
							StorageAccountType: "Premium_LRS",
							Caching:            "ReadOnly",
						},
					},
				},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &AzurePlan{}
	g.Expect(c.Create(context.TODO(), created)).To(gomega.Succeed())

	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.Succeed())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).To(gomega.Succeed())

	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.Succeed())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).To(gomega.Succeed())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}

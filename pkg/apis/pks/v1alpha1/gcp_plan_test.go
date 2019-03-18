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

func TestStorageGCPPlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-gcp-plan-name",
		Namespace: "fake-gcp-plan-namespace",
	}
	created := &GCPPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-gcp-plan-name",
			Namespace: "fake-gcp-plan-namespace",
		},
		Spec: GCPPlanSpec{
			Description: "fake-gcp-plan-description",
			ProviderSpec: GCPProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-gcp-provider-credentials-secret-name",
					Namespace: "fake-gcp-provider-credentials-secret-namespace",
				},
				Region: "fake-gcp-region",
			},
			ComputeSpec: GCPComputeSpec{
				MastersSpec: GCPComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-gcp-zone-1"},
				},
				WorkersSpec: GCPComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-gcp-zone-1"},
				},
			},
			NetworkSpec: GCPNetworkSpec{
				Name: "fake-gcp-network-name",
				DNS:  []string{"fake-gcp-dns-server"},
			},
			StorageSpec: GCPStorageSpec{
				MastersSpec: GCPStorageMastersSpec{
					Disks: []GCPDiskSpec{
						GCPDiskSpec{
							SizeGb: 1,
							Type:   "fake-gcp-disk-type-1",
						},
					},
				},
				WorkersSpec: GCPStorageWorkersSpec{
					Disks: []GCPDiskSpec{
						GCPDiskSpec{
							SizeGb: 1,
							Type:   "fake-gcp-disk-type-1",
						},
					},
				},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &GCPPlan{}
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

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

func TestStorageAWSPlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-aws-plan-name",
		Namespace: "fake-aws-plan-namespace",
	}
	created := &AWSPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-aws-plan-name",
			Namespace: "fake-aws-plan-namespace",
		},
		Spec: AWSPlanSpec{
			Description: "fake-aws-plan-description",
			ProviderSpec: AWSProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-aws-provider-credentials-secret-name",
					Namespace: "fake-aws-provider-credentials-secret-namespace",
				},
				Region: "fake-aws-region",
			},
			ComputeSpec: AWSComputeSpec{
				MastersSpec: AWSComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-aws-zone-1"},
				},
				WorkersSpec: AWSComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-aws-zone-1"},
				},
			},
			NetworkSpec: AWSNetworkSpec{
				VpcID: "fake-aws-vpc-id",
				DNS:   []string{"fake-aws-dns-server"},
			},
			StorageSpec: AWSStorageSpec{
				MastersSpec: AWSStorageMastersSpec{
					Disks: []AWSDiskSpec{
						AWSDiskSpec{
							SizeGb: 1,
						},
					},
				},
				WorkersSpec: AWSStorageWorkersSpec{
					Disks: []AWSDiskSpec{
						AWSDiskSpec{
							SizeGb: 1,
						},
					},
				},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &AWSPlan{}
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

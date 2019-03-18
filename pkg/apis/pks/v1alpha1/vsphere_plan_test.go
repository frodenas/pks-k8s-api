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

func TestStorageVSpherePlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-vsphere-plan-name",
		Namespace: "fake-vsphere-plan-namespace",
	}
	created := &VSpherePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-vsphere-plan-name",
			Namespace: "fake-vsphere-plan-namespace",
		},
		Spec: VSpherePlanSpec{
			Description: "fake-vsphere-plan-description",
			ProviderSpec: VSphereProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-vsphere-provider-credentials-secret-name",
					Namespace: "fake-vsphere-provider-credentials-secret-namespace",
				},
			},
			ComputeSpec: VSphereComputeSpec{
				MastersSpec: VSphereComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					VMFolder: "fake-vsphere-vm-folder",
					Zones: []VSphereZoneSpec{
						VSphereZoneSpec{
							Name:         "fake-vsphere-zone-name",
							Datacenter:   "fake-vsphere-zone-datacenter",
							Cluster:      "fake-vsphere-zone-cluster",
							ResourcePool: "fake-vsphere-zone-resource-pool",
						},
					},
				},
				WorkersSpec: VSphereComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					VMFolder: "fake-vsphere-vm-folder",
					Zones: []VSphereZoneSpec{
						VSphereZoneSpec{
							Name:         "fake-vsphere-zone-name",
							Datacenter:   "fake-vsphere-zone-datacenter",
							Cluster:      "fake-vsphere-zone-cluster",
							ResourcePool: "fake-vsphere-zone-resource-pool",
						},
					},
				},
			},
			NetworkSpec: VSphereNetworkSpec{
				DNS: []string{"fake-vsphere-dns-server"},
				DVSNetworkSpec: &VSphereDVSNetworkSpec{
					Name: "fake-vsphere-network-name",
				},
				NSXTNetworkSpec: &VSphereNSXTNetworkSpec{
					CredentialsSecretRef: corev1.SecretReference{
						Name:      "fake-vsphere-nsxt-credentials-secret-name",
						Namespace: "fake-vsphere-nsxt-credentials-secret-namespace",
					},
					T0RouterID:        "fake-vsphere-nsxt-t0-router-id",
					IPBlockIDs:        []string{"fake-vsphere-nsxt-ip-block-id"},
					FloatingIPPoolIDs: []string{"fake-vsphere-nsxt-floating-ip-pool-id"},
					NatMode:           true,
					LBSize:            "fake-vsphere-nsxt-lb-size",
					PodSubnetPrefix:   24,
				},
			},
			StorageSpec: VSphereStorageSpec{
				MastersSpec: VSphereStorageMastersSpec{
					Datastore: "fake-vsphere-datastore",
					Disks: []VSphereDiskSpec{
						VSphereDiskSpec{
							SizeGb: 1,
							Label:  "fake-vsphere-disk-label",
						},
					},
				},
				WorkersSpec: VSphereStorageWorkersSpec{
					Datastore: "fake-vsphere-datastore",
					Disks: []VSphereDiskSpec{
						VSphereDiskSpec{
							SizeGb: 1,
							Label:  "fake-vsphere-disk-label",
						},
					},
				},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &VSpherePlan{}
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

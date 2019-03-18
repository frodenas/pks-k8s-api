/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vsphereplan

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

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-vsphere-plan-name", Namespace: "fake-vsphere-plan-namespace"}}
var vSpherePlanKey = types.NamespacedName{Name: "fake-vsphere-plan-name", Namespace: "fake-vsphere-plan-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.VSpherePlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-vsphere-plan-name",
			Namespace: "fake-vsphere-plan-namespace",
		},
		Spec: pksv1alpha1.VSpherePlanSpec{
			Description: "fake-vsphere-plan-description",
			ProviderSpec: pksv1alpha1.VSphereProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-vsphere-provider-credentials-secret-name",
					Namespace: "fake-vsphere-provider-credentials-secret-namespace",
				},
			},
			ComputeSpec: pksv1alpha1.VSphereComputeSpec{
				MastersSpec: pksv1alpha1.VSphereComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					VMFolder: "fake-vsphere-vm-folder",
					Zones: []pksv1alpha1.VSphereZoneSpec{
						pksv1alpha1.VSphereZoneSpec{
							Name:         "fake-vsphere-zone-name",
							Datacenter:   "fake-vsphere-zone-datacenter",
							Cluster:      "fake-vsphere-zone-cluster",
							ResourcePool: "fake-vsphere-zone-resource-pool",
						},
					},
				},
				WorkersSpec: pksv1alpha1.VSphereComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					VMFolder: "fake-vsphere-vm-folder",
					Zones: []pksv1alpha1.VSphereZoneSpec{
						pksv1alpha1.VSphereZoneSpec{
							Name:         "fake-vsphere-zone-name",
							Datacenter:   "fake-vsphere-zone-datacenter",
							Cluster:      "fake-vsphere-zone-cluster",
							ResourcePool: "fake-vsphere-zone-resource-pool",
						},
					},
				},
			},
			NetworkSpec: pksv1alpha1.VSphereNetworkSpec{
				DNS: []string{"fake-vsphere-dns-server"},
				DVSNetworkSpec: &pksv1alpha1.VSphereDVSNetworkSpec{
					Name: "fake-vsphere-network-name",
				},
				NSXTNetworkSpec: &pksv1alpha1.VSphereNSXTNetworkSpec{
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
			StorageSpec: pksv1alpha1.VSphereStorageSpec{
				MastersSpec: pksv1alpha1.VSphereStorageMastersSpec{
					Datastore: "fake-vsphere-datastore",
					Disks: []pksv1alpha1.VSphereDiskSpec{
						pksv1alpha1.VSphereDiskSpec{
							SizeGb: 1,
							Label:  "fake-vsphere-disk-label",
						},
					},
				},
				WorkersSpec: pksv1alpha1.VSphereStorageWorkersSpec{
					Datastore: "fake-vsphere-datastore",
					Disks: []pksv1alpha1.VSphereDiskSpec{
						pksv1alpha1.VSphereDiskSpec{
							SizeGb: 1,
							Label:  "fake-vsphere-disk-label",
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

	// Create the VSphere Plan object
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	g.Eventually(func() error { return c.Get(context.TODO(), vSpherePlanKey, vSpherePlan) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), vSpherePlanKey, vSpherePlan)
		return vSpherePlan.Finalizers
	}, timeout).Should(gomega.ContainElement(VSpherePlanFinalizer))
}

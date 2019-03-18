/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package awsplan

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

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "fake-aws-plan-name", Namespace: "fake-aws-plan-namespace"}}
var awsPlanKey = types.NamespacedName{Name: "fake-aws-plan-name", Namespace: "fake-aws-plan-namespace"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &pksv1alpha1.AWSPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-aws-plan-name",
			Namespace: "fake-aws-plan-namespace",
		},
		Spec: pksv1alpha1.AWSPlanSpec{
			Description: "fake-aws-plan-description",
			ProviderSpec: pksv1alpha1.AWSProviderSpec{
				CredentialsSecretRef: corev1.SecretReference{
					Name:      "fake-aws-provider-credentials-secret-name",
					Namespace: "fake-aws-provider-credentials-secret-namespace",
				},
				Region: "fake-aws-region",
			},
			ComputeSpec: pksv1alpha1.AWSComputeSpec{
				MastersSpec: pksv1alpha1.AWSComputeMastersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-aws-zone-1"},
				},
				WorkersSpec: pksv1alpha1.AWSComputeWorkersSpec{
					Replicas: 1,
					NumCPUs:  2,
					MemoryMB: 3,
					Zones:    []string{"fake-aws-zone-1"},
				},
			},
			NetworkSpec: pksv1alpha1.AWSNetworkSpec{
				VpcID: "fake-aws-vpc-id",
				DNS:   []string{"fake-aws-dns-server"},
			},
			StorageSpec: pksv1alpha1.AWSStorageSpec{
				MastersSpec: pksv1alpha1.AWSStorageMastersSpec{
					Disks: []pksv1alpha1.AWSDiskSpec{
						pksv1alpha1.AWSDiskSpec{
							SizeGb: 1,
						},
					},
				},
				WorkersSpec: pksv1alpha1.AWSStorageWorkersSpec{
					Disks: []pksv1alpha1.AWSDiskSpec{
						pksv1alpha1.AWSDiskSpec{
							SizeGb: 1,
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

	// Create the AWS Plan object
	err = c.Create(context.TODO(), instance)
	g.Expect(err).To(gomega.Succeed())
	defer c.Delete(context.TODO(), instance)

	// Expect the Reconcile to be invoked
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	// Expect the object to exists
	awsPlan := &pksv1alpha1.AWSPlan{}
	g.Eventually(func() error { return c.Get(context.TODO(), awsPlanKey, awsPlan) }, timeout).Should(gomega.Succeed())

	// Expect to contain the finalizer
	g.Eventually(func() []string {
		c.Get(context.TODO(), awsPlanKey, awsPlan)
		return awsPlan.Finalizers
	}, timeout).Should(gomega.ContainElement(AWSPlanFinalizer))
}

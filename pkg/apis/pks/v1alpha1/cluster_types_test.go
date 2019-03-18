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

func TestStorageCluster(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-cluster-name",
		Namespace: "fake-cluster-namespace",
	}
	created := &Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-cluster-name",
			Namespace: "fake-cluster-namespace",
		},
		Spec: ClusterSpec{
			ExternalHostname:  "fake-cluster-external-hostname",
			Description:       "fake-cluster-description",
			NumWorkerReplicas: 32,
			ProvisionerSpec: ProvisionerSpec{
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
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &Cluster{}
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

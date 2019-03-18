/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageKubernetesProfile(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-kubernetes-profile-name",
		Namespace: "fake-kubernetes-profile-namespace",
	}
	created := &KubernetesProfile{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-kubernetes-profile-name",
			Namespace: "fake-kubernetes-profile-namespace",
		},
		Spec: KubernetesProfileSpec{
			Description: "fake-kubernetes-profile-description",
			Versions: KubernetesVersionsSpec{
				Master: "fake-kubernetes-profile-version-master",
				Worker: "fake-kubernetes-profile-version-worker",
			},
			NetworkSpec: KubernetesNetworkSpec{
				ServiceDomain:      "fake-kubernetes-profile-service-domain",
				ServicesCIDRBlocks: []string{"20.0.0.0/24"},
				PodsCIDRBlocks:     []string{"10.0.0.0/24"},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &KubernetesProfile{}
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

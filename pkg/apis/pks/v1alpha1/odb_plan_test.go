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

func TestStorageODBPlan(t *testing.T) {
	key := types.NamespacedName{
		Name:      "fake-odb-plan-name",
		Namespace: "fake-odb-plan-namespace",
	}
	created := &ODBPlan{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-odb-plan-name",
			Namespace: "fake-odb-plan-namespace",
		},
		Spec: ODBPlanSpec{
			Description: "fake-odb-plan-description",
			ServiceID:   "fake-odb-plan-service-id",
			PlanID:      "fake-odb-plan-plan-id",
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &ODBPlan{}
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

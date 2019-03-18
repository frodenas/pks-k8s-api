/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	"context"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	apitypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// IsUsingNSXT returns true if the Cluster specification uses NSX-T resources.
func IsUsingNSXT(k8sClient client.Client, instance *pksv1alpha1.Cluster) (bool, error) {
	// Return false if it is not using a vSphere plan.
	if instance.Spec.PlanRef.Kind != pksv1alpha1.VSpherePlanKind {
		return false, nil
	}

	// Retrieve the associated vSphere Plan.
	vSpherePlan := &pksv1alpha1.VSpherePlan{}
	if err := k8sClient.Get(context.TODO(), apitypes.NamespacedName{Name: instance.Spec.PlanRef.Name, Namespace: instance.Spec.PlanRef.Namespace}, vSpherePlan); err != nil {
		return false, err
	}

	// Return false if vSphere Plan Network Spec does not use NSX-T.
	if vSpherePlan.Spec.NetworkSpec.NSXTNetworkSpec == nil {
		return false, nil
	}

	return true, nil
}

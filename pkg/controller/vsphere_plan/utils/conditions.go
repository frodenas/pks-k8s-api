/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewVSpherePlanCondition creates a new vSphere Plan condition.
func NewVSpherePlanCondition(
	condType pksv1alpha1.VSpherePlanConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.VSpherePlanCondition {
	return &pksv1alpha1.VSpherePlanCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetVSpherePlanCondition returns the condition with the provided type.
func GetVSpherePlanCondition(
	status pksv1alpha1.VSpherePlanStatus,
	condType pksv1alpha1.VSpherePlanConditionType,
) *pksv1alpha1.VSpherePlanCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetVSpherePlanCondition updates the VSphere Plan to include the provided condition.
func SetVSpherePlanCondition(status *pksv1alpha1.VSpherePlanStatus, condition pksv1alpha1.VSpherePlanCondition) {
	currentCond := GetVSpherePlanCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveVSpherePlanCondition removes the vSphere Plan condition with the provided type.
func RemoveVSpherePlanCondition(status *pksv1alpha1.VSpherePlanStatus, condType pksv1alpha1.VSpherePlanConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllVSpherePlanConditionsTrue returns true if all VSphere Plan conditions are true.
func AreAllVSpherePlanConditionsTrue(status pksv1alpha1.VSpherePlanStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of vSphere Plan conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.VSpherePlanCondition,
	condType pksv1alpha1.VSpherePlanConditionType,
) []pksv1alpha1.VSpherePlanCondition {
	var newConditions []pksv1alpha1.VSpherePlanCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

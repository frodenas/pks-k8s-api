/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewGCPPlanCondition creates a new GCP Plan condition.
func NewGCPPlanCondition(
	condType pksv1alpha1.GCPPlanConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.GCPPlanCondition {
	return &pksv1alpha1.GCPPlanCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetGCPPlanCondition returns the condition with the provided type.
func GetGCPPlanCondition(
	status pksv1alpha1.GCPPlanStatus,
	condType pksv1alpha1.GCPPlanConditionType,
) *pksv1alpha1.GCPPlanCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetGCPPlanCondition updates the GCP Plan to include the provided condition.
func SetGCPPlanCondition(status *pksv1alpha1.GCPPlanStatus, condition pksv1alpha1.GCPPlanCondition) {
	currentCond := GetGCPPlanCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveGCPPlanCondition removes the GCP Plan condition with the provided type.
func RemoveGCPPlanCondition(status *pksv1alpha1.GCPPlanStatus, condType pksv1alpha1.GCPPlanConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllGCPPlanConditionsTrue returns true if all GCP Plan conditions are true.
func AreAllGCPPlanConditionsTrue(status pksv1alpha1.GCPPlanStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of GCP Plan conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.GCPPlanCondition,
	condType pksv1alpha1.GCPPlanConditionType,
) []pksv1alpha1.GCPPlanCondition {
	var newConditions []pksv1alpha1.GCPPlanCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

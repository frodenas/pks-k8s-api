/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAzurePlanCondition creates a new Azure Plan condition.
func NewAzurePlanCondition(
	condType pksv1alpha1.AzurePlanConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.AzurePlanCondition {
	return &pksv1alpha1.AzurePlanCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetAzurePlanCondition returns the condition with the provided type.
func GetAzurePlanCondition(
	status pksv1alpha1.AzurePlanStatus,
	condType pksv1alpha1.AzurePlanConditionType,
) *pksv1alpha1.AzurePlanCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetAzurePlanCondition updates the Azure Plan to include the provided condition.
func SetAzurePlanCondition(status *pksv1alpha1.AzurePlanStatus, condition pksv1alpha1.AzurePlanCondition) {
	currentCond := GetAzurePlanCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveAzurePlanCondition removes the AzureS Plan condition with the provided type.
func RemoveAzurePlanCondition(status *pksv1alpha1.AzurePlanStatus, condType pksv1alpha1.AzurePlanConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllAzurePlanConditionsTrue returns true if all Azure Plan conditions are true.
func AreAllAzurePlanConditionsTrue(status pksv1alpha1.AzurePlanStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of Azure Plan conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.AzurePlanCondition,
	condType pksv1alpha1.AzurePlanConditionType,
) []pksv1alpha1.AzurePlanCondition {
	var newConditions []pksv1alpha1.AzurePlanCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

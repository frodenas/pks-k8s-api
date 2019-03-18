/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewODBPlanCondition creates a new ODB Plan condition.
func NewODBPlanCondition(
	condType pksv1alpha1.ODBPlanConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.ODBPlanCondition {
	return &pksv1alpha1.ODBPlanCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetODBPlanCondition returns the condition with the provided type.
func GetODBPlanCondition(
	status pksv1alpha1.ODBPlanStatus,
	condType pksv1alpha1.ODBPlanConditionType,
) *pksv1alpha1.ODBPlanCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetODBPlanCondition updates the ODB Plan to include the provided condition.
func SetODBPlanCondition(status *pksv1alpha1.ODBPlanStatus, condition pksv1alpha1.ODBPlanCondition) {
	currentCond := GetODBPlanCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveODBPlanCondition removes the ODB Plan condition with the provided type.
func RemoveODBPlanCondition(status *pksv1alpha1.ODBPlanStatus, condType pksv1alpha1.ODBPlanConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllODBPlanConditionsTrue returns true if all ODB Plan conditions are true.
func AreAllODBPlanConditionsTrue(status pksv1alpha1.ODBPlanStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of ODB Plan conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.ODBPlanCondition,
	condType pksv1alpha1.ODBPlanConditionType,
) []pksv1alpha1.ODBPlanCondition {
	var newConditions []pksv1alpha1.ODBPlanCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

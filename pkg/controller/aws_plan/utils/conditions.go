/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAWSPlanCondition creates a new AWS Plan condition.
func NewAWSPlanCondition(
	condType pksv1alpha1.AWSPlanConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.AWSPlanCondition {
	return &pksv1alpha1.AWSPlanCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetAWSPlanCondition returns the condition with the provided type.
func GetAWSPlanCondition(
	status pksv1alpha1.AWSPlanStatus,
	condType pksv1alpha1.AWSPlanConditionType,
) *pksv1alpha1.AWSPlanCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetAWSPlanCondition updates the AWS Plan to include the provided condition.
func SetAWSPlanCondition(status *pksv1alpha1.AWSPlanStatus, condition pksv1alpha1.AWSPlanCondition) {
	currentCond := GetAWSPlanCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveAWSPlanCondition removes the AWS Plan condition with the provided type.
func RemoveAWSPlanCondition(status *pksv1alpha1.AWSPlanStatus, condType pksv1alpha1.AWSPlanConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllAWSPlanConditionsTrue returns true if all AWS Plan conditions are true.
func AreAllAWSPlanConditionsTrue(status pksv1alpha1.AWSPlanStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of AWS Plan conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.AWSPlanCondition,
	condType pksv1alpha1.AWSPlanConditionType,
) []pksv1alpha1.AWSPlanCondition {
	var newConditions []pksv1alpha1.AWSPlanCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

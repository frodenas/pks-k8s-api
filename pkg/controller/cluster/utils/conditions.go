/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewClusterCondition creates a new Cluster condition.
func NewClusterCondition(
	condType pksv1alpha1.ClusterConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
	rawData string,
) *pksv1alpha1.ClusterCondition {
	return &pksv1alpha1.ClusterCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
		RawData:            rawData,
	}
}

// GetClusterCondition returns the condition with the provided type.
func GetClusterCondition(
	status pksv1alpha1.ClusterStatus,
	condType pksv1alpha1.ClusterConditionType,
) *pksv1alpha1.ClusterCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetClusterCondition updates the Cluster to include the provided condition.
func SetClusterCondition(status *pksv1alpha1.ClusterStatus, condition pksv1alpha1.ClusterCondition) {
	currentCond := GetClusterCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveClusterCondition removes the Cluster condition with the provided type.
func RemoveClusterCondition(status *pksv1alpha1.ClusterStatus, condType pksv1alpha1.ClusterConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllClusterConditionsTrue returns true if all Cluster conditions are true.
func AreAllClusterConditionsTrue(status pksv1alpha1.ClusterStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of Cluster conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.ClusterCondition,
	condType pksv1alpha1.ClusterConditionType,
) []pksv1alpha1.ClusterCondition {
	var newConditions []pksv1alpha1.ClusterCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

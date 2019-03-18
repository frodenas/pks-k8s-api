/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewKubernetesProfileCondition creates a new Kubernetes Profile condition.
func NewKubernetesProfileCondition(
	condType pksv1alpha1.KubernetesProfileConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) *pksv1alpha1.KubernetesProfileCondition {
	return &pksv1alpha1.KubernetesProfileCondition{
		Type:               condType,
		Status:             status,
		LastUpdateTime:     metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetKubernetesProfileCondition returns the condition with the provided type.
func GetKubernetesProfileCondition(
	status pksv1alpha1.KubernetesProfileStatus,
	condType pksv1alpha1.KubernetesProfileConditionType,
) *pksv1alpha1.KubernetesProfileCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}

	return nil
}

// SetKubernetesProfileCondition updates the Kubernetes Profile to include the provided condition.
func SetKubernetesProfileCondition(status *pksv1alpha1.KubernetesProfileStatus, condition pksv1alpha1.KubernetesProfileCondition) {
	currentCond := GetKubernetesProfileCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status != condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveKubernetesProfileCondition removes the Kubernetes Profile condition with the provided type.
func RemoveKubernetesProfileCondition(status *pksv1alpha1.KubernetesProfileStatus, condType pksv1alpha1.KubernetesProfileConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// AreAllKubernetesProfileConditionsTrue returns true if all Kubernetes Profile conditions are true.
func AreAllKubernetesProfileConditionsTrue(status pksv1alpha1.KubernetesProfileStatus) bool {
	allConditionsTrue := true
	for _, condition := range status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			allConditionsTrue = false
			break
		}
	}
	return allConditionsTrue
}

// filterOutCondition returns a new slice of Kubernetes Profile conditions without conditions with the provided type.
func filterOutCondition(
	conditions []pksv1alpha1.KubernetesProfileCondition,
	condType pksv1alpha1.KubernetesProfileConditionType,
) []pksv1alpha1.KubernetesProfileCondition {
	var newConditions []pksv1alpha1.KubernetesProfileCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ODBPlanKind defines the ODBPlan Kind.
	ODBPlanKind = "ODBPlan"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ODBPlan is the Schema for the ODB Plans API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type ODBPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ODBPlanSpec   `json:"spec,omitempty"`
	Status ODBPlanStatus `json:"status,omitempty"`
}

// ODBPlanSpec represents a ODB Plan specification.
type ODBPlanSpec struct {
	// Description is the description of the ODB Plan.
	// +optional
	Description string `json:"description,omitempty"`

	// ServiceID is the identified of the Service.
	ServiceID string `json:"serviceId"`

	// PlanID is the identified of the Plan.
	PlanID string `json:"planId"`
}

// ODBPlanStatus defines the observed state of the ODB Plan.
type ODBPlanStatus struct {
	// The generation observed by the OD Plan controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a OD Plan current state.
	// +optional
	Conditions []ODBPlanCondition `json:"conditions,omitempty"`
}

// ODBPlanCondition describes the state of a ODB Plan at a certain point.
type ODBPlanCondition struct {
	// Type of ODB Plan condition.
	Type ODBPlanConditionType `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`

	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// The reason for the condition's last transition.
	Reason string `json:"reason"`

	// A human readable message indicating details about the transition.
	Message string `json:"message"`
}

// ODBPlanConditionType is the type of a ODB Plan condition
type ODBPlanConditionType string

// These are valid conditions of a ODB Plan.
const (
	// Validated means the ODB Plan has been validated.
	ODBPlanValidated ODBPlanConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ODBPlanList contains a list of ODB plans.
type ODBPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ODBPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ODBPlan{}, &ODBPlanList{})
}

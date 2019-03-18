/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// KubernetesProfileKind defines the KubernetesProfile Kind.
	KubernetesProfileKind = "KubernetesProfile"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubernetesProfile is the Schema for the Kubernetes Profiles API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type KubernetesProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubernetesProfileSpec   `json:"spec,omitempty"`
	Status KubernetesProfileStatus `json:"status,omitempty"`
}

// KubernetesProfileSpec represents a Kubernetes Profile specification.
type KubernetesProfileSpec struct {
	// Description is the description of the plant.
	// +optional
	Description string `json:"description,omitempty"`

	// Versions is the Kubernetes Versions specification.
	Versions KubernetesVersionsSpec `json:"versions"`

	// NetworkSpec
	NetworkSpec KubernetesNetworkSpec `json:"network"`
}

// KubernetesVersionsSpec represents a Kubernetes Versions specification.
type KubernetesVersionsSpec struct {
	// Master is the semantic version of the Kubernetes control plane to run.
	Master string `json:"master"`

	// Worker is the semantic version of Kubernetes kubelet to run.
	Worker string `json:"worker"`
}

// KubernetesNetworkSpec specifies the different networking parameters for kubernetes.
type KubernetesNetworkSpec struct {
	// Domain name for services.
	ServiceDomain string `json:"serviceDomain"`

	// The network ranges from which service VIPs are allocated.
	// +kubebuilder:validation:MinItems=1
	ServicesCIDRBlocks []string `json:"servicesCIDRBlocks"`

	// The network ranges from which Pod networks are allocated.
	// +optional
	PodsCIDRBlocks []string `json:"podsCIDRBlocks,omitempty"`
}

// KubernetesProfileStatus defines the observed state of the Kubernetes Profile.
type KubernetesProfileStatus struct {
	// The generation observed by the Kubernetes Profile controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a Kubernetes Profile current state.
	// +optional
	Conditions []KubernetesProfileCondition `json:"conditions,omitempty"`
}

// KubernetesProfileCondition describes the state of a Kubernetes Profile at a certain point.
type KubernetesProfileCondition struct {
	// Type of Kubernetes Profile condition.
	Type KubernetesProfileConditionType `json:"type"`

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

// KubernetesProfileConditionType is the type of a Kubernetes Profile condition
type KubernetesProfileConditionType string

// These are valid conditions of a Kubernetes Profile.
const (
	// Validated means the Kubernetes Profile has been validated.
	KubernetesProfileValidated KubernetesProfileConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubernetesProfileList contains a list of Kubernetes Profiles
type KubernetesProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubernetesProfile{}, &KubernetesProfileList{})
}

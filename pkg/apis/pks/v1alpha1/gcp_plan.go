/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// GCPPlanKind defines the GCPPlan Kind.
	GCPPlanKind = "GCPPlan"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GCPPlan is the Schema for the GCP Plans API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type GCPPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPPlanSpec   `json:"spec,omitempty"`
	Status GCPPlanStatus `json:"status,omitempty"`
}

// GCPPlanSpec represents a GCP Plan specification.
type GCPPlanSpec struct {
	// Description is the description of the GCP Plan.
	// +optional
	Description string `json:"description,omitempty"`

	// ProviderSpec is the GCP Provider specification.
	ProviderSpec GCPProviderSpec `json:"provider"`

	// ComputeSpec is a GCP Compute specification.
	ComputeSpec GCPComputeSpec `json:"compute"`

	// NetworkSpec is a GCP Network specification.
	NetworkSpec GCPNetworkSpec `json:"network"`

	// StorageSpec is a GCP Storage specification.
	StorageSpec GCPStorageSpec `json:"storage"`
}

// GCPProviderSpec represents a GCP Provider specification.
type GCPProviderSpec struct {
	// Secret containing vSphere credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`

	// Region is the GCP region name.
	Region string `json:"region"`
}

const (
	// GCPProviderCredentialsProjectKey sets the key for the GCP project value.
	GCPProviderCredentialsProjectKey = "project"

	// GCPProviderCredentialsJSONKeyKey sets the key for GCP JSON key value.
	GCPProviderCredentialsJSONKeyKey = "jsonKey"
)

// GCPComputeSpec represents a GCP Compute specification.
type GCPComputeSpec struct {
	// MastersSpec is the GCP Compute Masters specification.
	MastersSpec GCPComputeMastersSpec `json:"masters"`

	// WorkersSpec is the GCP Compute Workers specification.
	WorkersSpec GCPComputeWorkersSpec `json:"workers"`
}

// GCPComputeMastersSpec represents a GCP Compute Masters specification.
type GCPComputeMastersSpec struct {
	// Replicas is the number of master VMs.
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`

	// Zones are the zones where master VMs will be located.
	// +kubebuilder:validation:MinItems=1
	Zones []string `json:"zones"`
}

// GCPComputeWorkersSpec represents a GCP Compute Workers specification.
type GCPComputeWorkersSpec struct {
	// Replicas is the number of worker VMs.
	// +kubebuilder:validation:Minimum=0
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`

	// Zones are the zones where worker VMs will be located.
	// +kubebuilder:validation:MinItems=1
	Zones []string `json:"zones"`
}

// GCPNetworkSpec represents a GCP Network specification.
type GCPNetworkSpec struct {
	// Name is the name of the network to associate with the VMs.
	Name string `json:"name"`

	// DNS is a list of Domain Name Servers.
	// +kubebuilder:validation:MinItems=1
	DNS []string `json:"dns,omitempty"`
}

// GCPStorageSpec represents a GCP Storage specification.
type GCPStorageSpec struct {
	// MastersSpec is the GCP Storage Masters specification.
	MastersSpec GCPStorageMastersSpec `json:"masters"`

	// WorkersSpec is the GCP Storage Workers specification.
	WorkersSpec GCPStorageWorkersSpec `json:"workers"`
}

// GCPStorageMastersSpec represents a GCP Storage Masters specification.
type GCPStorageMastersSpec struct {
	// Disks are the disks to be attached to the master VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []GCPDiskSpec `json:"disks"`
}

// GCPStorageWorkersSpec represents a GCP Storage Workers specification.
type GCPStorageWorkersSpec struct {
	// Disks are the disks to be attached to the worker VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []GCPDiskSpec `json:"disks"`
}

// GCPDiskSpec represents a GCP Disk specification.
type GCPDiskSpec struct {
	// SizeGb is the size in Gb of the disk.
	// +kubebuilder:validation:Minimum=1
	SizeGb int64 `json:"sizeGb"`

	// Tye is the type of the Disk.
	// +optional
	Type string `json:"type,omitempty"`
}

// GCPPlanStatus defines the observed state of the GCP Plan.
type GCPPlanStatus struct {
	// The generation observed by the GCP Plan controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a GCP Plan current state.
	// +optional
	Conditions []GCPPlanCondition `json:"conditions,omitempty"`
}

// GCPPlanCondition describes the state of a GCP Plan at a certain point.
type GCPPlanCondition struct {
	// Type of GCP Plan condition.
	Type GCPPlanConditionType `json:"type"`

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

// GCPPlanConditionType is the type of a GCP Plan condition
type GCPPlanConditionType string

// These are valid conditions of a GCP Plan.
const (
	// Validated means the GCP Plan has been validated.
	GCPPlanValidated GCPPlanConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GCPPlanList contains a list of GCP plans.
type GCPPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GCPPlan{}, &GCPPlanList{})
}

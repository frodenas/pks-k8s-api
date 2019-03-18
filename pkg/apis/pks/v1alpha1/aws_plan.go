/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// AWSPlanKind defines the AWSPlan Kind.
	AWSPlanKind = "AWSPlan"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSPlan is the Schema for the AWS Plans API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type AWSPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSPlanSpec   `json:"spec,omitempty"`
	Status AWSPlanStatus `json:"status,omitempty"`
}

// AWSPlanSpec represents a AWS Plan specification.
type AWSPlanSpec struct {
	// Description is the description of the AWS Plan.
	// +optional
	Description string `json:"description,omitempty"`

	// ProviderSpec is the AWS Provider specification.
	ProviderSpec AWSProviderSpec `json:"provider"`

	// ComputeSpec is a AWS Compute specification.
	ComputeSpec AWSComputeSpec `json:"compute"`

	// NetworkSpec is a AWS Network specification.
	NetworkSpec AWSNetworkSpec `json:"network"`

	// StorageSpec is a AWS Storage specification.
	StorageSpec AWSStorageSpec `json:"storage"`
}

// AWSProviderSpec represents a AWS Provider specification.
type AWSProviderSpec struct {
	// Secret containing AWS credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`

	// Region is the AWS region name.
	Region string `json:"region"`
}

const (
	// AWSProviderCredentialsAccessKeyKey sets the key for the AWS access key value.
	AWSProviderCredentialsAccessKeyKey = "accessKey"

	// AWSProviderCredentialsSecretAccessKeyKey sets the key for AWS secret access key value.
	AWSProviderCredentialsSecretAccessKeyKey = "secretAccessKey"
)

// AWSComputeSpec represents a AWS Compute specification.
type AWSComputeSpec struct {
	// MastersSpec is the AWS Compute Masters specification.
	MastersSpec AWSComputeMastersSpec `json:"masters"`

	// WorkersSpec is the AWS Compute Workers specification.
	WorkersSpec AWSComputeWorkersSpec `json:"workers"`
}

// AWSComputeMastersSpec represents a AWS Compute Masters specification.
type AWSComputeMastersSpec struct {
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

// AWSComputeWorkersSpec represents a AWS Compute Workers specification.
type AWSComputeWorkersSpec struct {
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

// AWSNetworkSpec represents a AWS Network specification.
type AWSNetworkSpec struct {
	// VpcId is the ID of the VPC to be used to create network resources.
	VpcID string `json:"vpcId"`

	// DNS is a list of Domain Name Servers.
	// +kubebuilder:validation:MinItems=1
	DNS []string `json:"dns,omitempty"`
}

// AWSStorageSpec represents a AWS Storage specification.
type AWSStorageSpec struct {
	// MastersSpec is the AWS Storage Masters specification.
	MastersSpec AWSStorageMastersSpec `json:"masters"`

	// WorkersSpec is the AWS Storage Workers specification.
	WorkersSpec AWSStorageWorkersSpec `json:"workers"`
}

// AWSStorageMastersSpec represents a AWS Storage Masters specification.
type AWSStorageMastersSpec struct {
	// Disks are the disks to be attached to the master VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []AWSDiskSpec `json:"disks"`
}

// AWSStorageWorkersSpec represents a AWS Storage Workers specification.
type AWSStorageWorkersSpec struct {
	// Disks are the disks to be attached to the worker VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []AWSDiskSpec `json:"disks"`
}

// AWSDiskSpec represents a AWS Disk specification.
type AWSDiskSpec struct {
	// SizeGb is the size in Gb of the disk.
	// +kubebuilder:validation:Minimum=1
	SizeGb int64 `json:"sizeGb"`
}

// AWSPlanStatus defines the observed state of the AWS Plan.
type AWSPlanStatus struct {
	// The generation observed by the AWS Plan controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a AWS Plan current state.
	// +optional
	Conditions []AWSPlanCondition `json:"conditions,omitempty"`
}

// AWSPlanCondition describes the state of a AWS Plan at a certain point.
type AWSPlanCondition struct {
	// Type of AWS Plan condition.
	Type AWSPlanConditionType `json:"type"`

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

// AWSPlanConditionType is the type of a AWS Plan condition
type AWSPlanConditionType string

// These are valid conditions of a AWS Plan.
const (
	// Validated means the AWS Plan has been validated.
	AWSPlanValidated AWSPlanConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSPlanList contains a list of AWS plans.
type AWSPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSPlan{}, &AWSPlanList{})
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// AzurePlanKind defines the AzurePlan Kind.
	AzurePlanKind = "AzurePlan"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurePlan is the Schema for the Azure Plans API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type AzurePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzurePlanSpec   `json:"spec,omitempty"`
	Status AzurePlanStatus `json:"status,omitempty"`
}

// AzurePlanSpec represents a Azure Plan specification.
type AzurePlanSpec struct {
	// Description is the description of the Azure Plan.
	// +optional
	Description string `json:"description,omitempty"`

	// ProviderSpec is the Azure Provider specification.
	ProviderSpec AzureProviderSpec `json:"provider"`

	// ComputeSpec is a Azure Compute specification.
	ComputeSpec AzureComputeSpec `json:"compute"`

	// NetworkSpec is a Azure Network specification.
	NetworkSpec AzureNetworkSpec `json:"network"`

	// StorageSpec is a Azure Storage specification.
	StorageSpec AzureStorageSpec `json:"storage"`
}

// AzureProviderSpec represents a Azure Provider specification.
type AzureProviderSpec struct {
	// Secret containing Azure credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`

	// Environment is the Azure environment name.
	Environment string `json:"environment"`

	// Location is the Azure region name.
	Location string `json:"location"`

	// ResourceGroup is the Azure resource group name.
	ResourceGroup string `json:"resourceGroup"`
}

const (
	// AzureProviderCredentialsSubscriptionIDKey sets the key for Azure subscription id key value.
	AzureProviderCredentialsSubscriptionIDKey = "subscriptionId"

	// AzureProviderCredentialsTenantIDKey sets the key for Azure tenant id key value.
	AzureProviderCredentialsTenantIDKey = "tenantId"

	// AzureProviderCredentialsClientIDKey sets the key for Azure client id key value.
	AzureProviderCredentialsClientIDKey = "clientId"

	// AzureProviderCredentialsClientSecretyKey sets the key for Azure client secret key value.
	AzureProviderCredentialsClientSecretyKey = "clientSecret"
)

// AzureComputeSpec represents a Azure Compute specification.
type AzureComputeSpec struct {
	// MastersSpec is the Azure Compute Masters specification.
	MastersSpec AzureComputeMastersSpec `json:"masters"`

	// WorkersSpec is the Azure Compute Workers specification.
	WorkersSpec AzureComputeWorkersSpec `json:"workers"`
}

// AzureComputeMastersSpec represents a Azure Compute Masters specification.
type AzureComputeMastersSpec struct {
	// Replicas is the number of master VMs.
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`
}

// AzureComputeWorkersSpec represents a Azure Compute Workers specification.
type AzureComputeWorkersSpec struct {
	// Replicas is the number of worker VMs.
	// +kubebuilder:validation:Minimum=0
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`
}

// AzureNetworkSpec represents a Azure Network specification.
type AzureNetworkSpec struct {
	// Vnet is the name of the Azure Virtual Network to be used to create VMS.
	Vnet string `json:"vnet"`

	// Subnets configuration.
	Subnet string `json:"subnet"`

	// DNS is a list of Domain Name Servers.
	// +kubebuilder:validation:MinItems=1
	DNS []string `json:"dns,omitempty"`
}

// AzureStorageSpec represents a Azure Storage specification.
type AzureStorageSpec struct {
	// MastersSpec is the Azure Storage Masters specification.
	MastersSpec AzureStorageMastersSpec `json:"masters"`

	// WorkersSpec is the AzureS Storage Workers specification.
	WorkersSpec AzureStorageWorkersSpec `json:"workers"`
}

// AzureStorageMastersSpec represents a Azure Storage Masters specification.
type AzureStorageMastersSpec struct {
	// Disks are the disks to be attached to the master VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []AzureDiskSpec `json:"disks"`
}

// AzureStorageWorkersSpec represents a Azure Storage Workers specification.
type AzureStorageWorkersSpec struct {
	// Disks are the disks to be attached to the worker VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []AzureDiskSpec `json:"disks"`
}

// AzureDiskSpec represents a Azure Disk specification.
type AzureDiskSpec struct {
	// SizeGb is the size in Gb of the disk.
	// +kubebuilder:validation:Minimum=1
	SizeGb int64 `json:"sizeGb"`

	// StorageAccountType is the disk storage account type (Standard_LRS or Premium_LRS)
	// +kubebuilder:validation:Enum=Standard_LRS,Premium_LRS
	// +optional
	StorageAccountType string `json:"storageAccountType,omitempty"`

	// Caching is the type of the disk caching. It can be either None, ReadOnly or ReadWrite
	// +kubebuilder:validation:Enum=None,ReadOnly,ReadWrite
	// +optional
	Caching string `json:"caching,omitempty"`
}

// AzurePlanStatus defines the observed state of the Azure Plan.
type AzurePlanStatus struct {
	// The generation observed by the Azure Plan controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a Azure Plan current state.
	// +optional
	Conditions []AzurePlanCondition `json:"conditions,omitempty"`
}

// AzurePlanCondition describes the state of a Azure Plan at a certain point.
type AzurePlanCondition struct {
	// Type of Azure Plan condition.
	Type AzurePlanConditionType `json:"type"`

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

// AzurePlanConditionType is the type of a Azure Plan condition
type AzurePlanConditionType string

// These are valid conditions of a Azure Plan.
const (
	// Validated means the Azure Plan has been validated.
	AzurePlanValidated AzurePlanConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurePlanList contains a list of Azure plans.
type AzurePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzurePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzurePlan{}, &AzurePlanList{})
}

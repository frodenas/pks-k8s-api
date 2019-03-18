/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// VSpherePlanKind defines the VSpherePlan Kind.
	VSpherePlanKind = "VSpherePlan"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VSpherePlan is the Schema for the vSphere Plans API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Description",type="string",priority="0",JSONPath=".spec.description"
type VSpherePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VSpherePlanSpec   `json:"spec,omitempty"`
	Status VSpherePlanStatus `json:"status,omitempty"`
}

// VSpherePlanSpec represents a vSphere Plan specification.
type VSpherePlanSpec struct {
	// Description is the description of the vSphere Plan.
	// +optional
	Description string `json:"description,omitempty"`

	// ProviderSpec is the vSphere Provider specification.
	ProviderSpec VSphereProviderSpec `json:"provider"`

	// ComputeSpec is a vSphere Compute specification.
	ComputeSpec VSphereComputeSpec `json:"compute"`

	// NetworkSpec is a vSphere Network specification.
	NetworkSpec VSphereNetworkSpec `json:"network"`

	// StorageSpec is a vSphere Storage specification.
	StorageSpec VSphereStorageSpec `json:"storage"`
}

// VSphereProviderSpec represents a vSphere Provider specification.
type VSphereProviderSpec struct {
	// Secret containing vSphere credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`

	// Insecure determines whether communication with vCenter uses SSL validation.
	Insecure bool `json:"insecure"`
}

const (
	// VSphereProviderCredentialsVCenterURLKey sets the key for the vCenter URL value.
	VSphereProviderCredentialsVCenterURLKey = "vCenterURL"

	// VSphereProviderCredentialsVCenterUsernameKey sets the key for the vCenter user value.
	VSphereProviderCredentialsVCenterUsernameKey = "vCenterUsername"

	// VSphereProviderCredentialsVCenterPasswordKey sets the key for the vCenter password value.
	VSphereProviderCredentialsVCenterPasswordKey = "vCenterPassword"
)

// VSphereComputeSpec represents a vSphere Compute specification.
type VSphereComputeSpec struct {
	// MastersSpec is the vSphere Compute Masters specification.
	MastersSpec VSphereComputeMastersSpec `json:"masters"`

	// WorkersSpec is the vSphere Compute Workers specification.
	WorkersSpec VSphereComputeWorkersSpec `json:"workers"`
}

// VSphereComputeMastersSpec represents a vSphere Compute Masters specification.
type VSphereComputeMastersSpec struct {
	// Replicas is the number of master VMs.
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each master VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`

	// VMFolder is the vCenter folder where master VMs will be located.
	VMFolder string `json:"vmFolder"`

	// Zones are the zones where master VMs will be located.
	// +kubebuilder:validation:MinItems=1
	Zones []VSphereZoneSpec `json:"zones"`
}

// VSphereComputeWorkersSpec represents a vSphere Compute Workers specification.
type VSphereComputeWorkersSpec struct {
	// Replicas is the number of worker VMs.
	// +kubebuilder:validation:Minimum=0
	Replicas int32 `json:"replicas"`

	// NumCPUs is the number of CPUs to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=1
	NumCPUs int32 `json:"numCpus"`

	// MemoryMB is the amount of memory in Mb to be assigned to each worker VM.
	// +kubebuilder:validation:Minimum=2
	MemoryMB int64 `json:"memoryMb"`

	// VMFolder is the vCenter folder where master VMs will be located.
	VMFolder string `json:"vmFolder"`

	// Zones are the zones where worker VMs will be located.
	// +kubebuilder:validation:MinItems=1
	Zones []VSphereZoneSpec `json:"zones"`
}

// VSphereZoneSpec represents a vSphere Zone specification.
type VSphereZoneSpec struct {
	// Name is the name of the Zone.
	Name string `json:"name"`

	// Datacenter is the name of the vCenter datacenter associated with the Zone.
	Datacenter string `json:"datacenter"`

	// Cluster is the name of the vCenter cluster associated with the Zone.
	Cluster string `json:"cluster"`

	// ResourcePool is the name of the vCenter resource pool associated with the Zone.
	// +optional
	ResourcePool string `json:"resourcePool,omitempty"`
}

// VSphereNetworkSpec represents a VSphere Network specification.
type VSphereNetworkSpec struct {
	// DNS is a list of Domain Name Servers.
	// +kubebuilder:validation:MinItems=1
	DNS []string `json:"dns,omitempty"`

	// DVSNetworkSpec is the vSphere DVS Network specification.
	// +optional
	DVSNetworkSpec *VSphereDVSNetworkSpec `json:"dvs,omitempty"`

	// NSXTNetworkSpec is the vSphere NSX-T Network specification.
	// +optional
	NSXTNetworkSpec *VSphereNSXTNetworkSpec `json:"nsxt,omitempty"`
}

// VSphereDVSNetworkSpec represents a vSphere DVS Network specification.
type VSphereDVSNetworkSpec struct {
	// Name is the name of the vSphere Network.
	Name string `json:"name"`
}

// VSphereNSXTNetworkSpec represents a vSphere NSX-T Network specification.
type VSphereNSXTNetworkSpec struct {
	// Secret containing vSphere credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`

	// Insecure determines whether communication with NSX-T manager uses SSL validation.
	Insecure bool `json:"insecure"`

	// T0RouterID is the ID of the T0 Router.
	T0RouterID string `json:"t0RouterId"`

	// IPBlockID is the ID of the IP Block used to create VMs.
	// +kubebuilder:validation:MinItems=1
	IPBlockIDs []string `json:"ipBlockIds"`

	// FloatingIPPoolIDs are the IDs of the Floating IP Pools used to create VMs.
	// +kubebuilder:validation:MinItems=1
	FloatingIPPoolIDs []string `json:"floatingIPPoolIds"`

	// NatMode indicates if NAT should be used.
	NatMode bool `json:"natMode"`

	// LBSize is size of the Load Balancer.
	LBSize string `json:"lbSize"`

	// PodSubnetPrefix is prefix of the POD subnet.
	// +optional
	PodSubnetPrefix int `json:"podSubnetPrefix,omitempty"`
}

const (
	// VSphereNSXTCredentialsNSXTURLKey sets the key for the NSX-T Manager URL value.
	VSphereNSXTCredentialsNSXTURLKey = "nsxtURL"

	// VSphereNSXTCredentialsNSXTUsernameKey sets the key for the NSX-T Manager user value.
	VSphereNSXTCredentialsNSXTUsernameKey = "nsxtUsername"

	// VSphereNSXTCredentialsNSXTPasswordKey sets the key for the NSX-T Manager password value.
	VSphereNSXTCredentialsNSXTPasswordKey = "nsxtPassword"
)

// VSphereStorageSpec represents a vSphere Storage specification.
type VSphereStorageSpec struct {
	// MastersSpec is the vSphere Storage Masters specification.
	MastersSpec VSphereStorageMastersSpec `json:"masters"`

	// WorkersSpec is the vSphere Storage Workers specification.
	WorkersSpec VSphereStorageWorkersSpec `json:"workers"`
}

// VSphereStorageMastersSpec represents a vSphere Storage Masters specification.
type VSphereStorageMastersSpec struct {
	// Datastore is the name of the datastore to be used by the Worker VMs.
	Datastore string `json:"datastore"`

	// Disks are the disks to be attached to the master VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []VSphereDiskSpec `json:"disks"`
}

// VSphereStorageWorkersSpec represents a vSphere Storage Workers specification.
type VSphereStorageWorkersSpec struct {
	// Datastore is the name of the datastore to be used by the Worker VMs.
	Datastore string `json:"datastore"`

	// Disks are the disks to be attached to the worker VMs.
	// +kubebuilder:validation:MinItems=1
	Disks []VSphereDiskSpec `json:"disks"`
}

// VSphereDiskSpec represents a vSphere Disk specification.
type VSphereDiskSpec struct {
	// SizeGb is the size in Gb of the disk.
	// +kubebuilder:validation:Minimum=1
	SizeGb int64 `json:"sizeGb"`

	// Label is the label of the Disk.
	// +optional
	Label string `json:"label,omitempty"`
}

// VSpherePlanStatus defines the observed state of the VSphere Plan.
type VSpherePlanStatus struct {
	// The generation observed by the VSphere Plan controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a VSphere Plan current state.
	// +optional
	Conditions []VSpherePlanCondition `json:"conditions,omitempty"`
}

// VSpherePlanCondition describes the state of a VSphere Plan at a certain point.
type VSpherePlanCondition struct {
	// Type of VSphere Plan condition.
	Type VSpherePlanConditionType `json:"type"`

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

// VSpherePlanConditionType is the type of a VSphere Plan condition
type VSpherePlanConditionType string

// These are valid conditions of a VSphere Plan.
const (
	// Validated means the VSphere Plan has been validated.
	VSpherePlanValidated VSpherePlanConditionType = "Validated"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VSpherePlanList contains a list of VSphere plans.
type VSpherePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSpherePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VSpherePlan{}, &VSpherePlanList{})
}

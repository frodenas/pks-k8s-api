/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster is the Schema for the clusters API.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Plan",type="string",priority="0",JSONPath=".spec.planRef.name"
// +kubebuilder:printcolumn:name="Hostname",type="string",priority="0",JSONPath=".spec.externalHostname"
// +kubebuilder:printcolumn:name="Last Operation",type="string",priority="0",JSONPath=".status.lastOperation.type"
// +kubebuilder:printcolumn:name="Status",type="string",priority="0",JSONPath=".status.lastOperation.state"
// +kubebuilder:printcolumn:name="Description",type="string",priority="1",JSONPath=".spec.description"
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterSpec represents a Cluster specification.
type ClusterSpec struct {
	// ExternalHostName is the hostname from which to access the cluster Kubernetes API.
	ExternalHostname string `json:"externalHostname,omitempty"`

	// Description is the description of the cluster.
	// +optional
	Description string `json:"description,omitempty"`

	// NumWorkerReplicas is the number of worker VMs.
	// Setting this parameter overrides the number of worker VMs specified at the plan.
	// +optional
	NumWorkerReplicas int32 `json:"workerReplicas,omitempty"`

	// ProvisionerSpec is the Provisioner specification.
	ProvisionerSpec ProvisionerSpec `json:"provisioner"`

	// Plan is the plan resource associated with the cluster.
	PlanRef corev1.ObjectReference `json:"planRef"`

	// KubernetesProfileRef is the kubernetes profile resource associated with the cluster.
	KubernetesProfileRef corev1.ObjectReference `json:"kubernetesProfileRef"`
}

// ProvisionerSpec represents a Provisioner specification.
type ProvisionerSpec struct {
	// Type is the type of provisioner (supported ones are: DUMMY, BOSH, CAPI, ODB)
	// +kubebuilder:validation:Enum=DUMMY,BOSH,CAPI,ODB
	Type string `json:"type"`

	// Secret containing the provisioner credentials.
	CredentialsSecretRef corev1.SecretReference `json:"credentialsSecretRef"`
}

const (
	// DUMMYProvisioner is a Dummy Provisioner.
	DUMMYProvisioner = "DUMMY"

	// BOSHProvisioner is a BOSH Provisioner.
	BOSHProvisioner = "BOSH"

	// CAPIProvisioner is a Cluster API Provisioner.
	CAPIProvisioner = "CAPI"

	// ODBProvisioner is an On-Demand-Broker Provisioner.
	ODBProvisioner = "ODB"
)

const (
	// BOSHProvisionerCredentialsURLKey sets the key for the BOSH URL value.
	BOSHProvisionerCredentialsURLKey = "url"

	// BOSHProvisionerCredentialsClientIDKey sets the key for BOSH Client ID value.
	BOSHProvisionerCredentialsClientIDKey = "clientId"

	// BOSHProvisionerCredentialsClientSecretKey sets the key for BOSH Client Secret value.
	BOSHProvisionerCredentialsClientSecretKey = "clientSecret"

	// BOSHProvisionerCredentialsCACertKey sets the key for BOSH CA Certificate value.
	BOSHProvisionerCredentialsCACertKey = "CACert"
)

const (
	// ODBProvisionerCredentialsURLKey sets the key for the On-Demand-Broker URL value.
	ODBProvisionerCredentialsURLKey = "url"

	// ODBProvisionerCredentialsUsernameKey sets the key for On-Demand-Broker Username value.
	ODBProvisionerCredentialsUsernameKey = "username"

	// ODBProvisionerCredentialsPasswordKey sets the key for On-Demand-Broker Password value.
	ODBProvisionerCredentialsPasswordKey = "password"

	// ODBProvisionerCredentialsCACertKey sets the key for On-Demand-Broker CA Certificate value.
	ODBProvisionerCredentialsCACertKey = "CACert"
)

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	// The generation observed by the Cluster controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// LastOperation is last operation performed on a cluster.
	// +optional
	LastOperation *ClusterLastOperation `json:"lastOperation,omitempty"`

	// Represents the latest available observations of a Cluster current state.
	// +optional
	Conditions []ClusterCondition `json:"conditions,omitempty"`
}

// ClusterLastOperation represents the detail of the last performed operation on the Cluster object.
type ClusterLastOperation struct {
	// Description is the human-readable description of the last operation.
	// +optional
	Description string `json:"description,omitempty"`

	// StartTime is the timestamp at which LastOperation API was started.
	// +optional
	StartTime metav1.Time `json:"startTime,omitempty"`

	// LastUpdated is the timestamp at which LastOperation API was last-updated.
	// +optional
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`

	// State is the current status of the last performed operation.
	// +optional
	State ClusterLastOperationState `json:"state,omitempty"`

	// Type is the type of operation which was last performed.
	// +optional
	Type ClusterLastOperationType `json:"type,omitempty"`

	// ProvisionerData is a provisioner specific data.
	// +optional
	ProvisionerData string `json:"provisionerData,omitempty"`
}

// ClusterLastOperationState is the current status of the last performed operation.
type ClusterLastOperationState string

// These are the valid status of a last operation.
const (
	// ClusterLastOperationStateSucceed means the last operation performed succeed.
	ClusterLastOperationStateSucceeded ClusterLastOperationState = "Succeeded"

	// ClusterLastOperationStateInProgress means the last operation performed is still in progress.
	ClusterLastOperationStateInProgress ClusterLastOperationState = "InProgress"

	// ClusterLastOperationStateFailed means the last operation performed failed.
	ClusterLastOperationStateFailed ClusterLastOperationState = "Failed"
)

// ClusterLastOperationType is the type of operation which was last performed.
type ClusterLastOperationType string

// These are the valid types of a last operation.
const (
	// ClusterLastOperationTypeCreate means the last operation performed was a create.
	ClusterLastOperationTypeCreate ClusterLastOperationType = "Create"

	// ClusterLastOperationTypeUpdate means the last operation performed was an update.
	ClusterLastOperationTypeUpdate ClusterLastOperationType = "Update"

	// ClusterLastOperationTypeDelete means the last operation performed was a delete.
	ClusterLastOperationTypeDelete ClusterLastOperationType = "Delete"
)

// ClusterCondition describes the state of a Cluster at a certain point.
type ClusterCondition struct {
	// Type of Cluster condition.
	Type ClusterConditionType `json:"type"`

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

	// RawData is arbitrary metadata stored by controller.
	// +optional
	RawData string `json:"rawData,omitempty"`
}

// ClusterConditionType is the type of a Cluster condition
type ClusterConditionType string

// These are valid conditions of a Cluster.
const (
	// Validated means the Cluster has been validated.
	ClusterValidated ClusterConditionType = "Validated"

	// ClusterNSXTProvisioned means the NSXT-T resources for a Cluster have been provisioned.
	ClusterNSXTProvisioned ClusterConditionType = "NSXTProvisioned"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList contains a list of Clusters.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}

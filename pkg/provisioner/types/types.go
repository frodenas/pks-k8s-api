/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package provisionertypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProvisionerLastOperation represents the detail of the last performed operation by the provisioner.
type ProvisionerLastOperation struct {
	// Description is the human-readable description of the last operation.
	Description string `json:"description"`

	// LastUpdated is the timestamp at which LastOperation API was last-updated.
	LastUpdated metav1.Time `json:"lastUpdated"`

	// State is the current status of the last performed operation.
	State ProvisionerOperationState `json:"state"`

	// Type is the type of operation which was last performed.
	Type ProvisionerOperationType `json:"type"`

	// ProvisionerData is a provisioner arbitrary data.
	// +optional
	ProvisionerData string `json:"provisionerData,omitempty"`
}

// ProvisionerOperationState is the current status of the last performed operation.
type ProvisionerOperationState string

const (
	// ProvisionerOperationStateSucceeded means the last operation performed succeed.
	ProvisionerOperationStateSucceeded ProvisionerOperationState = "Succeeded"

	// ProvisionerOperationStateInProgress means the last operation performed is still in progress.
	ProvisionerOperationStateInProgress ProvisionerOperationState = "InProgress"

	// ProvisionerOperationStateFailed means the last operation performed failed.
	ProvisionerOperationStateFailed ProvisionerOperationState = "Failed"
)

// ProvisionerOperationType is the type of operation which was last performed.
type ProvisionerOperationType string

const (
	// ProvisionerOperationTypeCreate means the last operation performed was a create.
	ProvisionerOperationTypeCreate ProvisionerOperationType = "Create"

	// ProvisionerOperationTypeUpdate means the last operation performed was an update.
	ProvisionerOperationTypeUpdate ProvisionerOperationType = "Update"

	// ProvisionerOperationTypeDelete means the last operation performed was a delete.
	ProvisionerOperationTypeDelete ProvisionerOperationType = "Delete"
)

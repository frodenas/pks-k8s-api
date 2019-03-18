/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"github.com/cloudfoundry/bosh-cli/director"
)

// Client is a client to communicate with BOSH.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// DeleteDeployment deletes a deployment given a deployment name.
	DeleteDeployment(name string, force bool) (int, error)

	// Deploy deploys deployment given a deployment manifest.
	Deploy(name string, manifest []byte) (int, error)

	// GetDeployment gets a deployment given a deployment name.
	GetDeployment(name string) (bool, director.Deployment, error)

	// GetTask gets a BOSH task give a task id.
	GetTask(taskID int) (BoshTask, error)
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/director"
)

// GetDeployment gets a deployment given a deployment name.
func (bc *client) GetDeployment(name string) (bool, director.Deployment, error) {
	boshDirector, err := bc.Director(director.NewNoopTaskReporter())
	if err != nil {
		return false, nil, fmt.Errorf("failed to build director: %v", err)
	}

	deployments, err := boshDirector.Deployments()
	if err != nil {
		return false, nil, fmt.Errorf("failed to list deployments: %v", err)
	}

	for _, deployment := range deployments {
		if deployment.Name() == name {
			return true, deployment, nil
		}
	}

	return false, nil, nil
}

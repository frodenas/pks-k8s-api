/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/director"
)

// Deploy deploys deployment given a deployment manifest.
func (bc *client) Deploy(name string, manifest []byte) (int, error) {
	asyncTaskReporter := NewAsyncTaskReporter()
	boshDirector, err := bc.Director(asyncTaskReporter)
	if err != nil {
		return 0, fmt.Errorf("failed to build director: %v", err)
	}

	deployment, err := boshDirector.FindDeployment(name)
	if err != nil {
		return 0, fmt.Errorf("error when finding deployment `%s`: %v", name, err)
	}

	go func() {
		err = deployment.Update(manifest, director.UpdateOpts{})
		if err != nil {
			asyncTaskReporter.Err <- fmt.Errorf("error updating deployment `%s`: %v", name, err)
		}
	}()

	select {
	case err := <-asyncTaskReporter.Err:
		return 0, err
	case id := <-asyncTaskReporter.Task:
		return id, nil
	}
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"fmt"
)

// DeleteDeployment deletes a deployment given a deployment name.
func (bc *client) DeleteDeployment(name string, force bool) (int, error) {
	found, _, err := bc.GetDeployment(name)
	if err != nil {
		return 0, fmt.Errorf("error getting deployment `%s`: %v", name, err)
	}
	if !found {
		return 0, nil
	}

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
		err = deployment.Delete(force)
		if err != nil {
			asyncTaskReporter.Err <- fmt.Errorf("error deleting deployment `%s`: %v", name, err)
		}
	}()

	select {
	case err := <-asyncTaskReporter.Err:
		return 0, err
	case id := <-asyncTaskReporter.Task:
		return id, nil
	}
}

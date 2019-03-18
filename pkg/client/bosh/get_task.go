/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

import (
	"fmt"

	"github.com/cloudfoundry/bosh-cli/director"
)

// GetTask gets a BOSH task give a task id.
func (bc *client) GetTask(taskID int) (BoshTask, error) {
	director, err := bc.Director(director.NewNoopTaskReporter())
	if err != nil {
		return BoshTask{}, fmt.Errorf("failed to build director: %v", err)
	}

	task, err := director.FindTask(taskID)
	if err != nil {
		return BoshTask{}, fmt.Errorf("error when finding task id `%d`: %v", taskID, err)
	}

	return BoshTask{
		ID:          task.ID(),
		State:       task.State(),
		Description: task.Description(),
		Result:      task.Result(),
	}, nil
}

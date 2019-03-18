/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

// BoshTask represent a BOSH Task.
type BoshTask struct {
	ID          int
	State       string
	Description string
	Result      string
}

// TaskStateType is the state of a BOSH task.
type TaskStateType int

const (
	TaskQueued     = "queued"
	TaskProcessing = "processing"
	TaskDone       = "done"
	TaskError      = "error"
	TaskCancelled  = "cancelled"
	TaskCancelling = "cancelling"
	TaskTimeout    = "timeout"

	TaskComplete TaskStateType = iota
	TaskIncomplete
	TaskFailed
	TaskUnknown
)

// StateType return the state type of a BOSH task.
func (t BoshTask) StateType() TaskStateType {
	switch t.State {
	case TaskDone:
		return TaskComplete
	case TaskProcessing, TaskQueued, TaskCancelling:
		return TaskIncomplete
	case TaskCancelled, TaskError, TaskTimeout:
		return TaskFailed
	default:
		return TaskUnknown
	}
}

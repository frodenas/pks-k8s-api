/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package bosh

// AsyncTaskReporter is a BOSH asyncronous task reporter.
type AsyncTaskReporter struct {
	Task     chan int
	Err      chan error
	Finished chan bool
	State    string
}

// NewAsyncTaskReporter returns a new BOSH asyncronous task reporter.
func NewAsyncTaskReporter() *AsyncTaskReporter {
	return &AsyncTaskReporter{
		Task:     make(chan int, 1),
		Err:      make(chan error, 1),
		Finished: make(chan bool, 1),
	}
}

// TaskStarted returns a task started,
func (r *AsyncTaskReporter) TaskStarted(taskID int) {
	r.Task <- taskID
}

// TaskFinished returns a task finished.
func (r *AsyncTaskReporter) TaskFinished(taskID int, state string) {
	r.State = state
	r.Finished <- true
}

// TaskOutputChunk returns a Task output chunk.
func (r *AsyncTaskReporter) TaskOutputChunk(taskID int, chunk []byte) {}

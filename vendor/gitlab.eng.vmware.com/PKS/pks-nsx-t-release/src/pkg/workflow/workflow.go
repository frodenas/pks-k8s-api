/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package workflow

import (
	"log"
	"time"

	"github.com/cenkalti/backoff"
)

type SequentialWorkflows struct {
	workflows []Workflow
	backoff   backoff.BackOff
}

func NewSequentialWorkflows(wfs ...Workflow) Workflow {
	swf := SequentialWorkflows{}
	for _, wf := range wfs {
		swf.workflows = append(swf.workflows, wf)
	}
	return &swf
}

func NewSequentialWorkflowsWithBackOff(maxElapsedTime time.Duration, maxTries uint64, wfs ...Workflow) *SequentialWorkflows {
	swf := SequentialWorkflows{}
	for _, wf := range wfs {
		swf.workflows = append(swf.workflows, wf)
	}
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = maxElapsedTime
	swf.backoff = backoff.WithMaxRetries(bo, maxTries)
	return &swf
}

func (s *SequentialWorkflows) Run() error {
	var err error
	if s.backoff == nil {
		for _, step := range s.workflows {
			err = step.Run()
			if err != nil {
				return err
			}
		}
		return nil
	}
	logRetry := func(err error, next time.Duration) {
		log.Printf("retrying in %s, because %s", next, err)
	}
	for _, step := range s.workflows {
		err = backoff.RetryNotify(backoff.Operation(step.Run), s.backoff, logRetry)
		// if err isn't nil and we haven't exhausted our retries, keep trying
		for err != nil && s.backoff.NextBackOff() != backoff.Stop {
			err = backoff.RetryNotify(backoff.Operation(step.Run), s.backoff, logRetry)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type WorkflowFunc func() error

func (w WorkflowFunc) Run() error {
	return w()
}

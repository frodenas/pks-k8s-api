/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package workflow

// Workflow defines a workflow interface
type Workflow interface {
	Run() error
}

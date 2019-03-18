/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package controller

import (
	gcpplan "github.com/frodenas/pks-k8s-api/pkg/controller/gcp_plan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gcpplan.Add)
}

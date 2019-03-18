/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package controller

import (
	odbplan "github.com/frodenas/pks-k8s-api/pkg/controller/odb_plan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, odbplan.Add)
}

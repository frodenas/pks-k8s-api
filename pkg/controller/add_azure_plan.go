/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package controller

import (
	azureplan "github.com/frodenas/pks-k8s-api/pkg/controller/azure_plan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, azureplan.Add)
}

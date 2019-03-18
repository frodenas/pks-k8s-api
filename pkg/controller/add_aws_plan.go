/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package controller

import (
	awsplan "github.com/frodenas/pks-k8s-api/pkg/controller/aws_plan"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, awsplan.Add)
}

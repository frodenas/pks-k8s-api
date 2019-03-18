/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package webhook

import (
	server "github.com/frodenas/pks-k8s-api/pkg/webhook/admission_server"
)

func init() {
	// AddToManagerFuncs is a list of functions to create webhook servers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, server.Add)
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package controller

import (
	kubernetesprofile "github.com/frodenas/pks-k8s-api/pkg/controller/kubernetes_profile"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kubernetesprofile.Add)
}

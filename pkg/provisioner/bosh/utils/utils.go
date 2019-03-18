/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	"fmt"
)

// DeploymentName returns a deployment name given a cluster namespace and a name.
func DeploymentName(namespace string, name string) string {
	return fmt.Sprintf("pks-%s-%s", namespace, name)
}

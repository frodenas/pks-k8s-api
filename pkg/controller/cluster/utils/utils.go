/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	"fmt"
)

// ClusterName returns a cluster name given a namespace and a name.
func ClusterName(namespace string, name string) string {
	return fmt.Sprintf("%s-%s", namespace, name)
}

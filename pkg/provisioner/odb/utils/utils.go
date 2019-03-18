/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	"fmt"
)

// ServiceInstanceName returns a service instance name given a cluster namespace and a name.
func ServiceInstanceName(namespace string, name string) string {
	return fmt.Sprintf("pks-%s-%s", namespace, name)
}

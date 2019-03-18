/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import "strings"

// ResourceSplitter splits a resource.
func ResourceSplitter(resource string) string {
	splits := strings.Split(resource, "/")

	return splits[len(splits)-1]
}

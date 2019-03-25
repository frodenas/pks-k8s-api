/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package admissionserver

import (
	"fmt"

	"github.com/frodenas/pks-k8s-api/pkg/webhook/admission_server/vsphere_plan/validating"
)

func init() {
	for k, v := range validating.Builders {
		_, found := builderMap[k]
		if found {
			log.V(1).Info(fmt.Sprintf(
				"conflicting webhook builder names in builder map: %v", k))
		}
		builderMap[k] = v
	}

	for k, v := range validating.HandlerMap {
		_, found := HandlerMap[k]
		if found {
			log.V(1).Info(fmt.Sprintf(
				"conflicting webhook builder names in handler map: %v", k))
		}
		_, found = builderMap[k]
		if !found {
			log.V(1).Info(fmt.Sprintf(
				"can't find webhook builder name %q in builder map", k))
			continue
		}
		HandlerMap[k] = v
	}
}

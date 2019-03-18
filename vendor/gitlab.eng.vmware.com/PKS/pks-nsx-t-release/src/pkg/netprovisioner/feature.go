/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	version "github.com/hashicorp/go-version"
)

const (
	FEATURE_ALLOCATION_PROFILE = "Router Allocation Profile"
)

const (
	NSXT_2_3_BASE = "2.3"
	NSXT_2_2_BASE = "2.2"
	NSXT_2_1_BASE = "2.1"
)

func FeatureSupported(v1 string, feature string) (bool, error) {
	ver, err := version.NewVersion(v1)
	if err != nil {
		return false, err
	}
	switch feature {
	case FEATURE_ALLOCATION_PROFILE:
		nsxt23, err := version.NewVersion(NSXT_2_3_BASE)
		if err != nil {
			return false, err
		}
		return ver.GreaterThan(nsxt23) || ver.Equal(nsxt23), nil
	}
	return false, UnrecognizedFeatureError
}

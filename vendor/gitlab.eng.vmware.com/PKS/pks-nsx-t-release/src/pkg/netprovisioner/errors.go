/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	UnrecognizedFeatureError = Error("unrecognized feature")
)

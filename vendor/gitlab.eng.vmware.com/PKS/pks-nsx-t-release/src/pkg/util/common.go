/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package util

import (
	"reflect"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ParamsNilOrEmptyError = Error("Input params cannot be nil or empty")
)

// EnsureParams is a utility function to check
// if input param of a function is nil
func EnsureParams(params ...interface{}) error {
	for _, param := range params {
		t := reflect.ValueOf(param)
		switch t.Kind() {
		case reflect.String:
			if t.String() == "" {
				return ParamsNilOrEmptyError
			}
		case reflect.Slice:
			if t.Len() == 0 {
				return ParamsNilOrEmptyError
			}
		}
	}
	return nil
}

// ValidateFields validates if there is no empty fields
func ValidateFields(spec interface{}) bool {
	v := reflect.ValueOf(spec)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return false
		}
	}
	return true
}

// ExtractMetadataFromTags extracts metadata from tags
func ExtractMetadataFromTags(scope string, tags []*models.Tag) string {
	for _, tag := range tags {
		if tag.Scope == scope {
			return tag.Tag
		}
	}
	return ""
}

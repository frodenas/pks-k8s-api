/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package util

// StringPtr converts a string to string pointer
func StringPtr(s string) *string {
	return &s
}

// StringVal returns string from string pointer, nil returns ""
func StringVal(p *string) (s string) {
	if p != nil {
		s = *p
	}
	return
}

// IntPtr converts an int to an int pointer
func IntPtr(v int) *int {
	return &v
}

// IntVal returns int from int pointer, nil returns 0
func IntVal(p *int) (v int) {
	if p != nil {
		v = *p
	}
	return
}

// Int32Ptr converts an int32 to int32 pointer
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int32Val returns int32 from int64 pointer, nil returns 0
func Int32Val(p *int32) (v int32) {
	if p != nil {
		v = *p
	}
	return
}

// Int64Ptr converts an int64 to int64 pointer
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int64Val returns int64 from int64 pointer, nil returns 0
func Int64Val(p *int64) (v int64) {
	if p != nil {
		v = *p
	}
	return
}

// Uint64Ptr converts uint64 to uint64 pointer
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Uint64Val returns uint64 from uint64 pointer, nil returns 0
func Uint64Val(p *uint64) (v uint64) {
	if p != nil {
		v = *p
	}
	return
}

// BoolPtr converts bool into bool pointer
func BoolPtr(v bool) *bool {
	return &v
}

// BoolVal returns bool from bool pointer, nil returns false
func BoolVal(p *bool) (v bool) {
	if p != nil {
		v = *p
	}
	return
}

// Float64Ptr converts float64 to float64 pointer
func Float64Ptr(v float64) *float64 {
	return &v
}

// Float64Val returns float64 from float64 pointer, nil returns 0
func Float64Val(p *float64) (v float64) {
	if p != nil {
		v = *p
	}
	return
}

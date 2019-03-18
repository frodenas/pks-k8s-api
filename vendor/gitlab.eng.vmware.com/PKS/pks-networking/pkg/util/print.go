/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package util

import (
	"bytes"
	"encoding/json"
)

// PrettyPrint formats json string to properly indented json output
func PrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

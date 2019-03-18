/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package utils

import (
	corev1 "k8s.io/api/core/v1"
)

// GetSecretString gets a string value from a secret.
func GetSecretString(secret *corev1.Secret, key string) string {
	data, ok := secret.Data[key]
	if !ok {
		return ""
	}

	return string(data)
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package dummyprovisioner

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("provisioner.dummy")

// Provisioner is a Dummy provisioner.
type Provisioner struct {
	k8sClient client.Client
}

// NewProvisioner returns a new Provisioner.
func NewProvisioner(k8sClient client.Client, credentialsSecret *corev1.Secret) (*Provisioner, error) {
	return &Provisioner{
		k8sClient: k8sClient,
	}, nil
}

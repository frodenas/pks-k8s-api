/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshprovisioner

import (
	"fmt"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/client/bosh"
	"github.com/frodenas/pks-k8s-api/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("provisioner.bosh")

// Provisioner is a BOSH provisioner.
type Provisioner struct {
	k8sClient  client.Client
	boshClient bosh.Client
}

// NewProvisioner returns a new Provisioner.
func NewProvisioner(k8sClient client.Client, credentialsSecret *corev1.Secret) (*Provisioner, error) {
	boshURL, boshClientID, boshClientSecret, boshCACert, err := parseProvisionerCredentials(credentialsSecret)
	if err != nil {
		return nil, fmt.Errorf("error parsing BOSH provider credentials: %v", err)
	}

	boshClient, err := bosh.NewClient(boshURL, boshClientID, boshClientSecret, boshCACert)
	if err != nil {
		return nil, fmt.Errorf("error building a BOSH client: %v", err)
	}

	return &Provisioner{
		k8sClient:  k8sClient,
		boshClient: boshClient,
	}, nil
}

func parseProvisionerCredentials(credentialsSecret *corev1.Secret) (string, string, string, string, error) {
	boshURL := utils.GetSecretString(credentialsSecret, pksv1alpha1.BOSHProvisionerCredentialsURLKey)
	if boshURL == "" {
		return "", "", "", "", fmt.Errorf("BOSH URL is blank")
	}

	boshClientID := utils.GetSecretString(credentialsSecret, pksv1alpha1.BOSHProvisionerCredentialsClientIDKey)
	if boshClientID == "" {
		return "", "", "", "", fmt.Errorf("BOSH Client ID is blank")
	}

	boshClientSecret := utils.GetSecretString(credentialsSecret, pksv1alpha1.BOSHProvisionerCredentialsClientSecretKey)
	if boshClientSecret == "" {
		return "", "", "", "", fmt.Errorf("BOSH Client Secret is blank")
	}

	boshCACert := utils.GetSecretString(credentialsSecret, pksv1alpha1.BOSHProvisionerCredentialsCACertKey)
	if boshCACert == "" {
		return "", "", "", "", fmt.Errorf("BOSH CA certificate is blank")
	}

	return boshURL, boshClientID, boshClientSecret, boshCACert, nil
}

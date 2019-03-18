/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package odbprovisioner

import (
	"fmt"

	pksv1alpha1 "github.com/frodenas/pks-k8s-api/pkg/apis/pks/v1alpha1"
	"github.com/frodenas/pks-k8s-api/pkg/utils"
	osb "github.com/maplain/go-open-service-broker-client/v2"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("provisioner.odb")

// Provisioner is an On-Demand-Broker provisioner.
type Provisioner struct {
	k8sClient client.Client
	osbClient osb.Client
}

// NewProvisioner returns a new Provisioner.
func NewProvisioner(k8sClient client.Client, credentialsSecret *corev1.Secret) (*Provisioner, error) {
	odbURL, odbUsername, odbPassword, odbCACert, err := parseProvisionerCredentials(credentialsSecret)
	if err != nil {
		return nil, fmt.Errorf("error parsing ODB provisioner credentials: %v", err)
	}

	osbConfig := osb.DefaultClientConfiguration()
	osbConfig.Name = "pks-k8s-api"
	osbConfig.URL = odbURL
	osbConfig.APIVersion = osb.Version2_13()
	osbConfig.AuthConfig = &osb.AuthConfig{
		BasicAuthConfig: &osb.BasicAuthConfig{
			Username: odbUsername,
			Password: odbPassword,
		},
	}
	osbConfig.CAData = []byte(odbCACert)

	if odbCACert == "" {
		osbConfig.Insecure = true
	}

	osbClient, err := osb.NewClient(osbConfig)
	if err != nil {
		return nil, fmt.Errorf("error building an Open-Service-Broker client: %v", err)
	}

	return &Provisioner{
		k8sClient: k8sClient,
		osbClient: osbClient,
	}, nil
}

func parseProvisionerCredentials(credentialsSecret *corev1.Secret) (string, string, string, string, error) {
	odbURL := utils.GetSecretString(credentialsSecret, pksv1alpha1.ODBProvisionerCredentialsURLKey)
	if odbURL == "" {
		return "", "", "", "", fmt.Errorf("ODB URL is blank")
	}

	odbUsername := utils.GetSecretString(credentialsSecret, pksv1alpha1.ODBProvisionerCredentialsUsernameKey)
	if odbUsername == "" {
		return "", "", "", "", fmt.Errorf("ODB Username is blank")
	}

	odbPassword := utils.GetSecretString(credentialsSecret, pksv1alpha1.ODBProvisionerCredentialsPasswordKey)
	if odbPassword == "" {
		return "", "", "", "", fmt.Errorf("ODB Password is blank")
	}

	odbCACert := utils.GetSecretString(credentialsSecret, pksv1alpha1.ODBProvisionerCredentialsCACertKey)

	return odbURL, odbUsername, odbPassword, odbCACert, nil
}

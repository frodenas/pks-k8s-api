/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("client.azure")

type client struct {
	context       context.Context
	location      string
	resourceGroup string

	vnetClient network.VirtualNetworksClient
}

// NewClient returns a new Azure client given an Azure environment, location, resource group, subscription id, tenant id, client id, client secret, and a context.
func NewClient(ctx context.Context, environment string, location string, resourceGroup string, subscriptionID string, tenantID string, clientID string, clientSecret string) (Client, error) {
	// Create the Azure authorized.
	env, err := azure.EnvironmentFromName(environment)
	if err != nil {
		return nil, fmt.Errorf("an Azure Environment with name `%s` was not found: %v", environment, err)
	}

	oauthConfig, err := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, fmt.Errorf("creating the OAuth config: %v", err)
	}

	servicePrincipalToken, err := adal.NewServicePrincipalToken(*oauthConfig, clientID, clientSecret, env.ServiceManagementEndpoint)
	if err != nil {
		return nil, fmt.Errorf("creating the service principal token: %v", err)
	}

	// Create the Virtual Network Client.
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	vnetClient.BaseURI = env.ResourceManagerEndpoint

	return &client{
		context:       ctx,
		location:      location,
		resourceGroup: resourceGroup,
		vnetClient:    vnetClient,
	}, nil
}

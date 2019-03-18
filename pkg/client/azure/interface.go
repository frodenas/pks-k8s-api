/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azure

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
)

// Client represents an Azure Client.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// GetSubnet gets an Subnet object given a vnet and subnet names.
	GetSubnet(vnetName string, subnetName string) (*network.Subnet, error)

	// GetVnet gets an Vnet object given a vnet name.
	GetVnet(vnetName string) (*network.VirtualNetwork, error)
}

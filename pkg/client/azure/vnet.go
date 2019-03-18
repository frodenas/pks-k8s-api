/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
)

// GetVnet gets an Vnet object given a vnet name.
func (ac *client) GetVnet(vnetName string) (*network.VirtualNetwork, error) {
	log.Info(fmt.Sprintf("Getting Vnet `%s`", vnetName))

	vnet, err := ac.vnetClient.Get(ac.context, ac.resourceGroup, vnetName, "")
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Vnet `%s`", vnetName))
		return nil, err
	}

	return &vnet, nil
}

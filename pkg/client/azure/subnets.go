/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
)

// GetSubnet gets an Subnet object given a vnet and subnet names.
func (ac *client) GetSubnet(vnetName string, subnetName string) (*network.Subnet, error) {
	log.Info(fmt.Sprintf("Getting Subnet `%s/%s`", vnetName, subnetName))

	vnet, err := ac.GetVnet(vnetName)
	if err != nil {
		return nil, err
	}

	subnets := *vnet.Subnets
	for _, subnet := range subnets {
		if *subnet.Name == subnetName {
			return &subnet, nil
		}
	}

	return nil, fmt.Errorf("Subnet %s` not found in Vnet `%s`", subnetName, vnetName)
}

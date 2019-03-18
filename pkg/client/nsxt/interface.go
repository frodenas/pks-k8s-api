/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/networkmanager"
)

// Client is a client to communicate with a given NSX-T manager.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// CreateClusterNetwork creates the network resources for a cluster.
	CreateClusterNetwork(nm networkmanager.NetworkManager, name string, spec *networkmanager.NSXTClusterSpec) (networkmanager.NetworkInfo, error)

	// DeleteClusterNetwork creates the network resources for a cluster.
	DeleteClusterNetwork(nm networkmanager.NetworkManager, name string) error

	// GetIPBlock gets the IP block object given an IP block id.
	GetIPBlock(ipBlockID string) (*models.IPBlock, error)

	// GetIPPool gets the IP pool object given an IP block id.
	GetIPPool(ipPoolID string) (*models.IPPool, error)

	// GetLogicalRouter gets the logical router object given a logical router id.
	GetLogicalRouter(logicalRouterID string) (*models.LogicalRouter, error)

	// NewNetworkManager returns a new NSX-T network manager.
	NewNetworkManager(nsxtSpec *networkmanager.NSXTSpec) (networkmanager.NetworkManager, error)
}

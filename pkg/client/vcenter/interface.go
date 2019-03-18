/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"github.com/vmware/govmomi/object"
)

// Client is a client to communicate with a given vCenter server.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// GetComputeResource gets the compute resource object given a datacenter and cluster paths.
	GetComputeResource(datacenterPath string, clusterPath string) (*object.ComputeResource, error)

	// GetDataCenter gets the datacenter object given a datacenter path.
	GetDataCenter(datacenterPath string) (*object.Datacenter, error)

	// GetDatastore gets the datastore object given a datastore path.
	GetDatastore(datastorePath string) (*object.Datastore, error)

	// GetFolder gets the folder object given a folder path.
	GetFolder(folderPath string) (*object.Folder, error)

	// GetNetwork gets the network object given a network path.
	GetNetwork(networkPath string) (object.NetworkReference, error)

	// GetResourcePool gets the resource pool object given a datacenter and resource pool paths.
	GetResourcePool(datacenterPath string, resourcePoolPath string) (*object.ResourcePool, error)

	// IsVC returns true if we are connected to a vCenter.
	IsVC() bool
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcp

import (
	"google.golang.org/api/compute/v1"
)

// Client is a client to communicate with GCP.
//go:generate moq -out fakes/client.go -pkg fakes . Client
type Client interface {
	// GetDiskType gets a GCP Disk Type object given a zone and disk type name.
	GetDiskType(zone string, diskType string) (*compute.DiskType, error)

	// GetNetwork gets a GCP Network object given a network name.
	GetNetwork(network string) (*compute.Network, error)

	// GetZone gets a GCP Zone object given a region and zone names.
	GetZone(region string, zone string) (*compute.Zone, error)
}

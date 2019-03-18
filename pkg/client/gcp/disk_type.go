/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcp

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

// GetDiskType gets a GCP Disk Type object given a zone and disk type name.
func (gc *client) GetDiskType(zone string, diskType string) (*compute.DiskType, error) {
	log.Info(fmt.Sprintf("Getting DiskType `%s` from Zone `%s`", diskType, zone))

	n, err := gc.computeService.DiskTypes.Get(gc.project, zone, diskType).Do()
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting DiskType `%s` from Zone `%s`", diskType, zone))
		return nil, err
	}

	return n, nil
}

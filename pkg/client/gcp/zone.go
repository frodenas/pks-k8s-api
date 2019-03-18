/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcp

import (
	"fmt"

	"github.com/frodenas/pks-k8s-api/pkg/client/gcp/utils"
	"google.golang.org/api/compute/v1"
)

// GetZone gets a GCP Zone object given a region and zone names.
func (gc *client) GetZone(region string, zone string) (*compute.Zone, error) {
	log.Info(fmt.Sprintf("Getting Zone `%s/%s`", region, zone))

	z, err := gc.computeService.Zones.Get(gc.project, zone).Do()
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Zone `%s/%s`", region, zone))
		return nil, err
	}

	if utils.ResourceSplitter(z.Region) != region {
		return nil, fmt.Errorf("Zone `%s` does not belong to Region `%s`", zone, region)
	}

	return z, nil
}

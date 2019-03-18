/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package gcp

import (
	"fmt"

	"google.golang.org/api/compute/v1"
)

// GetNetwork gets a GCP Network object given a network name.
func (gc *client) GetNetwork(network string) (*compute.Network, error) {
	log.Info(fmt.Sprintf("Getting Network `%s`", network))

	n, err := gc.computeService.Networks.Get(gc.project, network).Do()
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Network `%s`", network))
		return nil, err
	}

	return n, nil
}

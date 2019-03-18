/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetNetwork gets the network object given a network path.
func (vc *client) GetNetwork(networkPath string) (object.NetworkReference, error) {
	log.Info(fmt.Sprintf("Getting Network `%s`", networkPath))

	network, err := vc.finder.Network(vc.context, networkPath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Network `%s`", networkPath))
		return nil, err
	}

	return network, nil
}

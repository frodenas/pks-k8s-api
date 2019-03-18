/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetDataCenter gets the datacenter object given a datacenter path.
func (vc *client) GetDataCenter(datacenterPath string) (*object.Datacenter, error) {
	log.Info(fmt.Sprintf("Getting Datacenter `%s`", datacenterPath))

	datacenter, err := vc.finder.Datacenter(vc.context, datacenterPath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Datacenter `%s`", datacenterPath))
		return nil, err
	}

	return datacenter, nil
}

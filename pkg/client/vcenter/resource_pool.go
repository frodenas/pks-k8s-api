/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetResourcePool gets the resource pool object given a datacenter and resource pool paths.
func (vc *client) GetResourcePool(datacenterPath string, resourcePoolPath string) (*object.ResourcePool, error) {
	log.Info(fmt.Sprintf("Getting ResourcePool `%s` from Datacenter `%s`", resourcePoolPath, datacenterPath))

	datacenter, err := vc.GetDataCenter(datacenterPath)
	if err != nil {
		return nil, err
	}

	vc.finder.SetDatacenter(datacenter)

	resourcePool, err := vc.finder.ResourcePool(vc.context, resourcePoolPath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting ResourcePool `%s` from Datacenter `%s`", resourcePoolPath, datacenterPath))
		return nil, err
	}

	return resourcePool, nil
}

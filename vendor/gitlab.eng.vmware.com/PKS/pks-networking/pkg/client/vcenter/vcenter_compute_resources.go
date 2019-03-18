/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"github.com/vmware/govmomi/object"
)

// GetComputeResource gets the compute resource object given a datacenter and clusterpath
func (vc *Client) GetComputeResource(datacenterPath string, clusterPath string) (*object.ComputeResource, error) {
	ctx := vc.context
	finder := vc.finder

	datacenter, datacenterErr := vc.GetDataCenter(datacenterPath)
	if datacenterErr != nil {
		return nil, datacenterErr
	}

	finder.SetDatacenter(datacenter)

	cluster, clusterErr := finder.ComputeResource(ctx, clusterPath)
	if clusterErr != nil {
		return nil, clusterErr
	}

	return cluster, nil
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"fmt"

	"github.com/vmware/govmomi/object"
)

// GetComputeResource gets the compute resource object given a datacenter and cluster paths.
func (vc *client) GetComputeResource(datacenterPath string, clusterPath string) (*object.ComputeResource, error) {
	log.Info(fmt.Sprintf("Getting ComputeResource `%s` from Datacenter `%s`", clusterPath, datacenterPath))

	datacenter, err := vc.GetDataCenter(datacenterPath)
	if err != nil {
		return nil, err
	}

	vc.finder.SetDatacenter(datacenter)

	cluster, err := vc.finder.ComputeResource(vc.context, clusterPath)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting ComputeResource `%s` from Datacenter `%s`", clusterPath, datacenterPath))
		return nil, err
	}

	return cluster, nil
}

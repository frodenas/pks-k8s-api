/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"github.com/vmware/govmomi/object"
)

// GetDataCenter gets the datacenter object given a datacenter path
func (vc *Client) GetDataCenter(datacenterPath string) (*object.Datacenter, error) {
	finder := vc.finder
	ctx := vc.context

	datacenter, datacenterErr := finder.Datacenter(ctx, datacenterPath)
	if datacenterErr != nil {
		return nil, datacenterErr
	}
	return datacenter, nil
}

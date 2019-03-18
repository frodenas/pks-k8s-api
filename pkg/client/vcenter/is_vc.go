/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package vcenter

// IsVC returns true if we are connected to a vCenter.
func (vc *client) IsVC() bool {
	return vc.client.IsVC()
}

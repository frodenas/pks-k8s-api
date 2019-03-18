/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package boshprovisioner

// BOSHProvisionerData is the BOSHProvisioner arbitrary data.
type BOSHProvisionerData struct {
	// TaskID is the identifier of the BOSH task.
	TaskID int `json:"TaskId"`
}

/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	nca "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/nsx_component_administration"
)

// ReadClusterStatus gets cluster status
func (nc *client) ReadClusterStatus() (*models.ClusterStatus, error) {
	params := nca.NewReadClusterStatusParams()
	res, err := nc.client.NsxComponentAdministration.ReadClusterStatus(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// GetIPPool gets the IP pool object given an IP block id.
func (nc *client) GetIPPool(ipPoolID string) (*models.IPPool, error) {
	log.Info(fmt.Sprintf("Getting IP Pool with ID `%s`", ipPoolID))

	ipBlock, err := nc.nsxtClient.ReadIPPool(ipPoolID)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting IP Pool with ID `%s`", ipPoolID))
		return nil, err
	}

	return ipBlock, nil
}

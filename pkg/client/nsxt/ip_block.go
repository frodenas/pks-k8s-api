/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// GetIPBlock gets the IP block object given an IP block id.
func (nc *client) GetIPBlock(ipBlockID string) (*models.IPBlock, error) {
	log.Info(fmt.Sprintf("Getting IP Block with ID `%s`", ipBlockID))

	ipBlock, err := nc.nsxtClient.ReadIPBlock(ipBlockID)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting IP Block with ID `%s`", ipBlockID))
		return nil, err
	}

	return ipBlock, nil
}

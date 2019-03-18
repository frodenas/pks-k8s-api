/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// GetLogicalRouter gets the logical router object given a logical router id.
func (nc *client) GetLogicalRouter(logicalRouterID string) (*models.LogicalRouter, error) {
	log.Info(fmt.Sprintf("Getting Logical Router with ID `%s`", logicalRouterID))

	logicalRouter, err := nc.nsxtClient.ReadLogicalRouter(logicalRouterID)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error getting Logical Router with ID `%s`", logicalRouterID))
		return nil, err
	}

	return logicalRouter, nil
}

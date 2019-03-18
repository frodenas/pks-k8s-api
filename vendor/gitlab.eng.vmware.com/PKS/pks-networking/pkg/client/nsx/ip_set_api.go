/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	grpo "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/grouping_objects"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ListIPSet will list all ip sets
func (nc *client) ListIPSets() (*models.IPSetListResult, error) {
	params := grpo.NewListIPSetsParams()
	res, err := nc.client.GroupingObjects.ListIPSets(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) DeleteIPSet(IPSetID string) error {
	params := grpo.NewDeleteIPSetParams().WithForce(util.BoolPtr(true)).WithIPSetID(IPSetID)
	_, err := nc.client.GroupingObjects.DeleteIPSet(params, nc.auth)
	if err != nil {
		return err
	}
	return err
}

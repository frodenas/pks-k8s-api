/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	grpo "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/grouping_objects"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ListNSGroups will list all ns groups
func (nc *client) ListNSGroups() (*models.NSGroupListResult, error) {
	params := grpo.NewListNSGroupsParams()
	res, err := nc.client.GroupingObjects.ListNSGroups(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) CreateNSGroup(grp *models.NSGroup) (*models.NSGroup, error) {
	params := grpo.NewCreateNSGroupParams().WithNSGroup(grp)
	res, err := nc.client.GroupingObjects.CreateNSGroup(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) DeleteNSGroup(NSGroupID string) error {
	params := grpo.NewDeleteNSGroupParams().WithForce(util.BoolPtr(true)).WithNsGroupID(NSGroupID)
	_, err := nc.client.GroupingObjects.DeleteNSGroup(params, nc.auth)
	if err != nil {
		return err
	}
	return err
}

func (nc *client) ReadNSGroup(nsGroupID string) (*models.NSGroup, error) {
	params := grpo.NewReadNSGroupParams().WithNsGroupID(nsGroupID)
	res, err := nc.client.GroupingObjects.ReadNSGroup(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

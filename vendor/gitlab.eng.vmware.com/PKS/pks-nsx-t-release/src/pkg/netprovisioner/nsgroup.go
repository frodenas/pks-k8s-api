/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

// CreateNSGroup creates a nsgroup with provided member criterions
// Note: only accepts NSGroupTagExpression for now because swagger 0.12.0
// doesn't 'discriminator' well, which means polymorphism is not supported
func (p *nsxNetworkProvisioner) CreateNSGroupWithCriteria(name string, membershipCriteria []*models.NSGroupTagExpression,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(name); err != nil {
		return "", err
	}

	req := &models.NSGroup{
		ManagedResource: models.ManagedResource{
			DisplayName: name + nsx.SuffixNSGroup,
			Tags:        tags,
		},
		MembershipCriteria: membershipCriteria,
	}
	p.log.Debugf("CreateNSGroupWithCriterion with spec: %+v\n", req)

	res, err := p.Client.CreateNSGroup(req)
	if err != nil {
		p.log.Errorf("Failed to CreateNSGroupWithCriterion: %+v\n", req)
		return "", err
	}
	p.log.Debugf("Successfully CreateNSGroupWithCriterion: %s\n", res.ID)
	return res.ID, nil
}

func (p *nsxNetworkProvisioner) DeleteNSGroup(id string) error {
	if err := util.EnsureParams(id); err != nil {
		return err
	}

	return p.Client.DeleteNSGroup(id)
}

// CheckNSGroup checks if the given NS Group ID is valid
func (p *nsxNetworkProvisioner) CheckNSGroup(nsGroupID string) error {
	if err := util.EnsureParams(nsGroupID); err != nil {
		return err
	}
	_, err := p.ReadNSGroup(nsGroupID)
	return err
}

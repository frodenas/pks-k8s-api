/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	svc "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/services"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ListFirewallSections will list all firewall sections
func (nc *client) ListFirewallSections() (*models.FirewallSectionListResult, error) {
	params := svc.NewListSectionsParams()
	res, err := nc.client.Services.ListSections(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteFirewallSection deletes firewall sections with cascade=true set, which
// means it will forcifully delete firewall section even and associated rules
func (nc *client) DeleteFirewallSection(FirewallSectionID string) error {
	params := svc.NewDeleteSectionParams().WithSectionID(FirewallSectionID).WithCascade(util.BoolPtr(true))
	_, err := nc.client.Services.DeleteSection(params, nc.auth)
	if err != nil {
		return err
	}
	return err
}

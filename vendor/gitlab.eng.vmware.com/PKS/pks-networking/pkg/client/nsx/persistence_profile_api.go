/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	svc "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/services"
)

// ListLoadBalancerPersistenceProfiles will list all load balancer persistence profiles
func (nc *client) ListLoadBalancerPersistenceProfiles() (*models.LbPersistenceProfileListResult, error) {
	params := svc.NewListLoadBalancerPersistenceProfilesParams()
	res, err := nc.client.Services.ListLoadBalancerPersistenceProfiles(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteLoadBalancerPersistenceProfile deletes load balancer persistence profile with provided id
func (nc *client) DeleteLoadBalancerPersistenceProfile(id string) error {
	params := svc.NewDeleteLoadBalancerPersistenceProfileParams().WithPersistenceProfileID(id)
	_, err := nc.client.Services.DeleteLoadBalancerPersistenceProfile(params, nc.auth)
	return err
}

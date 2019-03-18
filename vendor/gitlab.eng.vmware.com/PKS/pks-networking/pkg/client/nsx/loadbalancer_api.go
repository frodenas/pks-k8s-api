/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"errors"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	svc "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/services"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ListLoadBalancerServices will list all loadbalancer services
func (nc *client) ListLoadBalancerServices() (*models.LbServiceListResult, error) {
	params := svc.NewListLoadBalancerServicesParams()
	res, err := nc.client.Services.ListLoadBalancerServices(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) CreateLoadBalancerService(lbService *models.LbService) (*models.LbService, error) {
	params := svc.NewCreateLoadBalancerServiceParams().WithLbService(lbService)
	resOK, resCreated, err := nc.client.Services.CreateLoadBalancerService(params, nc.auth)
	if err != nil {
		return nil, err
	}
	if resOK != nil {
		return resOK.Payload, nil
	}
	if resCreated != nil {
		return resCreated.Payload, nil
	}
	return nil, errors.New("nsx/client: CreateLoadBalancerService: cannot have both resOK and resCreated nil")
}

// try to parse a *models.APIError, if succeeded, create a new error with error messages from APIError
// otherwise, return original error
func ParseAPIError(e *models.APIError) error {
	msg := e.RelatedAPIError.ErrorMessage + "\n"
	if e.RelatedErrors != nil {
		msg = msg + "relatedErrors error messages:\n"
		for _, re := range e.RelatedErrors {
			if re != nil {
				msg = msg + re.ErrorMessage + "\n"
			}
		}
	}
	return errors.New(msg)
}

func (nc *client) DeleteLoadBalancerService(LoadBalancerServiceID string) error {
	params := svc.NewDeleteLoadBalancerServiceParams().WithServiceID(LoadBalancerServiceID)
	_, err := nc.client.Services.DeleteLoadBalancerService(params, nc.auth)
	if err != nil {
		return err
	}
	return err
}

func (nc *client) ReadLoadBalancerService(LoadBalancerServiceID string) (*models.LbService, error) {
	params := svc.NewReadLoadBalancerServiceParams().WithServiceID(LoadBalancerServiceID)
	res, err := nc.client.Services.ReadLoadBalancerService(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, err
}

// ListLoadbalancerVirtualServers will list all loadbalancer virtual servers
func (nc *client) ListLoadBalancerVirtualServers() (*models.LbVirtualServerListResult, error) {
	params := svc.NewListLoadBalancerVirtualServersParams()
	res, err := nc.client.Services.ListLoadBalancerVirtualServers(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) CreateLoadBalancerVirtualServer(virtualServer *models.LbVirtualServer) (*models.LbVirtualServer, error) {
	params := svc.NewCreateLoadBalancerVirtualServerParams().WithLbVirtualServer(virtualServer)
	resOK, resCreated, err := nc.client.Services.CreateLoadBalancerVirtualServer(params, nc.auth)
	if err != nil {
		return nil, err
	}
	if resOK != nil {
		return resOK.Payload, nil
	}
	if resCreated != nil {
		return resCreated.Payload, nil
	}
	return nil, errors.New("nsx/client: CreateLoadBalancerVirtualServer: cannot have both resOK and resCreated nil")
}

func (nc *client) DeleteLoadBalancerVirtualServer(LoadBalancerVirtualServerID string) error {
	params := svc.NewDeleteLoadBalancerVirtualServerParams().WithVirtualServerID(LoadBalancerVirtualServerID)
	_, err := nc.client.Services.DeleteLoadBalancerVirtualServer(params, nc.auth)
	if err != nil {
		return err
	}
	return err
}

// ListLoadBalancerRules will list all loadbalancer rules
func (nc *client) ListLoadBalancerRules() (*models.LbRuleListResult, error) {
	params := svc.NewListLoadBalancerRulesParams()
	res, err := nc.client.Services.ListLoadBalancerRules(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) DeleteLoadBalancerRule(LoadBalancerRuleID string) error {
	params := svc.NewDeleteLoadBalancerRuleParams().WithRuleID(LoadBalancerRuleID)
	_, err := nc.client.Services.DeleteLoadBalancerRule(params, nc.auth)
	return err
}

// ListLoadBalancerPools will list all loadbalancer pools
func (nc *client) ListLoadBalancerPools() (*models.LbPoolListResult, error) {
	params := svc.NewListLoadBalancerPoolsParams()
	res, err := nc.client.Services.ListLoadBalancerPools(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) CreateLoadBalancerPool(pool *models.LbPool) (*models.LbPool, error) {
	params := svc.NewCreateLoadBalancerPoolParams().WithLbPool(pool)
	resOK, resCreated, err := nc.client.Services.CreateLoadBalancerPool(params, nc.auth)
	if err != nil {
		return nil, err
	}
	if resOK != nil {
		return resOK.Payload, nil
	}
	if resCreated != nil {
		return resCreated.Payload, nil
	}
	return nil, errors.New("nsx/client: CreateLoadBalancerPool: cannot have both resOK and resCreated nil")
}

func (nc *client) PerformPoolMemberAction(action, poolID string, settings *models.PoolMemberSettingList) (*models.LbPool, error) {
	params := svc.NewPerformPoolMemberActionParams().WithAction(action).WithPoolID(poolID).WithPoolMemberSettingList(settings)
	res, err := nc.client.Services.PerformPoolMemberAction(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) DeleteLoadBalancerPool(LoadBalancerPoolID string) error {
	params := svc.NewDeleteLoadBalancerPoolParams().WithPoolID(LoadBalancerPoolID)
	_, err := nc.client.Services.DeleteLoadBalancerPool(params, nc.auth)
	return err
}

// release associated external ip address on lb virtual server
func (nc *client) ReleaseLoadBalancerVirtualServerIP(vs *models.LbVirtualServer, readOnly bool) error {
	var err error
	if util.StringVal(vs.IPAddress) == "" {
		return err
	}
	if !ScopeExist(&vs.ManagedResource, NcpTagKeyExternalIPPool) {
		return err
	}
	poolID := util.StringVal(EvaluateTag(&vs.ManagedResource, NcpTagKeyExternalIPPool))
	if poolID == "" {
		return err
	}
	nc.VerboseInfo("external IP %s of lb virtual server %s allocated from external pool %s to be released\n", util.StringVal(vs.IPAddress), vs.ID, poolID)
	if readOnly {
		return nil
	}
	req := &models.AllocationIPAddress{}
	req.AllocationID = util.StringVal(vs.IPAddress)
	err = nc.ReleaseIPToIPPool(poolID, req)
	// don't return error here because multiple virtual servers might be
	// associated with the same IP
	if err != nil {
		nc.Error("failed to release external IP %s of lb virtual server %s to pool %s: %s", util.StringVal(vs.IPAddress), vs.ID, poolID, err.Error())
	} else {
		nc.VerboseInfo("external IP %s of lb virtual server %s allocated from external pool %s released successfully\n", util.StringVal(vs.IPAddress), vs.ID, poolID)
	}
	return nil
}

func (nc *client) ListLoadBalancerApplicationProfiles() (*models.LbAppProfileListResult, error) {
	params := svc.NewListLoadBalancerApplicationProfilesParams()
	res, err := nc.client.Services.ListLoadBalancerApplicationProfiles(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// only able to create tcpMonitor for now
func (nc *client) CreateLoadBalancerTcpMonitor(tcpMonitor *models.LbTCPMonitor) (*models.LbTCPMonitor, error) {
	params := svc.NewCreateLoadBalancerMonitorParams().WithLbMonitor(tcpMonitor)
	resOK, resCreated, err := nc.client.Services.CreateLoadBalancerMonitor(params, nc.auth)
	if err != nil {
		return nil, err
	}
	if resOK != nil {
		return resOK.Payload, nil
	}
	if resCreated != nil {
		return resCreated.Payload, nil
	}
	return nil, errors.New("nsx/client: CreateLoadBalancerTcpMonitor: cannot have both resOK and resCreated nil")
}

func (nc *client) DeleteLoadBalancerMonitor(LoadBalancerMonitorID string) error {
	params := svc.NewDeleteLoadBalancerMonitorParams().WithMonitorID(LoadBalancerMonitorID)
	_, err := nc.client.Services.DeleteLoadBalancerMonitor(params, nc.auth)
	return err
}

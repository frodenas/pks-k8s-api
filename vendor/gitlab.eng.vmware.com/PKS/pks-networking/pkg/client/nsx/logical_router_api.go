/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	lrs "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/logical_routing_and_services"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ListLogicalRouters list logical routers
func (nc *client) ListLogicalRouters() (*models.LogicalRouterListResult, error) {
	params := lrs.NewListLogicalRoutersParams()

	res, err := nc.client.LogicalRoutingAndServices.ListLogicalRouters(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// CreateLogicalRouter creates logical router
func (nc *client) CreateLogicalRouter(logicalRouterModel *models.LogicalRouter) (*models.LogicalRouter, error) {
	params := lrs.NewCreateLogicalRouterParams().WithLogicalRouter(logicalRouterModel)

	res, err := nc.client.LogicalRoutingAndServices.CreateLogicalRouter(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteLogicalRouter deletes the specified logical router
func (nc *client) DeleteLogicalRouter(LogicalRouterID string) error {
	params := lrs.NewDeleteLogicalRouterParams().WithForce(util.BoolPtr(true)).WithLogicalRouterID(LogicalRouterID)

	_, err := nc.client.LogicalRoutingAndServices.DeleteLogicalRouter(params, nc.auth)

	return err
}

// ReadLogicalRouter reads a logical router corresponding to an ID
func (nc *client) ReadLogicalRouter(LogicalRouterID string) (*models.LogicalRouter, error) {
	params := lrs.NewReadLogicalRouterParams().WithLogicalRouterID(LogicalRouterID)

	res, err := nc.client.LogicalRoutingAndServices.ReadLogicalRouter(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, err
}

// UpdateLogicalRouter updates a logical router
func (nc *client) UpdateLogicalRouter(logicalRouterModel *models.LogicalRouter) (*models.LogicalRouter, error) {
	params := lrs.NewUpdateLogicalRouterParams().WithLogicalRouter(logicalRouterModel).WithLogicalRouterID(logicalRouterModel.ID)

	res, err := nc.client.LogicalRoutingAndServices.UpdateLogicalRouter(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetAdvertisementConfig retrieves advertisement config of a router
func (nc *client) GetAdvertisementConfig(logicalRouterID string) (*models.AdvertisementConfig, error) {
	params := lrs.NewReadAdvertisementConfigParams().WithLogicalRouterID(logicalRouterID)

	res, err := nc.client.LogicalRoutingAndServices.ReadAdvertisementConfig(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// UpdateAdvertisementConfig updates advertisement config of a router
func (nc *client) UpdateAdvertisementConfig(logicalRouterID string, advertisementConfigModel *models.AdvertisementConfig) (*models.AdvertisementConfig, error) {
	params := lrs.NewUpdateAdvertisementConfigParams().WithLogicalRouterID(logicalRouterID).
		WithAdvertisementConfig(advertisementConfigModel)

	res, err := nc.client.LogicalRoutingAndServices.UpdateAdvertisementConfig(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// CreateLogicalRouterPort creates logical router port
func (nc *client) CreateLogicalRouterPort(logicalRouterPortModel *models.LogicalRouterPort) (*models.LogicalRouterPort, error) {
	params := lrs.NewCreateLogicalRouterPortParams().WithLogicalRouterPort(logicalRouterPortModel)

	res, err := nc.client.LogicalRoutingAndServices.CreateLogicalRouterPort(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListLogicalRouterPorts returns a list of router ports optionally associated
// with a specific router
func (nc *client) ListLogicalRouterPorts(routerID *string) (*models.LogicalRouterPortListResult, error) {
	params := lrs.NewListLogicalRouterPortsParams().WithLogicalRouterID(routerID)

	res, err := nc.client.LogicalRoutingAndServices.ListLogicalRouterPorts(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, err
}

// ListLogicalRouterPortsForSwitch returns a list of router ports optionally associated
// with a specific switch
func (nc *client) ListLogicalRouterPortsForSwitch(switchID *string) (*models.LogicalRouterPortListResult, error) {
	params := lrs.NewListLogicalRouterPortsParams().WithLogicalSwitchID(switchID)

	res, err := nc.client.LogicalRoutingAndServices.ListLogicalRouterPorts(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, err
}

func (nc *client) GetTier1LinkPort(routerID string) (*models.LogicalRouterPort, error) {
	ports, err := nc.ListLogicalRouterPorts(util.StringPtr(routerID))
	if err != nil {
		return nil, err
	}
	for _, port := range ports.Results {
		if port.ResourceType == PortTypeLinkPortOnTier1 {
			return port, nil
		}
	}
	return nil, nil
}

// DeleteLogicalRouterPort deletes the specified logical router port
func (nc *client) DeleteLogicalRouterPort(portID string) error {
	params := lrs.NewDeleteLogicalRouterPortParams().WithLogicalRouterPortID(portID).WithForce(util.BoolPtr(true))

	_, err := nc.client.LogicalRoutingAndServices.DeleteLogicalRouterPort(params, nc.auth)

	return err
}

// ListLogicalRoutersByType gets logical routers given a type
func (nc *client) ListLogicalRoutersByType(routerType string) (*models.LogicalRouterListResult, error) {
	params := lrs.NewListLogicalRoutersParams().WithRouterType(&routerType)
	res, err := nc.client.LogicalRoutingAndServices.ListLogicalRouters(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListT0LogicalRouters gets all T0 logical routers
func (nc *client) ListT0LogicalRouters() (*models.LogicalRouterListResult, error) {
	routerType := RouterTypeTier0
	res, err := nc.ListLogicalRoutersByType(routerType)
	return res, err
}

// ListT1LogicalRouters gets all T1 logical routers
func (nc *client) ListT1LogicalRouters() (*models.LogicalRouterListResult, error) {
	routerType := RouterTypeTier1
	res, err := nc.ListLogicalRoutersByType(routerType)
	return res, err
}

// AddNatRule adds nat rule to router
func (nc *client) AddNatRule(logicalRouterID string, natRuleModel *models.NatRule) (*models.NatRule, error) {
	params := lrs.NewAddNatRuleParams().WithLogicalRouterID(logicalRouterID).
		WithNatRule(natRuleModel)
	res, err := nc.client.LogicalRoutingAndServices.AddNatRule(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListNatRules lists nat rules of router
func (nc *client) ListNatRules(logicalRouterID string) (*models.NatRuleListResult, error) {
	params := lrs.NewListNatRulesParams().WithLogicalRouterID(logicalRouterID)
	res, err := nc.client.LogicalRoutingAndServices.ListNatRules(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetNatRule gets a NAT rule on a router given an ID
func (nc *client) GetNatRule(LogicalRouterID, natRuleID string) (*models.NatRule, error) {
	params := lrs.NewGetNatRuleParams().WithLogicalRouterID(LogicalRouterID).WithRuleID(natRuleID)

	res, err := nc.client.LogicalRoutingAndServices.GetNatRule(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, err
}

// DeleteNatRule deletes a nat rule on a router
func (nc *client) DeleteNatRule(logicalRouterID, natRuleID string) error {
	params := lrs.NewDeleteNatRuleParams().WithLogicalRouterID(logicalRouterID).
		WithRuleID(natRuleID)
	_, err := nc.client.LogicalRoutingAndServices.DeleteNatRule(params, nc.auth)
	if err != nil {
		return err
	}
	return nil
}

// TagRouter adds tag to a router component
func (nc *client) TagRouter(routerID string, tags []*models.Tag) (*models.LogicalRouter, error) {
	router, err := nc.ReadLogicalRouter(routerID)
	if err != nil {
		return nil, err
	}
	err = ValidateTags(router.ManagedResource, tags)
	if err != nil {
		return nil, err
	}

	router.Tags = append(router.Tags, tags...)
	res, err := nc.UpdateLogicalRouter(router)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (nc *client) RemoveRouterLinkPort(t1, t0 string, readOnly bool) error {
	nc.VerboseInfo("router link port from tier1 router %s to tier0 router %s to be removed\n", t1, t0)
	t1LinkPort, err := nc.GetTier1LinkPort(t1)
	if err != nil {
		return err
	}
	if t1LinkPort == nil {
		nc.Warn("could not get tier1 router link port for tier1 router %s\n", t1)
		return nil
	}
	nc.VerboseInfo("tier 1 link port %s on tier1 router %s to be removed\n", t1LinkPort.ID, t1)
	if !readOnly {
		err = nc.DeleteLogicalRouterPort(t1LinkPort.ID)
		if err != nil {
			return err
		}
		nc.VerboseInfo("tier 1 link port %s on router %s is removed successfully\n", t1LinkPort.ID, t1)
	}
	t0LinkPort, ok := t1LinkPort.LinkedLogicalRouterPortID.(map[string]interface{})
	if !ok {
		nc.Warn("could not get corresponding tier0 router link port for tier1 router link port %s", t1LinkPort.ID)
		return nil
	}
	t0LinkPortID, ok := t0LinkPort["target_id"].(string)
	if !ok {
		nc.Warn("could not get tier0 router link port for tier1 router link port %s", t1LinkPort.ID)
		return nil
	}
	nc.VerboseInfo("tier 0 link port %s on tier0 router %s to be removed\n", t0LinkPortID, t0)
	if !readOnly {
		err = nc.DeleteLogicalRouterPort(t0LinkPortID)
		if err != nil {
			return err
		}
		nc.VerboseInfo("tier 0 link port %s on router %s is removed successfully\n", t0LinkPortID, t0)
	}
	return nil
}

// release external ip address allocated to logical router
func (nc *client) ReleaseLogicalRouterExternalIP(router *models.LogicalRouter, readOnly bool) error {
	var err error
	if !ScopeExist(&router.ManagedResource, NcpTagKeyExternalIPPoolForSnat) {
		return err
	}
	poolID := util.StringVal(EvaluateTag(&router.ManagedResource, NcpTagKeyExternalIPPoolForSnat))
	if poolID == "" {
		return err
	}

	if !ScopeExist(&router.ManagedResource, NcpTagKeySnatIPForLogicalRouter) {
		return err
	}
	snatIP := util.StringVal(EvaluateTag(&router.ManagedResource, NcpTagKeySnatIPForLogicalRouter))
	if snatIP == "" {
		return err
	}
	nc.VerboseInfo("external IP %s of logical router %s allocated from external pool %s to be released\n", snatIP, router.ID, poolID)

	if !readOnly {
		req := &models.AllocationIPAddress{}
		req.AllocationID = snatIP
		err = nc.ReleaseIPToIPPool(poolID, req)
		if err != nil {
			nc.Error("failed to release external IP %s of logical router %s allocated from external pool %s\n", snatIP, router.ID, poolID)
			return nil
		}
		nc.VerboseInfo("external IP %s of logical router %s allocated from external pool %s is released successfully\n", snatIP, router.ID, poolID)
	}
	return nil
}

// release external ip address allocated to Nat Rule
func (nc *client) ReleaseNatRuleExternalIP(natRule *models.NatRule, readOnly bool) error {
	var err error
	if !ScopeExist(&natRule.ManagedResource, NcpTagKeyExternalIPPoolForSnat) {
		return err
	}
	poolID := util.StringVal(EvaluateTag(&natRule.ManagedResource, NcpTagKeyExternalIPPoolForSnat))
	if poolID == "" {
		return err
	}

	if !ScopeExist(&natRule.ManagedResource, NcpTagKeySnat) {
		return err
	}
	ncpSnat := util.StringVal(EvaluateTag(&natRule.ManagedResource, NcpTagKeySnat))
	if ncpSnat != "true" {
		// This is not a error. But a check to see if its a ncp created snat rule
		return nil
	}

	snatIP := natRule.TranslatedNetwork
	if snatIP == "" {
		return err
	}
	nc.VerboseInfo("external IP %s of logical router %s allocated from external pool %s to be released\n", snatIP, natRule.ID, poolID)

	if !readOnly {
		req := &models.AllocationIPAddress{}
		req.AllocationID = snatIP
		err = nc.ReleaseIPToIPPool(poolID, req)
		if err != nil {
			nc.Error("failed to release external IP %s of logical router %s allocated from external pool %s\n", snatIP, natRule.ID, poolID)
			return nil
		}
		nc.VerboseInfo("external IP %s of logical router %s allocated from external pool %s is released successfully\n", snatIP, natRule.ID, poolID)
	}
	return nil
}

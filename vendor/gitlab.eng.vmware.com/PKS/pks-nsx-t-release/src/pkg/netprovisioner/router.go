/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

type CreateT1RouterSpec struct {
	Name          string
	EdgeClusterID string
	LbSize        string
	NSXTVersion   string
	Tags          []*models.Tag
}

// CreateT1Router creates a T1 router in NSX
// func (p *nsxNetworkProvisioner) CreateT1Router(clusterName, edgeClusterID string,
//	serviceRouter bool, lbSize string, version string, tags []*models.Tag) (string, error) {
func (p *nsxNetworkProvisioner) CreateT1Router(spec CreateT1RouterSpec) (string, error) {

	if err := util.EnsureParams(spec.Name); err != nil {
		return "", err
	}

	req := &models.LogicalRouter{
		ManagedResource: models.ManagedResource{
			DisplayName: spec.Name + nsx.SuffixT1Router,
			Tags:        spec.Tags,
		},
		RouterType: util.StringPtr(nsx.RouterTypeTier1),
	}

	// Creating Service Router
	if spec.EdgeClusterID != "" {
		req.EdgeClusterID = spec.EdgeClusterID
		req.HighAvailabilityMode = nsx.HAActiveStandby
		// for nsx-t 2.3
		supported, err := FeatureSupported(spec.NSXTVersion, FEATURE_ALLOCATION_PROFILE)
		if err != nil {
			return "", err
		}
		if supported {
			if err := util.EnsureParams(spec.LbSize); err != nil {
				return "", err
			}
			req.AllocationProfile = &models.EdgeClusterMemberAllocationProfile{
				AllocationPool: &models.EdgeClusterMemberAllocationPool{
					AllocationSize:     util.StringPtr(spec.LbSize),
					AllocationPoolType: util.StringPtr(models.EdgeClusterMemberAllocationPoolAllocationPoolTypeLoadBalancerAllocationPool),
				},
			}
		}
	}

	p.log.Debugf("createT1Router with spec: %+v", req)

	t1Router, err := p.CreateLogicalRouter(req)
	if err != nil {
		p.log.Errorf("Failed to createT1Router: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createT1Router: %s", t1Router.ID)
	return t1Router.ID, nil
}

// CreateT0ToT1Port creates a port in T0 router
func (p *nsxNetworkProvisioner) CreateT0ToT1Port(clusterName, routerID string,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(clusterName, routerID); err != nil {
		return "", err
	}

	req := &models.LogicalRouterPort{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixT0ToT1Port,
			Tags:        tags,
		},
		ResourceType:    nsx.PortTypeLinkPortOnTier0,
		LogicalRouterID: util.StringPtr(routerID),
	}
	p.log.Debugf("createT0ToT1Port with spec: %+v", req)

	t0ToT1Port, err := p.CreateLogicalRouterPort(req)
	if err != nil {
		p.log.Errorf("Failed to createT0ToT1Port: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createT0ToT1Port: %s", t0ToT1Port.ID)
	return t0ToT1Port.ID, nil
}

// CreateT1ToT0Port creates a port in T1 router, then link it to existing T0 router port
func (p *nsxNetworkProvisioner) CreateT1ToT0Port(clusterName, routerID, portID string,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(clusterName, routerID, portID); err != nil {
		return "", err
	}

	req := &models.LogicalRouterPort{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixT1ToT0Port,
			Tags:        tags,
		},
		ResourceType:    nsx.PortTypeLinkPortOnTier1,
		LogicalRouterID: util.StringPtr(routerID),
		LinkedLogicalRouterPortID: &models.ResourceReference{
			TargetType: nsx.PortTypeLinkPortOnTier0,
			TargetID:   portID,
		},
	}
	p.log.Debugf("createT1ToT0Port with spec: %+v", req)

	t1ToT0Port, err := p.CreateLogicalRouterPort(req)
	if err != nil {
		p.log.Errorf("Failed to createT1ToT0Port: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createT1ToT0Port: %s", t1ToT0Port.ID)
	return t1ToT0Port.ID, nil
}

// EnableRouteAdvertisement updates router with all advertisement flags enabled
func (p *nsxNetworkProvisioner) EnableRouteAdvertisement(routerID string, lb bool) error {
	if err := util.EnsureParams(routerID); err != nil {
		return err
	}

	currentAdvertisement, err := p.GetAdvertisementConfig(routerID)
	if err != nil {
		p.log.Errorf("Failed getting current advertisement on router: %s", routerID)
		return err
	}

	req := &models.AdvertisementConfig{
		ManagedResource: models.ManagedResource{
			RevisionedResource: models.RevisionedResource{
				Revision: currentAdvertisement.Revision,
			},
		},
		Enabled: util.BoolPtr(true),
	}
	if lb {
		req.AdvertiseLbVip = util.BoolPtr(true)
	} else {
		req.AdvertiseNatRoutes = util.BoolPtr(true)
		req.AdvertiseNsxConnectedRoutes = util.BoolPtr(true)
		req.AdvertiseStaticRoutes = util.BoolPtr(true)
	}
	p.log.Debugf("enableRouteAdvertisement with spec: %+v", req)

	_, err = p.UpdateAdvertisementConfig(routerID, req)
	if err != nil {
		p.log.Errorf("Failed to enableRouteAdvertisement: %+v", req)
		return err
	}
	p.log.Debugln("Successfully enableRouteAdvertisement")
	return nil
}

// UntagT0Router removes the cluster name tag from the T0 router
func (p *nsxNetworkProvisioner) UntagT0Router(T0RouterID string, clusterName string) error {
	if err := util.EnsureParams(T0RouterID, clusterName); err != nil {
		return err
	}

	var T0 *models.LogicalRouter
	var err error
	if T0, err = p.ReadLogicalRouter(T0RouterID); err != nil {
		return err
	}
	nsx.RemoveTag(&T0.ManagedResource, models.Tag{
		Scope: nsx.NcpTagKeyCluster,
		Tag:   clusterName,
	})
	if _, err := p.UpdateLogicalRouter(T0); err != nil {
		return err
	}
	return nil
}

// DeleteT0ToT1Port deletes the port on the T0 router connected to the
// cluster T1 router
func (p *nsxNetworkProvisioner) DeleteT0ToT1Port(T0ToT1PortID string) error {
	if err := util.EnsureParams(T0ToT1PortID); err != nil {
		return err
	}
	return p.DeleteLogicalRouterPort(T0ToT1PortID)
}

// DeleteT1Router deletes the cluster T1 router provisioned by PKS
func (p *nsxNetworkProvisioner) DeleteT1Router(T1RouterID string) error {
	if err := util.EnsureParams(T1RouterID); err != nil {
		return err
	}
	return p.DeleteLogicalRouter(T1RouterID)
}

// CheckT0Router checks if the given routerID is valid and is of type T0
func (p *nsxNetworkProvisioner) CheckT0Router(routerID string) error {
	if err := util.EnsureParams(routerID); err != nil {
		return err
	}

	t0rRes, t0rErr := p.ReadLogicalRouter(routerID)
	if t0rErr != nil {
		return t0rErr
	}

	if util.StringVal(t0rRes.RouterType) != nsx.RouterTypeTier0 {
		t0rErr = fmt.Errorf("The Router %s is not a T0 router", routerID)
		return t0rErr
	}

	return nil
}

// ExtractEdgeClusterIDFromT0Router gets edge cluster ID from T0 router ID
func (p *nsxNetworkProvisioner) ExtractEdgeClusterIDFromT0Router(routerID string) (string, error) {
	if err := util.EnsureParams(routerID); err != nil {
		return "", err
	}

	res, err := p.ReadLogicalRouter(routerID)
	if err != nil {
		return "", err
	}
	return res.EdgeClusterID, nil
}

// CreateDnatRule creates DNAT rule for a router
func (p *nsxNetworkProvisioner) CreateDnatRule(routerID, clusterName, destinationIPAddress,
	translatedIPAddress string, tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(routerID, clusterName, destinationIPAddress, translatedIPAddress); err != nil {
		return "", err
	}

	req := &models.NatRule{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixNatRule,
			Tags:        tags,
		},
		Action:                  util.StringPtr(nsx.NatRuleDNAT),
		MatchDestinationNetwork: destinationIPAddress,
		TranslatedNetwork:       translatedIPAddress,
	}
	p.log.Debugf("Create DNAT rule on router %s with spec: %+v", routerID, req)

	res, err := p.AddNatRule(routerID, req)
	if err != nil {
		return "", err
	}
	p.log.Debugf("Successfully created DNAT rule for router %s with spec: %+v", routerID, req)

	return res.ID, nil
}

// CreateSnatRule creates SNAT rule for a router
func (p *nsxNetworkProvisioner) CreateSnatRule(routerID, clusterName, sourceIPAddress,
	translatedIPAddress string, tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(routerID, clusterName, sourceIPAddress, translatedIPAddress); err != nil {
		return "", err
	}

	req := &models.NatRule{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixNatRule,
			Tags:        tags,
		},
		Action:             util.StringPtr(nsx.NatRuleSNAT),
		MatchSourceNetwork: sourceIPAddress,
		TranslatedNetwork:  translatedIPAddress,
		NatPass:            util.BoolPtr(false),
	}
	p.log.Debugf("Create SNAT rule on router %s with spec: %+v", routerID, req)

	res, err := p.AddNatRule(routerID, req)
	if err != nil {
		return "", err
	}
	p.log.Debugf("Successfully created SNAT rule for router %s with spec: %+v", routerID, req)

	return res.ID, nil
}

// DeleteNatRule deletes a NAT rule
func (p *nsxNetworkProvisioner) DeleteNatRule(routerID, natRuleID string) error {

	if err := util.EnsureParams(routerID, natRuleID); err != nil {
		return err
	}

	p.log.Debugf("Deleting NAT rule %s on router %s", natRuleID, routerID)
	err := p.Client.DeleteNatRule(routerID, natRuleID)
	if err != nil {
		return err
	}
	p.log.Debugf("Successfully deleted NAT rule %s on router %s", natRuleID, routerID)

	return nil
}

// Get NAT rule from router
func (p *nsxNetworkProvisioner) GetNatRule(routerID, natRuleID string) (*models.NatRule, error) {
	if err := util.EnsureParams(routerID, natRuleID); err != nil {
		return nil, err
	}

	p.log.Debugf("Getting NAT rule %s on router %s", natRuleID, routerID)

	res, err := p.Client.GetNatRule(routerID, natRuleID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Extract floatingIP from NAT rule
func (p *nsxNetworkProvisioner) ExtractFloatingIPFromNatRule(routerID, natRuleID string) (string, error) {
	var floatingIP string

	p.log.Debugf("Extracting floating IP from NAT rule %s on router %s", natRuleID, routerID)

	res, err := p.GetNatRule(routerID, natRuleID)
	if err != nil {
		return "", err
	}

	if *res.Action == nsx.NatRuleSNAT {
		floatingIP = res.TranslatedNetwork
	} else if *res.Action == nsx.NatRuleDNAT {
		floatingIP = res.MatchDestinationNetwork
	}

	return floatingIP, nil
}

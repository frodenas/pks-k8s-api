/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"errors"

	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/resourcemanager"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

const (
	defaultMonitorFailCount = 3
	defaultMonitorRiseCount = 3
	// seconds
	defaultMonitorInterval = 10
	// seconds
	defaultMonitorTimeout = 10
)

func (p *nsxNetworkProvisioner) GetDefaultFastTCPProfile() (string, error) {
	for i := range resourcemanager.NewResource(p.Client).SetResourceType(nsx.ResourceTypeLbProfile).FilterBy(
		resourcemanager.NameEquals, nsx.NsxDefaultLbFastTcpProfileDisplayName).GetCollection() {
		m, err := nsx.GetManagedResource(i)
		if err != nil {
			return "", err
		}
		// return the first hit
		return m.ID, nil
	}
	return "", errors.New("GetDefaultFastTCPProfile(): nothing found")
}

func (p *nsxNetworkProvisioner) CreateServerPoolWithNSGroupAndActiveMonitors(poolName string, groupID string, activeMonitorIDs []string, tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(poolName); err != nil {
		return "", err
	}

	// by default, use Round_Robin for algorithm and AutoMap for snat translation
	req := &models.LbPool{
		ManagedResource: models.ManagedResource{
			DisplayName: poolName + nsx.SuffixLbPool,
			Tags:        tags,
		},
		ActiveMonitorIds: activeMonitorIDs,
		Algorithm:        util.StringPtr(nsx.AlgorithmRoundRobin),
		SnatTranslation: &models.LbSnatTranslation{
			Type: util.StringPtr(nsx.LbSnatTranslationAutoMap),
		},
		MemberGroup: &models.PoolMemberGroup{
			GroupingObject: &models.ResourceReference{
				TargetID:   groupID,
				TargetType: nsx.ResourceTypeNSGroup,
			},
			MaxIPListSize: util.Int64Ptr(nsx.LbMaxIPListSize),
		},
	}
	p.log.Debugf("CreateServerPoolWithNSGroup with spec: %+v\n", req)

	res, err := p.CreateLoadBalancerPool(req)
	if err != nil {
		p.log.Errorf("Failed to CreateServerPoolWithNSGroup: %+v\n", req)
		return "", err
	}
	p.log.Debugf("Successfully CreateServerPoolWithNSGroup: %s\n", res.ID)
	return res.ID, nil
}

func (p *nsxNetworkProvisioner) CreateVirtualServer(serverName, poolID string, fip strfmt.IPv4, port, profileID string,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(serverName, poolID, port, fip.String()); err != nil {
		return "", err
	}

	req := &models.LbVirtualServer{
		ManagedResource: models.ManagedResource{
			DisplayName: serverName + nsx.SuffixVirtualServer,
			Tags:        tags,
		},
		ApplicationProfileID: util.StringPtr(profileID),
		Enabled:              util.BoolPtr(true),
		IPProtocol:           util.StringPtr(nsx.TCP),
		PoolID:               poolID,
		IPAddress:            util.StringPtr(fip.String()),
		Port:                 util.StringPtr(port),
	}
	p.log.Debugf("CreateVirtualServer with spec: %+v\n", req)

	res, err := p.CreateLoadBalancerVirtualServer(req)
	if err != nil {
		p.log.Errorf("Failed to CreateVirtualServer: %+v\n", req)
		return "", err
	}
	p.log.Debugf("Successfully CreateVirtualServer: %s\n", res.ID)
	return res.ID, nil
}

func (p *nsxNetworkProvisioner) CreateLbService(lbName, lbSize, routerID, serverID string, tags []*models.Tag) (string, error) {
	if err := util.EnsureParams(lbName, routerID, serverID); err != nil {
		return "", err
	}

	req := &models.LbService{
		ManagedResource: models.ManagedResource{
			DisplayName:  lbName,
			Tags:         tags,
			ResourceType: nsx.ResourceTypeLbService,
		},
		Attachment: &models.ResourceReference{
			TargetID:   routerID,
			TargetType: nsx.ResourceTypeLogicalRouter,
		},
		VirtualServerIds: []string{
			serverID,
		},
		Size: util.StringPtr(lbSize),
	}
	p.log.Debugf("CreateLbService with spec: %+v\n", req)

	res, err := p.CreateLoadBalancerService(req)
	if err != nil {
		p.log.Errorf("Failed to CreateLbService: %+v\n", req)
		return "", err
	}
	p.log.Debugf("Successfully CreateLbService: %s\n", res.ID)
	return res.ID, nil
}

func (p *nsxNetworkProvisioner) DeleteLoadBalancerPool(id string) error {
	if err := util.EnsureParams(id); err != nil {
		return err
	}

	return p.Client.DeleteLoadBalancerPool(id)
}

func (p *nsxNetworkProvisioner) DeleteLoadBalancerVirtualServer(id string) error {
	if err := util.EnsureParams(id); err != nil {
		return err
	}

	return p.Client.DeleteLoadBalancerVirtualServer(id)
}

func (p *nsxNetworkProvisioner) DeleteLoadBalancerService(id string) error {
	if err := util.EnsureParams(id); err != nil {
		return err
	}

	return p.Client.DeleteLoadBalancerService(id)
}

func (p *nsxNetworkProvisioner) ReadLoadBalancerService(id string) (*models.LbService, error) {
	if err := util.EnsureParams(id); err != nil {
		return nil, err
	}

	return p.Client.ReadLoadBalancerService(id)
}

func (p *nsxNetworkProvisioner) DeleteLoadBalancerMonitor(id string) error {
	if err := util.EnsureParams(id); err != nil {
		return err
	}

	return p.Client.DeleteLoadBalancerMonitor(id)
}

func (p *nsxNetworkProvisioner) GetLoadbalancerByTag(lbScope, lbTag string) (*models.LbService, error) {
	lbs, err := p.ListLoadBalancerServices()
	if err != nil {
		return nil, err
	}
	for _, lb := range lbs.Results {
		for _, tag := range lb.Tags {
			if tag.Scope == lbScope && tag.Tag == lbTag {
				return lb, nil
			}
		}
	}
	return nil, nil
}

func (p *nsxNetworkProvisioner) CreateLbTcpMonitor(name, port string, tags []*models.Tag) (string, error) {
	if err := util.EnsureParams(name); err != nil {
		return "", err
	}

	req := &models.LbTCPMonitor{
		LbActiveMonitor: models.LbActiveMonitor{
			FallCount:   defaultMonitorFailCount,
			Interval:    defaultMonitorInterval,
			MonitorPort: port,
			RiseCount:   defaultMonitorRiseCount,
			Timeout:     defaultMonitorTimeout,
			LbMonitor: models.LbMonitor{
				ManagedResource: models.ManagedResource{
					DisplayName:  name,
					Tags:         tags,
					ResourceType: nsx.ResourceTypeLbTcpMonitor,
				},
				ResourceType: nsx.ResourceTypeLbTcpMonitor,
			},
		},
	}
	p.log.Debugf("CreateLbTcpMonitor with spec: %+v\n", req)

	res, err := p.CreateLoadBalancerTcpMonitor(req)
	if err != nil {
		p.log.Errorf("Failed to CreateLbTcpMonitor: %+v\n", req)
		return "", err
	}
	p.log.Debugf("Successfully CreateLbTcpMonitor: %s\n", res.ID)
	return res.ID, nil
}

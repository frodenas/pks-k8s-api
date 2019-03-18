/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"strconv"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

// GetAllResources collects nsx resources by searching tag clusterName
func (p *nsxNetworkProvisioner) GetAllResources(clusterName string,
	tag models.Tag) ([]*models.ManagedResource, error) {

	resources, err := p.SearchByTag(nil, tag)
	if err != nil {
		p.log.Errorf("Failed to get all resources from tag for cluster %s", clusterName)
		return nil, err
	}
	return resources.Results, nil
}

// GetMetadataFromSwitchTag gets metadata stored on a switch tag, i.e. floating IP and NAT mode
func (p *nsxNetworkProvisioner) GetMetadataFromSwitchTag(clusterSwitchID string) (string, bool, error) {
	natModeDefault := true

	if err := util.EnsureParams(clusterSwitchID); err != nil {
		return "", natModeDefault, err
	}

	p.log.Debugf("Getting floating IP from logical switch %s", clusterSwitchID)
	clusterSwitch, err := p.GetLogicalSwitch(clusterSwitchID)
	if err != nil {
		return "", natModeDefault, err
	}

	floatingIPAddress := util.ExtractMetadataFromTags(nsx.PksTagKeyFloatingIP, clusterSwitch.Tags)
	if floatingIPAddress == "" {
		p.log.Warnf("Floating IP is not found from logical switch %s", clusterSwitchID)
		return "", natModeDefault, nil
	}
	p.log.Debugf("Successfully get floating IP %s from logical switch %s", floatingIPAddress, clusterSwitchID)

	rawNatMode := util.ExtractMetadataFromTags(nsx.PksTagKeyNat, clusterSwitch.Tags)
	if rawNatMode == "" {
		p.log.Warnf("NAT mode is not found from logical switch %s, default to %s", clusterSwitchID, strconv.FormatBool(natModeDefault))
		return floatingIPAddress, natModeDefault, nil
	}

	natMode, err := strconv.ParseBool(rawNatMode)
	if err != nil {
		return "", natModeDefault, err
	}
	p.log.Debugf("Successfully got NAT mode %s from logical switch %s", rawNatMode, clusterSwitchID)

	return floatingIPAddress, natMode, nil
}

func (p *nsxNetworkProvisioner) GetIpamInfoFromSwitchTag(clusterSwitchID string) (string, string, error) {
	if err := util.EnsureParams(clusterSwitchID); err != nil {
		return "", "", err
	}

	p.log.Debugf("Getting Ipam info from logical switch %s", clusterSwitchID)
	clusterSwitch, err := p.GetLogicalSwitch(clusterSwitchID)
	if err != nil {
		return "", "", err
	}

	nodeIPBlockID := util.ExtractMetadataFromTags(nsx.PksTagKeyNodeIPBlock, clusterSwitch.Tags)
	lbFloatingIPPoolID := util.ExtractMetadataFromTags(nsx.PksTagKeyLBFloatingIPPool, clusterSwitch.Tags)

	return nodeIPBlockID, lbFloatingIPPoolID, nil
}

func (p *nsxNetworkProvisioner) GetT0RouterFromRouterTag(clusterRouterID string) (string, error) {
	if err := util.EnsureParams(clusterRouterID); err != nil {
		return "", err
	}

	p.log.Debugf("Getting T0 Router from logical router %s", clusterRouterID)
	clusterRouter, err := p.ReadLogicalRouter(clusterRouterID)
	if err != nil {
		return "", err
	}

	t0RouterID := util.ExtractMetadataFromTags(nsx.PksTagKeyT0Router, clusterRouter.Tags)
	if t0RouterID == "" {
		p.log.Warnf("T0 Router ID is not found from logical router %s", clusterRouterID)
		return "", nil
	}
	p.log.Debugf("Successfully got T0 Router ID %s from logical router %s", t0RouterID, clusterRouterID)

	return t0RouterID, nil
}

func (p *nsxNetworkProvisioner) GetMasterVMsNSGroupNameFromSwitchPortTag(portID string) (string, error) {
	if err := util.EnsureParams(portID); err != nil {
		return "", err
	}

	p.log.Debug("Getting Master VM NS Group from logical switch port %s", portID)

	clusterSwitchPort, err := p.GetLogicalPort(portID)
	if err != nil {
		return "", err
	}

	masterVMsNSGroupID := util.ExtractMetadataFromTags(nsx.PksTagKeyMasterVMsNSGroup, clusterSwitchPort.Tags)
	if masterVMsNSGroupID == "" {
		p.log.Warnf("Master VM NS Group ID is  not found from logical switch port %s", portID)
		return "", nil
	}

	masterVMsNSGroup, err := p.ReadNSGroup(masterVMsNSGroupID)
	if err != nil {
		return "", err
	}

	if masterVMsNSGroup != nil {
		return masterVMsNSGroup.DisplayName, nil
	}

	return "", nil
}

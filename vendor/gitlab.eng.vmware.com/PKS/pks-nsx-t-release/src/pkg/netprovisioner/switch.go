/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"net"
	"strings"

	"github.com/go-openapi/strfmt"

	"fmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/pool_management"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

// UpdateLogicalSwitchTags updates tags on a logical switch
func (p *nsxNetworkProvisioner) UpdateLogicalSwitchTags(id string, tags []*models.Tag) (string, error) {
	if err := util.EnsureParams(id); err != nil {
		return "", err
	}
	sw, err := p.GetLogicalSwitch(id)
	if err != nil {
		p.log.Errorf("Failed to UpdateLogicalSwitchTags: %+v", err.Error())
		return "", err
	}

	sw.Tags = append(sw.Tags, tags...)
	sw, err = p.UpdateLogicalSwitch(sw)
	if err != nil {
		p.log.Errorf("Failed to UpdateLogicalSwitchTags: %+v", err.Error())
		return "", err
	}
	return sw.ID, nil
}

// CreateLogicalSwitch creates a logical switch in NSX
func (p *nsxNetworkProvisioner) CreateLogicalSwitch(clusterName, overlayTransportZoneID string,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(clusterName, overlayTransportZoneID); err != nil {
		return "", err
	}

	req := &models.LogicalSwitch{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName,
			Tags:        tags,
		},
		TransportZoneID: util.StringPtr(overlayTransportZoneID),
		AdminState:      util.StringPtr(nsx.SwitchAdminStateUp),
		ReplicationMode: nsx.SwitchReplicationModeMtep,
	}
	p.log.Debugf("createLogicalSwitch with spec: %+v", req)

	logicalSwitch, err := p.Client.CreateLogicalSwitch(req)
	if err != nil {
		p.log.Errorf("Failed to createLogicalSwitch: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createLogicalSwitch: %s", logicalSwitch.ID)
	return logicalSwitch.ID, nil
}

// CreateSwitchToT1Port creates a logical port in a logical switch to connect to T1
func (p *nsxNetworkProvisioner) CreateSwitchToT1Port(clusterName, logicalSwitchID string,
	tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(clusterName, logicalSwitchID); err != nil {
		return "", err
	}

	req := &models.LogicalPort{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixSwitchToT1Port,
			Tags:        tags,
		},
		AdminState:      util.StringPtr(nsx.SwitchAdminStateUp),
		LogicalSwitchID: util.StringPtr(logicalSwitchID),
	}
	p.log.Debugf("createSwitchToT1Port with spec: %+v", req)

	switchToT1Port, err := p.CreateLogicalPort(req)
	if err != nil {
		p.log.Errorf("Failed to createSwitchToT1Port: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createSwitchToT1Port: %s", switchToT1Port.ID)
	return switchToT1Port.ID, nil
}

// CreateT1ToSwitchPort creates a logical port in T1 router to connect to switch
func (p *nsxNetworkProvisioner) CreateT1ToSwitchPort(clusterName, routerID, switchPortID string,
	gatewayIPAddress strfmt.IPv4, tags []*models.Tag) (string, error) {

	if err := util.EnsureParams(clusterName,
		routerID, switchPortID, gatewayIPAddress.String()); err != nil {
		return "", err
	}

	req := &models.LogicalRouterPort{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixT1ToSwitchPort,
			Tags:        tags,
		},
		ResourceType:    nsx.PortTypeDownLinkPort,
		LogicalRouterID: util.StringPtr(routerID),
		LinkedLogicalSwitchPortID: &models.ResourceReference{
			TargetType: nsx.PortTypeLinkPortOnSwitch,
			TargetID:   switchPortID,
		},
		Subnets: []*models.IPSubnet{
			{
				IPAddresses:  []strfmt.IPv4{gatewayIPAddress},
				PrefixLength: util.Int64Ptr(nsx.ClusterSubnetPrefixLength),
			},
		},
	}
	p.log.Debugf("createT1ToSwitchPort with spec: %+v", req)

	t1ToSwitchPort, err := p.CreateLogicalRouterPort(req)
	if err != nil {
		p.log.Errorf("Failed to createT1ToSwitchPort: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully createT1ToSwitchPort: %s", t1ToSwitchPort.ID)
	return t1ToSwitchPort.ID, nil
}

// AllocateSubnetBlock allocates a subnet block from IP block
func (p *nsxNetworkProvisioner) AllocateSubnetBlock(clusterName, ipBlockID string,
	tags []*models.Tag) (string, string, error) {

	if err := util.EnsureParams(clusterName, ipBlockID); err != nil {
		return "", "", err
	}

	req := &models.IPBlockSubnet{
		ManagedResource: models.ManagedResource{
			DisplayName: clusterName + nsx.SuffixSubnetBlock,
			Tags:        tags,
		},
		Size:    util.Int64Ptr(nsx.ClusterSubnetBlockSize),
		BlockID: util.StringPtr(ipBlockID),
	}
	p.log.Debugf("allocateSubnetBlock with spec: %+v", req)


	var subnetBlock *models.IPBlockSubnet

	err := util.RetryWithLock(util.Operation(func() error {
		var err error
		subnetBlock, err = p.AllocateSubnetFromIPBlock(req)
		return err
	}), IsValidError, p.log)
	
	if err != nil {
		return "", "", err
		
	}
	
	p.log.Debugf("Successfully allocateSubnetBlock: %s", subnetBlock.ID)
	
	return subnetBlock.ID, subnetBlock.Cidr, nil
}

// AllocateSubnetFromIPBlocks allocates a subnet block from a free IP block from the given IP  blocks
func (p *nsxNetworkProvisioner) AllocateSubnetFromIPBlocks(clusterName string, ipBlockIDs []string,
	tags []*models.Tag) (string, string, string, error) {

	var (
		subnetID, subnetCidr string
		err                  error
	)

	if err := util.EnsureParams(clusterName, ipBlockIDs); err != nil {
		return "", "", "", err
	}

	for _, ipBlockID := range ipBlockIDs {
		subnetID, subnetCidr, err = p.AllocateSubnetBlock(clusterName, ipBlockID, tags)
		if err != nil {
			if conflictErr, ok := err.(*pool_management.CreateIPBlockSubnetConflict); ok {
				if conflictErr.Payload.ErrorCode == nsx.NsxIPBlockExhaustionErrorCode {
					p.log.Warnf("Failed to allocate subnet block due to IP Block exhausted:  %s", conflictErr.Error())
					continue
				}
			}
			return "", "", "", err
		}
		return subnetID, subnetCidr, ipBlockID, nil
	}

	return "", "", "", fmt.Errorf("Insufficient free range/space in IP Block to allocate subnet")
}

// AllocateIPAddress allocates an IP address from subnet block
func (p *nsxNetworkProvisioner) AllocateIPAddress(ipSubnetBlockID string) (string, error) {
	if err := util.EnsureParams(ipSubnetBlockID); err != nil {
		return "", err
	}

	ipAddress, err := p.AllocateIPFromSubnetBlock(ipSubnetBlockID)
	if err != nil {
		p.log.Errorf("Failed to allocateIPAddress: %s", ipSubnetBlockID)
		return "", err
	}
	p.log.Debugf("Successfully allocateIPAddress: %s", ipAddress.AllocationID)
	return ipAddress.AllocationID, nil
}

// BuildIPAddress is a small helper function to statically
// create IP address based on specified CIDR
func (p *nsxNetworkProvisioner) BuildIPAddress(cidr string, IPPart string) string {
	if _, _, err := net.ParseCIDR(cidr); err != nil || IPPart == "" {
		return ""
	}
	nums := strings.Split(cidr, ".")
	return strings.Join(nums[:3], ".") + "." + IPPart
}

// DeleteClusterSwitch cleans up all resources related to logical switches for a kubernetes cluster
func (p *nsxNetworkProvisioner) DeleteClusterSwitch(LogicalSwitchID string) error {
	if err := util.EnsureParams(LogicalSwitchID); err != nil {
		return err
	}
	return p.DeleteLogicalSwitch(LogicalSwitchID)
}

// GetLogicalSwitchByTag returns the floating ip associated with the switch
func (p *nsxNetworkProvisioner) GetLogicalSwitchByTag(lsScope, lsTag string) (*models.LogicalSwitch, error) {
	lss, err := p.ListLogicalSwitches()
	if err != nil {
		return nil, err
	}
	for _, ls := range lss.Results {
		for _, tag := range ls.Tags {
			if tag.Scope == lsScope && tag.Tag == lsTag {
				return ls, nil
			}
		}
	}
	return nil, nil
}

// GetSwitchingProfileByTag returns the switching profile by tag
func (p *nsxNetworkProvisioner) GetSwitchingProfileByTag(spType, spScope, spTag string) (*models.BaseSwitchingProfile, error) {
	switchingProfiles, err := p.ListSwitchingProfilesByType(spType)
	if err != nil {
		return nil, err
	}

	for _, sp := range switchingProfiles.Results {
		for _, tag := range sp.Tags {
			if tag.Scope == spScope && tag.Tag == spTag {
				return sp, nil
			}
		}
	}
	return nil, nil
}

// CreateSpoofGuardSwitchingProfile creates a spoofguard switching profile
func (p *nsxNetworkProvisioner) CreateSpoofGuardSwitchingProfile(name, description string, tags []*models.Tag, whiteListProviders []string) (string, error) {
	if err := util.EnsureParams(name, description, tags); err != nil {
		return "", err
	}

	req := &models.SpoofGuardSwitchingProfile{
		ManagedResource: models.ManagedResource{
			DisplayName: name,
			Description: description,
			Tags:        tags,
		},
		ResourceType:       util.StringPtr(models.SwitchingProfileTypeIDEntryKeySpoofGuardSwitchingProfile),
		WhiteListProviders: whiteListProviders,
	}

	p.log.Debugf("CreateSpoofGuardSwitchingProfile with spec: %+v", req)

	switchingProfile, err := p.Client.CreateSpoofGuardSwitchingProfile(req)
	if err != nil {
		p.log.Errorf("Failed to create spoofguard switching profile: %+v", req)
		return "", err
	}
	p.log.Debugf("Successfully created spoof guard switching profile: %s", switchingProfile.ID)
	return switchingProfile.ID, nil
}

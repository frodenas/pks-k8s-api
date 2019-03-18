/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"fmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/pool_management"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
	"net"
)

// UntagIPBlock removes the cluster tag from the IP block
func (p *nsxNetworkProvisioner) UntagIPBlock(blockID string, clusterName string) error {
	var err error
	var IPBlock *models.IPBlock

	if err = util.EnsureParams(blockID, clusterName); err != nil {
		return err
	}
	if IPBlock, err = p.ReadIPBlock(blockID); err != nil {
		return err
	}
	nsx.RemoveTag(&IPBlock.ManagedResource, models.Tag{
		Scope: nsx.NcpTagKeyCluster,
		Tag:   clusterName,
	})
	if _, err = p.UpdateIPBlock(IPBlock); err != nil {
		return err
	}
	return nil
}

// UntagExternalBlocks removes the cluster tag from any IP blocks marked with
// an 'ncp/external':'true' tag
func (p *nsxNetworkProvisioner) UntagExternalBlocks(clusterName string) error {
	var searchResults *models.SearchResults
	var err error
	if err = util.EnsureParams(clusterName); err != nil {
		return err
	}

	clusterTag := models.Tag{
		Scope: nsx.NcpTagKeyCluster,
		Tag:   clusterName,
	}
	if searchResults, err = p.SearchByTag(util.StringPtr(nsx.SearchTypeIPBlock), clusterTag); err != nil {
		return err
	}
	for _, block := range searchResults.Results {
		if nsx.EvaluateTag(block, nsx.NcpTagKeyExternal) != nil {
			var extBlock *models.IPBlock
			if extBlock, err = p.ReadIPBlock(block.ID); err != nil {
				return err
			}
			nsx.RemoveTag(&extBlock.ManagedResource, clusterTag)
			if _, err = p.UpdateIPBlock(extBlock); err != nil {
				return err
			}
		}
	}
	return nil
}

// UntagExternalPools removes the cluster tag from any IP pools marked with
// an 'ncp/external':'true' tag
func (p *nsxNetworkProvisioner) UntagExternalPools(clusterName string) error {
	var searchResults *models.SearchResults
	var err error
	if err = util.EnsureParams(clusterName); err != nil {
		return err
	}

	clusterTag := models.Tag{
		Scope: nsx.NcpTagKeyCluster,
		Tag:   clusterName,
	}
	if searchResults, err = p.SearchByTag(util.StringPtr(nsx.SearchTypeIPPool), clusterTag); err != nil {
		return err
	}
	for _, pool := range searchResults.Results {
		if nsx.EvaluateTag(pool, nsx.NcpTagKeyExternal) != nil {
			var extPool *models.IPPool
			if extPool, err = p.ReadIPPool(pool.ID); err != nil {
				return err
			}
			nsx.RemoveTag(&extPool.ManagedResource, clusterTag)
			if _, err = p.UpdateIPPool(extPool); err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteIPBlockSubnet deletes the specified IP block subnet
func (p *nsxNetworkProvisioner) DeleteIPBlockSubnet(subnetID string) error {
	if err := util.EnsureParams(subnetID); err != nil {
		return err
	}

	return util.RetryWithLock(util.Operation(func() error {
		return p.Client.DeleteIPBlockSubnet(subnetID)
	}), IsValidError, p.log)
}

// CheckIPPool checks if the given IP pool ID is valid
func (p *nsxNetworkProvisioner) CheckIPPool(poolID string) error {
	if err := util.EnsureParams(poolID); err != nil {
		return err
	}
	_, err := p.ReadIPPool(poolID)
	return err
}

//CheckIPBlock checks if the given IP Block ID is valid
func (p *nsxNetworkProvisioner) CheckIPBlock(blockID string) error {
	if err := util.EnsureParams(blockID); err != nil {
		return err
	}
	_, err := p.ReadIPBlock(blockID)
	return err
}

//CheckIPBlockSubnetPrefix checks given prefix is the same with prefix of subnets in given IP Block ID
func (p *nsxNetworkProvisioner) CheckIPBlockSubnetPrefix(blockID string, prefix int) error {
	ipBlockSubnets, err := p.ListIPBlockSubnets(blockID)
	if err != nil {
		return err
	}

	for _, subnet := range ipBlockSubnets.Results {
		_, ipNet, err := net.ParseCIDR(subnet.Cidr)
		if err != nil {
			return err
		}
		ones, _ := ipNet.Mask.Size()
		if prefix != ones {
			return fmt.Errorf("prefix %d conflicts with existing subnet prefix %d", prefix, ones)
		}
	}
	return nil
}

// GetIPBlockSubnet retrieves block subnet given IP block ID
func (p *nsxNetworkProvisioner) GetIPBlockSubnet(clusterName, blockID string) (string, string, error) {
	if err := util.EnsureParams(clusterName, blockID); err != nil {
		return "", "", err
	}
	ipBlockSubnets, err := p.ListIPBlockSubnets(blockID)
	if err != nil {
		return "", "", err
	}

	for _, subnet := range ipBlockSubnets.Results {
		if subnet.DisplayName == clusterName+nsx.SuffixSubnetBlock {
			return subnet.ID, subnet.Cidr, nil
		}
	}
	p.log.Warnf("IP block subnet for cluster %s not found", clusterName)

	return "", "", nil
}

// AllocateFloatingIPAddress allocates an IP address from IP pool
func (p *nsxNetworkProvisioner) AllocateFloatingIPAddress(floatingIPPoolID string) (string, error) {
	if err := util.EnsureParams(floatingIPPoolID); err != nil {
		return "", err
	}

	p.log.Debugf("Allocating floating IP from pool %s", floatingIPPoolID)

	var allocatedFloatingIP *models.AllocationIPAddress

	err := util.RetryWithLock(util.Operation(func() error {
		var err error
		allocatedFloatingIP, err = p.AllocateIPFromIPPool(floatingIPPoolID)
		return err
	}), IsValidError, p.log)

	if err != nil {
		return "", err
	}

	p.log.Debugf("Successfully allocated floating IP %s from pool %s", allocatedFloatingIP.AllocationID, floatingIPPoolID)

	return allocatedFloatingIP.AllocationID, nil
}

// AllocateFloatingIPAddressFromIPPools allocates an IP address from a free IP pool from given IP pools
func (p *nsxNetworkProvisioner) AllocateFloatingIPAddressFromIPPools(floatingIPPoolIDs []string) (string, string, error) {
	if err := util.EnsureParams(floatingIPPoolIDs); err != nil {
		return "", "", err
	}

	for _, floatingIPPoolID := range floatingIPPoolIDs {
		allocatedFloatingIP, err := p.AllocateFloatingIPAddress(floatingIPPoolID)
		if err != nil {
			if conflictErr, ok := err.(*pool_management.AllocateOrReleaseFromIPPoolConflict); ok {
				if conflictErr.Payload.ErrorCode == nsx.NsxIPPoolExhaustionErrorCode {
					p.log.Warnf("Failed to allocate floating IP due to IP Pool exhausted:  %s", conflictErr.Error())
					continue
				}
			}
			return "", "", err
		}
		return allocatedFloatingIP, floatingIPPoolID, nil
	}

	return "", "", fmt.Errorf("Insufficient free IP's in IP pool to allocate IP")
}

// ReleaseFloatingIPAddress releases an IP address from in IP pool
func (p *nsxNetworkProvisioner) ReleaseFloatingIPAddress(floatingIPPoolID, floatingIPAddress string) error {
	if err := util.EnsureParams(floatingIPPoolID, floatingIPAddress); err != nil {
		return err
	}

	req := &models.AllocationIPAddress{
		AllocationID: floatingIPAddress,
	}
	p.log.Debugf("Release floating IP %s from pool %s", floatingIPAddress, floatingIPPoolID)

	err := util.RetryWithLock(util.Operation(func() error {
		return p.ReleaseIPToIPPool(floatingIPPoolID, req)
	}), IsValidError, p.log)

	if err != nil {
		return err
	}

	p.log.Debugf("Successfully released floating IP %s from pool %s", floatingIPAddress, floatingIPPoolID)

	return err

}

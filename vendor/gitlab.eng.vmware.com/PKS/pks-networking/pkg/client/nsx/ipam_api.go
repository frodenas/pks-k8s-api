/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	pm "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/pool_management"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/util"
)

// ReadIPBlock returns an IPBlock with the specified ID from the NSX Manager
func (nc *client) ReadIPBlock(blockID string) (*models.IPBlock, error) {
	params := pm.NewReadIPBlockParams().WithBlockID(blockID)

	res, err := nc.client.PoolManagement.ReadIPBlock(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// UpdateIPBlock updates an IPBlock's information to that of the
// parameter IPBlock
func (nc *client) UpdateIPBlock(IPBlock *models.IPBlock) (*models.IPBlock, error) {
	params := pm.NewUpdateIPBlockParams().WithIPBlock(IPBlock).WithBlockID(IPBlock.ID)

	res, err := nc.client.PoolManagement.UpdateIPBlock(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReadIPBlockSubnet reads info of a subnet block
func (nc *client) ReadIPBlockSubnet(subnetID string) (*models.IPBlockSubnet, error) {
	params := pm.NewReadIPBlockSubnetParams().WithSubnetID(subnetID)
	res, err := nc.client.PoolManagement.ReadIPBlockSubnet(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListIPBlockSubnets returns a slice of subnets corresponding to a specified IP Block.
func (nc *client) ListIPBlockSubnets(BlockID string) (*models.IPBlockSubnetListResult, error) {
	params := pm.NewListIPBlockSubnetsParams().WithBlockID(util.StringPtr(BlockID))
	res, err := nc.client.PoolManagement.ListIPBlockSubnets(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReadIPPool returns an IP pool with the specified ID
func (nc *client) ReadIPPool(poolID string) (*models.IPPool, error) {
	params := pm.NewReadIPPoolParams().WithPoolID(poolID)

	res, err := nc.client.PoolManagement.ReadIPPool(params, nc.auth)

	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// UpdateIPPool updates an IP pool
func (nc *client) UpdateIPPool(pool *models.IPPool) (*models.IPPool, error) {
	params := pm.NewUpdateIPPoolParams().WithIPPool(pool).WithPoolID(pool.ID)

	res, err := nc.client.PoolManagement.UpdateIPPool(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteIPBlockSubnet deletes the specified IP Block subnet.
// There must not be any allocated IP addresses from this block or
// the server will not perform the deletion.
func (nc *client) DeleteIPBlockSubnet(subnetID string) error {
	params := pm.NewDeleteIPBlockSubnetParams().WithSubnetID(subnetID)

	_, err := nc.client.PoolManagement.DeleteIPBlockSubnet(params, nc.auth)

	return err
}

// AllocateSubnetFromIPBlock creates a subnet block from given IP block
func (nc *client) AllocateSubnetFromIPBlock(ipBlockSubnetModel *models.IPBlockSubnet) (*models.IPBlockSubnet, error) {
	params := pm.NewCreateIPBlockSubnetParams().WithIPBlockSubnet(ipBlockSubnetModel)
	res, err := nc.client.PoolManagement.CreateIPBlockSubnet(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// AllocateIPFromSubnetBlock allocates an IP from given subnet block
func (nc *client) AllocateIPFromSubnetBlock(subnetID string) (*models.AllocationIPAddress, error) {
	params := pm.NewAllocateOrReleaseFromIPBlockSubnetParams().WithSubnetID(subnetID).
		WithAction(IPAddressActionAllocate)
	res, err := nc.client.PoolManagement.AllocateOrReleaseFromIPBlockSubnet(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (nc *client) ListIPPools() (*models.IPPoolListResult, error) {
	params := pm.NewListIPPoolsParams()
	res, err := nc.client.PoolManagement.ListIPPools(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ListIPPoolAllocations lists ip allocations from ip pool
func (nc *client) ListIPPoolAllocations(ipPoolID string) (*models.AllocationIPAddressListResult, error) {
	params := pm.NewListIPPoolAllocationsParams().WithPoolID(ipPoolID)
	res, err := nc.client.PoolManagement.ListIPPoolAllocations(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// AllocateIPFromIPPool allocates an IP from given IP pool
func (nc *client) AllocateIPFromIPPool(ipPoolID string) (*models.AllocationIPAddress, error) {
	params := pm.NewAllocateOrReleaseFromIPPoolParams().WithPoolID(ipPoolID).
		WithAction(IPAddressActionAllocate)
	res, err := nc.client.PoolManagement.AllocateOrReleaseFromIPPool(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReleaseIPToIPPool releases an IP in IP pool
func (nc *client) ReleaseIPToIPPool(ipPoolID string, allocationIPAddressModel *models.AllocationIPAddress) error {
	params := pm.NewAllocateOrReleaseFromIPPoolParams().WithPoolID(ipPoolID).
		WithAllocationIPAddress(allocationIPAddressModel).WithAction(IPAddressActionRelease)
	_, err := nc.client.PoolManagement.AllocateOrReleaseFromIPPool(params, nc.auth)
	if err != nil {
		return err
	}
	return nil
}

func (nc *client) DeleteIPPool(ipPoolID string, force bool) error {
	params := pm.NewDeleteIPPoolParams().WithPoolID(ipPoolID).WithForce(&force)
	_, err := nc.client.PoolManagement.DeleteIPPool(params, nc.auth)
	return err
}

// TagIPBlock adds tag to an IP block component
func (nc *client) TagIPBlock(ipBlockID string, tags []*models.Tag) (*models.IPBlock, error) {
	ipBlock, err := nc.ReadIPBlock(ipBlockID)
	if err != nil {
		return nil, err
	}
	err = ValidateTags(ipBlock.ManagedResource, tags)
	if err != nil {
		return nil, err
	}

	ipBlock.Tags = append(ipBlock.Tags, tags...)
	res, err := nc.UpdateIPBlock(ipBlock)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TagIPPool adds tag to an IP block component
func (nc *client) TagIPPool(ipPoolID string, tags []*models.Tag) (*models.IPPool, error) {
	ipPool, err := nc.ReadIPPool(ipPoolID)
	if err != nil {
		return nil, err
	}
	err = ValidateTags(ipPool.ManagedResource, tags)
	if err != nil {
		return nil, err
	}

	ipPool.Tags = append(ipPool.Tags, tags...)
	res, err := nc.UpdateIPPool(ipPool)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// release allocated ip addresses and delete ip pool
func (nc *client) CleanupIPPool(ipPoolID string, readOnly bool) error {
	allocations, err := nc.ListIPPoolAllocations(ipPoolID)
	if err != nil {
		return err
	}
	for _, allocation := range allocations.Results {
		nc.VerboseInfo("allocated ip %s to be removed\n", allocation.AllocationID)
		if !readOnly {
			req := &models.AllocationIPAddress{}
			req.AllocationID = allocation.AllocationID
			err = nc.ReleaseIPToIPPool(ipPoolID, req)
			if err != nil {
				nc.Debug("allocated ip %s cannot be removed due to %s\n", allocation.AllocationID, err)
				/* Continue here since release of IP's from IP Pool is a best effort operation.
				 * We are anyways force deleting the IP Pool
				 */
				continue
			}
			nc.VerboseInfo("allocated ip %s is removed successfully\n", allocation.AllocationID)
		}
	}
	nc.VerboseInfo("IP Pool %s to be removed\n", ipPoolID)
	if !readOnly {
		err = nc.DeleteIPPool(ipPoolID, true)
		if err != nil {
			return err
		}
		nc.VerboseInfo("IP Pool %s is removed successfully\n", ipPoolID)
	}
	return nil
}

/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"fmt"
	"strings"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	fab "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/fabric"
)

// ListFabricNodes  lists all the fabric nodes in a NSX
func (nc *client) ListFabricNodes() (*models.NodeListResult, error) {
	params := fab.NewListNodesParams()
	res, err := nc.client.Fabric.ListNodes(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReadFabricNode gets fabric node given a node id
func (nc *client) ReadFabricNode(nodeID string) (*models.Node, error) {
	params := fab.NewReadNodeParams().WithNodeID(nodeID)
	res, err := nc.client.Fabric.ReadNode(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetFabricNodeGivenParams gets the fabric node matching the parameters
func (nc *client) GetFabricNodeGivenParams(params *fab.ListNodesParams) (*models.Node, error) {
	res, err := nc.client.Fabric.ListNodes(params, nc.auth)
	if err != nil {
		return nil, err
	}
	if res.Payload.ResultCount > 0 {
		return res.Payload.Results[0], err
	}
	return nil, err
}

// GetFabricNodeMatchingIP gets the fabric node matching atleast one IP in the input list
func (nc *client) GetFabricNodeMatchingIP(NodeIPs []string) (*models.Node, error) {
	for _, nodeip := range NodeIPs {
		params := fab.NewListNodesParams().WithIPAddress(&nodeip)
		res, err := nc.GetFabricNodeGivenParams(params)
		if err == nil && res != nil {
			return res, err
		}
	}

	return nil, fmt.Errorf("Unable to find Fabric Node matching ip's %s",
		strings.Join(NodeIPs, ","))
}

// GetFabricNodeState gets fabric node's state
func (nc *client) GetFabricNodeState(nodeID string) (*models.ConfigurationState, error) {
	params := fab.NewGetFabricNodeStateParams().WithNodeID(nodeID)
	res, err := nc.client.Fabric.GetFabricNodeState(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// ReadFabricNodeStatus gets fabric node's status
func (nc *client) ReadFabricNodeStatus(nodeID string) (*models.NodeStatus, error) {
	params := fab.NewReadNodeStatusParams().WithNodeID(nodeID)
	res, err := nc.client.Fabric.ReadNodeStatus(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

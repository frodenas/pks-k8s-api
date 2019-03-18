/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
)

//FabricNode represents a fabric node
type FabricNode struct {
	MgmtIPs []string
}

//CheckFabricNodes checks if the given fabric nodes are valid and also check its state and status
func (p *nsxNetworkProvisioner) CheckFabricNodes(fabricNodes []*FabricNode) error {
	if err := util.EnsureParams(fabricNodes); err != nil {
		return err
	}
	for _, fabricNode := range fabricNodes {
		fabricNodeIPs := fabricNode.MgmtIPs

		p.log.Debugf("****Checking if Fabric Node %s is valid****", fabricNodeIPs)

		fabricnodeErr := p.CheckFabricNodeGivenMgmtIP(fabricNodeIPs)
		if fabricnodeErr != nil {
			p.log.Errorf("Fabric Node is invalid due to error %s", fabricnodeErr)
			return fabricnodeErr
		}

		p.log.Debugf("Fabric Node %s is valid", fabricNodeIPs)
	}
	return nil
}

//CheckFabricNodeGivenMgmtIP checks if the given fabric node is valid and also check its state and status
func (p *nsxNetworkProvisioner) CheckFabricNodeGivenMgmtIP(mgmtIPs []string) error {
	fabricnodeRes, fabricnodeErr := p.GetFabricNodeGivenMgmtIP(mgmtIPs)
	if fabricnodeErr != nil {
		return fabricnodeErr
	}

	fabricnodeStateErr := p.CheckFabricNodeState(fabricnodeRes.ID)
	if fabricnodeStateErr != nil {
		return fabricnodeStateErr
	}

	fabricnodeStatusErr := p.CheckFabricNodeStatus(fabricnodeRes.ID)
	if fabricnodeStatusErr != nil {
		return fabricnodeStatusErr
	}

	return nil
}

//GetFabricNodeGivenMgmtIP returns the fabric Node matching the management IP
func (p *nsxNetworkProvisioner) GetFabricNodeGivenMgmtIP(mgmtIPs []string) (*models.Node, error) {
	if err := util.EnsureParams(mgmtIPs); err != nil {
		return nil, err
	}
	fabricnodeRes, fabricnodeErr := p.GetFabricNodeMatchingIP(mgmtIPs)
	if fabricnodeErr != nil {
		return nil, fabricnodeErr
	}
	return fabricnodeRes, nil
}

//CheckFabricNodeState checks if the given fabric node's state is in ConfigurationStateSuccess
func (p *nsxNetworkProvisioner) CheckFabricNodeState(nodeID string) error {
	if err := util.EnsureParams(nodeID); err != nil {
		return err
	}
	nodestateRes, nodestateErr := p.GetFabricNodeState(nodeID)
	if nodestateErr != nil {
		return nodestateErr
	}

	if nodestateRes.State != models.ConfigurationStateStateSuccess {
		return fmt.Errorf("Fabric Node %s state %s is not in success",
			nodeID, nodestateRes.State)
	}

	return nil
}

//CheckFabricNodeStatus checks if the given fabric node's host node deployment status is Install successful
func (p *nsxNetworkProvisioner) CheckFabricNodeStatus(nodeID string) error {
	if err := util.EnsureParams(nodeID); err != nil {
		return err
	}
	nodestatusRes, nodestatusErr := p.ReadFabricNodeStatus(nodeID)
	if nodestatusErr != nil {
		return nodestatusErr
	}

	if nodestatusRes.HostNodeDeploymentStatus !=
		models.NodeStatusHostNodeDeploymentStatusINSTALLSUCCESSFUL {
		return fmt.Errorf("Fabric Node %s install status %s is not in success",
			nodeID, nodestatusRes.HostNodeDeploymentStatus)
	}
	return nil
}

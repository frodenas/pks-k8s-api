/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

import (
	"errors"

	np "gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/netprovisioner"
)

// nsxNetworkCreatePreCheck is a common function signature for all NSX objects create precheckers
type nsxNetworkCreatePreCheck func(clusterSpec *NSXTClusterSpec) error

// PreCheckCreateNetwork prechecks the network before instance creation
func (n *networkManager) PreCheckCreateNetwork(instanceID string, clusterSpec *NSXTClusterSpec) error {
	n.log.Debugln("PRECHECK CREATE CLUSTER NETWORK")

	err := n.validateNSXTSpec(clusterSpec)
	if err != nil {
		return err
	}

	err = n.validateClusterName(instanceID)
	if err != nil {
		return err
	}

	prechecks, prechecksErr := n.getCreateClusterNetworkPreCheckers()
	if prechecksErr != nil {
		return prechecksErr
	}

	for _, precheck := range prechecks {
		err := precheck(clusterSpec)
		if err != nil {
			return err
		}
	}

	return nil
}

func (n *networkManager) validateVcenterClusterSpec(vcspec *VcenterSpec) error {
	if n.vcp == nil {
		return errors.New("Vcenter Client not provided")
	}

	if vcspec == nil {
		return errors.New("Vcenter Cluster Spec not provided")
	}

	if len(vcspec.ComputeClusterPath) <= 0 {
		return errors.New("Vcenter Spec Compute Cluster Path not provided")
	}

	if len(vcspec.DataCenterPath) <= 0 {
		return errors.New("Vcenter Spec Data Center Path not provided")
	}

	return nil
}

func (n *networkManager) validateClusterName(clusterName string) error {
	if len(clusterName) <= 0 {
		return errors.New("Cluster Name not provided")
	}

	if len(clusterName) > 40 {
		return errors.New("Cluster Name exceeds 40 character limit")
	}

	return nil
}

func (n *networkManager) validateNSXTSpec(clusterSpec *NSXTClusterSpec) error {
	if n.nsxtSpec == nil {
		return errors.New("NSX Cluster Network Spec not provided")
	}

	if len(clusterSpec.T0RouterID) <= 0 {
		return errors.New("T0 Router ID not provided")
	}

	if len(clusterSpec.LbFloatingIPPoolIDs) <= 0 {
		return errors.New("Floating IP Pool ID not provided")
	}

	if len(clusterSpec.IPBlockIDs) <= 0 {
		return errors.New("IP Block ID not provided")
	}

	if len(clusterSpec.OverlayTransportZoneID) <= 0 {
		return errors.New("Overlay Transport Zone ID not provided")
	}

	if len(clusterSpec.EdgeClusterID) <= 0 {
		return errors.New("Edge Cluster ID not provided")
	}

	vcspecErr := n.validateVcenterClusterSpec(n.vcSpec)
	if vcspecErr != nil {
		return vcspecErr
	}

	return nil
}

func (n *networkManager) getFabricNodeIps(clusterSpec *NSXTClusterSpec) ([]*np.FabricNode, error) {
	var fabricNodes []*np.FabricNode
	esxiHosts, err := n.vcp.GetEsxiHostIPs(n.vcSpec.DataCenterPath, n.vcSpec.ComputeClusterPath)
	if err != nil {
		return nil, err
	}

	for _, esxiHost := range esxiHosts {
		fabricNode := &np.FabricNode{MgmtIPs: esxiHost.MgmtIPs}
		fabricNodes = append(fabricNodes, fabricNode)
	}

	return fabricNodes, err
}

func (n *networkManager) getCreateClusterNetworkPreCheckers() ([]nsxNetworkCreatePreCheck, error) {
	var prechecks []nsxNetworkCreatePreCheck

	prechecks = append(prechecks,
		n.preCheckCreateClusterStatus, n.preCheckCreateOverlayTransportZone,
		n.preCheckCreateT0Router, n.preCheckCreateIPPool,
		n.preCheckCreateIPBlock, n.preCheckCreateFabricNodes)

	return prechecks, nil
}

func (n *networkManager) preCheckCreateClusterStatus(clusterSpec *NSXTClusterSpec) error {
	n.log.Debugf("****Checking if NSX Cluster is stable****")

	err := n.np.CheckClusterStatus()
	if err != nil {
		n.log.Errorf("NSX Cluster stability check failed due to error %s", err)
		return err
	}

	n.log.Debugf("NSX Cluster is stable")
	return nil
}

func (n *networkManager) preCheckCreateOverlayTransportZone(clusterSpec *NSXTClusterSpec) error {
	overlayTransportZoneID := clusterSpec.OverlayTransportZoneID
	n.log.Debugf("****Checking if Overlay Transport Zone %s is valid****", overlayTransportZoneID)

	err := n.np.CheckTransportZone(overlayTransportZoneID)
	if err != nil {
		n.log.Errorf("Overlay Transport Zone %s is invalid due to error %s",
			overlayTransportZoneID, err)
		return err
	}

	n.log.Debugf("Overlay Transport Zone %s is valid", overlayTransportZoneID)
	return nil
}

func (n *networkManager) preCheckCreateT0Router(clusterSpec *NSXTClusterSpec) error {
	t0RouterID := clusterSpec.T0RouterID
	n.log.Debugf("****Checking if T0 Router %s is valid****", t0RouterID)

	err := n.np.CheckT0Router(t0RouterID)
	if err != nil {
		n.log.Errorf("T0 Router is invalid %s due to error %s", t0RouterID, err)
		return err
	}

	n.log.Debugf("T0 Router %s is valid", t0RouterID)
	return nil
}

func (n *networkManager) preCheckCreateIPPool(clusterSpec *NSXTClusterSpec) error {
	ipPoolID := clusterSpec.LbFloatingIPPoolIDs[0]
	n.log.Debugf("****Checking if Floating IP Pool %s is valid****", ipPoolID)

	err := n.np.CheckIPPool(ipPoolID)
	if err != nil {
		n.log.Errorf("Floating IP Pool %s is invalid due to error %s", ipPoolID, err)
		return err
	}

	n.log.Debugf("Floating IP Pool %s is valid", ipPoolID)
	return nil
}

func (n *networkManager) preCheckCreateIPBlock(clusterSpec *NSXTClusterSpec) error {
	ipblockID := clusterSpec.IPBlockIDs[0]
	n.log.Debugf("****Checking if IP Block %s is valid****", ipblockID)

	err := n.np.CheckIPBlock(ipblockID)
	if err != nil {
		n.log.Errorf("IP Block %s is invalid due to error %s", ipblockID, err)
		return err
	}

	n.log.Debugf("IP Block %s is valid", ipblockID)
	return nil
}

func (n *networkManager) preCheckCreateFabricNodes(clusterSpec *NSXTClusterSpec) error {
	fabricNodes, fabricNodesErr := n.getFabricNodeIps(clusterSpec)
	if fabricNodesErr != nil {
		n.log.Errorf("Failed to get Fabric Nodes due to error %s", fabricNodesErr)
		return fabricNodesErr
	}

	n.log.Debugf("****Checking if Fabric Nodes are valid****")

	err := n.np.CheckFabricNodes(fabricNodes)
	if err != nil {
		n.log.Errorf("Fabric Node check failed due teo error %s", err)
		return err
	}

	n.log.Debugf("Fabric Nodes are valid")
	return nil
}

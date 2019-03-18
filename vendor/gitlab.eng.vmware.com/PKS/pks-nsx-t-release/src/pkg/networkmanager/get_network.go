/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/workflow"
)

// GetResources returns the network resources of an instance
func (n *networkManager) GetResources(instanceID string) (CollectResourcesResp, error) {
	resp := &CollectResourcesResp{}
	collectResources := n.NewCollectResources(instanceID, resp, "CollectResources")

	err := collectResources.Run()
	return *resp, err
}

func (n *networkManager) GetNetwork(instanceID string) (NetworkInfo, error) {
	resp, err := n.GetResources(instanceID)

	if err != nil {
		return NetworkInfo{
			Status: NetworkUnknown, // in face of error, we're not sure if this network exist
		}, err
	}

	if resp.Num == 0 {
		return NetworkInfo{
			Status: NetworkNotFound,
		}, nil
	}

	return NetworkInfo{
		Cidr:                 resp.Cidr,
		Gateway:              resp.GatewayIPAddress,
		ExternalIP:           resp.MasterExternalIPAddress,
		LbServiceID:          resp.LbServiceID,
		LbSize:               resp.LbSize,
		SwitchName:           resp.LogicalSwitchName,
		T0RouterID:           resp.T0RouterID,
		MasterVMsNSGroupName: resp.MasterVMsNSGroupName,
		Status:               NetworkCreated, // if there is any resource associated with this network, we assert that this network is created before
	}, nil
}

type CollectResourcesResp struct {
	NatMode bool

	// number of resources
	Num int

	// logical router related info
	T0ToT1PortID     string
	T1RouterID       string
	T1ToSwitchPortID string
	T1ToT0PortID     string
	SnatRuleID       string
	T0RouterID       string

	// logical switch related info
	LogicalSwitchID      string
	LogicalSwitchName    string
	SwitchToT1PortID     string
	MasterVMsNSGroupName string

	// IPAM info
	IPBlockSubnetID         string
	Cidr                    string
	GatewayIPAddress        strfmt.IPv4
	MasterExternalIPAddress strfmt.IPv4
	SnatFloatingIPPoolID    string
	LBFloatingIPPoolID      string
	IPBlockID               string

	// Loadbalancer related info
	// optional
	LbServiceID       string
	LbSize            string
	LbVirtualServerID string
	LbPoolID          string
	LbNSGroupID       string
	LbTcpMonitorID    string

	// LoadBalancer logical router related info
	// optional
	LbT0ToT1PortID     string
	LbT1RouterID       string
	LbT1ToSwitchPortID string
	LbT1ToT0PortID     string
	// there is no nat rule on lb related logical router

	// LoadBalancer logical switch related info
	// optional
	LbLogicalSwitchID  string
	LbSwitchToT1PortID string
}

// NewCollectResources return a Workflow
func (n *networkManager) NewCollectResources(instanceID string, resp *CollectResourcesResp,
	logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.collectResources(instanceID, resp, n.log.WithField(projName, logField))
	})
}

// collectResources performs NSX resources collection from instance specifig tag
func (n *networkManager) collectResources(instanceID string, resp *CollectResourcesResp, log logrus.FieldLogger) error {

	log.Debugf("Collecting cluster %s resources\n", instanceID)

	resources, err := n.np.GetAllResources(
		instanceID,
		models.Tag{
			Scope: nsx.PksTagKeyCluster,
			Tag:   instanceID,
		},
	)
	if err != nil {
		log.Errorf("Failed collecting cluster %s resources\n", instanceID)
		return err
	}

	//Cluster Parameters like IPBlockID, T0RouterID, FloatingIPPoolID are retrieved
	err = n.getClusterParamsFromResources(resources, resp, log)
	if err != nil {
		log.Errorf("Failing collecting cluster %s parameters\n", instanceID)
		return err
	}

	for _, resource := range resources {
		log.Debugf("Resource for cluster %s found: %+v", instanceID, resource)
		// by default, count 1 resource
		resp.Num++
		switch resource.ResourceType {
		// loadbalancer
		case nsx.ResourceTypeLbVirtualServer:
			resp.LbVirtualServerID = resource.ID
		case nsx.ResourceTypeLbPool:
			resp.LbPoolID = resource.ID
		case nsx.ResourceTypeLbService:
			resp.LbServiceID = resource.ID
		case nsx.ResourceTypeNSGroup:
			resp.LbNSGroupID = resource.ID
		case nsx.ResourceTypeLbTcpMonitor:
			resp.LbTcpMonitorID = resource.ID
			// logical router
		case nsx.ResourceTypeLogicalRouter:
			if isLbResource(resource.DisplayName) {
				resp.LbT1RouterID = resource.ID
			} else {
				resp.T1RouterID = resource.ID
			}
		case nsx.ResourceTypeLogicalRouterLinkPortOnTIER0:
			if isLbResource(resource.DisplayName) {
				resp.LbT0ToT1PortID = resource.ID
			} else {
				resp.T0ToT1PortID = resource.ID
			}
		case nsx.ResourceTypeLogicalRouterDownlinkPort:
			if isLbResource(resource.DisplayName) {
				resp.LbT1ToSwitchPortID = resource.ID
			} else {
				resp.T1ToSwitchPortID = resource.ID
			}
		case nsx.ResourceTypeLogicalRouterLinkPortOnTIER1:
			if isLbResource(resource.DisplayName) {
				resp.LbT1ToT0PortID = resource.ID
			} else {
				resp.T1ToT0PortID = resource.ID
			}
			// logical switch
		case nsx.ResourceTypeLogicalSwitch:
			if isLbResource(resource.DisplayName) {
				resp.LbLogicalSwitchID = resource.ID
			} else {
				resp.LogicalSwitchID = resource.ID
				resp.LogicalSwitchName = resource.DisplayName
			}
		case nsx.ResourceTypeLogicalPort:
			if isLbResource(resource.DisplayName) {
				resp.LbSwitchToT1PortID = resource.ID
			} else {
				resp.SwitchToT1PortID = resource.ID
			}
		case nsx.ResourceTypeNatRule:
			//only process SNAT rule here
			natRule, err := n.np.GetNatRule(resp.T0RouterID, resource.ID)
			if err != nil {
				// don't return here
				log.Errorf("Failed to get NAT Rule %s details,  %s\n", resource.ID, err)
			}
			//only expects one SNAT rule here per cluster
			if natRule != nil && *natRule.Action == "SNAT" && resp.SnatRuleID == "" {
				resp.SnatRuleID = resource.ID
				snatFloatingIPPoolID := util.ExtractMetadataFromTags(nsx.PksTagKeySnatFloatingIpPool, natRule.Tags)
				if snatFloatingIPPoolID == "" {
					snatFloatingIPPoolID = n.nsxtSpec.FloatingIPPoolID
				}
				resp.SnatFloatingIPPoolID = snatFloatingIPPoolID
				log.Debugf("Found SNAT rule for cluster %s\n", instanceID)
			} else {
				// others don't count
				resp.Num--
			}
		default:
			// any resource that's not listed in above cases is not considered as part of `network`
			resp.Num--
			log.Debugf("Found unrecognized resources %s belong to cluster %s\n", resource.ResourceType, instanceID)
		}
	}

	if resp.Num == 0 {
		log.Warnf("There is no resource for cluster %s\n", instanceID)
		return nil
	}

	// IP subnet block is not tagged, need to find by name instead
	if resp.IPBlockID != "" {
		resp.IPBlockSubnetID, resp.Cidr, err = n.np.GetIPBlockSubnet(GetPKSResourceName(instanceID), resp.IPBlockID)
		if err != nil {
			return err
		}
		if resp.Cidr != "" {
			resp.GatewayIPAddress = strfmt.IPv4(n.np.BuildIPAddress(resp.Cidr, strconv.Itoa(nsx.ClusterIPPartT1Router)))
		}
	}

	if resp.LogicalSwitchID != "" {
		masterExternalIPAddress, natMode, err := n.np.GetMetadataFromSwitchTag(resp.LogicalSwitchID)
		if err != nil {
			return err
		}
		resp.MasterExternalIPAddress = strfmt.IPv4(masterExternalIPAddress)
		resp.NatMode = natMode
	}

	if resp.LbServiceID != "" {
		loadBalancer, err := n.np.ReadLoadBalancerService(resp.LbServiceID)
		if err != nil {
			return err
		}
		resp.LbSize = util.StringVal(loadBalancer.Size)
	}

	if resp.SwitchToT1PortID != "" {
		masterVMsNSGroupName, err := n.np.GetMasterVMsNSGroupNameFromSwitchPortTag(resp.SwitchToT1PortID)
		if err != nil {
			return err
		}
		resp.MasterVMsNSGroupName = masterVMsNSGroupName
	}

	return nil
}

func (n *networkManager) getClusterParamsFromResources(resources []*models.ManagedResource, resp *CollectResourcesResp,
	log logrus.FieldLogger) error {

	for _, resource := range resources {
		switch resource.ResourceType {
		case nsx.ResourceTypeLogicalSwitch:
			if !isLbResource(resource.DisplayName) {
				ipBlockID, lbFloatingIPPoolID, err := n.np.GetIpamInfoFromSwitchTag(resource.ID)
				if err != nil {
					log.Errorf("Error getting Node IP Block ID from switch %s tag: %v", resource.ID, err)
					return err
				}
				resp.IPBlockID = ipBlockID
				resp.LBFloatingIPPoolID = lbFloatingIPPoolID
			}
		case nsx.ResourceTypeLogicalRouter:
			if !isLbResource(resource.DisplayName) {
				t0RouterID, err := n.np.GetT0RouterFromRouterTag(resource.ID)
				if err != nil {
					log.Errorf("Error getting T0 Router ID from T1 router %s tag: %v", resource.ID, err)
					return err
				}
				resp.T0RouterID = t0RouterID
			}
		}
	}

	if resp.IPBlockID == "" {
		resp.IPBlockID = n.nsxtSpec.IPBlockID
	}

	if resp.LBFloatingIPPoolID == "" {
		resp.LBFloatingIPPoolID = n.nsxtSpec.FloatingIPPoolID
	}

	if resp.T0RouterID == "" {
		resp.T0RouterID = n.nsxtSpec.T0RouterID
	}

	return nil
}

/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package resourcemanager

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/printer"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/workflow"
)

type resourceManager struct {
	cluster    string
	t0RouterID string
	readOnly   bool
	nsx.Client
	resource Resource
	*printer.Printer
}

// a client to cleanup k8s cluster resources created by ncp
func NewResourceManager(nsxclient nsx.Client, readOnly bool) *resourceManager {
	res := &resourceManager{
		Client:   nsxclient,
		readOnly: readOnly,
		// print to stderr by default
		Printer: printer.New(os.Stderr),
	}
	res.resource = NewResource(nsxclient).SetReadOnly(readOnly).SetPrinter(res.Printer)
	return res
}

// change cluster name to cleanup another cluster using the same client
func (c *resourceManager) SetCluster(cluster string) ResourceManager {
	c.cluster = cluster
	return c
}

func (c *resourceManager) SetT0Router(t0RouterID string) ResourceManager {
	c.t0RouterID = t0RouterID
	return c
}

// Cleanup steps:
// 1. Cleanup firewall sections
// 2. Cleanup NSGroups
// 3. Cleanup ip sets
// 4. Cleanup loadbalancer services
// 5. Cleanup loadbalancer virtual servers
// 6. Cleanup loadbalancer rules
// 7. Cleanup loadbalancer pools
// 8. Cleanup link and ports between T1 and T0 router
// 9. Cleanup logical ports
// 10.Cleanup logical routers
// 11.Cleanup logical switches
// 12.Cleanup ip pools
func (c *resourceManager) CleanupAll() error {
	if c.t0RouterID == "" {
		return errors.New("T0 Router ID is empty")
	}
	return workflow.NewSequentialWorkflows(
		workflow.WorkflowFunc(c.CleanupFirewallSections),
		workflow.WorkflowFunc(c.CleanupNsGroups),
		workflow.WorkflowFunc(c.CleanupIpSets),
		workflow.WorkflowFunc(c.CleanupLbServices),
		workflow.WorkflowFunc(c.CleanupLbVirtualServers),
		workflow.WorkflowFunc(c.CleanupLbRules),
		workflow.WorkflowFunc(c.CleanupLbPools),
		workflow.WorkflowFunc(c.CleanupLbPersistenceProfiles),
		workflow.WorkflowFunc(c.CleanupRouterLinkPortsBetweenT0T1),
		workflow.WorkflowFunc(c.CleanupLogicalPorts),
		workflow.WorkflowFunc(c.CleanupLogicalRouters),
		workflow.WorkflowFunc(c.CleanupLogicalSwitches),
		workflow.WorkflowFunc(c.CleanupNatRuleOnT0),
		workflow.WorkflowFunc(c.CleanupIpPools),
		workflow.WorkflowFunc(c.CleanupL7ResourceCerts),
		workflow.WorkflowFunc(c.CleanupSpoofGuardSwitchingProfiles),
	).Run()
}

// simpleCleanupResource a specific implementation of ncp resources deletion logic
// this shouldn't be exposed
func (c *resourceManager) simpleCleanupResource(resourceType string) error {
	return c.resource.SetResourceType(resourceType).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		Cleanup()
}

// remove all ncp created load balancer persistence profiles
func (c *resourceManager) CleanupLbPersistenceProfiles() error {
	return c.simpleCleanupResource(nsx.ResourceTypePersistenceProfile)
}

// remove all ncp created firewall sections
func (c *resourceManager) CleanupFirewallSections() error {
	return c.simpleCleanupResource(nsx.ResourceTypeFirewallSection)
}

// remove all ncp created ns groups
func (c *resourceManager) CleanupNsGroups() error {
	return c.simpleCleanupResource(nsx.ResourceTypeNSGroup)
}

// remove all ip sets with tag ncp/cluster: [cluster]
func (c *resourceManager) CleanupIpSets() error {
	return c.simpleCleanupResource(nsx.ResourceTypeIPSet)
}

// remove all ncp created lb services
func (c *resourceManager) CleanupLbServices() error {
	return c.simpleCleanupResource(nsx.ResourceTypeLbService)
}

// remove all ncp created lb rules
func (c *resourceManager) CleanupLbRules() error {
	return c.simpleCleanupResource(nsx.ResourceTypeLbRule)
}

// remove all ncp created lb pools
func (c *resourceManager) CleanupLbPools() error {
	return c.simpleCleanupResource(nsx.ResourceTypeLbPool)
}

// remove all ncp created spoofguard switching profiles
func (c *resourceManager) CleanupSpoofGuardSwitchingProfiles() error {
	return c.simpleCleanupResource(nsx.ResourceTypeSpoofGuardSwitchingProfile)
}

// remove all ncp created logical ports with corresponding vif attachment
// detached
func (c *resourceManager) CleanupLogicalPorts() error {
	return c.simpleCleanupResource(nsx.ResourceTypeLogicalPort)
}

func (c *resourceManager) CleanupL7ResourceCerts() error {
	return c.simpleCleanupResource(nsx.ResourceTypeCertificateSelfSigned)
}

// remove all ncp created logical router ports
func (c *resourceManager) CleanupLogicalRouterPorts(id string) error {
	return c.resource.SetResourceType(nsx.ResourceTypeLogicalRouterPort).
		CollectBy(nsx.ResourceCollectFunc(func() ([]interface{}, error) {
			var err error
			ports, err := c.ListLogicalRouterPorts(util.StringPtr(id))
			if err != nil {
				return nil, err
			}
			var res []interface{}
			for _, p := range ports.Results {
				res = append(res, interface{}(p))
			}
			return res, nil
		})).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		Cleanup()
}

// remove all ncp created lb virtual servers
// for each lb virtual server, try to release its external ip address, if present
func (c *resourceManager) CleanupLbVirtualServers() error {
	return c.resource.SetResourceType(nsx.ResourceTypeLbVirtualServer).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.LbVirtualServer:
				return c.ReleaseLoadBalancerVirtualServerIP(obj, c.readOnly)
			default:
				return errors.New("CleanupLbVirtualServers(): unrecognized type")
			}
		})).
		Cleanup()
}

// remove all corresponding router link ports on t0 and ncp created t1 routers
// reference: https://github.com/openstack/vmware-nsxlib/blob/e69d8de2e66ecaa227cce04025265d1026b788a6/vmware_nsxlib/v3/router.py#L99
func (c *resourceManager) CleanupRouterLinkPortsBetweenT0T1() error {
	// silently return if t0 is empty
	if c.t0RouterID == "" {
		c.Warn("no t0 router found. skip CleanupRouterLinkPortsBetweenT0T1\n")
		return nil
	}
	return c.resource.SetResourceType(nsx.RouterTypeTier1).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.LogicalRouter:
				if util.StringVal(obj.RouterType) == nsx.RouterTypeTier1 {
					return c.RemoveRouterLinkPort(obj.ID, c.t0RouterID, c.readOnly)
				}
			default:
				return errors.New("CleanupRouterLinkPortsBetweenT0T1(): unrecognized type")
			}
			return nil
		})).
		Cleanup()
}

// removed all ncp created routers, but always skip tier0
func (c *resourceManager) CleanupLogicalRouters() error {
	return c.resource.SetResourceType(nsx.ResourceTypeLogicalRouter).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			var err error
			switch obj := i.(type) {
			case *models.LogicalRouter:
				{
					err = c.CleanupLogicalRouterPorts(obj.ID)
					if err != nil {
						return err
					}
					err = c.ReleaseLogicalRouterExternalIP(obj, c.readOnly)
					if err != nil {
						return err
					}
				}
			default:
				return errors.New(fmt.Sprintf("CleanupLogicalRouters(): unrecognized type"))
			}
			return nil
		})).
		Cleanup()
}

// CleanupNatRuleOnT0 will clean up all nat rules created on T0 router by ncp
func (c *resourceManager) CleanupNatRuleOnT0() error {
	// if T0 router ID is not set, simply skip
	if c.t0RouterID == "" {
		return nil
	}
	return c.resource.SetResourceType(nsx.ResourceTypeNatRule).
		CollectBy(nsx.ResourceCollectFunc(func() ([]interface{}, error) {
			lres, err := c.ListNatRules(c.t0RouterID)
			if err != nil {
				return nil, err
			}
			var res []interface{}
			for _, r := range lres.Results {
				res = append(res, interface{}(r))
			}
			return res, nil
		})).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.NatRule:
				natRuleID := obj.ID
				natRule, err := c.GetNatRule(c.t0RouterID, natRuleID)
				if err != nil {
					return err
				}
				err = c.ReleaseNatRuleExternalIP(natRule, c.readOnly)
				if err != nil {
					return err
				}
			default:
				return errors.New(fmt.Sprintf("CleanupNatRuleOnT0(): unrecognized type: %s", reflect.TypeOf(obj)))
			}
			return nil
		})).DeleteBy(nsx.ResourceDeleteFunc(func(id string) error {
		return c.DeleteNatRule(c.t0RouterID, id)
	})).Cleanup()
}

// deleteNatRuleOnMatchSourceNetwork deletes a nat rule based on provided
// matchSourceNetwork value
func (c *resourceManager) deleteNatRuleOnMatchSourceNetwork(logicalRouterID string, matchSourceNetwork string) error {
	return c.resource.SetResourceType(nsx.ResourceTypeNatRule).
		CollectBy(nsx.ResourceCollectFunc(func() ([]interface{}, error) {
			lres, err := c.ListNatRules(logicalRouterID)
			if err != nil {
				return nil, err
			}
			var res []interface{}
			for _, r := range lres.Results {
				res = append(res, interface{}(r))
			}
			return res, nil
		})).FilterBy(ResourceFilterFunc(func(i interface{}, args ...string) (bool, error) {
		switch obj := i.(type) {
		case *models.NatRule:
			return obj.MatchSourceNetwork == matchSourceNetwork, nil
		default:
			return false, errors.New(fmt.Sprintf("deleteNatRuleOnMatchSourceNetwork(): unrecognized type: %s", reflect.TypeOf(obj)))
		}
	})).PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
		switch obj := i.(type) {
		case *models.NatRule:
			natRuleID := obj.ID
			natRule, err := c.GetNatRule(logicalRouterID, natRuleID)
			if err != nil {
				return err
			}
			err = c.ReleaseNatRuleExternalIP(natRule, c.readOnly)
			if err != nil {
				return err
			}
		default:
			return errors.New(fmt.Sprintf("deleteNatRuleOnMatchSourceNetwork(): unrecognized type: %s", reflect.TypeOf(obj)))
		}
		return nil
	})).DeleteBy(nsx.ResourceDeleteFunc(func(id string) error {
		return c.DeleteNatRule(logicalRouterID, id)
	})).Cleanup()
}

func (c *resourceManager) CleanupLogicalSwitchPorts(id string) error {
	// overwrite collect function
	// no filter
	return c.resource.SetResourceType(nsx.ResourceTypeLogicalPort).
		CollectBy(nsx.ResourceCollectFunc(func() ([]interface{}, error) {
			lres, err := c.GetLogicalPortsForLogicalSwitch(id)
			if err != nil {
				return nil, err
			}
			var res []interface{}
			for _, r := range lres {
				res = append(res, interface{}(r))
			}
			return res, nil
		})).
		Cleanup()
}

func (c *resourceManager) CleanupRouterPortsForSwitch(id string) error {
	// no filter
	return c.resource.SetResourceType(nsx.ResourceTypeLogicalPort).
		CollectBy(nsx.ResourceCollectFunc(func() ([]interface{}, error) {
			lres, err := c.ListLogicalRouterPortsForSwitch(util.StringPtr(id))
			if err != nil {
				return nil, err
			}
			var res []interface{}
			for _, r := range lres.Results {
				res = append(res, interface{}(r))
			}
			return res, nil
		})).
		Cleanup()
}

func (c *resourceManager) CleanupIPPoolOnSwitch(id string) error {
	if id == "" {
		return nil
	}
	// get ip pool
	ipPool, err := c.ReadIPPool(id)
	if err != nil {
		c.Error("failed to get ip pool %s:%s", id, err.Error())
		return nil
	}
	subnetPtr := nsx.EvaluateTag(&ipPool.ManagedResource, nsx.NcpTagKeySubnet)
	subnetIDPtr := nsx.EvaluateTag(&ipPool.ManagedResource, nsx.NcpTagKeySubnetID)

	// cleanup nat rule for subnet on tier0 router
	if subnetPtr != nil {
		// if there is zero or multiple tier0 routers, silently continue
		if c.t0RouterID != "" {
			err = c.deleteNatRuleOnMatchSourceNetwork(c.t0RouterID, util.StringVal(subnetPtr))
			if err != nil {
				return err
			}
		}
	}

	// cleanup allocated ip addresses and ip pool
	err = c.CleanupIPPool(id, c.readOnly)
	if err != nil {
		return err
	}

	// cleanup subnet from ip block
	if subnetIDPtr != nil {
		c.VerboseInfo("IP Block subnet %s to be removed \n", util.StringVal(subnetIDPtr))
		if !c.readOnly {
			err = c.DeleteIPBlockSubnet(util.StringVal(subnetIDPtr))
			if err != nil {
				return err
			}
			c.VerboseInfo("IP Block subnet %s is removed successfully\n", util.StringVal(subnetIDPtr))
		}
	}
	return nil
}

// cleanup all ncp created logical switches
func (c *resourceManager) CleanupLogicalSwitches() error {
	return c.resource.SetResourceType(nsx.ResourceTypeLogicalSwitch).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.LogicalSwitch:
				{
					err := c.CleanupLogicalSwitchPorts(obj.ID)
					if err != nil {
						return err
					}
				}
			default:
				return errors.New(fmt.Sprintf("CleanupLogicalSwitches(): unrecognized type"))
			}
			return nil
		})).
		AfterDeleteBy(ResourceAfterDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.LogicalSwitch:
				{
					var err error
					err = c.CleanupRouterPortsForSwitch(obj.ID)
					if err != nil {
						return err
					}
					err = c.CleanupIPPoolOnSwitch(obj.IPPoolID)
					if err != nil {
						return err
					}
				}
			default:
				return errors.New(fmt.Sprintf("CleanupLogicalSwitches(): unrecognized type"))
			}
			return nil
		})).
		Cleanup()
}

// cleanup all ncp created ip pools but skip external ones
func (c *resourceManager) CleanupIpPools() error {
	return c.resource.SetResourceType(nsx.ResourceTypeIPPool).
		FilterBy(AdaptedIsNcpResource, c.cluster).
		FilterBy(NotNcpExternalResource).
		PreDeleteBy(ResourcePreDeleteFunc(func(i interface{}) error {
			switch obj := i.(type) {
			case *models.IPPool:
				{
					return c.CleanupIPPool(obj.ID, c.readOnly)
				}
			default:
				return errors.New(fmt.Sprintf("CleanupLogicalSwitches(): unrecognized type"))
			}
		})).
		Cleanup()
}

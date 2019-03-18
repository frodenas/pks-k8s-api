/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/workflow"
)

// DeleteNetwork deletes all networking resources created by PKS for a cluster.
// Functions defined here will delete corresponding resources if passed IDs are not empty, otherwise
// they'll simply just siliently return. So it works regardless of existences of loadbalancer related
// resources
func (n *networkManager) DeleteNetwork(instanceID string) error {

	// collect all cluster resources
	resp, err := n.GetResources(instanceID)
	if err != nil {
		return err
	}

	// return if nothing to delete
	if resp.Num == 0 {
		return nil
	}

	deleteT1Stack := n.NewDeleteT1Stack(&DeleteT1StackRequest{
		&DeleteT1RouterRequest{
			T0ToT1PortID:         resp.T0ToT1PortID,
			T1RouterID:           resp.T1RouterID,
			SnatRuleID:           resp.SnatRuleID,
			InstanceID:           instanceID,
			NatMode:              resp.NatMode,
			T0RouterID:           resp.T0RouterID,
			SnatFloatingIPPoolID: resp.SnatFloatingIPPoolID,
			forLB:                false,
		},
		resp.LogicalSwitchID,
	}, "DeleteT1Stack")
	deleteLbT1Stack := n.NewDeleteT1Stack(&DeleteT1StackRequest{
		&DeleteT1RouterRequest{
			T0ToT1PortID: resp.LbT0ToT1PortID,
			T1RouterID:   resp.LbT1RouterID,
			InstanceID:   instanceID,
			NatMode:      resp.NatMode,
			forLB:        true,
		},
		resp.LbLogicalSwitchID,
	}, "DeleteLbT1Stack")
	deleteIpam := n.NewDeleteIpam(&DeleteIpamRequest{
		IPBlockSubnetID:         resp.IPBlockSubnetID,
		MasterExternalIPAddress: resp.MasterExternalIPAddress,
		InstanceID:              instanceID,
		LBFloatingIPPoolID:      resp.LBFloatingIPPoolID,
	}, "DeleteIpam")
	deleteLbService := n.NewDeleteLbService(&DeleteLbServiceRequest{
		LbServiceID:       resp.LbServiceID,
		LbVirtualServerID: resp.LbVirtualServerID,
		LbPoolID:          resp.LbPoolID,
		LbNSGroupID:       resp.LbNSGroupID,
		LbTcpMonitorID:    resp.LbTcpMonitorID,
	}, "DeleteLbService")

	wf := workflow.NewSequentialWorkflows(deleteT1Stack, deleteLbT1Stack, deleteIpam, deleteLbService)
	if err := wf.Run(); err != nil {
		return err
	}

	return nil
}

// DeleteRouter implements Workflow. Its 'Run' method deletes
// router resources created for the cluster
type DeleteT1StackRequest struct {
	*DeleteT1RouterRequest
	SwitchID string
}

type DeleteT1RouterRequest struct {
	T0ToT1PortID         string
	T1RouterID           string
	SnatRuleID           string
	InstanceID           string
	NatMode              bool
	T0RouterID           string
	SnatFloatingIPPoolID string
	forLB                bool
}

// NewDeleteRouter returns a pointer to a new DeleteRouter struct
func (n *networkManager) NewDeleteT1Stack(req *DeleteT1StackRequest, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.deleteT1Stack(req, n.log.WithField(projName, logField))
	})
}

// Run implements cluster router deletion. T0 router is now shared tagged,
// thus we no longer need to remove NCP cluster specific tags.
func (n *networkManager) deleteT1Stack(req *DeleteT1StackRequest, log logrus.FieldLogger) error {
	var err error
	err = n.deleteT1Router(req.DeleteT1RouterRequest, log)
	if err != nil {
		return err
	}

	err = n.deleteSwitch(req.SwitchID)
	if err != nil {
		return err
	}

	return nil
}

func (n *networkManager) deleteT1Router(req *DeleteT1RouterRequest, log logrus.FieldLogger) error {
	var err error

	if req.NatMode && !req.forLB && req.SnatRuleID != "" {
		snatFloatingIP, err := n.np.ExtractFloatingIPFromNatRule(req.T0RouterID, req.SnatRuleID)
		if err != nil {
			return err
		}

		if snatFloatingIP != "" {
			if req.SnatFloatingIPPoolID != "" {
				if err = n.np.ReleaseFloatingIPAddress(req.SnatFloatingIPPoolID, snatFloatingIP); err != nil {
					return err
				}
			}
		} else {
			// don't return here because nat rule is not deleted yet
			log.Warnf("SNAT floating IP for %s is not found\n", req.InstanceID)
		}

		// NatRule should be deleted at the last in case anything wrong happens
		// before this point, we are still able to find out FIP by looking at nat
		// rule
		if err = n.np.DeleteNatRule(req.T0RouterID, req.SnatRuleID); err != nil {
			return err
		}
	}

	if req.T0ToT1PortID != "" {
		if err = n.np.DeleteT0ToT1Port(req.T0ToT1PortID); err != nil {
			return err
		}
	}

	if req.T1RouterID != "" {
		if err = n.np.DeleteT1Router(req.T1RouterID); err != nil {
			return err
		}
	}

	return nil
}

func (n *networkManager) deleteSwitch(switchID string) error {
	if switchID != "" {
		return n.np.DeleteClusterSwitch(switchID)
	}
	return nil
}

type DeleteIpamRequest struct {
	IPBlockSubnetID         string
	MasterExternalIPAddress strfmt.IPv4
	InstanceID              string
	LBFloatingIPPoolID      string
}

func (n *networkManager) NewDeleteIpam(req *DeleteIpamRequest, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.deleteIpam(req, n.log.WithField(projName, logField))
	})
}

// Run deletes all IPAM resources created by PKS. Since IP blocks and IP pools
// are currently shared tagged, we no longer need to remove NCP cluster specific tags.
func (n *networkManager) deleteIpam(req *DeleteIpamRequest, log logrus.FieldLogger) error {
	var err error

	if req.IPBlockSubnetID != "" {
		if err = n.np.DeleteIPBlockSubnet(req.IPBlockSubnetID); err != nil {
			return err
		}
	} else {
		log.Warnf(" block subnetID for %s is not found\n", req.InstanceID)
	}

	if req.MasterExternalIPAddress != "" {
		if req.LBFloatingIPPoolID != "" {
			if err = n.np.ReleaseFloatingIPAddress(req.LBFloatingIPPoolID, string(req.MasterExternalIPAddress)); err != nil {
				return err
			}
		}
	} else {
		log.Warnf("Floating IP for %s is not found\n", req.InstanceID)
	}

	return nil
}

type DeleteLbServiceRequest struct {
	LbServiceID       string
	LbVirtualServerID string
	LbPoolID          string
	LbNSGroupID       string
	LbTcpMonitorID    string
}

func (n *networkManager) NewDeleteLbService(req *DeleteLbServiceRequest, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.deleteLbService(req, n.log.WithField(projName, logField))
	})
}

func (n *networkManager) deleteLbService(req *DeleteLbServiceRequest, log logrus.FieldLogger) error {
	var err error
	if req.LbServiceID != "" {
		err = n.np.DeleteLoadBalancerService(req.LbServiceID)
		if err != nil {
			return err
		}
	}

	if req.LbVirtualServerID != "" {
		err = n.np.DeleteLoadBalancerVirtualServer(req.LbVirtualServerID)
		if err != nil {
			return err
		}
	}

	if req.LbPoolID != "" {
		err = n.np.DeleteLoadBalancerPool(req.LbPoolID)
		if err != nil {
			return err
		}
	}

	if req.LbTcpMonitorID != "" {
		err = n.np.DeleteLoadBalancerMonitor(req.LbTcpMonitorID)
		if err != nil {
			return err
		}
	}

	if req.LbNSGroupID != "" {
		err = n.np.DeleteNSGroup(req.LbNSGroupID)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/netprovisioner"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/util"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/workflow"
)

const (
	lbResourcePrefix  = "lb-"
	pksResourcePrefix = "pks-"
	projName          = "pks-networking"
	defaultK8sPort    = "8443"
	allocateFIP       = true
	allocateSubnet    = true
	defaultLbSize     = nsx.LbSizeSmall

	// public constants
	// export this variable so clients can decide whether to create cluster network with loadbalancer related resources
	WithLB = true
)

// PrecheckLoadBalancer only validates lbSize string for now
func (n *networkManager) PrecheckLoadBalancer(lbSize string) error {
	switch lbSize {
	case nsx.LbSizeSmall, nsx.LbSizeMedium, nsx.LbSizeLarge:
		return nil
	}
	return fmt.Errorf("unrecognizable loadbalancer size: %s. please choose one from following:%s, %s, %s", lbSize, nsx.LbSizeSmall, nsx.LbSizeMedium, nsx.LbSizeLarge)
}

// CreateLoadbalancer creates load balancer
func (n *networkManager) CreateLoadbalancer(instanceID string, clusterSpec *NSXTClusterSpec) (string, string, error) {
	var err error

	if clusterSpec.LBSize != "" {
		if err = n.PrecheckLoadBalancer(clusterSpec.LBSize); err != nil {
			return "", "", err
		}
	}

	if err := n.populateNsxClusterParams(clusterSpec); err != nil {
		return "", "", err
	}

	if clusterSpec.LBSize == "" {
		clusterSpec.LBSize = defaultLbSize
	}

	if clusterSpec.LBName == "" || clusterSpec.LBFloatingIP == "" {
		return "", "", fmt.Errorf("Cluster %s spec missing lb parameters", instanceID)
	}

	createLbStackResp := &CreateT1StackResp{}
	createLbStackRequest := &CreateT1StackRequest{
		Name:                   clusterSpec.LBName,
		InstanceID:             instanceID,
		AllocateSubnet:         !allocateSubnet,    // don't allocate subnet for lb t1 router
		AllocateFIP:            !allocateFIP,       // don't allocate floating ip
		ForLb:                  true,               // this t1 router is used for load balancer
		LbSize:                 clusterSpec.LBSize, // size of load balancer if this t1 router is used for LB service
		IPBlockIDs:             clusterSpec.IPBlockIDs,
		LbFloatingIPPoolIDs:    clusterSpec.LbFloatingIPPoolIDs,
		SnatFloatingIPPoolIDs:  clusterSpec.SnatFloatingIPPoolIDs,
		T0RouterID:             clusterSpec.T0RouterID,
		EdgeClusterID:          clusterSpec.EdgeClusterID,
		OverlayTransportZoneID: clusterSpec.OverlayTransportZoneID,
		NatMode:                *clusterSpec.NatMode,
	}
	createLbStack := n.NewCreateT1Stack(createLbStackRequest, createLbStackResp, "CreateLbT1Stack")
	if err = createLbStack.Run(); err != nil {
		return "", "", err
	}

	// run this step after previous one finish so that values in response are
	// populated
	createLbServiceResp := &CreateLbServiceResp{}
	createLbServiceReq := &CreateLbServiceReq{
		InstanceID: instanceID,
		LbSize:     clusterSpec.LBSize,
		RouterID:   createLbStackResp.Router.RouterID,
		Fip:        strfmt.IPv4(clusterSpec.LBFloatingIP), // this floating ip will be used on virtual server
		Name:       clusterSpec.LBName,
		MembershipCriteria: []*models.NSGroupTagExpression{
			&models.NSGroupTagExpression{
				// default scopeOp: equals
				// default tagOp: equals
				Scope:             nsx.PksTagKeyK8sMasterVM,
				Tag:               instanceID,
				TargetType:        util.StringPtr(nsx.ResourceTypeLogicalPort),
				NSGroupExpression: models.NSGroupExpression{ResourceType: util.StringPtr(nsx.ResourceTypeNSGroupTagExpression)},
			},
		},
	}

	createLbService := n.NewCreateLbService(createLbServiceReq, createLbServiceResp, "CreateLbService")
	if err = createLbService.Run(); err != nil {
		return "", "", err
	}
	return createLbServiceResp.LbID, createLbServiceReq.LbSize, nil
}

// CreateNetwork provisions required network components for
// a provided instance
func (n *networkManager) CreateNetwork(instanceID string, clusterSpec *NSXTClusterSpec) (NetworkInfo, error) {
	var err error
	var results NetworkInfo

	if !util.ValidateFields(*n.nsxtSpec) {
		return results, fmt.Errorf("Network nsx-t spec contains empty field")
	}

	if err := n.populateNsxClusterParams(clusterSpec); err != nil {
		return results, err
	}

	info, err := n.GetNetwork(instanceID)
	if err != nil {
		return results, err
	}

	// return if there is already resource associated with cluster
	if info.Status != NetworkNotFound {
		return results, fmt.Errorf("Resources for instance %s exist, must specify different instanceID. Collected resources:%+v", instanceID, info)
	}

	// create t1 router and switch job
	createStackResp := &CreateT1StackResp{}
	createStackRequest := &CreateT1StackRequest{
		Name:                   GetPKSResourceName(instanceID),
		InstanceID:             instanceID,
		AllocateSubnet:         allocateSubnet,
		AllocateFIP:            allocateFIP,
		ForLb:                  false, // this t1 router is used for k8s cluster
		IPBlockIDs:             clusterSpec.IPBlockIDs,
		T0RouterID:             clusterSpec.T0RouterID,
		EdgeClusterID:          clusterSpec.EdgeClusterID,
		OverlayTransportZoneID: clusterSpec.OverlayTransportZoneID,
		NatMode:                *clusterSpec.NatMode,
		LbFloatingIPPoolIDs:    clusterSpec.LbFloatingIPPoolIDs,
		SnatFloatingIPPoolIDs:  clusterSpec.SnatFloatingIPPoolIDs,
		MasterVMsNSGroupID:     clusterSpec.MasterVMsNSGroupID,
	}
	createStack := n.NewCreateT1Stack(createStackRequest, createStackResp, "CreateT1Stack")

	wf := workflow.NewSequentialWorkflows(createStack)
	if err = wf.Run(); err != nil {
		return results, err
	}

	results.Cidr = createStackResp.Switch.Cidr
	results.Gateway = createStackResp.Switch.GatewayIP
	results.ExternalIP = createStackResp.ExternalIPAddress
	clusterSpec.LBFloatingIP = string(createStackResp.ExternalIPAddress)

	// created switch has the name specified in creation request
	results.SwitchName = createStackRequest.Name
	results.T0RouterID = clusterSpec.T0RouterID
	results.MasterVMsNSGroupName = createStackResp.Switch.MasterVMsNSGroupName

	// create load balancer related resources conditionally
	if clusterSpec.WithLB {
		if clusterSpec.LBName == "" {
			clusterSpec.LBName = GetLbResourceName(instanceID)
		}
		lbID, lbSize, err := n.CreateLoadbalancer(instanceID, clusterSpec)
		if err != nil {
			return results, err
		}
		results.LbName = GetLbResourceName(instanceID)
		results.LbServiceID = lbID
		results.LbSize = lbSize
	}
	results.Status = NetworkCreated

	return results, nil
}

type CreateLbServiceReq struct {
	InstanceID         string
	LbSize             string
	Name               string
	MembershipCriteria []*models.NSGroupTagExpression
	Fip                strfmt.IPv4
	RouterID           string
}

// only LbID is required at this point if in future we need to provide this ID to ncp for sharing
type CreateLbServiceResp struct {
	LbID string
}

// This is only meant to create load balancer resources for k8s master vms that is on 8443 port
func (n *networkManager) NewCreateLbService(req *CreateLbServiceReq, resp *CreateLbServiceResp, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.createLbService(req, resp, n.log.WithField(projName, logField))
	})
}

func (n *networkManager) createLbService(req *CreateLbServiceReq, resp *CreateLbServiceResp, log logrus.FieldLogger) error {
	tags := []*models.Tag{
		{
			Scope: nsx.PksTagKeyCluster,
			Tag:   req.InstanceID,
		},
	}

	profileID, err := n.np.GetDefaultFastTCPProfile()
	if err != nil {
		return err
	}

	nsgroupID, err := n.np.CreateNSGroupWithCriteria(req.Name, req.MembershipCriteria, tags)
	if err != nil {
		return err
	}

	tcpMonitorID, err := n.np.CreateLbTcpMonitor(req.Name, defaultK8sPort, tags)
	if err != nil {
		return err
	}

	poolID, err := n.np.CreateServerPoolWithNSGroupAndActiveMonitors(req.Name, nsgroupID, []string{tcpMonitorID}, tags)
	if err != nil {
		return err
	}

	serverID, err := n.np.CreateVirtualServer(req.Name, poolID, req.Fip, defaultK8sPort, profileID, tags)
	if err != nil {
		return err
	}

	lbID, err := n.np.CreateLbService(req.Name, req.LbSize, req.RouterID, serverID, tags)
	if err != nil {
		return err
	}

	resp.LbID = lbID
	return nil
}

type CreateT1StackRequest struct {
	Name                   string
	InstanceID             string
	AllocateSubnet         bool
	AllocateFIP            bool
	ForLb                  bool
	LbSize                 string
	IPBlockIDs             []string
	LbFloatingIPPoolIDs    []string
	SnatFloatingIPPoolIDs  []string
	T0RouterID             string
	EdgeClusterID          string
	OverlayTransportZoneID string
	NatMode                bool
	MasterVMsNSGroupID     string
}

// SnatRuleID is not returned here because get_cluster_network will be able to fetch this information based on tags
type CreateT1StackResp struct {
	Router            *createT1RouterResp
	Switch            *createSwitchResp
	ExternalIPAddress strfmt.IPv4 // optional
}

func (n *networkManager) NewCreateT1Stack(req *CreateT1StackRequest, resp *CreateT1StackResp, logField string) workflow.Workflow {
	return workflow.WorkflowFunc(func() error {
		return n.createT1Stack(req, resp, n.log.WithField(projName, logField))
	})
}

func (n *networkManager) createT1Stack(req *CreateT1StackRequest, resp *CreateT1StackResp, log logrus.FieldLogger) error {
	var err error

	rreq := &createT1RouterReq{
		Name:                   req.Name,
		InstanceID:             req.InstanceID,
		ForLb:                  req.ForLb,
		LbSize:                 req.LbSize,
		IPBlockIDs:             req.IPBlockIDs,
		T0RouterID:             req.T0RouterID,
		EdgeClusterID:          req.EdgeClusterID,
		OverlayTransportZoneID: req.OverlayTransportZoneID,
		NatMode:                req.NatMode,
	}
	rresp := &createT1RouterResp{}
	err = n.createT1Router(rreq, rresp, log)
	if err != nil {
		return err
	}

	sreq := &createSwitchRequest{
		Name:                   req.Name,
		RouterID:               rresp.RouterID,
		AllocateSubnet:         req.AllocateSubnet,
		InstanceID:             req.InstanceID,
		IPBlockIDs:             req.IPBlockIDs,
		OverlayTransportZoneID: req.OverlayTransportZoneID,
		NatMode:                req.NatMode,
		MasterVMsNSGroupID:     req.MasterVMsNSGroupID,
	}
	sresp := &createSwitchResp{}
	err = n.createSwitch(sreq, sresp, log)
	if err != nil {
		return err
	}

	if req.AllocateFIP {
		externalIPAddress, _, err := n.allocateFloatingIP(sresp.SwitchID, req.LbFloatingIPPoolIDs, log)
		if err != nil {
			return err
		}
		resp.ExternalIPAddress = strfmt.IPv4(externalIPAddress)
	}

	// SNAT rule is required in Nat mode to give external access to K8s nodes
	// don't create snat rule for loadbalancer T1 router
	if req.NatMode && !req.ForLb {
		err := n.createSnatRule(req.Name, sresp.Cidr, req.InstanceID, req.T0RouterID, req.SnatFloatingIPPoolIDs, req.NatMode, log)
		if err != nil {
			return err
		}
	}

	resp.Router = rresp
	resp.Switch = sresp
	return nil
}

type createT1RouterReq struct {
	Name                   string
	InstanceID             string
	ForLb                  bool
	LbSize                 string
	IPBlockIDs             []string
	T0RouterID             string
	EdgeClusterID          string
	OverlayTransportZoneID string
	NatMode                bool
}

type createT1RouterResp struct {
	RouterID string
}

// Run in CreateT1Router provisions complete creation of a T1 router that includes router creation, stitch to
// T0, and enable route advertisement
func (n *networkManager) createT1Router(req *createT1RouterReq, resp *createT1RouterResp, log logrus.FieldLogger) error {
	clusterRouterTags := []*models.Tag{
		{
			Scope: nsx.PksTagKeyCluster,
			Tag:   req.InstanceID,
		},
	}

	spec := netprovisioner.CreateT1RouterSpec{
		Name:        req.Name,
		NSXTVersion: n.nsxtSpec.NSXTVersion,
		Tags:        clusterRouterTags,
	}

	if !req.ForLb {
		clusterT0RouterTags := append(clusterRouterTags, &models.Tag{
			Scope: nsx.PksTagKeyT0Router,
			Tag:   req.T0RouterID,
		})
		spec.Tags = clusterT0RouterTags
	} else {
		// needs to create service router for LB
		spec.EdgeClusterID = req.EdgeClusterID
		spec.LbSize = req.LbSize
	}

	t1RouterID, err := n.np.CreateT1Router(spec)
	if err != nil {
		return err
	}

	t0ToT1PortID, err := n.np.CreateT0ToT1Port(req.Name, req.T0RouterID, clusterRouterTags)
	if err != nil {
		return err
	}

	// ignore t1ToT0PortID here
	_, err = n.np.CreateT1ToT0Port(req.Name, t1RouterID, t0ToT1PortID, clusterRouterTags)
	if err != nil {
		return err
	}

	err = n.np.EnableRouteAdvertisement(t1RouterID, req.ForLb)
	if err != nil {
		return err
	}

	resp.RouterID = t1RouterID
	return nil
}

type createSwitchRequest struct {
	RouterID               string
	Name                   string
	AllocateSubnet         bool
	InstanceID             string
	IPBlockIDs             []string
	OverlayTransportZoneID string
	NatMode                bool
	MasterVMsNSGroupID     string
}

type createSwitchResp struct {
	// switch related info
	SwitchID string
	// switch subnet configurations
	Cidr                 string
	GatewayIP            strfmt.IPv4
	MasterVMsNSGroupName string
}

// Run in CreateSwitch provisions complete creation of a logical switch for Kubernetes cluster
// that includes switch creation, stitch to T1
// if allocateSubnet is not true, default cidr will be used
func (n *networkManager) createSwitch(req *createSwitchRequest, resp *createSwitchResp, log logrus.FieldLogger) error {
	var (
		subnetCidr, subnetBlockID, subnetIPBlockID string
		gateway                                    strfmt.IPv4
		clusterSwitchWithIPBlockTags               []*models.Tag
		clusterSwitchT1PortTags                    []*models.Tag
		masterVMsNSGroupName                       string
		err                                        error
	)

	clusterSwitchTags := []*models.Tag{
		{
			Scope: nsx.PksTagKeyCluster,
			Tag:   req.InstanceID,
		},
		{
			Scope: nsx.PksTagKeyNat,
			Tag:   strconv.FormatBool(req.NatMode),
		},
	}

	// ignore subnetID here
	if req.AllocateSubnet {
		// IP for T1 router is built up based on the subnetCidr with
		// assumption of prefix length 24, this must be further discussed
		// when prefix is not 24. It's decided to go this way to ease
		// the process of cleanup NSX components
		subnetBlockID, subnetCidr, subnetIPBlockID, err = n.np.AllocateSubnetFromIPBlocks(req.Name, req.IPBlockIDs,
			clusterSwitchTags)
		if err != nil {
			return err
		}
		gateway = strfmt.IPv4(n.np.BuildIPAddress(subnetCidr, strconv.Itoa(nsx.ClusterIPPartT1Router)))
		clusterSwitchWithIPBlockTags = append(clusterSwitchTags, &models.Tag{
			Scope: nsx.PksTagKeyNodeIPBlock,
			Tag:   subnetIPBlockID,
		})
	} else {
		_, subnetCidr, gateway = "", nsx.DefaultT1RouterSubnetCidr, nsx.DefaultT1RouterSubnetGateway
		clusterSwitchWithIPBlockTags = clusterSwitchTags
	}

	switchID, err := n.np.CreateLogicalSwitch(req.Name, req.OverlayTransportZoneID, clusterSwitchWithIPBlockTags)
	if err != nil {
		// Need to cleanup subnet block if creation of logical switch fails since we use the IP block ID tag on
		// logical switch during delete. If this is not cleaned up during delete, we wouldnt be able to find the
		// ip subnet block ID
		// #nosec G104 
		_ = n.cleanupSubnetBlock(req, subnetBlockID)
		return err
	}

	if req.MasterVMsNSGroupID != "" {
		clusterSwitchT1PortTags = append(clusterSwitchTags, &models.Tag{
			Scope: nsx.PksTagKeyMasterVMsNSGroup,
			Tag:   req.MasterVMsNSGroupID,
		})
		nsGroup, err := n.np.ReadNSGroup(req.MasterVMsNSGroupID)
		if err != nil {
			return err
		}
		masterVMsNSGroupName = nsGroup.DisplayName
	} else {
		clusterSwitchT1PortTags = clusterSwitchTags
	}

	switchToT1PortID, err := n.np.CreateSwitchToT1Port(req.Name, switchID, clusterSwitchT1PortTags)
	if err != nil {
		return err
	}

	// ignore t1ToSwitchPortID here
	_, err = n.np.CreateT1ToSwitchPort(req.Name, req.RouterID, switchToT1PortID, gateway, clusterSwitchTags)
	if err != nil {
		return err
	}

	resp.SwitchID = switchID
	resp.Cidr = subnetCidr
	resp.GatewayIP = gateway
	resp.MasterVMsNSGroupName = masterVMsNSGroupName
	return nil
}

func (n *networkManager) cleanupSubnetBlock(req *createSwitchRequest, subnetBlockID string) error {
	if req.AllocateSubnet {
		if err := n.np.DeleteIPBlockSubnet(subnetBlockID); err != nil {
			return fmt.Errorf("Failed to cleanup subnet block %s: %v", subnetBlockID, err)
		}
	}

	return nil
}

// allocateFloatingIP allocates a floating IP from shared Floating IP Pool and tags given logical switch with allocated
// floating IP
func (n *networkManager) allocateFloatingIP(switchID string, floatingIPPoolIDs []string, log logrus.FieldLogger) (string, string, error) {
	masterExternalIPAddress, floatingIPPoolID, err := n.np.AllocateFloatingIPAddressFromIPPools(floatingIPPoolIDs)
	if err != nil {
		return "", "", err
	}

	clusterSwitchFloatingIPTags := []*models.Tag{
		&models.Tag{
			Scope: nsx.PksTagKeyFloatingIP,
			Tag:   masterExternalIPAddress,
		},
		{
			Scope: nsx.PksTagKeyLBFloatingIPPool,
			Tag:   floatingIPPoolID,
		},
	}

	switchID, err = n.np.UpdateLogicalSwitchTags(switchID, clusterSwitchFloatingIPTags)
	if err != nil {
		if fipErr := n.np.ReleaseFloatingIPAddress(floatingIPPoolID, masterExternalIPAddress); fipErr != nil {
			log.Errorf("Failed to release external IP: %s\n", fipErr)
		}
		return "", "", err
	}
	return masterExternalIPAddress, floatingIPPoolID, nil
}

// createSnatRule will firstly allocate a floating IP from shared floating IP Pool with instanceID, natmode tags, then create
// a snat rule for given sources on shared T0 router
func (n *networkManager) createSnatRule(name, cidr, instanceID, t0RouterID string, floatingIPPoolIDs []string, natMode bool, log logrus.FieldLogger) error {
	tags := []*models.Tag{
		{
			Scope: nsx.PksTagKeyCluster,
			Tag:   instanceID,
		},
		{
			Scope: nsx.PksTagKeyNat,
			Tag:   strconv.FormatBool(natMode),
		},
	}
	// SNAT rule is required in Nat mode to give external access to K8s nodes
	if natMode {
		snatFloatingIPAddress, snatFloatingIPPoolID, err := n.np.AllocateFloatingIPAddressFromIPPools(floatingIPPoolIDs)
		if err != nil {
			return err
		}

		tags = append(tags, &models.Tag{
			Scope: nsx.PksTagKeySnatFloatingIpPool,
			Tag:   snatFloatingIPPoolID,
		})

		// ignore snatRuleID here
		_, err = n.np.CreateSnatRule(t0RouterID, name, cidr, snatFloatingIPAddress, tags)
		if err != nil {
			if fipErr := n.np.ReleaseFloatingIPAddress(snatFloatingIPPoolID, snatFloatingIPAddress); fipErr != nil {
				log.Errorf("Failed to release SNAT floating IP: %s\n", fipErr)
			}
			return err
		}
		return nil
	}
	return nil
}

func GetLbResourceName(name string) string {
	return lbResourcePrefix + GetPKSResourceName(name)
}

func isLbResource(name string) bool {
	return strings.HasPrefix(name, lbResourcePrefix)
}

func GetPKSResourceName(name string) string {
	return pksResourcePrefix + name
}

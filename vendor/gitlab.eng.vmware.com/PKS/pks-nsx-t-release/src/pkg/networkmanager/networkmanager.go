/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/vcenter"
	np "gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/netprovisioner"
	vcp "gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/vcprovisioner"
)

// LogLevel is the type of log level in package level
type LogLevel string

// Log levels
const (
	LogPanic LogLevel = "panic"
	LogFatal LogLevel = "fatal"
	LogError LogLevel = "error"
	LogWarn  LogLevel = "warn"
	LogInfo  LogLevel = "info"
	LogDebug LogLevel = "debug"

	NetworkCreated NetworkStatus = iota
	NetworkNotFound
	NetworkUnknown
)

type NetworkStatus int

// networkManager implements ClusterNetworkManager interface
type networkManager struct {
	np                np.NsxNetworkProvisioner
	vcp               *vcp.VcenterProvisioner
	nsxtSpec          *NSXTSpec
	vcSpec            *VcenterSpec
	t1RouterSpec      *T1RouterSpec
	logicalSwitchSpec *LogicalSwitchSpec
	ipamSpec          *IPAMSpec
	log               logrus.FieldLogger
}

// verify networkManager implements NetworkManager interface
var _ NetworkManager = &networkManager{}

// NSXTSpec holds static configuration of a nsx-t environment
// note: don't put dynamic runtime variable in here
type NSXTSpec struct {
	EdgeClusterID          string
	T0RouterID             string
	OverlayTransportZoneID string
	IPBlockID              string
	FloatingIPPoolID       string
	NatMode                bool
	NSXTVersion            string
}

// VcenterSpec defines configuration of a vcenter
type VcenterSpec struct {
	DataCenterPath     string
	ComputeClusterPath string
}

// NSXTClusterSpec holds cluster specific NSX-T configurations
type NSXTClusterSpec struct {
	T0RouterID            string
	IPBlockIDs            []string
	LbFloatingIPPoolIDs   []string
	SnatFloatingIPPoolIDs []string
	MasterVMsNSGroupID    string

	// NAT parameters
	NatMode *bool

	// LB parameters
	WithLB       bool
	LBSize       string
	LBName       string
	LBFloatingIP string

	EdgeClusterID          string
	OverlayTransportZoneID string
}

// T1RouterSpec represents spec of the NSX T1 router
type T1RouterSpec struct {
	ClusterT1RouterID      string
	ClusterLogicalSwitchID string
	T0ToT1PortID           string
	T1ToT0PortID           string
	SnatRuleID             string
	DnatRuleID             string
	Cidr                   string
	AllocatedCidrs         []string
}

// LogicalSwitchSpec represents spec of a NSX switch
type LogicalSwitchSpec struct {
	LogicalSwitchID  string
	SwitchToT1PortID string
	T1ToSwitchPortID string
	Cidr             string
	IPPoolID         string
	GatewayIPAddress strfmt.IPv4
}

// IPAMSpec represents the spec of IPAM resources
type IPAMSpec struct {
	ClusterBlockSubnetID    string
	MasterExternalIPAddress strfmt.IPv4
	MasterInternalIPAddress strfmt.IPv4
}

// NetworkInfo represents spec of network info about an instance
type NetworkInfo struct {
	Cidr                 string
	Gateway              strfmt.IPv4
	ExternalIP           strfmt.IPv4
	LbServiceID          string
	LbName               string
	LbSize               string
	Status               NetworkStatus
	SwitchName           string
	T0RouterID           string
	MasterVMsNSGroupName string
}

// utcFormatter is used to convert timestamp to UTC
type utcFormatter struct {
	logrus.TextFormatter
}

// Format is used to convert timestamp to UTC
func (u utcFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.TextFormatter.Format(e)
}

// NewNetworkManager creates a new networkManager
// return concrete implementation
// NetworkManager here is multi-thread safe because all its fields are static, bound to one nsx-t environment
func NewNetworkManager(nsxtSpec *NSXTSpec,
	vcspec *VcenterSpec, vcclient *vcenter.Client,
	nsxNetworkProvisioner np.NsxNetworkProvisioner,
	loglevel LogLevel, optionalSpecs ...interface{}) (*networkManager, error) {

	log := logrus.New()
	log.Formatter = utcFormatter{logrus.TextFormatter{FullTimestamp: true}}
	level, err := logrus.ParseLevel(string(loglevel))
	if err != nil {
		defaultLogLevel := logrus.InfoLevel
		log.Warnf("Invalid log level %s, will fall back to %s\n", loglevel, defaultLogLevel.String())
		log.SetLevel(defaultLogLevel)
	} else {
		log.SetLevel(level)
	}
	logEntry := log.WithField("pks-networking", "networkManager")
	nsxNetworkProvisioner.SetLogger(logEntry)
	nm := &networkManager{
		np:                nsxNetworkProvisioner,
		vcp:               vcp.NewVcenterProvisioner(vcclient, logEntry),
		nsxtSpec:          nsxtSpec,
		t1RouterSpec:      &T1RouterSpec{},
		logicalSwitchSpec: &LogicalSwitchSpec{},
		ipamSpec:          &IPAMSpec{},
		vcSpec:            vcspec,
		log:               logEntry,
	}

	err = populateNsxParams(nm)
	if err != nil {
		return nil, err
	}

	for _, optionalSpec := range optionalSpecs {
		switch t := optionalSpec.(type) {
		case *T1RouterSpec:
			nm.t1RouterSpec = t
		case *LogicalSwitchSpec:
			nm.logicalSwitchSpec = t
		case *IPAMSpec:
			nm.ipamSpec = t
		}
	}

	return nm, nil
}

// populateNsxParams injects nsxtVersion to nsxtSpec
func populateNsxParams(n *networkManager) error {
	nodeProperties, err := n.np.ReadNodeProperties()
	if err != nil {
		n.log.Errorf("Failed to get node properties")
		return err
	}
	n.nsxtSpec.NSXTVersion = nodeProperties.NodeVersion
	if n.nsxtSpec.T0RouterID != "" {
		edgeClusterID, overlayTransportZoneID, err := n.getNSXProvisioningParams(n.nsxtSpec.T0RouterID)
		if err != nil {
			n.log.Errorf("Failed to get NSX provisioning properties: %v", err)
			return err
		}
		n.nsxtSpec.EdgeClusterID = edgeClusterID
		n.nsxtSpec.OverlayTransportZoneID = overlayTransportZoneID
		n.log.Debugf("Populated spec with edge cluster %s and transport zone %s", edgeClusterID, overlayTransportZoneID)
	}

	return nil
}

// populateNsxClusterParams injects edgeClusterID and transportZoneID to NSXTClusterSpec
func (n *networkManager) populateNsxClusterParams(nsxtClusterSpec *NSXTClusterSpec) error {
	if nsxtClusterSpec.T0RouterID != "" && (nsxtClusterSpec.EdgeClusterID == "" || nsxtClusterSpec.OverlayTransportZoneID == "") {
		edgeClusterID, overlayTransportZoneID, err := n.getNSXProvisioningParams(nsxtClusterSpec.T0RouterID)
		if err != nil {
			n.log.Errorf("Failed to get NSX provisioning properties: %v", err)
			return err
		}
		nsxtClusterSpec.EdgeClusterID = edgeClusterID
		nsxtClusterSpec.OverlayTransportZoneID = overlayTransportZoneID
		n.log.Debugf("Populated cluster spec with edge cluster %s and transport zone %s", edgeClusterID, overlayTransportZoneID)
	}

	if nsxtClusterSpec.T0RouterID == "" {
		nsxtClusterSpec.T0RouterID = n.nsxtSpec.T0RouterID
		nsxtClusterSpec.EdgeClusterID = n.nsxtSpec.EdgeClusterID
		nsxtClusterSpec.OverlayTransportZoneID = n.nsxtSpec.OverlayTransportZoneID
	}

	if nsxtClusterSpec.IPBlockIDs == nil || len(nsxtClusterSpec.IPBlockIDs) == 0 {
		nsxtClusterSpec.IPBlockIDs = append(nsxtClusterSpec.IPBlockIDs, n.nsxtSpec.IPBlockID)
	}

	if nsxtClusterSpec.LbFloatingIPPoolIDs == nil || len(nsxtClusterSpec.LbFloatingIPPoolIDs) == 0 {
		nsxtClusterSpec.LbFloatingIPPoolIDs = append(nsxtClusterSpec.LbFloatingIPPoolIDs, n.nsxtSpec.FloatingIPPoolID)
	}

	if nsxtClusterSpec.SnatFloatingIPPoolIDs == nil || len(nsxtClusterSpec.SnatFloatingIPPoolIDs) == 0 {
		nsxtClusterSpec.SnatFloatingIPPoolIDs = nsxtClusterSpec.LbFloatingIPPoolIDs
	}

	if nsxtClusterSpec.NatMode == nil {
		nsxtClusterSpec.NatMode = &n.nsxtSpec.NatMode
	}

	return nil
}

// getNSXProvisioningParams gets edge cluster ID and overlay transport zone ID given t0 router ID
func (n *networkManager) getNSXProvisioningParams(t0RouterID string) (edgeClusterID string, transportZoneID string, err error) {
	edgeClusterID, err = n.np.ExtractEdgeClusterIDFromT0Router(t0RouterID)
	if err != nil {
		n.log.Errorf("Failed to extract edge cluster ID from router %s", t0RouterID)
		return "", "", err
	}

	transportZoneID, err = n.np.ExtractTransportZoneIDFromEdgeCluster(edgeClusterID)
	if err != nil {
		n.log.Errorf("Failed to extract transport zone ID from edge cluster %s", edgeClusterID)
		return "", "", err
	}

	return edgeClusterID, transportZoneID, nil
}

// bindLogger sets logger for the implementation of NsxNetworkProvisioner in this repo
// not set logger if not present
// returns a provisioner instance with different logger
func bindLogger(p np.NsxNetworkProvisioner, log logrus.FieldLogger) np.NsxNetworkProvisioner {
	p.SetLogger(log)
	return p
}

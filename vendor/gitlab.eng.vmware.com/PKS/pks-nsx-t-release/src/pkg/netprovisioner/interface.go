/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"github.com/Sirupsen/logrus"
	"github.com/go-openapi/strfmt"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// NsxNetworkProvisioner represents provisioner to create/delete NSX resources
type NsxNetworkProvisioner interface {
	SetLogger(logrus.FieldLogger)
	Logger() logrus.FieldLogger
	CheckClusterStatus() error
	CheckTransportZone(string) error
	CheckT0Router(string) error
	CheckIPPool(string) error
	CheckIPBlock(string) error
	CheckIPBlockSubnetPrefix(string, int) error
	CheckFabricNodes([]*FabricNode) error
	CheckFabricNodeState(string) error
	CheckFabricNodeStatus(string) error
	CheckNSGroup(string) error

	GetFabricNodeGivenMgmtIP([]string) (*models.Node, error)

	ExtractEdgeClusterIDFromT0Router(string) (string, error)
	ExtractTransportZoneIDFromEdgeCluster(string) (string, error)
	UntagTransportZone(string, string) error

	GetAllResources(string, models.Tag) ([]*models.ManagedResource, error)

	CreateT1Router(CreateT1RouterSpec) (string, error)
	CreateT0ToT1Port(string, string, []*models.Tag) (string, error)
	CreateT1ToT0Port(string, string, string, []*models.Tag) (string, error)
	EnableRouteAdvertisement(string, bool) error
	CreateDnatRule(string, string, string, string, []*models.Tag) (string, error)
	CreateSnatRule(string, string, string, string, []*models.Tag) (string, error)
	DeleteNatRule(string, string) error
	GetNatRule(string, string) (*models.NatRule, error)
	ExtractFloatingIPFromNatRule(string, string) (string, error)

	AllocateFloatingIPAddress(string) (string, error)
	AllocateFloatingIPAddressFromIPPools([]string) (string, string, error)
	CreateLogicalSwitch(string, string, []*models.Tag) (string, error)
	UpdateLogicalSwitchTags(string, []*models.Tag) (string, error)
	CreateSwitchToT1Port(string, string, []*models.Tag) (string, error)
	AllocateSubnetBlock(string, string, []*models.Tag) (string, string, error)
	AllocateSubnetFromIPBlocks(string, []string, []*models.Tag) (string, string, string, error)
	AllocateIPAddress(string) (string, error)
	BuildIPAddress(string, string) string
	CreateT1ToSwitchPort(string, string, string, strfmt.IPv4, []*models.Tag) (string, error)
	GetSwitchingProfileByTag(string, string, string) (*models.BaseSwitchingProfile, error)
	CreateSpoofGuardSwitchingProfile(string, string, []*models.Tag, []string) (string, error)

	GetIPBlockSubnet(string, string) (string, string, error)
	GetMetadataFromSwitchTag(string) (string, bool, error)
	GetIpamInfoFromSwitchTag(string) (string, string, error)
	GetT0RouterFromRouterTag(clusterRouterID string) (string, error)
	GetMasterVMsNSGroupNameFromSwitchPortTag(string) (string, error)
	DeleteT0ToT1Port(string) error
	DeleteT1Router(string) error
	UntagT0Router(string, string) error

	DeleteClusterSwitch(string) error
	GetLogicalSwitchByTag(lsScope, lsTag string) (*models.LogicalSwitch, error)

	DeleteIPBlockSubnet(string) error
	ReleaseFloatingIPAddress(string, string) error
	UntagIPBlock(string, string) error
	UntagExternalBlocks(string) error
	UntagExternalPools(string) error

	GetDefaultFastTCPProfile() (string, error)
	CreateServerPoolWithNSGroupAndActiveMonitors(string, string, []string, []*models.Tag) (string, error)
	CreateVirtualServer(string, string, strfmt.IPv4, string, string, []*models.Tag) (string, error)
	CreateLbService(string, string, string, string, []*models.Tag) (string, error)
	DeleteLoadBalancerService(string) error
	ReadLoadBalancerService(string) (*models.LbService, error)
	DeleteLoadBalancerVirtualServer(string) error
	DeleteLoadBalancerPool(string) error
	GetLoadbalancerByTag(string, string) (*models.LbService, error)
	CreateLbTcpMonitor(string, string, []*models.Tag) (string, error)
	DeleteLoadBalancerMonitor(string) error

	//nsgroup
	CreateNSGroupWithCriteria(string, []*models.NSGroupTagExpression, []*models.Tag) (string, error)
	DeleteNSGroup(string) error
	ReadNSGroup(string) (*models.NSGroup, error)

	// trust-management
	AddCertificateImport(trustObjectData *models.TrustObjectData) (*models.CertificateList, error)
	DeleteCertificate(certID string) error
	GetCertificates() (*models.CertificateList, error)
	GetCertificate(certID string) (*models.Certificate, error)
	RegisterPrincipalIdentity(principalIdentity *models.PrincipalIdentity) (*models.PrincipalIdentity, error)
	DeletePrincipalIdentity(principalIdentityID string) error
	GetPrincipalIdentities() (*models.PrincipalIdentityList, error)

	// node service
	ReadNodeProperties() (*models.NodeProperties, error)
}

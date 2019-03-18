/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	fab "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/fabric"
)

// Client represents all nsx client wrappers
type Client interface {
	Validate() error
	WithInsecure(bool) Client
	WithRootCA(string) Client
	WithClientCert(string, string) Client
	WithClientCertFromFile(string, string) Client
	WithBasicAuth(string, string) Client
	WithOverWriteHeader() Client
	WithDebug() Client

	ReadClusterStatus() (*models.ClusterStatus, error)

	ListFabricNodes() (*models.NodeListResult, error)
	ReadFabricNode(string) (*models.Node, error)
	GetFabricNodeGivenParams(*fab.ListNodesParams) (*models.Node, error)
	GetFabricNodeMatchingIP([]string) (*models.Node, error)
	GetFabricNodeState(string) (*models.ConfigurationState, error)
	ReadFabricNodeStatus(string) (*models.NodeStatus, error)

	ReadIPBlock(string) (*models.IPBlock, error)
	UpdateIPBlock(*models.IPBlock) (*models.IPBlock, error)
	ReadIPBlockSubnet(string) (*models.IPBlockSubnet, error)
	ListIPBlockSubnets(string) (*models.IPBlockSubnetListResult, error)
	ReadIPPool(string) (*models.IPPool, error)
	UpdateIPPool(*models.IPPool) (*models.IPPool, error)
	ListIPPoolAllocations(string) (*models.AllocationIPAddressListResult, error)
	DeleteIPBlockSubnet(string) error
	AllocateSubnetFromIPBlock(*models.IPBlockSubnet) (*models.IPBlockSubnet, error)
	AllocateIPFromSubnetBlock(string) (*models.AllocationIPAddress, error)
	AllocateIPFromIPPool(string) (*models.AllocationIPAddress, error)
	ReleaseIPToIPPool(string, *models.AllocationIPAddress) error
	TagIPBlock(string, []*models.Tag) (*models.IPBlock, error)
	TagIPPool(string, []*models.Tag) (*models.IPPool, error)
	ListIPPools() (*models.IPPoolListResult, error)
	DeleteIPPool(string, bool) error
	CleanupIPPool(string, bool) error

	ListLogicalRouters() (*models.LogicalRouterListResult, error)
	CreateLogicalRouter(*models.LogicalRouter) (*models.LogicalRouter, error)
	DeleteLogicalRouter(string) error
	ReadLogicalRouter(string) (*models.LogicalRouter, error)
	UpdateLogicalRouter(*models.LogicalRouter) (*models.LogicalRouter, error)
	GetAdvertisementConfig(string) (*models.AdvertisementConfig, error)
	UpdateAdvertisementConfig(string, *models.AdvertisementConfig) (*models.AdvertisementConfig, error)
	CreateLogicalRouterPort(*models.LogicalRouterPort) (*models.LogicalRouterPort, error)
	ListLogicalRouterPorts(*string) (*models.LogicalRouterPortListResult, error)
	DeleteLogicalRouterPort(string) error
	ListLogicalRoutersByType(string) (*models.LogicalRouterListResult, error)
	ListT0LogicalRouters() (*models.LogicalRouterListResult, error)
	ListT1LogicalRouters() (*models.LogicalRouterListResult, error)
	AddNatRule(string, *models.NatRule) (*models.NatRule, error)
	ListNatRules(string) (*models.NatRuleListResult, error)
	GetNatRule(string, string) (*models.NatRule, error)
	DeleteNatRule(string, string) error
	TagRouter(string, []*models.Tag) (*models.LogicalRouter, error)
	ListLogicalRouterPortsForSwitch(*string) (*models.LogicalRouterPortListResult, error)
	GetTier1LinkPort(routerID string) (*models.LogicalRouterPort, error)
	RemoveRouterLinkPort(string, string, bool) error
	ReleaseLogicalRouterExternalIP(*models.LogicalRouter, bool) error
	ReleaseNatRuleExternalIP(*models.NatRule, bool) error

	CreateLogicalSwitch(*models.LogicalSwitch) (*models.LogicalSwitch, error)
	UpdateLogicalSwitch(*models.LogicalSwitch) (*models.LogicalSwitch, error)
	GetLogicalSwitch(string) (*models.LogicalSwitch, error)
	DeleteLogicalSwitch(string) error
	ListLogicalSwitches() (*models.LogicalSwitchListResult, error)
	CreateLogicalPort(*models.LogicalPort) (*models.LogicalPort, error)
	GetLogicalSwitchGivenName(string) (*models.LogicalSwitch, error)
	ListLogicalPorts() (*models.LogicalPortListResult, error)
	GetLogicalPort(string) (*models.LogicalPort, error)
	GetLogicalPortsForLogicalSwitch(string) ([]*models.LogicalPort, error)
	DeleteLogicalPort(string) error
	CreateSpoofGuardSwitchingProfile(*models.SpoofGuardSwitchingProfile) (*models.SpoofGuardSwitchingProfile, error)
	ListSwitchingProfilesByType(string) (*models.SwitchingProfilesListResult, error)
	DeleteSwitchingProfile(string) error

	SearchByTag(*string, models.Tag) (*models.SearchResults, error)

	ListTranportZones() (*models.TransportZoneListResult, error)
	GetTransportZone(string) (*models.TransportZone, error)
	UpdateTransportZone(*models.TransportZone) (*models.TransportZone, error)
	TagTransportZone(string, []*models.Tag) (*models.TransportZone, error)
	GetTransportNode(string) (*models.TransportNode, error)

	ReadEdgeCluster(string) (*models.EdgeCluster, error)

	// firewall section
	ListFirewallSections() (*models.FirewallSectionListResult, error)
	DeleteFirewallSection(string) error

	// ip set
	DeleteIPSet(string) error
	ListIPSets() (*models.IPSetListResult, error)

	// ns group
	ListNSGroups() (*models.NSGroupListResult, error)
	DeleteNSGroup(string) error
	CreateNSGroup(*models.NSGroup) (*models.NSGroup, error)
	ReadNSGroup(string) (*models.NSGroup, error)

	// load balancer
	ListLoadBalancerServices() (*models.LbServiceListResult, error)
	CreateLoadBalancerService(*models.LbService) (*models.LbService, error)
	DeleteLoadBalancerService(string) error
	ReadLoadBalancerService(string) (*models.LbService, error)
	ListLoadBalancerVirtualServers() (*models.LbVirtualServerListResult, error)
	CreateLoadBalancerVirtualServer(*models.LbVirtualServer) (*models.LbVirtualServer, error)
	DeleteLoadBalancerVirtualServer(string) error
	ListLoadBalancerRules() (*models.LbRuleListResult, error)
	CreateLoadBalancerPool(pool *models.LbPool) (*models.LbPool, error)
	PerformPoolMemberAction(string, string, *models.PoolMemberSettingList) (*models.LbPool, error)
	DeleteLoadBalancerRule(string) error
	ListLoadBalancerPools() (*models.LbPoolListResult, error)
	DeleteLoadBalancerPool(string) error
	ReleaseLoadBalancerVirtualServerIP(*models.LbVirtualServer, bool) error
	ListLoadBalancerApplicationProfiles() (*models.LbAppProfileListResult, error)
	CreateLoadBalancerTcpMonitor(*models.LbTCPMonitor) (*models.LbTCPMonitor, error)
	DeleteLoadBalancerMonitor(string) error

	ListLoadBalancerPersistenceProfiles() (*models.LbPersistenceProfileListResult, error)
	DeleteLoadBalancerPersistenceProfile(id string) error

	// registry
	ResourceDeleteFunc(string) ResourceDeleteFunc
	ResourceCollectFunc(string) ResourceCollectFunc

	// trust-management
	AddCertificateImport(trustObjectData *models.TrustObjectData) (*models.CertificateList, error)
	DeleteCertificate(certID string) error
	GetCertificates() (*models.CertificateList, error)
	GetCertificate(certID string) (*models.Certificate, error)
	RegisterPrincipalIdentity(principalIdentity *models.PrincipalIdentity) (*models.PrincipalIdentity, error)
	DeletePrincipalIdentity(principalIdentityID string) error
	GetPrincipalIdentities() (*models.PrincipalIdentityList, error)

	// node service
	CreateProxyServiceApplyCertificate(certificateID string) error
	ReadNodeProperties() (*models.NodeProperties, error)
}

type ResourceCollectFunc func() ([]interface{}, error)
type ResourceDeleteFunc func(string) error

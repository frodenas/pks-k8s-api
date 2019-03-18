/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

const (
	basePath    = "/api/v1"
	contentType = "application/json"
)

var scheme = []string{"https"}

// NSX router types
const (
	RouterTypeTier0 = "TIER0"
	RouterTypeTier1 = "TIER1"
)

// Tag related constants
const (
	PksTagKeyK8sMasterVM            = "pks/k8smastervm"
	PksTagKeyCluster                = "pks/cluster"
	PksTagKeyFloatingIP             = "pks/floating_ip"
	PksTagValueFloatingIPDefault    = "none"
	PksTagKeyNoNat                  = "pks/no_nat"
	PksTagKeyNat                    = "pks/nat"
	PksTagSuperuser                 = "pks/superuser"
	PksTagCACert                    = "pks/cacert"
	PksTagKeyNodeIPBlock            = "pks/node_ip_block"
	PksTagKeySnatFloatingIpPool     = "pks/snat_fip_pool"
	PksTagKeyLBFloatingIPPool       = "pks/lb_fip_pool"
	PksTagKeyT0Router               = "pks/t0_router"
	PksTagKeyMasterVMsNSGroup       = "pks/master_vms_nsgroup"
	NcpTagKeyCluster                = "ncp/cluster"
	NcpTagKeyVersion                = "ncp/version"
	NcpTagKeyNode                   = "ncp/node_name"
	NcpTagKeySharedResource         = "ncp/shared_resource"
	NcpTagKeyExternal               = "ncp/external"
	NcpTagValueExternal             = "true"
	NsxMaxTagsAllowed               = 15
	NcpTagKeyExternalIPPool         = "ext_pool_id"
	NcpTagKeyExternalIPPoolForSnat  = "ncp/extpoolid"
	NcpTagKeySnatIPForLogicalRouter = "ncp/snat_ip"
	NcpTagKeySnat                   = "ncp/snat"
	NcpTagKeySubnet                 = "ncp/subnet"
	NcpTagKeySubnetID               = "ncp/subnet_id"
	NcpTagKeyOwner                  = "owner"
	NcpTagValueOwner                = "NCP"
	SpoofGuardPortBindings          = "LPORT_BINDINGS"
)

// HighAvailability types
const (
	HAActiveStandby = "ACTIVE_STANDBY"
	HAActiveActive  = "ACTIVE_ACTIVE"
)

// Display name suffix
const (
	SuffixT1Router       = "-cluster-router"
	SuffixT1ToT0Port     = "-t1-to-t0-port"
	SuffixT0ToT1Port     = "-t0-to-t1-port"
	SuffixLogicalSwitch  = "-cluster-switch"
	SuffixSwitchToT1Port = "-switch-to-t1-port"
	SuffixT1ToSwitchPort = "-t1-to-switch-port"
	SuffixSubnetBlock    = "-subnet-block"
	SuffixNatRule        = "-nat-rule"
	SuffixLbPool         = "-lb-pool"
	SuffixVirtualServer  = "-virtual-server"
	SuffixLbService      = "-lb-service"
	SuffixNSGroup        = "-nsgroup"
)

// Available port types on a logical router
const (
	PortTypeDownLinkPort    = "LogicalRouterDownLinkPort"
	PortTypeLinkPortOnTier0 = "LogicalRouterLinkPortOnTIER0"
	PortTypeLinkPortOnTier1 = "LogicalRouterLinkPortOnTIER1"
	PortTypeUpLinkPort      = "LogicalRouterUpLinkPort"
)

// Available port types on a logical switch.
const (
	PortTypeLinkPortOnSwitch = "LogicalPort"
)

// Replication modes for NSX logical switch
const (
	// Stands for Mutilple Tunnel EndPoint. Each VNI will nominate one host Tunnel Endpoint as MTEP,
	// which replicates BUM traffic to other hosts within the same VNI
	SwitchReplicationModeMtep = "MTEP"

	// Hosts create a copy of each BUM frame and send copy to each tunnel endpoint that it knows for each VNI
	SwitchReplicationModeSource = "SOURCE"
)

// Admin states for NSX logical switch
const (
	// Being managed by nsx manager
	SwitchAdminStateUp = "UP"

	// Not being managed by nsx manager
	SwitchAdminStateDown = "DOWN"
)

// Allocate release related constant
const (
	IPAddressActionAllocate = "ALLOCATE"
	IPAddressActionRelease  = "RELEASE"
)

// Cluster subnet related constant
const (
	ClusterSubnetPrefixLength = 24
	ClusterSubnetBlockSize    = 256
	ClusterIPPartT1Router     = 1
)

// Searchable resource types
const (
	SearchTypeIPBlock = "IPBlock"
	SearchTypeIPPool  = "IPPool"
)

// NSX resource type
const (
	ResourceTypeLogicalRouter                = "LogicalRouter"
	ResourceTypeLogicalSwitch                = "LogicalSwitch"
	ResourceTypeLogicalPort                  = "LogicalPort"
	ResourceTypeLogicalRouterLinkPortOnTIER0 = "LogicalRouterLinkPortOnTIER0"
	ResourceTypeLogicalRouterDownlinkPort    = "LogicalRouterDownLinkPort"
	ResourceTypeLogicalRouterLinkPortOnTIER1 = "LogicalRouterLinkPortOnTIER1"
	ResourceTypeNatRule                      = "NatRule"
	ResourceTypeFirewallSection              = "FirewallSection"
	ResourceTypeNSGroup                      = "NSGroup"
	ResourceTypeLbService                    = "LbService"
	ResourceTypeLbVirtualServer              = "LbVirtualServer"
	ResourceTypeLbPool                       = "LbPool"
	ResourceTypeLbRule                       = "LbRule"
	ResourceTypeIPSet                        = "IPSet"
	ResourceTypeLogicalRouterPort            = "LogicalRouterPort"
	ResourceTypeIPPool                       = "IPPool"
	ResourceTypeLbFastTcpProfile             = "LbFastTcpProfile"
	ResourceTypeLbProfile                    = "LbProfile"
	ResourceTypeNSGroupTagExpression         = "NSGroupTagExpression"
	ResourceTypeLbTcpMonitor                 = "LbTcpMonitor"
	ResourceTypePrincipalIdentity            = "PrincipalIdentity"
	ResourceTypeCertificateSelfSigned        = "certificate_self_signed"
	ResourceTypeSpoofGuardSwitchingProfile   = "SpoofGuardSwitchingProfile"
	ResourceTypePersistenceProfile           = "LbPersistenceProfile"
	ResourceTypeLbSourceIpPersistenceProfile = "LbSourceIpPersistenceProfile"
	ResourceTypeLbCookiePersistenceProfile   = "LbCookiePersistenceProfile"
)

// Transport zone types
const (
	TransportZoneTypeOverlay = "OVERLAY"
	TransportZoneTypeVLAN    = "VLAN"
)

// NAT rule types
const (
	NatRuleDNAT = "DNAT"
	NatRuleSNAT = "SNAT"
)

// Load balancer
const (
	NsxDefaultLbFastTcpProfileDisplayName = "nsx-default-lb-fast-tcp-profile"
	AlgorithmRoundRobin                   = "ROUND_ROBIN"
	TCP                                   = "TCP"
	LbSizeSmall                           = "SMALL"
	LbSizeMedium                          = "MEDIUM"
	LbSizeLarge                           = "LARGE"
	LbSnatTranslationAutoMap              = "LbSnatAutoMap"
	LbSnatTranslationIpPool               = "LbSnatIpPool"
	LbMaxIPListSize                       = int64(5)
)

// Default Values
const (
	DefaultT1RouterSubnetCidr    = "169.254.169.1/24"
	DefaultT1RouterSubnetGateway = "169.254.169.1"
)

// Trust-management related constants
const (
	// permission groups
	ReadOnlyApiUsers  = "read_only_api_users"
	ReadWriteApiUsers = "read_write_api_users"
	Superusers        = "superusers"
	// Display name
	SuperuserDisplayName = "pks-nsx-superuser"
)

// NSX error codes
const (
	NsxIPBlockExhaustionErrorCode = 5137
	NsxIPPoolExhaustionErrorCode  = 5109
)

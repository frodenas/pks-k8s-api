/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package resourcemanager

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/printer"
)

type ResourceManager interface {
	CleanupFirewallSections() error
	CleanupNsGroups() error
	CleanupIpSets() error
	CleanupLbServices() error
	CleanupLbVirtualServers() error
	CleanupLbRules() error
	CleanupLbPools() error
	CleanupRouterLinkPortsBetweenT0T1() error
	CleanupLogicalPorts() error
	CleanupLogicalRouters() error
	CleanupLogicalRouterPorts(id string) error
	CleanupLogicalSwitches() error
	CleanupRouterPortsForSwitch(id string) error
	CleanupIPPoolOnSwitch(id string) error
	CleanupSpoofGuardSwitchingProfiles() error
	CleanupIpPools() error
	CleanupAll() error

	SetCluster(cluster string) ResourceManager
	SetT0Router(t0RouterID string) ResourceManager
}

type Resource interface {
	SetResourceType(string) Resource
	GetResourceType() string

	CollectBy(nsx.ResourceCollectFunc) Resource
	FilterBy(ResourceFilterFunc, ...string) Resource
	PreDeleteBy(ResourcePreDeleteFunc) Resource
	DeleteBy(f nsx.ResourceDeleteFunc) Resource
	AfterDeleteBy(ResourceAfterDeleteFunc) Resource
	SetPrinter(*printer.Printer) Resource

	GetCollection() ResourceCollection

	Cleanup() error
	SetReadOnly(bool) Resource
}

type ResourceCollection <-chan interface{}
type ResourceFilterFunc func(interface{}, ...string) (bool, error)
type ResourcePreDeleteFunc func(interface{}) error
type ResourceAfterDeleteFunc func(interface{}) error

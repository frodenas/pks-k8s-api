/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package networkmanager

// NetworkManager manages lifecycle of networks
type NetworkManager interface {
	// if withLb is true, CreateNetwork should create loadbalancer related resources in
	// target network
	CreateNetwork(instanceID string, clusterSpec *NSXTClusterSpec) (NetworkInfo, error)
	CreateLoadbalancer(instanceID string, clusterSpec *NSXTClusterSpec) (string, string, error)
	PrecheckLoadBalancer(lbSize string) error
	GetNetwork(instanceID string) (NetworkInfo, error)
	GetResources(instanceID string) (CollectResourcesResp, error)

	DeleteNetwork(instanceID string) error
	PreCheckCreateNetwork(instanceID string, clusterSpec *NSXTClusterSpec) error

	// CreateGlobalResources creates resources pertaining to all clusters like spoofguard switching profile
	CreateGlobalResources() error
}

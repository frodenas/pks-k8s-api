/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenter

import (
	"errors"

	"github.com/vmware/govmomi/object"
)

// GetEsxiHostsInCluster gets all the esxi host objects belonging to a cluster
func (vc *Client) GetEsxiHostsInCluster(datacenterPath string, clusterPath string) ([]*object.HostSystem, error) {
	ctx := vc.context

	cluster, clusterErr := vc.GetComputeResource(datacenterPath, clusterPath)
	if clusterErr != nil {
		return nil, clusterErr
	}

	hosts, hostsErr := cluster.Hosts(ctx)
	if hostsErr != nil {
		return nil, hostsErr
	}

	return hosts, nil
}

// GetEsxiHostManagementIPs gets all the management ips of a given host
func (vc *Client) GetEsxiHostManagementIPs(host *object.HostSystem) ([]string, error) {
	var managementIps []string
	ctx := vc.context

	if host == nil {
		return nil, errors.New("host system is nil")
	}

	mgmtIPs, mgmtIPErr := host.ManagementIPs(ctx)
	if mgmtIPErr != nil {
		return nil, mgmtIPErr
	}

	for _, mgmtIP := range mgmtIPs {
		managementIps = append(managementIps, mgmtIP.String())
	}

	return managementIps, nil
}

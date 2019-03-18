/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenterprovisioner

import (
	"fmt"
)

// EsxiHost represents a esxi host
type EsxiHost struct {
	MgmtIPs []string
}

//GetEsxiHostIPs gets all the management IPs of all esxi hosts in a compute cluster
func (vcp *VcenterProvisioner) GetEsxiHostIPs(dcPath string, computeResourcePath string) ([]EsxiHost, error) {
	var esxiHosts []EsxiHost

	hosts, hostsErr := vcp.client.GetEsxiHostsInCluster(dcPath, computeResourcePath)
	if hostsErr != nil {
		return nil, hostsErr
	}

	for _, host := range hosts {
		mgmtips, mgmtipsErr := vcp.client.GetEsxiHostManagementIPs(host)
		if mgmtipsErr != nil {
			return nil, mgmtipsErr
		}

		if mgmtips == nil || len(mgmtips) <= 0 {
			return nil, fmt.Errorf("Failed to find mgmt ips for esxi host %s", host.Name())
		}

		esxiHosts = append(esxiHosts, EsxiHost{mgmtips})
	}

	return esxiHosts, nil
}

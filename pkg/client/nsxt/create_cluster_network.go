/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/networkmanager"
)

// CreateClusterNetwork creates the network resources for a cluster.
func (nc *client) CreateClusterNetwork(nm networkmanager.NetworkManager, name string, spec *networkmanager.NSXTClusterSpec) (networkmanager.NetworkInfo, error) {
	log.Info(fmt.Sprintf("Creating Network for Cluster `%s`", name))

	networkInfo, err := nm.CreateNetwork(name, spec)
	if err != nil {
		log.Error(err, fmt.Sprintf("Error creating Network for Cluster `%s`", name))

		errDel := nc.DeleteClusterNetwork(nm, name)
		if errDel != nil {
			log.Error(errDel, fmt.Sprintf("Error deleting Network for Cluster `%s`", name))
		}
		return networkInfo, err
	}

	return networkInfo, nil
}

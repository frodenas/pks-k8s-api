/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/networkmanager"
)

// DeleteClusterNetwork creates the network resources for a cluster.
func (nc *client) DeleteClusterNetwork(nm networkmanager.NetworkManager, name string) error {
	log.Info(fmt.Sprintf("Deleting Network for Cluster `%s`", name))

	if err := nm.DeleteNetwork(name); err != nil {
		log.Error(err, fmt.Sprintf("Error deleting Network for Cluster `%s`", name))
		return err
	}

	return nil
}

/*
 * Copyright (c) 2019 VMware, Inc. All rights reserved.
 */

package nsxt

import (
	"fmt"

	"gitlab.eng.vmware.com/PKS/pks-nsx-t-release/src/pkg/networkmanager"
)

// NewNetworkManager returns a new NSX-T network manager.
func (nc *client) NewNetworkManager(nsxtSpec *networkmanager.NSXTSpec) (networkmanager.NetworkManager, error) {
	networkManager, err := networkmanager.NewNetworkManager(nsxtSpec, nil, nil, nc.nsxtProvisioner, "INFO")
	if err != nil {
		return nil, fmt.Errorf("error creating a NSX-T network manager: %v", err)
	}

	return networkManager, nil
}

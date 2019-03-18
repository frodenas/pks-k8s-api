/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"errors"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// CheckClusterStatus checks if controller status is stable
func (p *nsxNetworkProvisioner) CheckClusterStatus() error {
	csRes, csErr := p.ReadClusterStatus()
	if csErr != nil {
		p.log.Errorf("Unable to check NSX Cluster stability due to error %s", csErr)
		return csErr
	}

	if csRes.ControlClusterStatus.Status != models.ControllerClusterStatusStatusSTABLE {
		csErr = errors.New("NSX Control Cluster is unstable")
		return csErr
	}

	if csRes.MgmtClusterStatus.Status != models.ManagementClusterStatusStatusSTABLE {
		csErr = errors.New("NSX Management Cluster is unstable")
		return csErr
	}

	return nil
}

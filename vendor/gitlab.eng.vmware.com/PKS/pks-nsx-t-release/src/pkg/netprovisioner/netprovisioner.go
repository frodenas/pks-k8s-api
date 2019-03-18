/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package netprovisioner

import (
	"github.com/Sirupsen/logrus"
	"gitlab.eng.vmware.com/PKS/pks-networking/pkg/client/nsx"
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/pool_management"
)


// NsxNetworkProvisioner provides all the methods to operate on NSX network resources
type nsxNetworkProvisioner struct {
	nsx.Client
	log logrus.FieldLogger
}

// SetLogger sets a private logger for the provisioner
func (np *nsxNetworkProvisioner) SetLogger(log logrus.FieldLogger) {
	np.log = log
}

// Logger returns provisioner's logger
func (np *nsxNetworkProvisioner) Logger() logrus.FieldLogger {
	return np.log
}

// NewNsxNetworkProvisioner creates a new NsxNetworkProvisioner
func NewNsxNetworkProvisioner(c nsx.Client, log logrus.FieldLogger) (*nsxNetworkProvisioner, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return &nsxNetworkProvisioner{
		Client: c,
		log:    log,
	}, nil
}

// IsValidError identifies if the error is a valid error for retry
func IsValidError(err error) bool {	
	switch err.(type) {
	case *pool_management.AllocateOrReleaseFromIPBlockSubnetConflict:
		conflictErr, _ := err.(*pool_management.AllocateOrReleaseFromIPPoolConflict)
		if conflictErr.Payload.ErrorCode == nsx.NsxIPPoolExhaustionErrorCode {
			break
		}
		return true
	case *pool_management.CreateIPBlockSubnetConflict:
		conflictErr, _ := err.(*pool_management.CreateIPBlockSubnetConflict)
		if conflictErr.Payload.ErrorCode == nsx.NsxIPBlockExhaustionErrorCode  {
			break
		}
		return true
	case *pool_management.DeleteIPBlockSubnetConflict:
		return true
	case *pool_management.UpdateIPBlockConflict: //not used yet
		return true
	}
	return false
}
/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package vcenterprovisioner

import (
	"github.com/Sirupsen/logrus"
	"github.com/vmware/govmomi/object"
)

type vcenterAPIClient interface {
	GetEsxiHostsInCluster(datacenterPath string, clusterPath string) ([]*object.HostSystem, error)
	GetEsxiHostManagementIPs(host *object.HostSystem) ([]string, error)
}

// VcenterProvisioner provides all the methods to operate on Vcenter resources
type VcenterProvisioner struct {
	client vcenterAPIClient
	log    logrus.FieldLogger
}

// NewVcenterProvisioner creates a new VcenterProvisioner
func NewVcenterProvisioner(c vcenterAPIClient, log logrus.FieldLogger) *VcenterProvisioner {
	return &VcenterProvisioner{
		client: c,
		log:    log,
	}
}

// GetLogger returns provisioner's logger
func (c *VcenterProvisioner) GetLogger() logrus.FieldLogger {
	return c.log
}

/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"strings"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	nca "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/nsx_component_administration"
)

// CreateProxyServiceApplyCertificate adds a new certificate
func (nc *client) CreateProxyServiceApplyCertificate(certificateID string) error {
	params := nca.NewCreateProxyServiceApplyCertificateActionApplyCertificateParams().WithAction("apply_certificate").WithCertificateID(certificateID)
	_, err := nc.client.NsxComponentAdministration.CreateProxyServiceApplyCertificateActionApplyCertificate(params, nc.auth)
	if err != nil && strings.HasPrefix(err.Error(), "no consumer") {
		err = nil
	}
	return err
}

// ReadNodeProperties reads NSX manager properties
func (nc *client) ReadNodeProperties() (*models.NodeProperties, error) {
	params := nca.NewReadNodePropertiesParams()
	res, err := nc.client.NsxComponentAdministration.ReadNodeProperties(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

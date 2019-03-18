/*
 * Copyright (c) 2018 VMware, Inc. All rights reserved.
 */

package nsx

import (
	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
	nca "gitlab.eng.vmware.com/PKS/pks-networking/gen/nsx/nsx_component_administration"
)

// AddCertificateImport adds a new certificate
func (nc *client) AddCertificateImport(trustObjectData *models.TrustObjectData) (*models.CertificateList, error) {
	params := nca.NewAddCertificateImportParams().WithTrustObjectData(trustObjectData).WithAction("import")
	res, err := nc.client.NsxComponentAdministration.AddCertificateImport(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeleteCertificate deletes certificate for the given certificate ID
func (nc *client) DeleteCertificate(certID string) error {
	params := nca.NewDeleteCertificateParams().WithCertID(certID)
	_, err := nc.client.NsxComponentAdministration.DeleteCertificate(params, nc.auth)
	return err
}

// GetCertificates returns all the user facing components certificates
func (nc *client) GetCertificates() (*models.CertificateList, error) {
	params := nca.NewGetCertificatesParams()
	res, err := nc.client.NsxComponentAdministration.GetCertificates(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetCertificate shows certificate data for the given certificate ID
func (nc *client) GetCertificate(certID string) (*models.Certificate, error) {
	params := nca.NewGetCertificateParams().WithCertID(certID)
	res, err := nc.client.NsxComponentAdministration.GetCertificate(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// RegisterPrincipalIdentity registers a name certificate combination
func (nc *client) RegisterPrincipalIdentity(principalIdentity *models.PrincipalIdentity) (*models.PrincipalIdentity, error) {
	params := nca.NewRegisterPrincipalIdentityParams().WithPrincipalIdentity(principalIdentity)
	res, err := nc.client.NsxComponentAdministration.RegisterPrincipalIdentity(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// DeletePrincipalIdentity deletes a principal identity
func (nc *client) DeletePrincipalIdentity(principalIdentityID string) error {
	params := nca.NewDeletePrincipalIdentityParams().WithPrincipalIdentityID(principalIdentityID)
	_, err := nc.client.NsxComponentAdministration.DeletePrincipalIdentity(params, nc.auth)
	return err
}

// GetPrincipalIdentities returns the list of principal identities
func (nc *client) GetPrincipalIdentities() (*models.PrincipalIdentityList, error) {
	params := nca.NewGetPrincipalIdentitiesParams()
	res, err := nc.client.NsxComponentAdministration.GetPrincipalIdentities(params, nc.auth)
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

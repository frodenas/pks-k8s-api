// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LbPersistenceProfile lb persistence profile
// swagger:model LbPersistenceProfile

type LbPersistenceProfile struct {
	ManagedResource

	// Persistence shared flag for the associated virtual servers
	//
	// If persistence shared flag is not set in the cookie persistence profile
	// bound to a virtual server, it defaults to cookie persistence that is
	// private to each virtual server and is qualified by the pool. This is
	// accomplished by load balancer inserting a cookie with name in the
	// format &lt;name&gt;.&lt;virtual_server_id&gt;.&lt;pool_id&gt;.
	// If persistence shared flag is set in the cookie persistence profile, in
	// cookie insert mode, cookie persistence could be shared across multiple
	// virtual servers that are bound to the same pools. The cookie name would
	// be changed to &lt;name&gt;.&lt;profile-id&gt;.&lt;pool-id&gt;.
	// If persistence shared flag is not set in the sourceIp persistence
	// profile bound to a virtual server, each virtual server that the profile
	// is bound to maintains its own private persistence table.
	// If persistence shared flag is set in the sourceIp persistence profile,
	// all virtual servers the profile is bound to share the same persistence
	// table.
	//
	PersistenceShared *bool `json:"persistence_shared,omitempty"`

	// Source-ip persistence ensures all connections from a client (identified by
	// IP address) are sent to the same backend server for a specified period.
	// Cookie persistence allows related client connections, identified by the
	// same cookie in HTTP requests, to be redirected to the same server.
	//
	// Required: true
	ResourceType string `json:"resource_type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LbPersistenceProfile) UnmarshalJSON(raw []byte) error {

	var aO0 ManagedResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedResource = aO0

	var data struct {
		PersistenceShared *bool `json:"persistence_shared,omitempty"`

		ResourceType string `json:"resource_type"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.PersistenceShared = data.PersistenceShared

	m.ResourceType = data.ResourceType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LbPersistenceProfile) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ManagedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		PersistenceShared *bool `json:"persistence_shared,omitempty"`

		ResourceType string `json:"resource_type"`
	}

	data.PersistenceShared = m.PersistenceShared

	data.ResourceType = m.ResourceType

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this lb persistence profile
func (m *LbPersistenceProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ManagedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var lbPersistenceProfileTypeResourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LbCookiePersistenceProfile","LbSourceIpPersistenceProfile"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lbPersistenceProfileTypeResourceTypePropEnum = append(lbPersistenceProfileTypeResourceTypePropEnum, v)
	}
}

// property enum
func (m *LbPersistenceProfile) validateResourceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, lbPersistenceProfileTypeResourceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LbPersistenceProfile) validateResourceType(formats strfmt.Registry) error {

	if err := validate.RequiredString("resource_type", "body", string(m.ResourceType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateResourceTypeEnum("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LbPersistenceProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LbPersistenceProfile) UnmarshalBinary(b []byte) error {
	var res LbPersistenceProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

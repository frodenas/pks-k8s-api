// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPSet Set of one or more IP addresses
//
// IPSet is used to group individual IP addresses, range of IP addresses or subnets.
// An IPSet is a homogeneous group of IP addresses, either of type IPv4 or of type
// IPv6. IPSets can be used as source or destination in firewall rules. These can
// also be used as members of NSGroups.
//
// swagger:model IPSet

type IPSet struct {
	ManagedResource

	// IP addresses
	//
	// IP addresses
	// Max Items: 100
	IPAddresses []string `json:"ip_addresses"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IPSet) UnmarshalJSON(raw []byte) error {

	var aO0 ManagedResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedResource = aO0

	var data struct {
		IPAddresses []string `json:"ip_addresses,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.IPAddresses = data.IPAddresses

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IPSet) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ManagedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		IPAddresses []string `json:"ip_addresses,omitempty"`
	}

	data.IPAddresses = m.IPAddresses

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this IP set
func (m *IPSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ManagedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSet) validateIPAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAddresses) { // not required
		return nil
	}

	iIPAddressesSize := int64(len(m.IPAddresses))

	if err := validate.MaxItems("ip_addresses", "body", iIPAddressesSize, 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPSet) UnmarshalBinary(b []byte) error {
	var res IPSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

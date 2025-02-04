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

// IPSubnet IP subnet
// swagger:model IPSubnet

type IPSubnet struct {

	// IPv4 Addresses
	//
	// IPv4 Addresses
	// Required: true
	// Max Items: 1
	// Min Items: 1
	IPAddresses []strfmt.IPv4 `json:"ip_addresses"`

	// Subnet Prefix Length
	//
	// Subnet Prefix Length
	// Required: true
	// Maximum: 32
	// Minimum: 1
	PrefixLength *int64 `json:"prefix_length"`
}

/* polymorph IPSubnet ip_addresses false */

/* polymorph IPSubnet prefix_length false */

// Validate validates this IP subnet
func (m *IPSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrefixLength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPSubnet) validateIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("ip_addresses", "body", m.IPAddresses); err != nil {
		return err
	}

	iIPAddressesSize := int64(len(m.IPAddresses))

	if err := validate.MinItems("ip_addresses", "body", iIPAddressesSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("ip_addresses", "body", iIPAddressesSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *IPSubnet) validatePrefixLength(formats strfmt.Registry) error {

	if err := validate.Required("prefix_length", "body", m.PrefixLength); err != nil {
		return err
	}

	if err := validate.MinimumInt("prefix_length", "body", int64(*m.PrefixLength), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("prefix_length", "body", int64(*m.PrefixLength), 32, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPSubnet) UnmarshalBinary(b []byte) error {
	var res IPSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

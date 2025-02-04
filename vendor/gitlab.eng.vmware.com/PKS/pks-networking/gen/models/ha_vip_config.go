// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HaVipConfig ha vip config
// swagger:model HaVipConfig

type HaVipConfig struct {

	// Flag to enable this ha vip config.
	//
	// Flag to enable this ha vip config.
	Enabled *bool `json:"enabled,omitempty"`

	// Floating IP address subnets
	//
	// Array of IP address subnets which will be used as floating IP addresses. | Note - this configuration is applicable only for Active-Standby LogicalRouter. | For Active-Active LogicalRouter this configuration will be rejected.
	// Required: true
	// Max Items: 1
	// Min Items: 1
	HaVipSubnets []*VIPSubnet `json:"ha_vip_subnets"`

	// Identifiers of uplink ports for providing redundancy
	//
	// Identifiers of logical router uplink ports which are to be paired to provide | redundancy. Floating IP will be owned by one of these uplink ports (depending upon | which node is Active).
	// Required: true
	RedundantUplinkPortIds []string `json:"redundant_uplink_port_ids"`
}

/* polymorph HaVipConfig enabled false */

/* polymorph HaVipConfig ha_vip_subnets false */

/* polymorph HaVipConfig redundant_uplink_port_ids false */

// Validate validates this ha vip config
func (m *HaVipConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHaVipSubnets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRedundantUplinkPortIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HaVipConfig) validateHaVipSubnets(formats strfmt.Registry) error {

	if err := validate.Required("ha_vip_subnets", "body", m.HaVipSubnets); err != nil {
		return err
	}

	iHaVipSubnetsSize := int64(len(m.HaVipSubnets))

	if err := validate.MinItems("ha_vip_subnets", "body", iHaVipSubnetsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("ha_vip_subnets", "body", iHaVipSubnetsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.HaVipSubnets); i++ {

		if swag.IsZero(m.HaVipSubnets[i]) { // not required
			continue
		}

		if m.HaVipSubnets[i] != nil {

			if err := m.HaVipSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ha_vip_subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HaVipConfig) validateRedundantUplinkPortIds(formats strfmt.Registry) error {

	if err := validate.Required("redundant_uplink_port_ids", "body", m.RedundantUplinkPortIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HaVipConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HaVipConfig) UnmarshalBinary(b []byte) error {
	var res HaVipConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

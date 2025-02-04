// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EdgeClusterMemberAllocationProfile edge cluster member allocation profile
// swagger:model EdgeClusterMemberAllocationProfile

type EdgeClusterMemberAllocationProfile struct {

	// Edge Cluster Member Allocation Pool for logical router
	//
	// Logical router allocation can be tracked for specific services and
	// services may have their own hard limits and allocation sizes. For
	// example load balancer pool should be specified if load balancer
	// service will be attached to logical router.
	//
	AllocationPool *EdgeClusterMemberAllocationPool `json:"allocation_pool,omitempty"`
}

/* polymorph EdgeClusterMemberAllocationProfile allocation_pool false */

// Validate validates this edge cluster member allocation profile
func (m *EdgeClusterMemberAllocationProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocationPool(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeClusterMemberAllocationProfile) validateAllocationPool(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocationPool) { // not required
		return nil
	}

	if m.AllocationPool != nil {

		if err := m.AllocationPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation_pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeClusterMemberAllocationProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeClusterMemberAllocationProfile) UnmarshalBinary(b []byte) error {
	var res EdgeClusterMemberAllocationProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

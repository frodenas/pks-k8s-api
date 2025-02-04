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

// RevisionedResource revisioned resource
// swagger:model RevisionedResource

type RevisionedResource struct {
	Resource

	// Generation of this resource config
	//
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	// Required: true
	Revision *int64 `json:"_revision"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RevisionedResource) UnmarshalJSON(raw []byte) error {

	var aO0 Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Resource = aO0

	var data struct {
		Revision *int64 `json:"_revision"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Revision = data.Revision

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RevisionedResource) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Revision *int64 `json:"_revision"`
	}

	data.Revision = m.Revision

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this revisioned resource
func (m *RevisionedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevisionedResource) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("_revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RevisionedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevisionedResource) UnmarshalBinary(b []byte) error {
	var res RevisionedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

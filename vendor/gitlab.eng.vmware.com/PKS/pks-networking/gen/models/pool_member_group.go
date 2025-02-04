// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoolMemberGroup pool member group
// swagger:model PoolMemberGroup

type PoolMemberGroup struct {

	// List of customized pool member settings
	//
	// The list is used to show the customized pool member settings. User can
	// only user pool member action API to update the admin state for a specific
	// IP address.
	//
	CustomizedMembers []*PoolMemberSetting `json:"customized_members"`

	// Grouping object resource reference
	//
	// Load balancer pool support grouping object as dynamic pool members.
	// The IP list of the grouping object such as NSGroup would be used as
	// pool member IP setting.
	//
	// Required: true
	GroupingObject *ResourceReference `json:"grouping_object"`

	// Filter of ipv4 or ipv6 address of grouping object IP list
	//
	// Ip revision filter is used to filter IPv4 or IPv6 addresses from the
	// grouping object.
	// If the filter is not specified, both IPv4 and IPv6 addresses would be
	// used as server IPs.
	// The link local and loopback addresses would be always filtered out.
	//
	IPRevisionFilter *string `json:"ip_revision_filter,omitempty"`

	// Maximum number of grouping object IP address list
	//
	// The size is used to define the maximum number of grouping object IP
	// address list. These IP addresses would be used as pool members.
	// If the grouping object includes more than certain number of
	// IP addresses, the redundant parts would be ignored and those IP
	// addresses would not be treated as pool members.
	//
	// Required: true
	// Minimum: 0
	MaxIPListSize *int64 `json:"max_ip_list_size"`
}

/* polymorph PoolMemberGroup customized_members false */

/* polymorph PoolMemberGroup grouping_object false */

/* polymorph PoolMemberGroup ip_revision_filter false */

/* polymorph PoolMemberGroup max_ip_list_size false */

// Validate validates this pool member group
func (m *PoolMemberGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomizedMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupingObject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPRevisionFilter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxIPListSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolMemberGroup) validateCustomizedMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomizedMembers) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomizedMembers); i++ {

		if swag.IsZero(m.CustomizedMembers[i]) { // not required
			continue
		}

		if m.CustomizedMembers[i] != nil {

			if err := m.CustomizedMembers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customized_members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolMemberGroup) validateGroupingObject(formats strfmt.Registry) error {

	if err := validate.Required("grouping_object", "body", m.GroupingObject); err != nil {
		return err
	}

	if m.GroupingObject != nil {

		if err := m.GroupingObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grouping_object")
			}
			return err
		}
	}

	return nil
}

var poolMemberGroupTypeIPRevisionFilterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPV4","IPV6","IPV4_IPV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poolMemberGroupTypeIPRevisionFilterPropEnum = append(poolMemberGroupTypeIPRevisionFilterPropEnum, v)
	}
}

const (
	// PoolMemberGroupIPRevisionFilterIPV4 captures enum value "IPV4"
	PoolMemberGroupIPRevisionFilterIPV4 string = "IPV4"
	// PoolMemberGroupIPRevisionFilterIPV6 captures enum value "IPV6"
	PoolMemberGroupIPRevisionFilterIPV6 string = "IPV6"
	// PoolMemberGroupIPRevisionFilterIPV4IPV6 captures enum value "IPV4_IPV6"
	PoolMemberGroupIPRevisionFilterIPV4IPV6 string = "IPV4_IPV6"
)

// prop value enum
func (m *PoolMemberGroup) validateIPRevisionFilterEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, poolMemberGroupTypeIPRevisionFilterPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PoolMemberGroup) validateIPRevisionFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.IPRevisionFilter) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPRevisionFilterEnum("ip_revision_filter", "body", *m.IPRevisionFilter); err != nil {
		return err
	}

	return nil
}

func (m *PoolMemberGroup) validateMaxIPListSize(formats strfmt.Registry) error {

	if err := validate.Required("max_ip_list_size", "body", m.MaxIPListSize); err != nil {
		return err
	}

	if err := validate.MinimumInt("max_ip_list_size", "body", int64(*m.MaxIPListSize), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolMemberGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolMemberGroup) UnmarshalBinary(b []byte) error {
	var res PoolMemberGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

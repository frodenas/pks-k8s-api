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

// LogicalRouter logical router
// swagger:model LogicalRouter

type LogicalRouter struct {
	ManagedResource

	// Logical Router Configuration
	//
	// Contains config properties for tier0 routers
	AdvancedConfig *LogicalRouterConfig `json:"advanced_config,omitempty"`

	// Edge Cluster Member Allocation Profile
	//
	// Configurations options to auto allocate edge cluster members for
	// logical router. Auto allocation is supported only for TIER1 and pick
	// least utilized member post current assignment for next allocation.
	//
	AllocationProfile *EdgeClusterMemberAllocationProfile `json:"allocation_profile,omitempty"`

	// Identifier of the edge cluster for this Logical Router
	//
	// Used for tier0 routers
	EdgeClusterID string `json:"edge_cluster_id,omitempty"`

	// Member indices of the edge node on the cluster
	//
	// For stateful services, the logical router should be associated with
	// edge cluster. For TIER 1 logical router, for manual placement of
	// service router within the cluster, edge cluster member indices needs
	// to be provided else same will be auto-allocated. You can provide
	// maximum two indices for HA ACTIVE_STANDBY. For TIER0 logical router
	// this property is no use and placement is derived from logical router
	// uplink or loopback port.
	//
	EdgeClusterMemberIndices []int64 `json:"edge_cluster_member_indices"`

	// Failover mode for active-standby logical router instances.
	//
	// This failover mode determines, whether the preferred service router instance
	// for given logical router will preempt the peer.
	// Note - It can be specified if and only if logical router is ACTIVE_STANDBY and
	// NON_PREEMPTIVE mode is supported only for a Tier1 logical router. For Tier0 ACTIVE_STANDBY logical router,
	// failover mode is always PREEMPTIVE, i.e. once the preferred node comes up
	// after a failure, it will preempt the peer causing failover from current active to preferred node. For ACTIVE_ACTIVE logical routers, this field must not be populated.
	//
	FailoverMode string `json:"failover_mode,omitempty"`

	// LR Firewall Section References
	//
	// List of Firewall sections related to Logical Router.
	// Read Only: true
	FirewallSections []*ResourceReference `json:"firewall_sections"`

	// High availability mode
	//
	// High availability mode
	HighAvailabilityMode string `json:"high_availability_mode,omitempty"`

	// Preferred edge cluster member index in active standby mode
	// for pre-emptive failover
	//
	//
	// Used for tier0 routers only
	// Minimum: 0
	PreferredEdgeClusterMemberIndex *int64 `json:"preferred_edge_cluster_member_index,omitempty"`

	// Type of Logical Router
	//
	// Type of Logical Router
	// Required: true
	RouterType *string `json:"router_type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LogicalRouter) UnmarshalJSON(raw []byte) error {

	var aO0 ManagedResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedResource = aO0

	var data struct {
		AdvancedConfig *LogicalRouterConfig `json:"advanced_config,omitempty"`

		AllocationProfile *EdgeClusterMemberAllocationProfile `json:"allocation_profile,omitempty"`

		EdgeClusterID string `json:"edge_cluster_id,omitempty"`

		EdgeClusterMemberIndices []int64 `json:"edge_cluster_member_indices,omitempty"`

		FailoverMode string `json:"failover_mode,omitempty"`

		FirewallSections []*ResourceReference `json:"firewall_sections,omitempty"`

		HighAvailabilityMode string `json:"high_availability_mode,omitempty"`

		PreferredEdgeClusterMemberIndex *int64 `json:"preferred_edge_cluster_member_index,omitempty"`

		RouterType *string `json:"router_type"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.AdvancedConfig = data.AdvancedConfig

	m.AllocationProfile = data.AllocationProfile

	m.EdgeClusterID = data.EdgeClusterID

	m.EdgeClusterMemberIndices = data.EdgeClusterMemberIndices

	m.FailoverMode = data.FailoverMode

	m.FirewallSections = data.FirewallSections

	m.HighAvailabilityMode = data.HighAvailabilityMode

	m.PreferredEdgeClusterMemberIndex = data.PreferredEdgeClusterMemberIndex

	m.RouterType = data.RouterType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LogicalRouter) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ManagedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		AdvancedConfig *LogicalRouterConfig `json:"advanced_config,omitempty"`

		AllocationProfile *EdgeClusterMemberAllocationProfile `json:"allocation_profile,omitempty"`

		EdgeClusterID string `json:"edge_cluster_id,omitempty"`

		EdgeClusterMemberIndices []int64 `json:"edge_cluster_member_indices,omitempty"`

		FailoverMode string `json:"failover_mode,omitempty"`

		FirewallSections []*ResourceReference `json:"firewall_sections,omitempty"`

		HighAvailabilityMode string `json:"high_availability_mode,omitempty"`

		PreferredEdgeClusterMemberIndex *int64 `json:"preferred_edge_cluster_member_index,omitempty"`

		RouterType *string `json:"router_type"`
	}

	data.AdvancedConfig = m.AdvancedConfig

	data.AllocationProfile = m.AllocationProfile

	data.EdgeClusterID = m.EdgeClusterID

	data.EdgeClusterMemberIndices = m.EdgeClusterMemberIndices

	data.FailoverMode = m.FailoverMode

	data.FirewallSections = m.FirewallSections

	data.HighAvailabilityMode = m.HighAvailabilityMode

	data.PreferredEdgeClusterMemberIndex = m.PreferredEdgeClusterMemberIndex

	data.RouterType = m.RouterType

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this logical router
func (m *LogicalRouter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ManagedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvancedConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocationProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeClusterMemberIndices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailoverMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirewallSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHighAvailabilityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredEdgeClusterMemberIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouterType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogicalRouter) validateAdvancedConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvancedConfig) { // not required
		return nil
	}

	if m.AdvancedConfig != nil {

		if err := m.AdvancedConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("advanced_config")
			}
			return err
		}
	}

	return nil
}

func (m *LogicalRouter) validateAllocationProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocationProfile) { // not required
		return nil
	}

	if m.AllocationProfile != nil {

		if err := m.AllocationProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation_profile")
			}
			return err
		}
	}

	return nil
}

func (m *LogicalRouter) validateEdgeClusterMemberIndices(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeClusterMemberIndices) { // not required
		return nil
	}

	return nil
}

var logicalRouterTypeFailoverModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PREEMPTIVE","NON_PREEMPTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logicalRouterTypeFailoverModePropEnum = append(logicalRouterTypeFailoverModePropEnum, v)
	}
}

// property enum
func (m *LogicalRouter) validateFailoverModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logicalRouterTypeFailoverModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogicalRouter) validateFailoverMode(formats strfmt.Registry) error {

	if swag.IsZero(m.FailoverMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateFailoverModeEnum("failover_mode", "body", m.FailoverMode); err != nil {
		return err
	}

	return nil
}

func (m *LogicalRouter) validateFirewallSections(formats strfmt.Registry) error {

	if swag.IsZero(m.FirewallSections) { // not required
		return nil
	}

	for i := 0; i < len(m.FirewallSections); i++ {

		if swag.IsZero(m.FirewallSections[i]) { // not required
			continue
		}

		if m.FirewallSections[i] != nil {

			if err := m.FirewallSections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("firewall_sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var logicalRouterTypeHighAvailabilityModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE_ACTIVE","ACTIVE_STANDBY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logicalRouterTypeHighAvailabilityModePropEnum = append(logicalRouterTypeHighAvailabilityModePropEnum, v)
	}
}

// property enum
func (m *LogicalRouter) validateHighAvailabilityModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logicalRouterTypeHighAvailabilityModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogicalRouter) validateHighAvailabilityMode(formats strfmt.Registry) error {

	if swag.IsZero(m.HighAvailabilityMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateHighAvailabilityModeEnum("high_availability_mode", "body", m.HighAvailabilityMode); err != nil {
		return err
	}

	return nil
}

func (m *LogicalRouter) validatePreferredEdgeClusterMemberIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.PreferredEdgeClusterMemberIndex) { // not required
		return nil
	}

	if err := validate.MinimumInt("preferred_edge_cluster_member_index", "body", int64(*m.PreferredEdgeClusterMemberIndex), 0, false); err != nil {
		return err
	}

	return nil
}

var logicalRouterTypeRouterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TIER0","TIER1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logicalRouterTypeRouterTypePropEnum = append(logicalRouterTypeRouterTypePropEnum, v)
	}
}

// property enum
func (m *LogicalRouter) validateRouterTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logicalRouterTypeRouterTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogicalRouter) validateRouterType(formats strfmt.Registry) error {

	if err := validate.Required("router_type", "body", m.RouterType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRouterTypeEnum("router_type", "body", *m.RouterType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogicalRouter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogicalRouter) UnmarshalBinary(b []byte) error {
	var res LogicalRouter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

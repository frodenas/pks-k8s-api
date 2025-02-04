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

// LbPool lb pool
// swagger:model LbPool

type LbPool struct {
	ManagedResource

	// active monitor identifier list
	//
	// In case of active healthchecks, load balancer itself initiates new
	// connections (or sends ICMP ping) to the servers periodically to check
	// their health, completely independent of any data traffic. Active
	// healthchecks are disabled by default and can be enabled for a server
	// pool by binding a health monitor to the pool. Currently, only one active
	// health monitor can be configured per server pool.
	//
	ActiveMonitorIds []string `json:"active_monitor_ids"`

	// Load balancing algorithm, configurable per pool controls how the
	// incoming connections are distributed among the members.
	//
	Algorithm *string `json:"algorithm,omitempty"`

	// Load balancer member setting with grouping object
	//
	// Load balancer pool support grouping object as dynamic pool members.
	// When member group is defined, members setting should not be specified.
	//
	MemberGroup *PoolMemberGroup `json:"member_group,omitempty"`

	// load balancer pool members
	//
	// Server pool consists of one or more pool members. Each pool member
	// is identified, typically, by an IP address and a port.
	//
	Members []*PoolMember `json:"members"`

	// minimum number of active pool members to consider pool as active
	//
	// A pool is considered active if there are at least certain
	// minimum number of members.
	//
	// Minimum: 1
	MinActiveMembers int64 `json:"min_active_members,omitempty"`

	// passive monitor identifier
	//
	// Passive healthchecks are disabled by default and can be enabled by
	// attaching a passive health monitor to a server pool.
	// Each time a client connection to a pool member fails, its failed count
	// is incremented. For pools bound to L7 virtual servers, a connection is
	// considered to be failed and failed count is incremented if any TCP
	// connection errors (e.g. TCP RST or failure to send data) or SSL
	// handshake failures occur. For pools bound to L4 virtual servers, if no
	// response is received to a TCP SYN sent to the pool member or if a TCP
	// RST is received in response to a TCP SYN, then the pool member is
	// considered to have failed and the failed count is incremented.
	//
	PassiveMonitorID string `json:"passive_monitor_id,omitempty"`

	// snat translation configuration
	//
	// Depending on the topology, Source NAT (SNAT) may be required to ensure
	// traffic from the server destined to the client is received by the load
	// balancer. SNAT can be enabled per pool. If SNAT is not enabled for a
	// pool, then load balancer uses the client IP and port (spoofing) while
	// establishing connections to the servers. This is referred to as no-SNAT
	// or TRANSPARENT mode.
	//
	SnatTranslation *LbSnatTranslation `json:"snat_translation,omitempty"`

	// TCP multiplexing enable flag
	//
	// TCP multiplexing allows the same TCP connection between load balancer
	// and the backend server to be used for sending multiple client requests
	// from different client TCP connections.
	//
	TCPMultiplexingEnabled *bool `json:"tcp_multiplexing_enabled,omitempty"`

	// maximum number of TCP connections for multiplexing
	//
	// The maximum number of TCP connections per pool that are idly kept alive
	// for sending future client requests.
	//
	// Minimum: 0
	TCPMultiplexingNumber *int64 `json:"tcp_multiplexing_number,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LbPool) UnmarshalJSON(raw []byte) error {

	var aO0 ManagedResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ManagedResource = aO0

	var data struct {
		ActiveMonitorIds []string `json:"active_monitor_ids,omitempty"`

		Algorithm *string `json:"algorithm,omitempty"`

		MemberGroup *PoolMemberGroup `json:"member_group,omitempty"`

		Members []*PoolMember `json:"members,omitempty"`

		MinActiveMembers int64 `json:"min_active_members,omitempty"`

		PassiveMonitorID string `json:"passive_monitor_id,omitempty"`

		SnatTranslation *LbSnatTranslation `json:"snat_translation,omitempty"`

		TCPMultiplexingEnabled *bool `json:"tcp_multiplexing_enabled,omitempty"`

		TCPMultiplexingNumber *int64 `json:"tcp_multiplexing_number,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.ActiveMonitorIds = data.ActiveMonitorIds

	m.Algorithm = data.Algorithm

	m.MemberGroup = data.MemberGroup

	m.Members = data.Members

	m.MinActiveMembers = data.MinActiveMembers

	m.PassiveMonitorID = data.PassiveMonitorID

	m.SnatTranslation = data.SnatTranslation

	m.TCPMultiplexingEnabled = data.TCPMultiplexingEnabled

	m.TCPMultiplexingNumber = data.TCPMultiplexingNumber

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LbPool) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ManagedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		ActiveMonitorIds []string `json:"active_monitor_ids,omitempty"`

		Algorithm *string `json:"algorithm,omitempty"`

		MemberGroup *PoolMemberGroup `json:"member_group,omitempty"`

		Members []*PoolMember `json:"members,omitempty"`

		MinActiveMembers int64 `json:"min_active_members,omitempty"`

		PassiveMonitorID string `json:"passive_monitor_id,omitempty"`

		SnatTranslation *LbSnatTranslation `json:"snat_translation,omitempty"`

		TCPMultiplexingEnabled *bool `json:"tcp_multiplexing_enabled,omitempty"`

		TCPMultiplexingNumber *int64 `json:"tcp_multiplexing_number,omitempty"`
	}

	data.ActiveMonitorIds = m.ActiveMonitorIds

	data.Algorithm = m.Algorithm

	data.MemberGroup = m.MemberGroup

	data.Members = m.Members

	data.MinActiveMembers = m.MinActiveMembers

	data.PassiveMonitorID = m.PassiveMonitorID

	data.SnatTranslation = m.SnatTranslation

	data.TCPMultiplexingEnabled = m.TCPMultiplexingEnabled

	data.TCPMultiplexingNumber = m.TCPMultiplexingNumber

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this lb pool
func (m *LbPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ManagedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveMonitorIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinActiveMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnatTranslation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPMultiplexingNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LbPool) validateActiveMonitorIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ActiveMonitorIds) { // not required
		return nil
	}

	return nil
}

var lbPoolTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ROUND_ROBIN","WEIGHTED_ROUND_ROBIN","LEAST_CONNECTION","IP_HASH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lbPoolTypeAlgorithmPropEnum = append(lbPoolTypeAlgorithmPropEnum, v)
	}
}

// property enum
func (m *LbPool) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, lbPoolTypeAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LbPool) validateAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.Algorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgorithmEnum("algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *LbPool) validateMemberGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberGroup) { // not required
		return nil
	}

	if m.MemberGroup != nil {

		if err := m.MemberGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member_group")
			}
			return err
		}
	}

	return nil
}

func (m *LbPool) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {

		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {

			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LbPool) validateMinActiveMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.MinActiveMembers) { // not required
		return nil
	}

	if err := validate.MinimumInt("min_active_members", "body", int64(m.MinActiveMembers), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *LbPool) validateSnatTranslation(formats strfmt.Registry) error {

	if swag.IsZero(m.SnatTranslation) { // not required
		return nil
	}

	if m.SnatTranslation != nil {

		if err := m.SnatTranslation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snat_translation")
			}
			return err
		}
	}

	return nil
}

func (m *LbPool) validateTCPMultiplexingNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.TCPMultiplexingNumber) { // not required
		return nil
	}

	if err := validate.MinimumInt("tcp_multiplexing_number", "body", int64(*m.TCPMultiplexingNumber), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LbPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LbPool) UnmarshalBinary(b []byte) error {
	var res LbPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

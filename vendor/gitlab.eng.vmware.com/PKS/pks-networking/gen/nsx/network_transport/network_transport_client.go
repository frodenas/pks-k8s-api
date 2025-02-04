// Code generated by go-swagger; DO NOT EDIT.

package network_transport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new network transport API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network transport API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTransportNode gets a transport node

Returns information about a specified transport node.
*/
func (a *Client) GetTransportNode(params *GetTransportNodeParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransportNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransportNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransportNode",
		Method:             "GET",
		PathPattern:        "/transport-nodes/{transportnode-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransportNodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransportNodeOK), nil

}

/*
GetTransportZone gets a transport zone

Returns information about a single transport zone.
*/
func (a *Client) GetTransportZone(params *GetTransportZoneParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransportZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransportZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransportZone",
		Method:             "GET",
		PathPattern:        "/transport-zones/{zone-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransportZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransportZoneOK), nil

}

/*
ListEdgeClusters lists edge clusters

Returns information about the configured edge clusters, which enable you to
group together transport nodes of the type EdgeNode and apply fabric
profiles to all members of the edge cluster. Each edge node can participate
in only one edge cluster.

*/
func (a *Client) ListEdgeClusters(params *ListEdgeClustersParams, authInfo runtime.ClientAuthInfoWriter) (*ListEdgeClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEdgeClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListEdgeClusters",
		Method:             "GET",
		PathPattern:        "/edge-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListEdgeClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEdgeClustersOK), nil

}

/*
ListTransportZones lists transport zones

Returns information about configured transport zones. NSX requires at
least one transport zone. NSX uses transport zones to provide connectivity
based on the topology of the underlying network, trust zones, or
organizational separations. For example, you might have hypervisors that
use one network for management traffic and a different network for VM
traffic. This architecture would require two transport zones. The
combination of transport zones plus transport connectors enables NSX to
form tunnels between hypervisors. Transport zones define which interfaces
on the hypervisors can communicate with which other interfaces on other
hypervisors to establish overlay tunnels or provide connectivity to a VLAN.
A logical switch can be in one (and only one) transport zone. This means
that all of a switch's interfaces must be in the same transport zone.
However, each hypervisor virtual switch (OVS or VDS) has multiple
interfaces (connectors), and each connector can be attached to a different
logical switch. For example, on a single hypervisor with two connectors,
connector A can be attached to logical switch 1 in transport zone A, while
connector B is attached to logical switch 2 in transport zone B. In this
way, a single hypervisor can participate in multiple transport zones. The
API for creating a transport zone requires that a single host switch be
specified for each transport zone, and multiple transport zones can share
the same host switch.

*/
func (a *Client) ListTransportZones(params *ListTransportZonesParams, authInfo runtime.ClientAuthInfoWriter) (*ListTransportZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTransportZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTransportZones",
		Method:             "GET",
		PathPattern:        "/transport-zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTransportZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTransportZonesOK), nil

}

/*
ReadEdgeCluster reads edge cluster

Returns information about the specified edge cluster.
*/
func (a *Client) ReadEdgeCluster(params *ReadEdgeClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ReadEdgeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadEdgeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadEdgeCluster",
		Method:             "GET",
		PathPattern:        "/edge-clusters/{edge-cluster-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadEdgeClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadEdgeClusterOK), nil

}

/*
UpdateTransportZone updates a transport zone

Updates an existing transport zone. Modifiable parameters are
transport_type (VLAN or OVERLAY), description, and display_name. The
request must include the existing host_switch_name.

*/
func (a *Client) UpdateTransportZone(params *UpdateTransportZoneParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTransportZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTransportZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateTransportZone",
		Method:             "PUT",
		PathPattern:        "/transport-zones/{zone-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTransportZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTransportZoneOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

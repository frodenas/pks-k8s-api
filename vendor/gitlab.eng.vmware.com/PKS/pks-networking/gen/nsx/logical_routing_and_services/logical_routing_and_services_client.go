// Code generated by go-swagger; DO NOT EDIT.

package logical_routing_and_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new logical routing and services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for logical routing and services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddNatRule adds a n a t rule in a specific logical router

Add a NAT rule in a specific logical router.

*/
func (a *Client) AddNatRule(params *AddNatRuleParams, authInfo runtime.ClientAuthInfoWriter) (*AddNatRuleCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNatRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddNatRule",
		Method:             "POST",
		PathPattern:        "/logical-routers/{logical-router-id}/nat/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddNatRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddNatRuleCreated), nil

}

/*
CreateLogicalRouter creates a logical router

Creates a logical router. The required parameters are router_type (TIER0 or
TIER1) and edge_cluster_id (TIER0 only). Optional parameters include
internal and external transit network addresses.

*/
func (a *Client) CreateLogicalRouter(params *CreateLogicalRouterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLogicalRouterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLogicalRouterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateLogicalRouter",
		Method:             "POST",
		PathPattern:        "/logical-routers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLogicalRouterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateLogicalRouterCreated), nil

}

/*
CreateLogicalRouterPort creates a logical router port

Creates a logical router port. The required parameters include resource_type
(LogicalRouterUpLinkPort, LogicalRouterDownLinkPort, LogicalRouterLinkPort,
LogicalRouterLoopbackPort); and logical_router_id (the router to which each
logical router port is assigned). The service_bindings parameter is optional.

*/
func (a *Client) CreateLogicalRouterPort(params *CreateLogicalRouterPortParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLogicalRouterPortCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLogicalRouterPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateLogicalRouterPort",
		Method:             "POST",
		PathPattern:        "/logical-router-ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLogicalRouterPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateLogicalRouterPortCreated), nil

}

/*
DeleteLogicalRouter deletes a logical router

Deletes the specified logical router. You must delete associated logical
router ports before you can delete a logical router. Otherwise use force
delete which will delete all related ports and other entities associated
with that LR. To force delete logical router pass force=true in query param.

*/
func (a *Client) DeleteLogicalRouter(params *DeleteLogicalRouterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLogicalRouterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLogicalRouterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLogicalRouter",
		Method:             "DELETE",
		PathPattern:        "/logical-routers/{logical-router-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLogicalRouterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLogicalRouterOK), nil

}

/*
DeleteLogicalRouterPort deletes a logical router port

Deletes the specified logical router port. You must delete logical router
ports before you can delete the associated logical router. To Delete Tier0
router link port you must have to delete attached tier1 router link port,
otherwise pass "force=true" as query param to force delete the Tier0
router link port.

*/
func (a *Client) DeleteLogicalRouterPort(params *DeleteLogicalRouterPortParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLogicalRouterPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLogicalRouterPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLogicalRouterPort",
		Method:             "DELETE",
		PathPattern:        "/logical-router-ports/{logical-router-port-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLogicalRouterPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLogicalRouterPortOK), nil

}

/*
DeleteNatRule deletes a specific n a t rule from a logical router

Delete a specific NAT rule from a logical router

*/
func (a *Client) DeleteNatRule(params *DeleteNatRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNatRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNatRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNatRule",
		Method:             "DELETE",
		PathPattern:        "/logical-routers/{logical-router-id}/nat/rules/{rule-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNatRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNatRuleOK), nil

}

/*
GetNatRule gets a specific n a t rule from a given logical router

Get a specific NAT rule from a given logical router

*/
func (a *Client) GetNatRule(params *GetNatRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetNatRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNatRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNatRule",
		Method:             "GET",
		PathPattern:        "/logical-routers/{logical-router-id}/nat/rules/{rule-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNatRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNatRuleOK), nil

}

/*
ListLogicalRouterPorts lists logical router ports

Returns information about all logical router ports. Information includes the
resource_type (LogicalRouterUpLinkPort, LogicalRouterDownLinkPort,
LogicalRouterLinkPort, LogicalRouterLoopbackPort); logical_router_id
(the router to which each logical router port is assigned);
and any service_bindings (such as DHCP relay service).
The GET request can include a query parameter (logical_router_id
or logical_switch_id).

*/
func (a *Client) ListLogicalRouterPorts(params *ListLogicalRouterPortsParams, authInfo runtime.ClientAuthInfoWriter) (*ListLogicalRouterPortsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLogicalRouterPortsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListLogicalRouterPorts",
		Method:             "GET",
		PathPattern:        "/logical-router-ports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListLogicalRouterPortsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListLogicalRouterPortsOK), nil

}

/*
ListLogicalRouters lists logical routers

Returns information about all logical routers, including the UUID, internal
and external transit network addresses, and the router type (TIER0 or
TIER1). You can get information for only TIER0 routers or only the TIER1
routers by including the router_type query parameter.

*/
func (a *Client) ListLogicalRouters(params *ListLogicalRoutersParams, authInfo runtime.ClientAuthInfoWriter) (*ListLogicalRoutersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLogicalRoutersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListLogicalRouters",
		Method:             "GET",
		PathPattern:        "/logical-routers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListLogicalRoutersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListLogicalRoutersOK), nil

}

/*
ListNatRules lists n a t rules of the logical router

Returns paginated list of all user defined NAT rules of the specific logical router

*/
func (a *Client) ListNatRules(params *ListNatRulesParams, authInfo runtime.ClientAuthInfoWriter) (*ListNatRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNatRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListNatRules",
		Method:             "GET",
		PathPattern:        "/logical-routers/{logical-router-id}/nat/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNatRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNatRulesOK), nil

}

/*
ReadAdvertisementConfig reads the advertisement configuration on a logical router

Returns information about the routes to be advertised by the specified
TIER1 logical router.

*/
func (a *Client) ReadAdvertisementConfig(params *ReadAdvertisementConfigParams, authInfo runtime.ClientAuthInfoWriter) (*ReadAdvertisementConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadAdvertisementConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadAdvertisementConfig",
		Method:             "GET",
		PathPattern:        "/logical-routers/{logical-router-id}/routing/advertisement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadAdvertisementConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadAdvertisementConfigOK), nil

}

/*
ReadLogicalRouter reads logical router

Returns information about the specified logical router.
*/
func (a *Client) ReadLogicalRouter(params *ReadLogicalRouterParams, authInfo runtime.ClientAuthInfoWriter) (*ReadLogicalRouterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadLogicalRouterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadLogicalRouter",
		Method:             "GET",
		PathPattern:        "/logical-routers/{logical-router-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadLogicalRouterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadLogicalRouterOK), nil

}

/*
ReadLogicalRouterPort reads logical router port

Returns information about the specified logical router port.
*/
func (a *Client) ReadLogicalRouterPort(params *ReadLogicalRouterPortParams, authInfo runtime.ClientAuthInfoWriter) (*ReadLogicalRouterPortOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadLogicalRouterPortParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReadLogicalRouterPort",
		Method:             "GET",
		PathPattern:        "/logical-router-ports/{logical-router-port-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReadLogicalRouterPortReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReadLogicalRouterPortOK), nil

}

/*
UpdateAdvertisementConfig updates the advertisement configuration on a logical router

Modifies the route advertisement configuration on the specified logical router.

*/
func (a *Client) UpdateAdvertisementConfig(params *UpdateAdvertisementConfigParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAdvertisementConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdvertisementConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateAdvertisementConfig",
		Method:             "PUT",
		PathPattern:        "/logical-routers/{logical-router-id}/routing/advertisement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdvertisementConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAdvertisementConfigOK), nil

}

/*
UpdateLogicalRouter updates a logical router

Modifies the specified logical router. Modifiable attributes include the
internal_transit_network, external_transit_networks, and edge_cluster_id
(for TIER0 routers).

*/
func (a *Client) UpdateLogicalRouter(params *UpdateLogicalRouterParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLogicalRouterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLogicalRouterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateLogicalRouter",
		Method:             "PUT",
		PathPattern:        "/logical-routers/{logical-router-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLogicalRouterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateLogicalRouterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SearchByTag searches for resources by tag

Search for resources by tags
*/
func (a *Client) SearchByTag(params *SearchByTagParams, authInfo runtime.ClientAuthInfoWriter) (*SearchByTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchByTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SearchByTag",
		Method:             "GET",
		PathPattern:        "/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchByTagReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchByTagOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

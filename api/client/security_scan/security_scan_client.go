// Code generated by go-swagger; DO NOT EDIT.

package security_scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new security scan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security scan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSecurityScanResult adds a new security scan result

Adds a new security scan result into Testra app
*/
func (a *Client) CreateSecurityScanResult(params *CreateSecurityScanResultParams) (*CreateSecurityScanResultCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSecurityScanResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSecurityScanResult",
		Method:             "POST",
		PathPattern:        "/projects/{projectId}/executions/{executionId}/scan-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSecurityScanResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSecurityScanResultCreated), nil

}

/*
GetSecurityScanResult gets security scan results

Returns security scan result for the given project and execution

*/
func (a *Client) GetSecurityScanResult(params *GetSecurityScanResultParams) (*GetSecurityScanResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSecurityScanResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSecurityScanResult",
		Method:             "GET",
		PathPattern:        "/projects/{projectId}/executions/{executionId}/scan-results",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSecurityScanResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSecurityScanResultOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

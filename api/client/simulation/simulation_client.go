// Code generated by go-swagger; DO NOT EDIT.

package simulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new simulation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for simulation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSimulation creates a new simulation

Adds a new simulation into Testra app. It takes a JSON object containing a name that was not used before.
*/
func (a *Client) CreateSimulation(params *CreateSimulationParams) (*CreateSimulationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSimulationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSimulation",
		Method:             "POST",
		PathPattern:        "/projects/{projectId}/executions/{executionId}/simulations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSimulationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSimulationCreated), nil

}

/*
DeleteSimulation deletes simulation info
*/
func (a *Client) DeleteSimulation(params *DeleteSimulationParams) (*DeleteSimulationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSimulationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSimulation",
		Method:             "DELETE",
		PathPattern:        "/projects/{projectId}/executions/{executionId}/simulation/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSimulationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSimulationOK), nil

}

/*
GetSimulations lists of all simulations

Returns list of all simulations within given project and execution

*/
func (a *Client) GetSimulations(params *GetSimulationsParams) (*GetSimulationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSimulationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSimulations",
		Method:             "GET",
		PathPattern:        "/projects/{projectId}/executions/{executionId}/simulations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSimulationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSimulationsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
// Code generated by go-swagger; DO NOT EDIT.

package simulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSimulationParams creates a new DeleteSimulationParams object
// with the default values initialized.
func NewDeleteSimulationParams() *DeleteSimulationParams {
	var ()
	return &DeleteSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSimulationParamsWithTimeout creates a new DeleteSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSimulationParamsWithTimeout(timeout time.Duration) *DeleteSimulationParams {
	var ()
	return &DeleteSimulationParams{

		timeout: timeout,
	}
}

// NewDeleteSimulationParamsWithContext creates a new DeleteSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSimulationParamsWithContext(ctx context.Context) *DeleteSimulationParams {
	var ()
	return &DeleteSimulationParams{

		Context: ctx,
	}
}

// NewDeleteSimulationParamsWithHTTPClient creates a new DeleteSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSimulationParamsWithHTTPClient(client *http.Client) *DeleteSimulationParams {
	var ()
	return &DeleteSimulationParams{
		HTTPClient: client,
	}
}

/*DeleteSimulationParams contains all the parameters to send to the API endpoint
for the delete simulation operation typically these are written to a http.Request
*/
type DeleteSimulationParams struct {

	/*ExecutionID
	  Execution Id

	*/
	ExecutionID string
	/*ID
	  Simulation Id

	*/
	ID string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete simulation params
func (o *DeleteSimulationParams) WithTimeout(timeout time.Duration) *DeleteSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete simulation params
func (o *DeleteSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete simulation params
func (o *DeleteSimulationParams) WithContext(ctx context.Context) *DeleteSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete simulation params
func (o *DeleteSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete simulation params
func (o *DeleteSimulationParams) WithHTTPClient(client *http.Client) *DeleteSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete simulation params
func (o *DeleteSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the delete simulation params
func (o *DeleteSimulationParams) WithExecutionID(executionID string) *DeleteSimulationParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the delete simulation params
func (o *DeleteSimulationParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithID adds the id to the delete simulation params
func (o *DeleteSimulationParams) WithID(id string) *DeleteSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete simulation params
func (o *DeleteSimulationParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete simulation params
func (o *DeleteSimulationParams) WithProjectID(projectID string) *DeleteSimulationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete simulation params
func (o *DeleteSimulationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

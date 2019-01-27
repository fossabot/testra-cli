// Code generated by go-swagger; DO NOT EDIT.

package scenario

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

// NewDeleteScenarioParams creates a new DeleteScenarioParams object
// with the default values initialized.
func NewDeleteScenarioParams() *DeleteScenarioParams {
	var ()
	return &DeleteScenarioParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScenarioParamsWithTimeout creates a new DeleteScenarioParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScenarioParamsWithTimeout(timeout time.Duration) *DeleteScenarioParams {
	var ()
	return &DeleteScenarioParams{

		timeout: timeout,
	}
}

// NewDeleteScenarioParamsWithContext creates a new DeleteScenarioParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScenarioParamsWithContext(ctx context.Context) *DeleteScenarioParams {
	var ()
	return &DeleteScenarioParams{

		Context: ctx,
	}
}

// NewDeleteScenarioParamsWithHTTPClient creates a new DeleteScenarioParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScenarioParamsWithHTTPClient(client *http.Client) *DeleteScenarioParams {
	var ()
	return &DeleteScenarioParams{
		HTTPClient: client,
	}
}

/*DeleteScenarioParams contains all the parameters to send to the API endpoint
for the delete scenario operation typically these are written to a http.Request
*/
type DeleteScenarioParams struct {

	/*ID
	  Scenario Id

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

// WithTimeout adds the timeout to the delete scenario params
func (o *DeleteScenarioParams) WithTimeout(timeout time.Duration) *DeleteScenarioParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scenario params
func (o *DeleteScenarioParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scenario params
func (o *DeleteScenarioParams) WithContext(ctx context.Context) *DeleteScenarioParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scenario params
func (o *DeleteScenarioParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scenario params
func (o *DeleteScenarioParams) WithHTTPClient(client *http.Client) *DeleteScenarioParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scenario params
func (o *DeleteScenarioParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete scenario params
func (o *DeleteScenarioParams) WithID(id string) *DeleteScenarioParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete scenario params
func (o *DeleteScenarioParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete scenario params
func (o *DeleteScenarioParams) WithProjectID(projectID string) *DeleteScenarioParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete scenario params
func (o *DeleteScenarioParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScenarioParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

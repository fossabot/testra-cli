// Code generated by go-swagger; DO NOT EDIT.

package execution

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

// NewDeleteExecutionParams creates a new DeleteExecutionParams object
// with the default values initialized.
func NewDeleteExecutionParams() *DeleteExecutionParams {
	var ()
	return &DeleteExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExecutionParamsWithTimeout creates a new DeleteExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteExecutionParamsWithTimeout(timeout time.Duration) *DeleteExecutionParams {
	var ()
	return &DeleteExecutionParams{

		timeout: timeout,
	}
}

// NewDeleteExecutionParamsWithContext creates a new DeleteExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteExecutionParamsWithContext(ctx context.Context) *DeleteExecutionParams {
	var ()
	return &DeleteExecutionParams{

		Context: ctx,
	}
}

// NewDeleteExecutionParamsWithHTTPClient creates a new DeleteExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteExecutionParamsWithHTTPClient(client *http.Client) *DeleteExecutionParams {
	var ()
	return &DeleteExecutionParams{
		HTTPClient: client,
	}
}

/*DeleteExecutionParams contains all the parameters to send to the API endpoint
for the delete execution operation typically these are written to a http.Request
*/
type DeleteExecutionParams struct {

	/*ID
	  Execution Id

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

// WithTimeout adds the timeout to the delete execution params
func (o *DeleteExecutionParams) WithTimeout(timeout time.Duration) *DeleteExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete execution params
func (o *DeleteExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete execution params
func (o *DeleteExecutionParams) WithContext(ctx context.Context) *DeleteExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete execution params
func (o *DeleteExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete execution params
func (o *DeleteExecutionParams) WithHTTPClient(client *http.Client) *DeleteExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete execution params
func (o *DeleteExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete execution params
func (o *DeleteExecutionParams) WithID(id string) *DeleteExecutionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete execution params
func (o *DeleteExecutionParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete execution params
func (o *DeleteExecutionParams) WithProjectID(projectID string) *DeleteExecutionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete execution params
func (o *DeleteExecutionParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

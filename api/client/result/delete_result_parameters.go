// Code generated by go-swagger; DO NOT EDIT.

package result

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

// NewDeleteResultParams creates a new DeleteResultParams object
// with the default values initialized.
func NewDeleteResultParams() *DeleteResultParams {
	var ()
	return &DeleteResultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResultParamsWithTimeout creates a new DeleteResultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteResultParamsWithTimeout(timeout time.Duration) *DeleteResultParams {
	var ()
	return &DeleteResultParams{

		timeout: timeout,
	}
}

// NewDeleteResultParamsWithContext creates a new DeleteResultParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteResultParamsWithContext(ctx context.Context) *DeleteResultParams {
	var ()
	return &DeleteResultParams{

		Context: ctx,
	}
}

// NewDeleteResultParamsWithHTTPClient creates a new DeleteResultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteResultParamsWithHTTPClient(client *http.Client) *DeleteResultParams {
	var ()
	return &DeleteResultParams{
		HTTPClient: client,
	}
}

/*DeleteResultParams contains all the parameters to send to the API endpoint
for the delete result operation typically these are written to a http.Request
*/
type DeleteResultParams struct {

	/*ExecutionID
	  Execution Id

	*/
	ExecutionID string
	/*ID
	  Result Id

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

// WithTimeout adds the timeout to the delete result params
func (o *DeleteResultParams) WithTimeout(timeout time.Duration) *DeleteResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete result params
func (o *DeleteResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete result params
func (o *DeleteResultParams) WithContext(ctx context.Context) *DeleteResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete result params
func (o *DeleteResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete result params
func (o *DeleteResultParams) WithHTTPClient(client *http.Client) *DeleteResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete result params
func (o *DeleteResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the delete result params
func (o *DeleteResultParams) WithExecutionID(executionID string) *DeleteResultParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the delete result params
func (o *DeleteResultParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithID adds the id to the delete result params
func (o *DeleteResultParams) WithID(id string) *DeleteResultParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete result params
func (o *DeleteResultParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete result params
func (o *DeleteResultParams) WithProjectID(projectID string) *DeleteResultParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete result params
func (o *DeleteResultParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

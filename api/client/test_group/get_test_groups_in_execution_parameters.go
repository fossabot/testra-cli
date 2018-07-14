// Code generated by go-swagger; DO NOT EDIT.

package test_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTestGroupsInExecutionParams creates a new GetTestGroupsInExecutionParams object
// with the default values initialized.
func NewGetTestGroupsInExecutionParams() *GetTestGroupsInExecutionParams {
	var ()
	return &GetTestGroupsInExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestGroupsInExecutionParamsWithTimeout creates a new GetTestGroupsInExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestGroupsInExecutionParamsWithTimeout(timeout time.Duration) *GetTestGroupsInExecutionParams {
	var ()
	return &GetTestGroupsInExecutionParams{

		timeout: timeout,
	}
}

// NewGetTestGroupsInExecutionParamsWithContext creates a new GetTestGroupsInExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestGroupsInExecutionParamsWithContext(ctx context.Context) *GetTestGroupsInExecutionParams {
	var ()
	return &GetTestGroupsInExecutionParams{

		Context: ctx,
	}
}

// NewGetTestGroupsInExecutionParamsWithHTTPClient creates a new GetTestGroupsInExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestGroupsInExecutionParamsWithHTTPClient(client *http.Client) *GetTestGroupsInExecutionParams {
	var ()
	return &GetTestGroupsInExecutionParams{
		HTTPClient: client,
	}
}

/*GetTestGroupsInExecutionParams contains all the parameters to send to the API endpoint
for the get test groups in execution operation typically these are written to a http.Request
*/
type GetTestGroupsInExecutionParams struct {

	/*ExecutionID
	  Execution Id

	*/
	ExecutionID string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) WithTimeout(timeout time.Duration) *GetTestGroupsInExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) WithContext(ctx context.Context) *GetTestGroupsInExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) WithHTTPClient(client *http.Client) *GetTestGroupsInExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) WithExecutionID(executionID string) *GetTestGroupsInExecutionParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithProjectID adds the projectID to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) WithProjectID(projectID string) *GetTestGroupsInExecutionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get test groups in execution params
func (o *GetTestGroupsInExecutionParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestGroupsInExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param executionId
	if err := r.SetPathParam("executionId", o.ExecutionID); err != nil {
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
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

// NewGetResultsParams creates a new GetResultsParams object
// with the default values initialized.
func NewGetResultsParams() *GetResultsParams {
	var ()
	return &GetResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResultsParamsWithTimeout creates a new GetResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResultsParamsWithTimeout(timeout time.Duration) *GetResultsParams {
	var ()
	return &GetResultsParams{

		timeout: timeout,
	}
}

// NewGetResultsParamsWithContext creates a new GetResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResultsParamsWithContext(ctx context.Context) *GetResultsParams {
	var ()
	return &GetResultsParams{

		Context: ctx,
	}
}

// NewGetResultsParamsWithHTTPClient creates a new GetResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResultsParamsWithHTTPClient(client *http.Client) *GetResultsParams {
	var ()
	return &GetResultsParams{
		HTTPClient: client,
	}
}

/*GetResultsParams contains all the parameters to send to the API endpoint
for the get results operation typically these are written to a http.Request
*/
type GetResultsParams struct {

	/*ExecutionID
	  Execution Id

	*/
	ExecutionID string
	/*ProjectID
	  Project Id

	*/
	ProjectID string
	/*Status
	  Filter test results by the given result status

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get results params
func (o *GetResultsParams) WithTimeout(timeout time.Duration) *GetResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get results params
func (o *GetResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get results params
func (o *GetResultsParams) WithContext(ctx context.Context) *GetResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get results params
func (o *GetResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get results params
func (o *GetResultsParams) WithHTTPClient(client *http.Client) *GetResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get results params
func (o *GetResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get results params
func (o *GetResultsParams) WithExecutionID(executionID string) *GetResultsParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get results params
func (o *GetResultsParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithProjectID adds the projectID to the get results params
func (o *GetResultsParams) WithProjectID(projectID string) *GetResultsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get results params
func (o *GetResultsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithStatus adds the status to the get results params
func (o *GetResultsParams) WithStatus(status *string) *GetResultsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get results params
func (o *GetResultsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

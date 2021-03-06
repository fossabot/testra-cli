// Code generated by go-swagger; DO NOT EDIT.

package security_scan

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

	models "github.com/testra/testra-cli/api/models"
)

// NewCreateSecurityScanResultParams creates a new CreateSecurityScanResultParams object
// with the default values initialized.
func NewCreateSecurityScanResultParams() *CreateSecurityScanResultParams {
	var ()
	return &CreateSecurityScanResultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecurityScanResultParamsWithTimeout creates a new CreateSecurityScanResultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSecurityScanResultParamsWithTimeout(timeout time.Duration) *CreateSecurityScanResultParams {
	var ()
	return &CreateSecurityScanResultParams{

		timeout: timeout,
	}
}

// NewCreateSecurityScanResultParamsWithContext creates a new CreateSecurityScanResultParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSecurityScanResultParamsWithContext(ctx context.Context) *CreateSecurityScanResultParams {
	var ()
	return &CreateSecurityScanResultParams{

		Context: ctx,
	}
}

// NewCreateSecurityScanResultParamsWithHTTPClient creates a new CreateSecurityScanResultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSecurityScanResultParamsWithHTTPClient(client *http.Client) *CreateSecurityScanResultParams {
	var ()
	return &CreateSecurityScanResultParams{
		HTTPClient: client,
	}
}

/*CreateSecurityScanResultParams contains all the parameters to send to the API endpoint
for the create security scan result operation typically these are written to a http.Request
*/
type CreateSecurityScanResultParams struct {

	/*Body*/
	Body *models.ScanResultRequest
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

// WithTimeout adds the timeout to the create security scan result params
func (o *CreateSecurityScanResultParams) WithTimeout(timeout time.Duration) *CreateSecurityScanResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create security scan result params
func (o *CreateSecurityScanResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create security scan result params
func (o *CreateSecurityScanResultParams) WithContext(ctx context.Context) *CreateSecurityScanResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create security scan result params
func (o *CreateSecurityScanResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create security scan result params
func (o *CreateSecurityScanResultParams) WithHTTPClient(client *http.Client) *CreateSecurityScanResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create security scan result params
func (o *CreateSecurityScanResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create security scan result params
func (o *CreateSecurityScanResultParams) WithBody(body *models.ScanResultRequest) *CreateSecurityScanResultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create security scan result params
func (o *CreateSecurityScanResultParams) SetBody(body *models.ScanResultRequest) {
	o.Body = body
}

// WithExecutionID adds the executionID to the create security scan result params
func (o *CreateSecurityScanResultParams) WithExecutionID(executionID string) *CreateSecurityScanResultParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the create security scan result params
func (o *CreateSecurityScanResultParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithProjectID adds the projectID to the create security scan result params
func (o *CreateSecurityScanResultParams) WithProjectID(projectID string) *CreateSecurityScanResultParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create security scan result params
func (o *CreateSecurityScanResultParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecurityScanResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

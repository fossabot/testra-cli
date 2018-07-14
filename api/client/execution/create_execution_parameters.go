// Code generated by go-swagger; DO NOT EDIT.

package execution

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

	models "github.com/testra-tech/testra-cli/api/models"
)

// NewCreateExecutionParams creates a new CreateExecutionParams object
// with the default values initialized.
func NewCreateExecutionParams() *CreateExecutionParams {
	var ()
	return &CreateExecutionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExecutionParamsWithTimeout creates a new CreateExecutionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateExecutionParamsWithTimeout(timeout time.Duration) *CreateExecutionParams {
	var ()
	return &CreateExecutionParams{

		timeout: timeout,
	}
}

// NewCreateExecutionParamsWithContext creates a new CreateExecutionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateExecutionParamsWithContext(ctx context.Context) *CreateExecutionParams {
	var ()
	return &CreateExecutionParams{

		Context: ctx,
	}
}

// NewCreateExecutionParamsWithHTTPClient creates a new CreateExecutionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateExecutionParamsWithHTTPClient(client *http.Client) *CreateExecutionParams {
	var ()
	return &CreateExecutionParams{
		HTTPClient: client,
	}
}

/*CreateExecutionParams contains all the parameters to send to the API endpoint
for the create execution operation typically these are written to a http.Request
*/
type CreateExecutionParams struct {

	/*Body*/
	Body *models.ExecutionRequest
	/*ProjectID
	  Execution Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create execution params
func (o *CreateExecutionParams) WithTimeout(timeout time.Duration) *CreateExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create execution params
func (o *CreateExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create execution params
func (o *CreateExecutionParams) WithContext(ctx context.Context) *CreateExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create execution params
func (o *CreateExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create execution params
func (o *CreateExecutionParams) WithHTTPClient(client *http.Client) *CreateExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create execution params
func (o *CreateExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create execution params
func (o *CreateExecutionParams) WithBody(body *models.ExecutionRequest) *CreateExecutionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create execution params
func (o *CreateExecutionParams) SetBody(body *models.ExecutionRequest) {
	o.Body = body
}

// WithProjectID adds the projectID to the create execution params
func (o *CreateExecutionParams) WithProjectID(projectID string) *CreateExecutionParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create execution params
func (o *CreateExecutionParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
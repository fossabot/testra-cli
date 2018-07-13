// Code generated by go-swagger; DO NOT EDIT.

package testcase

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

// NewCreateTestcaseParams creates a new CreateTestcaseParams object
// with the default values initialized.
func NewCreateTestcaseParams() *CreateTestcaseParams {
	var ()
	return &CreateTestcaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTestcaseParamsWithTimeout creates a new CreateTestcaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTestcaseParamsWithTimeout(timeout time.Duration) *CreateTestcaseParams {
	var ()
	return &CreateTestcaseParams{

		timeout: timeout,
	}
}

// NewCreateTestcaseParamsWithContext creates a new CreateTestcaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTestcaseParamsWithContext(ctx context.Context) *CreateTestcaseParams {
	var ()
	return &CreateTestcaseParams{

		Context: ctx,
	}
}

// NewCreateTestcaseParamsWithHTTPClient creates a new CreateTestcaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTestcaseParamsWithHTTPClient(client *http.Client) *CreateTestcaseParams {
	var ()
	return &CreateTestcaseParams{
		HTTPClient: client,
	}
}

/*CreateTestcaseParams contains all the parameters to send to the API endpoint
for the create testcase operation typically these are written to a http.Request
*/
type CreateTestcaseParams struct {

	/*Body*/
	Body *models.TestcaseRequest
	/*ProjectID
	  Testcase Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create testcase params
func (o *CreateTestcaseParams) WithTimeout(timeout time.Duration) *CreateTestcaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create testcase params
func (o *CreateTestcaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create testcase params
func (o *CreateTestcaseParams) WithContext(ctx context.Context) *CreateTestcaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create testcase params
func (o *CreateTestcaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create testcase params
func (o *CreateTestcaseParams) WithHTTPClient(client *http.Client) *CreateTestcaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create testcase params
func (o *CreateTestcaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create testcase params
func (o *CreateTestcaseParams) WithBody(body *models.TestcaseRequest) *CreateTestcaseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create testcase params
func (o *CreateTestcaseParams) SetBody(body *models.TestcaseRequest) {
	o.Body = body
}

// WithProjectID adds the projectID to the create testcase params
func (o *CreateTestcaseParams) WithProjectID(projectID string) *CreateTestcaseParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create testcase params
func (o *CreateTestcaseParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTestcaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

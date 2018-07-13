// Code generated by go-swagger; DO NOT EDIT.

package scenario

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

// NewGetScenarioParams creates a new GetScenarioParams object
// with the default values initialized.
func NewGetScenarioParams() *GetScenarioParams {
	var ()
	return &GetScenarioParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScenarioParamsWithTimeout creates a new GetScenarioParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScenarioParamsWithTimeout(timeout time.Duration) *GetScenarioParams {
	var ()
	return &GetScenarioParams{

		timeout: timeout,
	}
}

// NewGetScenarioParamsWithContext creates a new GetScenarioParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScenarioParamsWithContext(ctx context.Context) *GetScenarioParams {
	var ()
	return &GetScenarioParams{

		Context: ctx,
	}
}

// NewGetScenarioParamsWithHTTPClient creates a new GetScenarioParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScenarioParamsWithHTTPClient(client *http.Client) *GetScenarioParams {
	var ()
	return &GetScenarioParams{
		HTTPClient: client,
	}
}

/*GetScenarioParams contains all the parameters to send to the API endpoint
for the get scenario operation typically these are written to a http.Request
*/
type GetScenarioParams struct {

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

// WithTimeout adds the timeout to the get scenario params
func (o *GetScenarioParams) WithTimeout(timeout time.Duration) *GetScenarioParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scenario params
func (o *GetScenarioParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scenario params
func (o *GetScenarioParams) WithContext(ctx context.Context) *GetScenarioParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scenario params
func (o *GetScenarioParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scenario params
func (o *GetScenarioParams) WithHTTPClient(client *http.Client) *GetScenarioParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scenario params
func (o *GetScenarioParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get scenario params
func (o *GetScenarioParams) WithID(id string) *GetScenarioParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get scenario params
func (o *GetScenarioParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the get scenario params
func (o *GetScenarioParams) WithProjectID(projectID string) *GetScenarioParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get scenario params
func (o *GetScenarioParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScenarioParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

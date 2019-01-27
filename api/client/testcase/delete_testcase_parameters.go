// Code generated by go-swagger; DO NOT EDIT.

package testcase

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

// NewDeleteTestcaseParams creates a new DeleteTestcaseParams object
// with the default values initialized.
func NewDeleteTestcaseParams() *DeleteTestcaseParams {
	var ()
	return &DeleteTestcaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTestcaseParamsWithTimeout creates a new DeleteTestcaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTestcaseParamsWithTimeout(timeout time.Duration) *DeleteTestcaseParams {
	var ()
	return &DeleteTestcaseParams{

		timeout: timeout,
	}
}

// NewDeleteTestcaseParamsWithContext creates a new DeleteTestcaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTestcaseParamsWithContext(ctx context.Context) *DeleteTestcaseParams {
	var ()
	return &DeleteTestcaseParams{

		Context: ctx,
	}
}

// NewDeleteTestcaseParamsWithHTTPClient creates a new DeleteTestcaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTestcaseParamsWithHTTPClient(client *http.Client) *DeleteTestcaseParams {
	var ()
	return &DeleteTestcaseParams{
		HTTPClient: client,
	}
}

/*DeleteTestcaseParams contains all the parameters to send to the API endpoint
for the delete testcase operation typically these are written to a http.Request
*/
type DeleteTestcaseParams struct {

	/*ID
	  Testcase Id

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

// WithTimeout adds the timeout to the delete testcase params
func (o *DeleteTestcaseParams) WithTimeout(timeout time.Duration) *DeleteTestcaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete testcase params
func (o *DeleteTestcaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete testcase params
func (o *DeleteTestcaseParams) WithContext(ctx context.Context) *DeleteTestcaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete testcase params
func (o *DeleteTestcaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete testcase params
func (o *DeleteTestcaseParams) WithHTTPClient(client *http.Client) *DeleteTestcaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete testcase params
func (o *DeleteTestcaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete testcase params
func (o *DeleteTestcaseParams) WithID(id string) *DeleteTestcaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete testcase params
func (o *DeleteTestcaseParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete testcase params
func (o *DeleteTestcaseParams) WithProjectID(projectID string) *DeleteTestcaseParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete testcase params
func (o *DeleteTestcaseParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTestcaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

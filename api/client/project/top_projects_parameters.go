// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTopProjectsParams creates a new TopProjectsParams object
// with the default values initialized.
func NewTopProjectsParams() *TopProjectsParams {
	var ()
	return &TopProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTopProjectsParamsWithTimeout creates a new TopProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTopProjectsParamsWithTimeout(timeout time.Duration) *TopProjectsParams {
	var ()
	return &TopProjectsParams{

		timeout: timeout,
	}
}

// NewTopProjectsParamsWithContext creates a new TopProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTopProjectsParamsWithContext(ctx context.Context) *TopProjectsParams {
	var ()
	return &TopProjectsParams{

		Context: ctx,
	}
}

// NewTopProjectsParamsWithHTTPClient creates a new TopProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTopProjectsParamsWithHTTPClient(client *http.Client) *TopProjectsParams {
	var ()
	return &TopProjectsParams{
		HTTPClient: client,
	}
}

/*TopProjectsParams contains all the parameters to send to the API endpoint
for the top projects operation typically these are written to a http.Request
*/
type TopProjectsParams struct {

	/*Size
	  No of projects to retrieve

	*/
	Size int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the top projects params
func (o *TopProjectsParams) WithTimeout(timeout time.Duration) *TopProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top projects params
func (o *TopProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top projects params
func (o *TopProjectsParams) WithContext(ctx context.Context) *TopProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top projects params
func (o *TopProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top projects params
func (o *TopProjectsParams) WithHTTPClient(client *http.Client) *TopProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top projects params
func (o *TopProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSize adds the size to the top projects params
func (o *TopProjectsParams) WithSize(size int64) *TopProjectsParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the top projects params
func (o *TopProjectsParams) SetSize(size int64) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *TopProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param size
	qrSize := o.Size
	qSize := swag.FormatInt64(qrSize)
	if qSize != "" {
		if err := r.SetQueryParam("size", qSize); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

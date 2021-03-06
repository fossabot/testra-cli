// Code generated by go-swagger; DO NOT EDIT.

package scenario

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

// NewGetScenariosParams creates a new GetScenariosParams object
// with the default values initialized.
func NewGetScenariosParams() *GetScenariosParams {
	var ()
	return &GetScenariosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScenariosParamsWithTimeout creates a new GetScenariosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScenariosParamsWithTimeout(timeout time.Duration) *GetScenariosParams {
	var ()
	return &GetScenariosParams{

		timeout: timeout,
	}
}

// NewGetScenariosParamsWithContext creates a new GetScenariosParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScenariosParamsWithContext(ctx context.Context) *GetScenariosParams {
	var ()
	return &GetScenariosParams{

		Context: ctx,
	}
}

// NewGetScenariosParamsWithHTTPClient creates a new GetScenariosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScenariosParamsWithHTTPClient(client *http.Client) *GetScenariosParams {
	var ()
	return &GetScenariosParams{
		HTTPClient: client,
	}
}

/*GetScenariosParams contains all the parameters to send to the API endpoint
for the get scenarios operation typically these are written to a http.Request
*/
type GetScenariosParams struct {

	/*FeatureID
	  Feature Id

	*/
	FeatureID *string
	/*ProjectID
	  Project Id

	*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scenarios params
func (o *GetScenariosParams) WithTimeout(timeout time.Duration) *GetScenariosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scenarios params
func (o *GetScenariosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scenarios params
func (o *GetScenariosParams) WithContext(ctx context.Context) *GetScenariosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scenarios params
func (o *GetScenariosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scenarios params
func (o *GetScenariosParams) WithHTTPClient(client *http.Client) *GetScenariosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scenarios params
func (o *GetScenariosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureID adds the featureID to the get scenarios params
func (o *GetScenariosParams) WithFeatureID(featureID *string) *GetScenariosParams {
	o.SetFeatureID(featureID)
	return o
}

// SetFeatureID adds the featureId to the get scenarios params
func (o *GetScenariosParams) SetFeatureID(featureID *string) {
	o.FeatureID = featureID
}

// WithProjectID adds the projectID to the get scenarios params
func (o *GetScenariosParams) WithProjectID(projectID string) *GetScenariosParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get scenarios params
func (o *GetScenariosParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScenariosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FeatureID != nil {

		// query param featureId
		var qrFeatureID string
		if o.FeatureID != nil {
			qrFeatureID = *o.FeatureID
		}
		qFeatureID := qrFeatureID
		if qFeatureID != "" {
			if err := r.SetQueryParam("featureId", qFeatureID); err != nil {
				return err
			}
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

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TestcaseRequest TestCaseRequest
// swagger:model TestcaseRequest
type TestcaseRequest struct {

	// class name
	// Required: true
	ClassName *string `json:"className"`

	// manual
	Manual bool `json:"manual,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this testcase request
func (m *TestcaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestcaseRequest) validateClassName(formats strfmt.Registry) error {

	if err := validate.Required("className", "body", m.ClassName); err != nil {
		return err
	}

	return nil
}

func (m *TestcaseRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TestcaseRequest) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestcaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestcaseRequest) UnmarshalBinary(b []byte) error {
	var res TestcaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

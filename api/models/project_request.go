// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectRequest Project Request
// swagger:model ProjectRequest
type ProjectRequest struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// project type
	// Enum: [TEST SIMULATION SECURITY]
	ProjectType string `json:"projectType,omitempty"`
}

// Validate validates this project request
func (m *ProjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ProjectRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var projectRequestTypeProjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TEST","SIMULATION","SECURITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectRequestTypeProjectTypePropEnum = append(projectRequestTypeProjectTypePropEnum, v)
	}
}

const (

	// ProjectRequestProjectTypeTEST captures enum value "TEST"
	ProjectRequestProjectTypeTEST string = "TEST"

	// ProjectRequestProjectTypeSIMULATION captures enum value "SIMULATION"
	ProjectRequestProjectTypeSIMULATION string = "SIMULATION"

	// ProjectRequestProjectTypeSECURITY captures enum value "SECURITY"
	ProjectRequestProjectTypeSECURITY string = "SECURITY"
)

// prop value enum
func (m *ProjectRequest) validateProjectTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, projectRequestTypeProjectTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProjectRequest) validateProjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProjectTypeEnum("projectType", "body", m.ProjectType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectRequest) UnmarshalBinary(b []byte) error {
	var res ProjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

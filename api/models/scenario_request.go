// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScenarioRequest Scenario Request
// swagger:model ScenarioRequest
type ScenarioRequest struct {

	// after
	After []*TestStep `json:"after"`

	// background steps
	BackgroundSteps []*TestStep `json:"backgroundSteps"`

	// before
	Before []*TestStep `json:"before"`

	// feature description
	// Required: true
	FeatureDescription *string `json:"featureDescription"`

	// feature name
	// Required: true
	FeatureName *string `json:"featureName"`

	// manual
	Manual bool `json:"manual,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// steps
	// Required: true
	Steps []*TestStep `json:"steps"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this scenario request
func (m *ScenarioRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackgroundSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScenarioRequest) validateAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.After) { // not required
		return nil
	}

	for i := 0; i < len(m.After); i++ {
		if swag.IsZero(m.After[i]) { // not required
			continue
		}

		if m.After[i] != nil {
			if err := m.After[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("after" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScenarioRequest) validateBackgroundSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.BackgroundSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.BackgroundSteps); i++ {
		if swag.IsZero(m.BackgroundSteps[i]) { // not required
			continue
		}

		if m.BackgroundSteps[i] != nil {
			if err := m.BackgroundSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backgroundSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScenarioRequest) validateBefore(formats strfmt.Registry) error {

	if swag.IsZero(m.Before) { // not required
		return nil
	}

	for i := 0; i < len(m.Before); i++ {
		if swag.IsZero(m.Before[i]) { // not required
			continue
		}

		if m.Before[i] != nil {
			if err := m.Before[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("before" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScenarioRequest) validateFeatureDescription(formats strfmt.Registry) error {

	if err := validate.Required("featureDescription", "body", m.FeatureDescription); err != nil {
		return err
	}

	return nil
}

func (m *ScenarioRequest) validateFeatureName(formats strfmt.Registry) error {

	if err := validate.Required("featureName", "body", m.FeatureName); err != nil {
		return err
	}

	return nil
}

func (m *ScenarioRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ScenarioRequest) validateSteps(formats strfmt.Registry) error {

	if err := validate.Required("steps", "body", m.Steps); err != nil {
		return err
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScenarioRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScenarioRequest) UnmarshalBinary(b []byte) error {
	var res ScenarioRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
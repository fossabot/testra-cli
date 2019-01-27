// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Simulation Simulation
// swagger:model Simulation
type Simulation struct {

	// execution Id
	ExecutionID string `json:"executionId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// scenarios
	Scenarios []*SimulationScenario `json:"scenarios"`
}

// Validate validates this simulation
func (m *Simulation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScenarios(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Simulation) validateScenarios(formats strfmt.Registry) error {

	if swag.IsZero(m.Scenarios) { // not required
		return nil
	}

	for i := 0; i < len(m.Scenarios); i++ {
		if swag.IsZero(m.Scenarios[i]) { // not required
			continue
		}

		if m.Scenarios[i] != nil {
			if err := m.Scenarios[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scenarios" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Simulation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Simulation) UnmarshalBinary(b []byte) error {
	var res Simulation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
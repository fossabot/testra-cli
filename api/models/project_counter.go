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

// ProjectCounter ProjectCounter
// swagger:model ProjectCounter
type ProjectCounter struct {

	// simulations count
	// Required: true
	SimulationsCount *int64 `json:"simulationsCount"`

	// test cases count
	// Required: true
	TestCasesCount *int64 `json:"testCasesCount"`

	// test executions count
	// Required: true
	TestExecutionsCount *int64 `json:"testExecutionsCount"`

	// test results count
	// Required: true
	TestResultsCount *int64 `json:"testResultsCount"`

	// test scenarios count
	// Required: true
	TestScenariosCount *int64 `json:"testScenariosCount"`

	// vulnerability alerts count
	// Required: true
	VulnerabilityAlertsCount *int64 `json:"vulnerabilityAlertsCount"`
}

// Validate validates this project counter
func (m *ProjectCounter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSimulationsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestCasesCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestExecutionsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestResultsCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTestScenariosCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilityAlertsCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectCounter) validateSimulationsCount(formats strfmt.Registry) error {

	if err := validate.Required("simulationsCount", "body", m.SimulationsCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectCounter) validateTestCasesCount(formats strfmt.Registry) error {

	if err := validate.Required("testCasesCount", "body", m.TestCasesCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectCounter) validateTestExecutionsCount(formats strfmt.Registry) error {

	if err := validate.Required("testExecutionsCount", "body", m.TestExecutionsCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectCounter) validateTestResultsCount(formats strfmt.Registry) error {

	if err := validate.Required("testResultsCount", "body", m.TestResultsCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectCounter) validateTestScenariosCount(formats strfmt.Registry) error {

	if err := validate.Required("testScenariosCount", "body", m.TestScenariosCount); err != nil {
		return err
	}

	return nil
}

func (m *ProjectCounter) validateVulnerabilityAlertsCount(formats strfmt.Registry) error {

	if err := validate.Required("vulnerabilityAlertsCount", "body", m.VulnerabilityAlertsCount); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectCounter) UnmarshalBinary(b []byte) error {
	var res ProjectCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

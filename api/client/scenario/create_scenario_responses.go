// Code generated by go-swagger; DO NOT EDIT.

package scenario

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra-tech/testra-cli/api/models"
)

// CreateScenarioReader is a Reader for the CreateScenario structure.
type CreateScenarioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScenarioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateScenarioCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewCreateScenarioNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateScenarioDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateScenarioCreated creates a CreateScenarioCreated with default headers values
func NewCreateScenarioCreated() *CreateScenarioCreated {
	return &CreateScenarioCreated{}
}

/*CreateScenarioCreated handles this case with default header values.

Successful creation of Scenario
*/
type CreateScenarioCreated struct {
	Payload *models.Scenario
}

func (o *CreateScenarioCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/scenarios][%d] createScenarioCreated  %+v", 201, o.Payload)
}

func (o *CreateScenarioCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Scenario)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScenarioNotFound creates a CreateScenarioNotFound with default headers values
func NewCreateScenarioNotFound() *CreateScenarioNotFound {
	return &CreateScenarioNotFound{}
}

/*CreateScenarioNotFound handles this case with default header values.

Project not found
*/
type CreateScenarioNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CreateScenarioNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/scenarios][%d] createScenarioNotFound  %+v", 404, o.Payload)
}

func (o *CreateScenarioNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScenarioDefault creates a CreateScenarioDefault with default headers values
func NewCreateScenarioDefault(code int) *CreateScenarioDefault {
	return &CreateScenarioDefault{
		_statusCode: code,
	}
}

/*CreateScenarioDefault handles this case with default header values.

Any other errors
*/
type CreateScenarioDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create scenario default response
func (o *CreateScenarioDefault) Code() int {
	return o._statusCode
}

func (o *CreateScenarioDefault) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/scenarios][%d] createScenario default  %+v", o._statusCode, o.Payload)
}

func (o *CreateScenarioDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

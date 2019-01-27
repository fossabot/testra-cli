// Code generated by go-swagger; DO NOT EDIT.

package scenario

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra/testra-cli/api/models"
)

// GetScenarioReader is a Reader for the GetScenario structure.
type GetScenarioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScenarioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetScenarioOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetScenarioNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetScenarioDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScenarioOK creates a GetScenarioOK with default headers values
func NewGetScenarioOK() *GetScenarioOK {
	return &GetScenarioOK{}
}

/*GetScenarioOK handles this case with default header values.

Scenario Info
*/
type GetScenarioOK struct {
	Payload *models.Scenario
}

func (o *GetScenarioOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/scenarios/{id}][%d] getScenarioOK  %+v", 200, o.Payload)
}

func (o *GetScenarioOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Scenario)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScenarioNotFound creates a GetScenarioNotFound with default headers values
func NewGetScenarioNotFound() *GetScenarioNotFound {
	return &GetScenarioNotFound{}
}

/*GetScenarioNotFound handles this case with default header values.

Scenario/Project not found
*/
type GetScenarioNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetScenarioNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/scenarios/{id}][%d] getScenarioNotFound  %+v", 404, o.Payload)
}

func (o *GetScenarioNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScenarioDefault creates a GetScenarioDefault with default headers values
func NewGetScenarioDefault(code int) *GetScenarioDefault {
	return &GetScenarioDefault{
		_statusCode: code,
	}
}

/*GetScenarioDefault handles this case with default header values.

Any other errors
*/
type GetScenarioDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get scenario default response
func (o *GetScenarioDefault) Code() int {
	return o._statusCode
}

func (o *GetScenarioDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/scenarios/{id}][%d] getScenario default  %+v", o._statusCode, o.Payload)
}

func (o *GetScenarioDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

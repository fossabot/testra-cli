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

// DeleteScenarioReader is a Reader for the DeleteScenario structure.
type DeleteScenarioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScenarioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteScenarioOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteScenarioNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteScenarioDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteScenarioOK creates a DeleteScenarioOK with default headers values
func NewDeleteScenarioOK() *DeleteScenarioOK {
	return &DeleteScenarioOK{}
}

/*DeleteScenarioOK handles this case with default header values.

Successful deletion of given Scenario
*/
type DeleteScenarioOK struct {
	Payload *models.Scenario
}

func (o *DeleteScenarioOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/scenarios/{id}][%d] deleteScenarioOK  %+v", 200, o.Payload)
}

func (o *DeleteScenarioOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Scenario)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScenarioNotFound creates a DeleteScenarioNotFound with default headers values
func NewDeleteScenarioNotFound() *DeleteScenarioNotFound {
	return &DeleteScenarioNotFound{}
}

/*DeleteScenarioNotFound handles this case with default header values.

Scenario not found
*/
type DeleteScenarioNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteScenarioNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/scenarios/{id}][%d] deleteScenarioNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScenarioNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScenarioDefault creates a DeleteScenarioDefault with default headers values
func NewDeleteScenarioDefault(code int) *DeleteScenarioDefault {
	return &DeleteScenarioDefault{
		_statusCode: code,
	}
}

/*DeleteScenarioDefault handles this case with default header values.

Any other errors
*/
type DeleteScenarioDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete scenario default response
func (o *DeleteScenarioDefault) Code() int {
	return o._statusCode
}

func (o *DeleteScenarioDefault) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/scenarios/{id}][%d] deleteScenario default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteScenarioDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

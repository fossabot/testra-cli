// Code generated by go-swagger; DO NOT EDIT.

package result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra-tech/testra-cli/api/models"
)

// DeleteResultReader is a Reader for the DeleteResult structure.
type DeleteResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteResultOK creates a DeleteResultOK with default headers values
func NewDeleteResultOK() *DeleteResultOK {
	return &DeleteResultOK{}
}

/*DeleteResultOK handles this case with default header values.

Successful deletion of given Result
*/
type DeleteResultOK struct {
	Payload *models.TestResult
}

func (o *DeleteResultOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/executions/{executionId}/results/{id}][%d] deleteResultOK  %+v", 200, o.Payload)
}

func (o *DeleteResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResultNotFound creates a DeleteResultNotFound with default headers values
func NewDeleteResultNotFound() *DeleteResultNotFound {
	return &DeleteResultNotFound{}
}

/*DeleteResultNotFound handles this case with default header values.

Test result not found
*/
type DeleteResultNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteResultNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/executions/{executionId}/results/{id}][%d] deleteResultNotFound  %+v", 404, o.Payload)
}

func (o *DeleteResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResultDefault creates a DeleteResultDefault with default headers values
func NewDeleteResultDefault(code int) *DeleteResultDefault {
	return &DeleteResultDefault{
		_statusCode: code,
	}
}

/*DeleteResultDefault handles this case with default header values.

Any other errors
*/
type DeleteResultDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete result default response
func (o *DeleteResultDefault) Code() int {
	return o._statusCode
}

func (o *DeleteResultDefault) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/executions/{executionId}/results/{id}][%d] deleteResult default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

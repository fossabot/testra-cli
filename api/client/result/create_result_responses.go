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

// CreateResultReader is a Reader for the CreateResult structure.
type CreateResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateResultCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateResultCreated creates a CreateResultCreated with default headers values
func NewCreateResultCreated() *CreateResultCreated {
	return &CreateResultCreated{}
}

/*CreateResultCreated handles this case with default header values.

Successful creation of Result
*/
type CreateResultCreated struct {
	Payload *models.TestResult
}

func (o *CreateResultCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/executions/{executionId}/results][%d] createResultCreated  %+v", 201, o.Payload)
}

func (o *CreateResultCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResultDefault creates a CreateResultDefault with default headers values
func NewCreateResultDefault(code int) *CreateResultDefault {
	return &CreateResultDefault{
		_statusCode: code,
	}
}

/*CreateResultDefault handles this case with default header values.

Any other errors
*/
type CreateResultDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create result default response
func (o *CreateResultDefault) Code() int {
	return o._statusCode
}

func (o *CreateResultDefault) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/executions/{executionId}/results][%d] createResult default  %+v", o._statusCode, o.Payload)
}

func (o *CreateResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
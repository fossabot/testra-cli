// Code generated by go-swagger; DO NOT EDIT.

package testcase

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra/testra-cli/api/models"
)

// CreateTestcaseReader is a Reader for the CreateTestcase structure.
type CreateTestcaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTestcaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTestcaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateTestcaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTestcaseCreated creates a CreateTestcaseCreated with default headers values
func NewCreateTestcaseCreated() *CreateTestcaseCreated {
	return &CreateTestcaseCreated{}
}

/*CreateTestcaseCreated handles this case with default header values.

Successful creation of Testcase
*/
type CreateTestcaseCreated struct {
	Payload *models.Testcase
}

func (o *CreateTestcaseCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/testcases][%d] createTestcaseCreated  %+v", 201, o.Payload)
}

func (o *CreateTestcaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Testcase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTestcaseDefault creates a CreateTestcaseDefault with default headers values
func NewCreateTestcaseDefault(code int) *CreateTestcaseDefault {
	return &CreateTestcaseDefault{
		_statusCode: code,
	}
}

/*CreateTestcaseDefault handles this case with default header values.

Any other errors
*/
type CreateTestcaseDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create testcase default response
func (o *CreateTestcaseDefault) Code() int {
	return o._statusCode
}

func (o *CreateTestcaseDefault) Error() string {
	return fmt.Sprintf("[POST /projects/{projectId}/testcases][%d] createTestcase default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTestcaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

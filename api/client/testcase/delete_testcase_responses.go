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

// DeleteTestcaseReader is a Reader for the DeleteTestcase structure.
type DeleteTestcaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTestcaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTestcaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteTestcaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteTestcaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTestcaseOK creates a DeleteTestcaseOK with default headers values
func NewDeleteTestcaseOK() *DeleteTestcaseOK {
	return &DeleteTestcaseOK{}
}

/*DeleteTestcaseOK handles this case with default header values.

Successful deletion of given Testcase
*/
type DeleteTestcaseOK struct {
	Payload *models.Testcase
}

func (o *DeleteTestcaseOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/testcases/{id}][%d] deleteTestcaseOK  %+v", 200, o.Payload)
}

func (o *DeleteTestcaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Testcase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTestcaseNotFound creates a DeleteTestcaseNotFound with default headers values
func NewDeleteTestcaseNotFound() *DeleteTestcaseNotFound {
	return &DeleteTestcaseNotFound{}
}

/*DeleteTestcaseNotFound handles this case with default header values.

Testcase not found
*/
type DeleteTestcaseNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteTestcaseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/testcases/{id}][%d] deleteTestcaseNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTestcaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTestcaseDefault creates a DeleteTestcaseDefault with default headers values
func NewDeleteTestcaseDefault(code int) *DeleteTestcaseDefault {
	return &DeleteTestcaseDefault{
		_statusCode: code,
	}
}

/*DeleteTestcaseDefault handles this case with default header values.

Any other errors
*/
type DeleteTestcaseDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete testcase default response
func (o *DeleteTestcaseDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTestcaseDefault) Error() string {
	return fmt.Sprintf("[DELETE /projects/{projectId}/testcases/{id}][%d] deleteTestcase default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTestcaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

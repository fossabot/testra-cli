// Code generated by go-swagger; DO NOT EDIT.

package result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra/testra-cli/api/models"
)

// GetResultReader is a Reader for the GetResult structure.
type GetResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultOK creates a GetResultOK with default headers values
func NewGetResultOK() *GetResultOK {
	return &GetResultOK{}
}

/*GetResultOK handles this case with default header values.

Successful update of given Result
*/
type GetResultOK struct {
	Payload *models.TestResult
}

func (o *GetResultOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/executions/{executionId}/results/{id}][%d] getResultOK  %+v", 200, o.Payload)
}

func (o *GetResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultNotFound creates a GetResultNotFound with default headers values
func NewGetResultNotFound() *GetResultNotFound {
	return &GetResultNotFound{}
}

/*GetResultNotFound handles this case with default header values.

Test result not found
*/
type GetResultNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetResultNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/executions/{executionId}/results/{id}][%d] getResultNotFound  %+v", 404, o.Payload)
}

func (o *GetResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultDefault creates a GetResultDefault with default headers values
func NewGetResultDefault(code int) *GetResultDefault {
	return &GetResultDefault{
		_statusCode: code,
	}
}

/*GetResultDefault handles this case with default header values.

Any other errors
*/
type GetResultDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get result default response
func (o *GetResultDefault) Code() int {
	return o._statusCode
}

func (o *GetResultDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/executions/{executionId}/results/{id}][%d] getResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

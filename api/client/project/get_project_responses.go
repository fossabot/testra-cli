// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra/testra-cli/api/models"
)

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*GetProjectOK handles this case with default header values.

Successful update of given Project
*/
type GetProjectOK struct {
	Payload *models.Project
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*GetProjectNotFound handles this case with default header values.

Project not found
*/
type GetProjectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDefault creates a GetProjectDefault with default headers values
func NewGetProjectDefault(code int) *GetProjectDefault {
	return &GetProjectDefault{
		_statusCode: code,
	}
}

/*GetProjectDefault handles this case with default header values.

Any other errors
*/
type GetProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get project default response
func (o *GetProjectDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectDefault) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] getProject default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

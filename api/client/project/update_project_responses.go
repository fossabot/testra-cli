// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra-tech/testra-cli/api/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*UpdateProjectOK handles this case with default header values.

Successful update of given Project
*/
type UpdateProjectOK struct {
	Payload *models.Project
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/*UpdateProjectNotFound handles this case with default header values.

Project not found
*/
type UpdateProjectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectDefault creates a UpdateProjectDefault with default headers values
func NewUpdateProjectDefault(code int) *UpdateProjectDefault {
	return &UpdateProjectDefault{
		_statusCode: code,
	}
}

/*UpdateProjectDefault handles this case with default header values.

Any other errors
*/
type UpdateProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update project default response
func (o *UpdateProjectDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectDefault) Error() string {
	return fmt.Sprintf("[PUT /projects/{id}][%d] updateProject default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
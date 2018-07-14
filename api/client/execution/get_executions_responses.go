// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra-tech/testra-cli/api/models"
)

// GetExecutionsReader is a Reader for the GetExecutions structure.
type GetExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExecutionsOK creates a GetExecutionsOK with default headers values
func NewGetExecutionsOK() *GetExecutionsOK {
	return &GetExecutionsOK{}
}

/*GetExecutionsOK handles this case with default header values.

Successful Response
*/
type GetExecutionsOK struct {
	Payload []*models.Execution
}

func (o *GetExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/executions][%d] getExecutionsOK  %+v", 200, o.Payload)
}

func (o *GetExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
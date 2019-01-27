// Code generated by go-swagger; DO NOT EDIT.

package execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra/testra-cli/api/models"
)

// RecentExecutionsReader is a Reader for the RecentExecutions structure.
type RecentExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecentExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecentExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRecentExecutionsOK creates a RecentExecutionsOK with default headers values
func NewRecentExecutionsOK() *RecentExecutionsOK {
	return &RecentExecutionsOK{}
}

/*RecentExecutionsOK handles this case with default header values.

Successful Response
*/
type RecentExecutionsOK struct {
	Payload []*models.Execution
}

func (o *RecentExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /executions/recents][%d] recentExecutionsOK  %+v", 200, o.Payload)
}

func (o *RecentExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

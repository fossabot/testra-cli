// Code generated by go-swagger; DO NOT EDIT.

package counter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/testra-tech/testra-cli/api/models"
)

// GetCountersReader is a Reader for the GetCounters structure.
type GetCountersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCountersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCountersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCountersOK creates a GetCountersOK with default headers values
func NewGetCountersOK() *GetCountersOK {
	return &GetCountersOK{}
}

/*GetCountersOK handles this case with default header values.

Successful Response
*/
type GetCountersOK struct {
	Payload *models.Counter
}

func (o *GetCountersOK) Error() string {
	return fmt.Sprintf("[GET /counters][%d] getCountersOK  %+v", 200, o.Payload)
}

func (o *GetCountersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Counter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

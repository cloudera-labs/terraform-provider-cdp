// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// DeleteEnvironmentReader is a Reader for the DeleteEnvironment structure.
type DeleteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteEnvironmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEnvironmentOK creates a DeleteEnvironmentOK with default headers values
func NewDeleteEnvironmentOK() *DeleteEnvironmentOK {
	return &DeleteEnvironmentOK{}
}

/*DeleteEnvironmentOK handles this case with default header values.

Expected response to a valid request.
*/
type DeleteEnvironmentOK struct {
	Payload models.DeleteEnvironmentResponse
}

func (o *DeleteEnvironmentOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/deleteEnvironment][%d] deleteEnvironmentOK  %+v", 200, o.Payload)
}

func (o *DeleteEnvironmentOK) GetPayload() models.DeleteEnvironmentResponse {
	return o.Payload
}

func (o *DeleteEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEnvironmentDefault creates a DeleteEnvironmentDefault with default headers values
func NewDeleteEnvironmentDefault(code int) *DeleteEnvironmentDefault {
	return &DeleteEnvironmentDefault{
		_statusCode: code,
	}
}

/*DeleteEnvironmentDefault handles this case with default header values.

The default response on an error.
*/
type DeleteEnvironmentDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete environment default response
func (o *DeleteEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEnvironmentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/deleteEnvironment][%d] deleteEnvironment default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEnvironmentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

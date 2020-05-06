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

// GetKeytabReader is a Reader for the GetKeytab structure.
type GetKeytabReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKeytabReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKeytabOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKeytabDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKeytabOK creates a GetKeytabOK with default headers values
func NewGetKeytabOK() *GetKeytabOK {
	return &GetKeytabOK{}
}

/*GetKeytabOK handles this case with default header values.

Expected response to a valid request.
*/
type GetKeytabOK struct {
	Payload *models.GetKeytabResponse
}

func (o *GetKeytabOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getKeytab][%d] getKeytabOK  %+v", 200, o.Payload)
}

func (o *GetKeytabOK) GetPayload() *models.GetKeytabResponse {
	return o.Payload
}

func (o *GetKeytabOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetKeytabResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKeytabDefault creates a GetKeytabDefault with default headers values
func NewGetKeytabDefault(code int) *GetKeytabDefault {
	return &GetKeytabDefault{
		_statusCode: code,
	}
}

/*GetKeytabDefault handles this case with default header values.

The default response on an error.
*/
type GetKeytabDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get keytab default response
func (o *GetKeytabDefault) Code() int {
	return o._statusCode
}

func (o *GetKeytabDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getKeytab][%d] getKeytab default  %+v", o._statusCode, o.Payload)
}

func (o *GetKeytabDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKeytabDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

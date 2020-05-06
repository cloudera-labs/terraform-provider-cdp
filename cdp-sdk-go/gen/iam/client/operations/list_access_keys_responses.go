// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// ListAccessKeysReader is a Reader for the ListAccessKeys structure.
type ListAccessKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccessKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAccessKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAccessKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAccessKeysOK creates a ListAccessKeysOK with default headers values
func NewListAccessKeysOK() *ListAccessKeysOK {
	return &ListAccessKeysOK{}
}

/*ListAccessKeysOK handles this case with default header values.

Expected response to a valid request.
*/
type ListAccessKeysOK struct {
	Payload *models.ListAccessKeysResponse
}

func (o *ListAccessKeysOK) Error() string {
	return fmt.Sprintf("[POST /iam/listAccessKeys][%d] listAccessKeysOK  %+v", 200, o.Payload)
}

func (o *ListAccessKeysOK) GetPayload() *models.ListAccessKeysResponse {
	return o.Payload
}

func (o *ListAccessKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAccessKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccessKeysDefault creates a ListAccessKeysDefault with default headers values
func NewListAccessKeysDefault(code int) *ListAccessKeysDefault {
	return &ListAccessKeysDefault{
		_statusCode: code,
	}
}

/*ListAccessKeysDefault handles this case with default header values.

The default response on an error.
*/
type ListAccessKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list access keys default response
func (o *ListAccessKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListAccessKeysDefault) Error() string {
	return fmt.Sprintf("[POST /iam/listAccessKeys][%d] listAccessKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListAccessKeysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAccessKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

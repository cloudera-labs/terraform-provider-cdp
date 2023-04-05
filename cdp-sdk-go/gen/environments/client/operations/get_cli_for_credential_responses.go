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

// GetCliForCredentialReader is a Reader for the GetCliForCredential structure.
type GetCliForCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCliForCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCliForCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCliForCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCliForCredentialOK creates a GetCliForCredentialOK with default headers values
func NewGetCliForCredentialOK() *GetCliForCredentialOK {
	return &GetCliForCredentialOK{}
}

/*
GetCliForCredentialOK describes a response with status code 200, with default header values.

Create credential CLI command for valid request.
*/
type GetCliForCredentialOK struct {
	Payload *models.GetCliForCredentialResponse
}

// IsSuccess returns true when this get cli for credential o k response has a 2xx status code
func (o *GetCliForCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cli for credential o k response has a 3xx status code
func (o *GetCliForCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cli for credential o k response has a 4xx status code
func (o *GetCliForCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cli for credential o k response has a 5xx status code
func (o *GetCliForCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cli for credential o k response a status code equal to that given
func (o *GetCliForCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cli for credential o k response
func (o *GetCliForCredentialOK) Code() int {
	return 200
}

func (o *GetCliForCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForCredential][%d] getCliForCredentialOK  %+v", 200, o.Payload)
}

func (o *GetCliForCredentialOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForCredential][%d] getCliForCredentialOK  %+v", 200, o.Payload)
}

func (o *GetCliForCredentialOK) GetPayload() *models.GetCliForCredentialResponse {
	return o.Payload
}

func (o *GetCliForCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCliForCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCliForCredentialDefault creates a GetCliForCredentialDefault with default headers values
func NewGetCliForCredentialDefault(code int) *GetCliForCredentialDefault {
	return &GetCliForCredentialDefault{
		_statusCode: code,
	}
}

/*
GetCliForCredentialDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetCliForCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get cli for credential default response has a 2xx status code
func (o *GetCliForCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cli for credential default response has a 3xx status code
func (o *GetCliForCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cli for credential default response has a 4xx status code
func (o *GetCliForCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cli for credential default response has a 5xx status code
func (o *GetCliForCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cli for credential default response a status code equal to that given
func (o *GetCliForCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cli for credential default response
func (o *GetCliForCredentialDefault) Code() int {
	return o._statusCode
}

func (o *GetCliForCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForCredential][%d] getCliForCredential default  %+v", o._statusCode, o.Payload)
}

func (o *GetCliForCredentialDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getCliForCredential][%d] getCliForCredential default  %+v", o._statusCode, o.Payload)
}

func (o *GetCliForCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCliForCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
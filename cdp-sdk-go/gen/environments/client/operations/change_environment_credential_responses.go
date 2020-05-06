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

// ChangeEnvironmentCredentialReader is a Reader for the ChangeEnvironmentCredential structure.
type ChangeEnvironmentCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeEnvironmentCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeEnvironmentCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeEnvironmentCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeEnvironmentCredentialOK creates a ChangeEnvironmentCredentialOK with default headers values
func NewChangeEnvironmentCredentialOK() *ChangeEnvironmentCredentialOK {
	return &ChangeEnvironmentCredentialOK{}
}

/*ChangeEnvironmentCredentialOK handles this case with default header values.

Expected response to a valid request.
*/
type ChangeEnvironmentCredentialOK struct {
	Payload *models.ChangeEnvironmentCredentialResponse
}

func (o *ChangeEnvironmentCredentialOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredentialOK  %+v", 200, o.Payload)
}

func (o *ChangeEnvironmentCredentialOK) GetPayload() *models.ChangeEnvironmentCredentialResponse {
	return o.Payload
}

func (o *ChangeEnvironmentCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeEnvironmentCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeEnvironmentCredentialDefault creates a ChangeEnvironmentCredentialDefault with default headers values
func NewChangeEnvironmentCredentialDefault(code int) *ChangeEnvironmentCredentialDefault {
	return &ChangeEnvironmentCredentialDefault{
		_statusCode: code,
	}
}

/*ChangeEnvironmentCredentialDefault handles this case with default header values.

The default response on an error.
*/
type ChangeEnvironmentCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the change environment credential default response
func (o *ChangeEnvironmentCredentialDefault) Code() int {
	return o._statusCode
}

func (o *ChangeEnvironmentCredentialDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredential default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeEnvironmentCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ChangeEnvironmentCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

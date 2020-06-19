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

// CreateLdapProviderReader is a Reader for the CreateLdapProvider structure.
type CreateLdapProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLdapProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLdapProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateLdapProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateLdapProviderOK creates a CreateLdapProviderOK with default headers values
func NewCreateLdapProviderOK() *CreateLdapProviderOK {
	return &CreateLdapProviderOK{}
}

/*CreateLdapProviderOK handles this case with default header values.

Expected response to a valid request.
*/
type CreateLdapProviderOK struct {
	Payload *models.CreateLdapProviderResponse
}

func (o *CreateLdapProviderOK) Error() string {
	return fmt.Sprintf("[POST /iam/createLdapProvider][%d] createLdapProviderOK  %+v", 200, o.Payload)
}

func (o *CreateLdapProviderOK) GetPayload() *models.CreateLdapProviderResponse {
	return o.Payload
}

func (o *CreateLdapProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateLdapProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapProviderDefault creates a CreateLdapProviderDefault with default headers values
func NewCreateLdapProviderDefault(code int) *CreateLdapProviderDefault {
	return &CreateLdapProviderDefault{
		_statusCode: code,
	}
}

/*CreateLdapProviderDefault handles this case with default header values.

The default response on an error.
*/
type CreateLdapProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create ldap provider default response
func (o *CreateLdapProviderDefault) Code() int {
	return o._statusCode
}

func (o *CreateLdapProviderDefault) Error() string {
	return fmt.Sprintf("[POST /iam/createLdapProvider][%d] createLdapProvider default  %+v", o._statusCode, o.Payload)
}

func (o *CreateLdapProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLdapProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

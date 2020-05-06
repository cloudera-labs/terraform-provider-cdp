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

// UpdateSamlProviderReader is a Reader for the UpdateSamlProvider structure.
type UpdateSamlProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSamlProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSamlProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSamlProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSamlProviderOK creates a UpdateSamlProviderOK with default headers values
func NewUpdateSamlProviderOK() *UpdateSamlProviderOK {
	return &UpdateSamlProviderOK{}
}

/*UpdateSamlProviderOK handles this case with default header values.

Expected response to a valid request.
*/
type UpdateSamlProviderOK struct {
	Payload *models.UpdateSamlProviderResponse
}

func (o *UpdateSamlProviderOK) Error() string {
	return fmt.Sprintf("[POST /iam/updateSamlProvider][%d] updateSamlProviderOK  %+v", 200, o.Payload)
}

func (o *UpdateSamlProviderOK) GetPayload() *models.UpdateSamlProviderResponse {
	return o.Payload
}

func (o *UpdateSamlProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateSamlProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlProviderDefault creates a UpdateSamlProviderDefault with default headers values
func NewUpdateSamlProviderDefault(code int) *UpdateSamlProviderDefault {
	return &UpdateSamlProviderDefault{
		_statusCode: code,
	}
}

/*UpdateSamlProviderDefault handles this case with default header values.

The default response on an error.
*/
type UpdateSamlProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update saml provider default response
func (o *UpdateSamlProviderDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSamlProviderDefault) Error() string {
	return fmt.Sprintf("[POST /iam/updateSamlProvider][%d] updateSamlProvider default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSamlProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSamlProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

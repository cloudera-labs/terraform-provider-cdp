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

// ExtractCatalogReader is a Reader for the ExtractCatalog structure.
type ExtractCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtractCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtractCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtractCatalogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtractCatalogOK creates a ExtractCatalogOK with default headers values
func NewExtractCatalogOK() *ExtractCatalogOK {
	return &ExtractCatalogOK{}
}

/*
ExtractCatalogOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ExtractCatalogOK struct {
	Payload *models.ExtractCatalogResponse
}

// IsSuccess returns true when this extract catalog o k response has a 2xx status code
func (o *ExtractCatalogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extract catalog o k response has a 3xx status code
func (o *ExtractCatalogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extract catalog o k response has a 4xx status code
func (o *ExtractCatalogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extract catalog o k response has a 5xx status code
func (o *ExtractCatalogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extract catalog o k response a status code equal to that given
func (o *ExtractCatalogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extract catalog o k response
func (o *ExtractCatalogOK) Code() int {
	return 200
}

func (o *ExtractCatalogOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/extractCatalog][%d] extractCatalogOK  %+v", 200, o.Payload)
}

func (o *ExtractCatalogOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/extractCatalog][%d] extractCatalogOK  %+v", 200, o.Payload)
}

func (o *ExtractCatalogOK) GetPayload() *models.ExtractCatalogResponse {
	return o.Payload
}

func (o *ExtractCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtractCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractCatalogDefault creates a ExtractCatalogDefault with default headers values
func NewExtractCatalogDefault(code int) *ExtractCatalogDefault {
	return &ExtractCatalogDefault{
		_statusCode: code,
	}
}

/*
ExtractCatalogDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ExtractCatalogDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this extract catalog default response has a 2xx status code
func (o *ExtractCatalogDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extract catalog default response has a 3xx status code
func (o *ExtractCatalogDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extract catalog default response has a 4xx status code
func (o *ExtractCatalogDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extract catalog default response has a 5xx status code
func (o *ExtractCatalogDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extract catalog default response a status code equal to that given
func (o *ExtractCatalogDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extract catalog default response
func (o *ExtractCatalogDefault) Code() int {
	return o._statusCode
}

func (o *ExtractCatalogDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/extractCatalog][%d] extractCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractCatalogDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/extractCatalog][%d] extractCatalog default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractCatalogDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExtractCatalogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
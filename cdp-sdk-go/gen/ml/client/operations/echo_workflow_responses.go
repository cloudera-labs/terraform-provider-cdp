// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// EchoWorkflowReader is a Reader for the EchoWorkflow structure.
type EchoWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EchoWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEchoWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEchoWorkflowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEchoWorkflowOK creates a EchoWorkflowOK with default headers values
func NewEchoWorkflowOK() *EchoWorkflowOK {
	return &EchoWorkflowOK{}
}

/*
EchoWorkflowOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type EchoWorkflowOK struct {
	Payload *models.EchoWorkflowResponse
}

// IsSuccess returns true when this echo workflow o k response has a 2xx status code
func (o *EchoWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this echo workflow o k response has a 3xx status code
func (o *EchoWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this echo workflow o k response has a 4xx status code
func (o *EchoWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this echo workflow o k response has a 5xx status code
func (o *EchoWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this echo workflow o k response a status code equal to that given
func (o *EchoWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the echo workflow o k response
func (o *EchoWorkflowOK) Code() int {
	return 200
}

func (o *EchoWorkflowOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/echoWorkflow][%d] echoWorkflowOK  %+v", 200, o.Payload)
}

func (o *EchoWorkflowOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/echoWorkflow][%d] echoWorkflowOK  %+v", 200, o.Payload)
}

func (o *EchoWorkflowOK) GetPayload() *models.EchoWorkflowResponse {
	return o.Payload
}

func (o *EchoWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EchoWorkflowResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEchoWorkflowDefault creates a EchoWorkflowDefault with default headers values
func NewEchoWorkflowDefault(code int) *EchoWorkflowDefault {
	return &EchoWorkflowDefault{
		_statusCode: code,
	}
}

/*
EchoWorkflowDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type EchoWorkflowDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this echo workflow default response has a 2xx status code
func (o *EchoWorkflowDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this echo workflow default response has a 3xx status code
func (o *EchoWorkflowDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this echo workflow default response has a 4xx status code
func (o *EchoWorkflowDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this echo workflow default response has a 5xx status code
func (o *EchoWorkflowDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this echo workflow default response a status code equal to that given
func (o *EchoWorkflowDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the echo workflow default response
func (o *EchoWorkflowDefault) Code() int {
	return o._statusCode
}

func (o *EchoWorkflowDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/echoWorkflow][%d] echoWorkflow default  %+v", o._statusCode, o.Payload)
}

func (o *EchoWorkflowDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/echoWorkflow][%d] echoWorkflow default  %+v", o._statusCode, o.Payload)
}

func (o *EchoWorkflowDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EchoWorkflowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
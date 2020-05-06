// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewAddUserToGroupParams creates a new AddUserToGroupParams object
// with the default values initialized.
func NewAddUserToGroupParams() *AddUserToGroupParams {
	var ()
	return &AddUserToGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserToGroupParamsWithTimeout creates a new AddUserToGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddUserToGroupParamsWithTimeout(timeout time.Duration) *AddUserToGroupParams {
	var ()
	return &AddUserToGroupParams{

		timeout: timeout,
	}
}

// NewAddUserToGroupParamsWithContext creates a new AddUserToGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddUserToGroupParamsWithContext(ctx context.Context) *AddUserToGroupParams {
	var ()
	return &AddUserToGroupParams{

		Context: ctx,
	}
}

// NewAddUserToGroupParamsWithHTTPClient creates a new AddUserToGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddUserToGroupParamsWithHTTPClient(client *http.Client) *AddUserToGroupParams {
	var ()
	return &AddUserToGroupParams{
		HTTPClient: client,
	}
}

/*AddUserToGroupParams contains all the parameters to send to the API endpoint
for the add user to group operation typically these are written to a http.Request
*/
type AddUserToGroupParams struct {

	/*Input*/
	Input *models.AddUserToGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add user to group params
func (o *AddUserToGroupParams) WithTimeout(timeout time.Duration) *AddUserToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user to group params
func (o *AddUserToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user to group params
func (o *AddUserToGroupParams) WithContext(ctx context.Context) *AddUserToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user to group params
func (o *AddUserToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user to group params
func (o *AddUserToGroupParams) WithHTTPClient(client *http.Client) *AddUserToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user to group params
func (o *AddUserToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the add user to group params
func (o *AddUserToGroupParams) WithInput(input *models.AddUserToGroupRequest) *AddUserToGroupParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the add user to group params
func (o *AddUserToGroupParams) SetInput(input *models.AddUserToGroupRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

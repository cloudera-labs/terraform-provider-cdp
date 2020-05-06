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

// NewUpdateGroupParams creates a new UpdateGroupParams object
// with the default values initialized.
func NewUpdateGroupParams() *UpdateGroupParams {
	var ()
	return &UpdateGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGroupParamsWithTimeout creates a new UpdateGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGroupParamsWithTimeout(timeout time.Duration) *UpdateGroupParams {
	var ()
	return &UpdateGroupParams{

		timeout: timeout,
	}
}

// NewUpdateGroupParamsWithContext creates a new UpdateGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGroupParamsWithContext(ctx context.Context) *UpdateGroupParams {
	var ()
	return &UpdateGroupParams{

		Context: ctx,
	}
}

// NewUpdateGroupParamsWithHTTPClient creates a new UpdateGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGroupParamsWithHTTPClient(client *http.Client) *UpdateGroupParams {
	var ()
	return &UpdateGroupParams{
		HTTPClient: client,
	}
}

/*UpdateGroupParams contains all the parameters to send to the API endpoint
for the update group operation typically these are written to a http.Request
*/
type UpdateGroupParams struct {

	/*Input*/
	Input *models.UpdateGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update group params
func (o *UpdateGroupParams) WithTimeout(timeout time.Duration) *UpdateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update group params
func (o *UpdateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update group params
func (o *UpdateGroupParams) WithContext(ctx context.Context) *UpdateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update group params
func (o *UpdateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update group params
func (o *UpdateGroupParams) WithHTTPClient(client *http.Client) *UpdateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update group params
func (o *UpdateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update group params
func (o *UpdateGroupParams) WithInput(input *models.UpdateGroupRequest) *UpdateGroupParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update group params
func (o *UpdateGroupParams) SetInput(input *models.UpdateGroupRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

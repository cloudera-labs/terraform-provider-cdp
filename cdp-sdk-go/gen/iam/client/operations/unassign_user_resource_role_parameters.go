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

// NewUnassignUserResourceRoleParams creates a new UnassignUserResourceRoleParams object
// with the default values initialized.
func NewUnassignUserResourceRoleParams() *UnassignUserResourceRoleParams {
	var ()
	return &UnassignUserResourceRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnassignUserResourceRoleParamsWithTimeout creates a new UnassignUserResourceRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnassignUserResourceRoleParamsWithTimeout(timeout time.Duration) *UnassignUserResourceRoleParams {
	var ()
	return &UnassignUserResourceRoleParams{

		timeout: timeout,
	}
}

// NewUnassignUserResourceRoleParamsWithContext creates a new UnassignUserResourceRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnassignUserResourceRoleParamsWithContext(ctx context.Context) *UnassignUserResourceRoleParams {
	var ()
	return &UnassignUserResourceRoleParams{

		Context: ctx,
	}
}

// NewUnassignUserResourceRoleParamsWithHTTPClient creates a new UnassignUserResourceRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnassignUserResourceRoleParamsWithHTTPClient(client *http.Client) *UnassignUserResourceRoleParams {
	var ()
	return &UnassignUserResourceRoleParams{
		HTTPClient: client,
	}
}

/*UnassignUserResourceRoleParams contains all the parameters to send to the API endpoint
for the unassign user resource role operation typically these are written to a http.Request
*/
type UnassignUserResourceRoleParams struct {

	/*Input*/
	Input *models.UnassignUserResourceRoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) WithTimeout(timeout time.Duration) *UnassignUserResourceRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) WithContext(ctx context.Context) *UnassignUserResourceRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) WithHTTPClient(client *http.Client) *UnassignUserResourceRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) WithInput(input *models.UnassignUserResourceRoleRequest) *UnassignUserResourceRoleParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the unassign user resource role params
func (o *UnassignUserResourceRoleParams) SetInput(input *models.UnassignUserResourceRoleRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UnassignUserResourceRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

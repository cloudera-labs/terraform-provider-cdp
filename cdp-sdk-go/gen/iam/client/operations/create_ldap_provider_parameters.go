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

// NewCreateLdapProviderParams creates a new CreateLdapProviderParams object
// with the default values initialized.
func NewCreateLdapProviderParams() *CreateLdapProviderParams {
	var ()
	return &CreateLdapProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLdapProviderParamsWithTimeout creates a new CreateLdapProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLdapProviderParamsWithTimeout(timeout time.Duration) *CreateLdapProviderParams {
	var ()
	return &CreateLdapProviderParams{

		timeout: timeout,
	}
}

// NewCreateLdapProviderParamsWithContext creates a new CreateLdapProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLdapProviderParamsWithContext(ctx context.Context) *CreateLdapProviderParams {
	var ()
	return &CreateLdapProviderParams{

		Context: ctx,
	}
}

// NewCreateLdapProviderParamsWithHTTPClient creates a new CreateLdapProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLdapProviderParamsWithHTTPClient(client *http.Client) *CreateLdapProviderParams {
	var ()
	return &CreateLdapProviderParams{
		HTTPClient: client,
	}
}

/*CreateLdapProviderParams contains all the parameters to send to the API endpoint
for the create ldap provider operation typically these are written to a http.Request
*/
type CreateLdapProviderParams struct {

	/*Input*/
	Input *models.CreateLdapProviderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create ldap provider params
func (o *CreateLdapProviderParams) WithTimeout(timeout time.Duration) *CreateLdapProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create ldap provider params
func (o *CreateLdapProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create ldap provider params
func (o *CreateLdapProviderParams) WithContext(ctx context.Context) *CreateLdapProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create ldap provider params
func (o *CreateLdapProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create ldap provider params
func (o *CreateLdapProviderParams) WithHTTPClient(client *http.Client) *CreateLdapProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create ldap provider params
func (o *CreateLdapProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create ldap provider params
func (o *CreateLdapProviderParams) WithInput(input *models.CreateLdapProviderRequest) *CreateLdapProviderParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create ldap provider params
func (o *CreateLdapProviderParams) SetInput(input *models.CreateLdapProviderRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLdapProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

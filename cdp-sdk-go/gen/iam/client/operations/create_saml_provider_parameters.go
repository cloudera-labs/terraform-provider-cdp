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

// NewCreateSamlProviderParams creates a new CreateSamlProviderParams object
// with the default values initialized.
func NewCreateSamlProviderParams() *CreateSamlProviderParams {
	var ()
	return &CreateSamlProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSamlProviderParamsWithTimeout creates a new CreateSamlProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSamlProviderParamsWithTimeout(timeout time.Duration) *CreateSamlProviderParams {
	var ()
	return &CreateSamlProviderParams{

		timeout: timeout,
	}
}

// NewCreateSamlProviderParamsWithContext creates a new CreateSamlProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSamlProviderParamsWithContext(ctx context.Context) *CreateSamlProviderParams {
	var ()
	return &CreateSamlProviderParams{

		Context: ctx,
	}
}

// NewCreateSamlProviderParamsWithHTTPClient creates a new CreateSamlProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSamlProviderParamsWithHTTPClient(client *http.Client) *CreateSamlProviderParams {
	var ()
	return &CreateSamlProviderParams{
		HTTPClient: client,
	}
}

/*CreateSamlProviderParams contains all the parameters to send to the API endpoint
for the create saml provider operation typically these are written to a http.Request
*/
type CreateSamlProviderParams struct {

	/*Input*/
	Input *models.CreateSamlProviderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create saml provider params
func (o *CreateSamlProviderParams) WithTimeout(timeout time.Duration) *CreateSamlProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create saml provider params
func (o *CreateSamlProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create saml provider params
func (o *CreateSamlProviderParams) WithContext(ctx context.Context) *CreateSamlProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create saml provider params
func (o *CreateSamlProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create saml provider params
func (o *CreateSamlProviderParams) WithHTTPClient(client *http.Client) *CreateSamlProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create saml provider params
func (o *CreateSamlProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create saml provider params
func (o *CreateSamlProviderParams) WithInput(input *models.CreateSamlProviderRequest) *CreateSamlProviderParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create saml provider params
func (o *CreateSamlProviderParams) SetInput(input *models.CreateSamlProviderRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSamlProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

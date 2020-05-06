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

// NewDeleteSamlProviderParams creates a new DeleteSamlProviderParams object
// with the default values initialized.
func NewDeleteSamlProviderParams() *DeleteSamlProviderParams {
	var ()
	return &DeleteSamlProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSamlProviderParamsWithTimeout creates a new DeleteSamlProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSamlProviderParamsWithTimeout(timeout time.Duration) *DeleteSamlProviderParams {
	var ()
	return &DeleteSamlProviderParams{

		timeout: timeout,
	}
}

// NewDeleteSamlProviderParamsWithContext creates a new DeleteSamlProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSamlProviderParamsWithContext(ctx context.Context) *DeleteSamlProviderParams {
	var ()
	return &DeleteSamlProviderParams{

		Context: ctx,
	}
}

// NewDeleteSamlProviderParamsWithHTTPClient creates a new DeleteSamlProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSamlProviderParamsWithHTTPClient(client *http.Client) *DeleteSamlProviderParams {
	var ()
	return &DeleteSamlProviderParams{
		HTTPClient: client,
	}
}

/*DeleteSamlProviderParams contains all the parameters to send to the API endpoint
for the delete saml provider operation typically these are written to a http.Request
*/
type DeleteSamlProviderParams struct {

	/*Input*/
	Input *models.DeleteSamlProviderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete saml provider params
func (o *DeleteSamlProviderParams) WithTimeout(timeout time.Duration) *DeleteSamlProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saml provider params
func (o *DeleteSamlProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saml provider params
func (o *DeleteSamlProviderParams) WithContext(ctx context.Context) *DeleteSamlProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saml provider params
func (o *DeleteSamlProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saml provider params
func (o *DeleteSamlProviderParams) WithHTTPClient(client *http.Client) *DeleteSamlProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saml provider params
func (o *DeleteSamlProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete saml provider params
func (o *DeleteSamlProviderParams) WithInput(input *models.DeleteSamlProviderRequest) *DeleteSamlProviderParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete saml provider params
func (o *DeleteSamlProviderParams) SetInput(input *models.DeleteSamlProviderRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSamlProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

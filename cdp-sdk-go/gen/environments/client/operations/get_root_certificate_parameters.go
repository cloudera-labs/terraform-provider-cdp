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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewGetRootCertificateParams creates a new GetRootCertificateParams object
// with the default values initialized.
func NewGetRootCertificateParams() *GetRootCertificateParams {
	var ()
	return &GetRootCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRootCertificateParamsWithTimeout creates a new GetRootCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRootCertificateParamsWithTimeout(timeout time.Duration) *GetRootCertificateParams {
	var ()
	return &GetRootCertificateParams{

		timeout: timeout,
	}
}

// NewGetRootCertificateParamsWithContext creates a new GetRootCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRootCertificateParamsWithContext(ctx context.Context) *GetRootCertificateParams {
	var ()
	return &GetRootCertificateParams{

		Context: ctx,
	}
}

// NewGetRootCertificateParamsWithHTTPClient creates a new GetRootCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRootCertificateParamsWithHTTPClient(client *http.Client) *GetRootCertificateParams {
	var ()
	return &GetRootCertificateParams{
		HTTPClient: client,
	}
}

/*GetRootCertificateParams contains all the parameters to send to the API endpoint
for the get root certificate operation typically these are written to a http.Request
*/
type GetRootCertificateParams struct {

	/*Input*/
	Input *models.GetRootCertificateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get root certificate params
func (o *GetRootCertificateParams) WithTimeout(timeout time.Duration) *GetRootCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get root certificate params
func (o *GetRootCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get root certificate params
func (o *GetRootCertificateParams) WithContext(ctx context.Context) *GetRootCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get root certificate params
func (o *GetRootCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get root certificate params
func (o *GetRootCertificateParams) WithHTTPClient(client *http.Client) *GetRootCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get root certificate params
func (o *GetRootCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get root certificate params
func (o *GetRootCertificateParams) WithInput(input *models.GetRootCertificateRequest) *GetRootCertificateParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get root certificate params
func (o *GetRootCertificateParams) SetInput(input *models.GetRootCertificateRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetRootCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

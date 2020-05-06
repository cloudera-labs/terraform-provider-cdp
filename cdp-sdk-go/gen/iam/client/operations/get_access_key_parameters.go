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

// NewGetAccessKeyParams creates a new GetAccessKeyParams object
// with the default values initialized.
func NewGetAccessKeyParams() *GetAccessKeyParams {
	var ()
	return &GetAccessKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccessKeyParamsWithTimeout creates a new GetAccessKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccessKeyParamsWithTimeout(timeout time.Duration) *GetAccessKeyParams {
	var ()
	return &GetAccessKeyParams{

		timeout: timeout,
	}
}

// NewGetAccessKeyParamsWithContext creates a new GetAccessKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccessKeyParamsWithContext(ctx context.Context) *GetAccessKeyParams {
	var ()
	return &GetAccessKeyParams{

		Context: ctx,
	}
}

// NewGetAccessKeyParamsWithHTTPClient creates a new GetAccessKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccessKeyParamsWithHTTPClient(client *http.Client) *GetAccessKeyParams {
	var ()
	return &GetAccessKeyParams{
		HTTPClient: client,
	}
}

/*GetAccessKeyParams contains all the parameters to send to the API endpoint
for the get access key operation typically these are written to a http.Request
*/
type GetAccessKeyParams struct {

	/*Input*/
	Input *models.GetAccessKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get access key params
func (o *GetAccessKeyParams) WithTimeout(timeout time.Duration) *GetAccessKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get access key params
func (o *GetAccessKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get access key params
func (o *GetAccessKeyParams) WithContext(ctx context.Context) *GetAccessKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get access key params
func (o *GetAccessKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get access key params
func (o *GetAccessKeyParams) WithHTTPClient(client *http.Client) *GetAccessKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get access key params
func (o *GetAccessKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get access key params
func (o *GetAccessKeyParams) WithInput(input *models.GetAccessKeyRequest) *GetAccessKeyParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get access key params
func (o *GetAccessKeyParams) SetInput(input *models.GetAccessKeyRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccessKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewDescribeAllowedInstanceTypesParams creates a new DescribeAllowedInstanceTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeAllowedInstanceTypesParams() *DescribeAllowedInstanceTypesParams {
	return &DescribeAllowedInstanceTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeAllowedInstanceTypesParamsWithTimeout creates a new DescribeAllowedInstanceTypesParams object
// with the ability to set a timeout on a request.
func NewDescribeAllowedInstanceTypesParamsWithTimeout(timeout time.Duration) *DescribeAllowedInstanceTypesParams {
	return &DescribeAllowedInstanceTypesParams{
		timeout: timeout,
	}
}

// NewDescribeAllowedInstanceTypesParamsWithContext creates a new DescribeAllowedInstanceTypesParams object
// with the ability to set a context for a request.
func NewDescribeAllowedInstanceTypesParamsWithContext(ctx context.Context) *DescribeAllowedInstanceTypesParams {
	return &DescribeAllowedInstanceTypesParams{
		Context: ctx,
	}
}

// NewDescribeAllowedInstanceTypesParamsWithHTTPClient creates a new DescribeAllowedInstanceTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeAllowedInstanceTypesParamsWithHTTPClient(client *http.Client) *DescribeAllowedInstanceTypesParams {
	return &DescribeAllowedInstanceTypesParams{
		HTTPClient: client,
	}
}

/*
DescribeAllowedInstanceTypesParams contains all the parameters to send to the API endpoint

	for the describe allowed instance types operation.

	Typically these are written to a http.Request.
*/
type DescribeAllowedInstanceTypesParams struct {

	// Input.
	Input models.DescribeAllowedInstanceTypesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe allowed instance types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeAllowedInstanceTypesParams) WithDefaults() *DescribeAllowedInstanceTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe allowed instance types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeAllowedInstanceTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) WithTimeout(timeout time.Duration) *DescribeAllowedInstanceTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) WithContext(ctx context.Context) *DescribeAllowedInstanceTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) WithHTTPClient(client *http.Client) *DescribeAllowedInstanceTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) WithInput(input models.DescribeAllowedInstanceTypesRequest) *DescribeAllowedInstanceTypesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the describe allowed instance types params
func (o *DescribeAllowedInstanceTypesParams) SetInput(input models.DescribeAllowedInstanceTypesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeAllowedInstanceTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
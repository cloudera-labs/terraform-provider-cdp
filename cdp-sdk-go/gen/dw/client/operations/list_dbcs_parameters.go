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

// NewListDbcsParams creates a new ListDbcsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDbcsParams() *ListDbcsParams {
	return &ListDbcsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDbcsParamsWithTimeout creates a new ListDbcsParams object
// with the ability to set a timeout on a request.
func NewListDbcsParamsWithTimeout(timeout time.Duration) *ListDbcsParams {
	return &ListDbcsParams{
		timeout: timeout,
	}
}

// NewListDbcsParamsWithContext creates a new ListDbcsParams object
// with the ability to set a context for a request.
func NewListDbcsParamsWithContext(ctx context.Context) *ListDbcsParams {
	return &ListDbcsParams{
		Context: ctx,
	}
}

// NewListDbcsParamsWithHTTPClient creates a new ListDbcsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDbcsParamsWithHTTPClient(client *http.Client) *ListDbcsParams {
	return &ListDbcsParams{
		HTTPClient: client,
	}
}

/*
ListDbcsParams contains all the parameters to send to the API endpoint

	for the list dbcs operation.

	Typically these are written to a http.Request.
*/
type ListDbcsParams struct {

	// Input.
	Input *models.ListDbcsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list dbcs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDbcsParams) WithDefaults() *ListDbcsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list dbcs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDbcsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list dbcs params
func (o *ListDbcsParams) WithTimeout(timeout time.Duration) *ListDbcsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dbcs params
func (o *ListDbcsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dbcs params
func (o *ListDbcsParams) WithContext(ctx context.Context) *ListDbcsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dbcs params
func (o *ListDbcsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dbcs params
func (o *ListDbcsParams) WithHTTPClient(client *http.Client) *ListDbcsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dbcs params
func (o *ListDbcsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list dbcs params
func (o *ListDbcsParams) WithInput(input *models.ListDbcsRequest) *ListDbcsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list dbcs params
func (o *ListDbcsParams) SetInput(input *models.ListDbcsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListDbcsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
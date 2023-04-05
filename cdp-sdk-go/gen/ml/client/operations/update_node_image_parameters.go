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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// NewUpdateNodeImageParams creates a new UpdateNodeImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNodeImageParams() *UpdateNodeImageParams {
	return &UpdateNodeImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNodeImageParamsWithTimeout creates a new UpdateNodeImageParams object
// with the ability to set a timeout on a request.
func NewUpdateNodeImageParamsWithTimeout(timeout time.Duration) *UpdateNodeImageParams {
	return &UpdateNodeImageParams{
		timeout: timeout,
	}
}

// NewUpdateNodeImageParamsWithContext creates a new UpdateNodeImageParams object
// with the ability to set a context for a request.
func NewUpdateNodeImageParamsWithContext(ctx context.Context) *UpdateNodeImageParams {
	return &UpdateNodeImageParams{
		Context: ctx,
	}
}

// NewUpdateNodeImageParamsWithHTTPClient creates a new UpdateNodeImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNodeImageParamsWithHTTPClient(client *http.Client) *UpdateNodeImageParams {
	return &UpdateNodeImageParams{
		HTTPClient: client,
	}
}

/*
UpdateNodeImageParams contains all the parameters to send to the API endpoint

	for the update node image operation.

	Typically these are written to a http.Request.
*/
type UpdateNodeImageParams struct {

	// Input.
	Input *models.UpdateNodeImageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update node image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNodeImageParams) WithDefaults() *UpdateNodeImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update node image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNodeImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update node image params
func (o *UpdateNodeImageParams) WithTimeout(timeout time.Duration) *UpdateNodeImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update node image params
func (o *UpdateNodeImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update node image params
func (o *UpdateNodeImageParams) WithContext(ctx context.Context) *UpdateNodeImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update node image params
func (o *UpdateNodeImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update node image params
func (o *UpdateNodeImageParams) WithHTTPClient(client *http.Client) *UpdateNodeImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update node image params
func (o *UpdateNodeImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update node image params
func (o *UpdateNodeImageParams) WithInput(input *models.UpdateNodeImageRequest) *UpdateNodeImageParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update node image params
func (o *UpdateNodeImageParams) SetInput(input *models.UpdateNodeImageRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNodeImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
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

// NewGetCloudPlatformPrerequisitesParams creates a new GetCloudPlatformPrerequisitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCloudPlatformPrerequisitesParams() *GetCloudPlatformPrerequisitesParams {
	return &GetCloudPlatformPrerequisitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudPlatformPrerequisitesParamsWithTimeout creates a new GetCloudPlatformPrerequisitesParams object
// with the ability to set a timeout on a request.
func NewGetCloudPlatformPrerequisitesParamsWithTimeout(timeout time.Duration) *GetCloudPlatformPrerequisitesParams {
	return &GetCloudPlatformPrerequisitesParams{
		timeout: timeout,
	}
}

// NewGetCloudPlatformPrerequisitesParamsWithContext creates a new GetCloudPlatformPrerequisitesParams object
// with the ability to set a context for a request.
func NewGetCloudPlatformPrerequisitesParamsWithContext(ctx context.Context) *GetCloudPlatformPrerequisitesParams {
	return &GetCloudPlatformPrerequisitesParams{
		Context: ctx,
	}
}

// NewGetCloudPlatformPrerequisitesParamsWithHTTPClient creates a new GetCloudPlatformPrerequisitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCloudPlatformPrerequisitesParamsWithHTTPClient(client *http.Client) *GetCloudPlatformPrerequisitesParams {
	return &GetCloudPlatformPrerequisitesParams{
		HTTPClient: client,
	}
}

/*
GetCloudPlatformPrerequisitesParams contains all the parameters to send to the API endpoint

	for the get cloud platform prerequisites operation.

	Typically these are written to a http.Request.
*/
type GetCloudPlatformPrerequisitesParams struct {

	// Input.
	Input *models.GetCloudPlatformPrerequisitesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cloud platform prerequisites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudPlatformPrerequisitesParams) WithDefaults() *GetCloudPlatformPrerequisitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cloud platform prerequisites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudPlatformPrerequisitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) WithTimeout(timeout time.Duration) *GetCloudPlatformPrerequisitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) WithContext(ctx context.Context) *GetCloudPlatformPrerequisitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) WithHTTPClient(client *http.Client) *GetCloudPlatformPrerequisitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) WithInput(input *models.GetCloudPlatformPrerequisitesRequest) *GetCloudPlatformPrerequisitesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get cloud platform prerequisites params
func (o *GetCloudPlatformPrerequisitesParams) SetInput(input *models.GetCloudPlatformPrerequisitesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudPlatformPrerequisitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
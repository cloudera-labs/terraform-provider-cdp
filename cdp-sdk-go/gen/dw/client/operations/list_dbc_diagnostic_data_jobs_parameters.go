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

// NewListDbcDiagnosticDataJobsParams creates a new ListDbcDiagnosticDataJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDbcDiagnosticDataJobsParams() *ListDbcDiagnosticDataJobsParams {
	return &ListDbcDiagnosticDataJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDbcDiagnosticDataJobsParamsWithTimeout creates a new ListDbcDiagnosticDataJobsParams object
// with the ability to set a timeout on a request.
func NewListDbcDiagnosticDataJobsParamsWithTimeout(timeout time.Duration) *ListDbcDiagnosticDataJobsParams {
	return &ListDbcDiagnosticDataJobsParams{
		timeout: timeout,
	}
}

// NewListDbcDiagnosticDataJobsParamsWithContext creates a new ListDbcDiagnosticDataJobsParams object
// with the ability to set a context for a request.
func NewListDbcDiagnosticDataJobsParamsWithContext(ctx context.Context) *ListDbcDiagnosticDataJobsParams {
	return &ListDbcDiagnosticDataJobsParams{
		Context: ctx,
	}
}

// NewListDbcDiagnosticDataJobsParamsWithHTTPClient creates a new ListDbcDiagnosticDataJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDbcDiagnosticDataJobsParamsWithHTTPClient(client *http.Client) *ListDbcDiagnosticDataJobsParams {
	return &ListDbcDiagnosticDataJobsParams{
		HTTPClient: client,
	}
}

/*
ListDbcDiagnosticDataJobsParams contains all the parameters to send to the API endpoint

	for the list dbc diagnostic data jobs operation.

	Typically these are written to a http.Request.
*/
type ListDbcDiagnosticDataJobsParams struct {

	// Input.
	Input *models.ListDbcDiagnosticDataJobsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list dbc diagnostic data jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDbcDiagnosticDataJobsParams) WithDefaults() *ListDbcDiagnosticDataJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list dbc diagnostic data jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDbcDiagnosticDataJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) WithTimeout(timeout time.Duration) *ListDbcDiagnosticDataJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) WithContext(ctx context.Context) *ListDbcDiagnosticDataJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) WithHTTPClient(client *http.Client) *ListDbcDiagnosticDataJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) WithInput(input *models.ListDbcDiagnosticDataJobsRequest) *ListDbcDiagnosticDataJobsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list dbc diagnostic data jobs params
func (o *ListDbcDiagnosticDataJobsParams) SetInput(input *models.ListDbcDiagnosticDataJobsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListDbcDiagnosticDataJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
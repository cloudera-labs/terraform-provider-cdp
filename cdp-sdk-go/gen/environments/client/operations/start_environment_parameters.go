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

// NewStartEnvironmentParams creates a new StartEnvironmentParams object
// with the default values initialized.
func NewStartEnvironmentParams() *StartEnvironmentParams {
	var ()
	return &StartEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartEnvironmentParamsWithTimeout creates a new StartEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartEnvironmentParamsWithTimeout(timeout time.Duration) *StartEnvironmentParams {
	var ()
	return &StartEnvironmentParams{

		timeout: timeout,
	}
}

// NewStartEnvironmentParamsWithContext creates a new StartEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartEnvironmentParamsWithContext(ctx context.Context) *StartEnvironmentParams {
	var ()
	return &StartEnvironmentParams{

		Context: ctx,
	}
}

// NewStartEnvironmentParamsWithHTTPClient creates a new StartEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartEnvironmentParamsWithHTTPClient(client *http.Client) *StartEnvironmentParams {
	var ()
	return &StartEnvironmentParams{
		HTTPClient: client,
	}
}

/*StartEnvironmentParams contains all the parameters to send to the API endpoint
for the start environment operation typically these are written to a http.Request
*/
type StartEnvironmentParams struct {

	/*Input*/
	Input *models.StartEnvironmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start environment params
func (o *StartEnvironmentParams) WithTimeout(timeout time.Duration) *StartEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start environment params
func (o *StartEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start environment params
func (o *StartEnvironmentParams) WithContext(ctx context.Context) *StartEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start environment params
func (o *StartEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start environment params
func (o *StartEnvironmentParams) WithHTTPClient(client *http.Client) *StartEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start environment params
func (o *StartEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the start environment params
func (o *StartEnvironmentParams) WithInput(input *models.StartEnvironmentRequest) *StartEnvironmentParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the start environment params
func (o *StartEnvironmentParams) SetInput(input *models.StartEnvironmentRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *StartEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

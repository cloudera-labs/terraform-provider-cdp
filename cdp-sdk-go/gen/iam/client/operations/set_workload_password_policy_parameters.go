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

// NewSetWorkloadPasswordPolicyParams creates a new SetWorkloadPasswordPolicyParams object
// with the default values initialized.
func NewSetWorkloadPasswordPolicyParams() *SetWorkloadPasswordPolicyParams {
	var ()
	return &SetWorkloadPasswordPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetWorkloadPasswordPolicyParamsWithTimeout creates a new SetWorkloadPasswordPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetWorkloadPasswordPolicyParamsWithTimeout(timeout time.Duration) *SetWorkloadPasswordPolicyParams {
	var ()
	return &SetWorkloadPasswordPolicyParams{

		timeout: timeout,
	}
}

// NewSetWorkloadPasswordPolicyParamsWithContext creates a new SetWorkloadPasswordPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetWorkloadPasswordPolicyParamsWithContext(ctx context.Context) *SetWorkloadPasswordPolicyParams {
	var ()
	return &SetWorkloadPasswordPolicyParams{

		Context: ctx,
	}
}

// NewSetWorkloadPasswordPolicyParamsWithHTTPClient creates a new SetWorkloadPasswordPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetWorkloadPasswordPolicyParamsWithHTTPClient(client *http.Client) *SetWorkloadPasswordPolicyParams {
	var ()
	return &SetWorkloadPasswordPolicyParams{
		HTTPClient: client,
	}
}

/*SetWorkloadPasswordPolicyParams contains all the parameters to send to the API endpoint
for the set workload password policy operation typically these are written to a http.Request
*/
type SetWorkloadPasswordPolicyParams struct {

	/*Input*/
	Input *models.SetWorkloadPasswordPolicyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) WithTimeout(timeout time.Duration) *SetWorkloadPasswordPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) WithContext(ctx context.Context) *SetWorkloadPasswordPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) WithHTTPClient(client *http.Client) *SetWorkloadPasswordPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) WithInput(input *models.SetWorkloadPasswordPolicyRequest) *SetWorkloadPasswordPolicyParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the set workload password policy params
func (o *SetWorkloadPasswordPolicyParams) SetInput(input *models.SetWorkloadPasswordPolicyRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SetWorkloadPasswordPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

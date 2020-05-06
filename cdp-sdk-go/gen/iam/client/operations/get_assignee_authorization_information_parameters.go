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

// NewGetAssigneeAuthorizationInformationParams creates a new GetAssigneeAuthorizationInformationParams object
// with the default values initialized.
func NewGetAssigneeAuthorizationInformationParams() *GetAssigneeAuthorizationInformationParams {
	var ()
	return &GetAssigneeAuthorizationInformationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssigneeAuthorizationInformationParamsWithTimeout creates a new GetAssigneeAuthorizationInformationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssigneeAuthorizationInformationParamsWithTimeout(timeout time.Duration) *GetAssigneeAuthorizationInformationParams {
	var ()
	return &GetAssigneeAuthorizationInformationParams{

		timeout: timeout,
	}
}

// NewGetAssigneeAuthorizationInformationParamsWithContext creates a new GetAssigneeAuthorizationInformationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssigneeAuthorizationInformationParamsWithContext(ctx context.Context) *GetAssigneeAuthorizationInformationParams {
	var ()
	return &GetAssigneeAuthorizationInformationParams{

		Context: ctx,
	}
}

// NewGetAssigneeAuthorizationInformationParamsWithHTTPClient creates a new GetAssigneeAuthorizationInformationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssigneeAuthorizationInformationParamsWithHTTPClient(client *http.Client) *GetAssigneeAuthorizationInformationParams {
	var ()
	return &GetAssigneeAuthorizationInformationParams{
		HTTPClient: client,
	}
}

/*GetAssigneeAuthorizationInformationParams contains all the parameters to send to the API endpoint
for the get assignee authorization information operation typically these are written to a http.Request
*/
type GetAssigneeAuthorizationInformationParams struct {

	/*Input*/
	Input *models.GetAssigneeAuthorizationInformationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) WithTimeout(timeout time.Duration) *GetAssigneeAuthorizationInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) WithContext(ctx context.Context) *GetAssigneeAuthorizationInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) WithHTTPClient(client *http.Client) *GetAssigneeAuthorizationInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) WithInput(input *models.GetAssigneeAuthorizationInformationRequest) *GetAssigneeAuthorizationInformationParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get assignee authorization information params
func (o *GetAssigneeAuthorizationInformationParams) SetInput(input *models.GetAssigneeAuthorizationInformationRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssigneeAuthorizationInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

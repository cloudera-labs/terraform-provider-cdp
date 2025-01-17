// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeDbcDiagnosticDataJobResponse Response object for the describeDbcDiagnosticDataJobRequest method.
//
// swagger:model DescribeDbcDiagnosticDataJobResponse
type DescribeDbcDiagnosticDataJobResponse struct {

	// The job belonging to the particular Database Catalog
	Job *DbcDiagnosticDataJob `json:"job,omitempty"`
}

// Validate validates this describe dbc diagnostic data job response
func (m *DescribeDbcDiagnosticDataJobResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDbcDiagnosticDataJobResponse) validateJob(formats strfmt.Registry) error {
	if swag.IsZero(m.Job) { // not required
		return nil
	}

	if m.Job != nil {
		if err := m.Job.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe dbc diagnostic data job response based on the context it is used
func (m *DescribeDbcDiagnosticDataJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDbcDiagnosticDataJobResponse) contextValidateJob(ctx context.Context, formats strfmt.Registry) error {

	if m.Job != nil {
		if err := m.Job.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("job")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("job")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDbcDiagnosticDataJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDbcDiagnosticDataJobResponse) UnmarshalBinary(b []byte) error {
	var res DescribeDbcDiagnosticDataJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

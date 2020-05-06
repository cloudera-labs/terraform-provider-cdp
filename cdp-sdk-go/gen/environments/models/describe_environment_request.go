// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DescribeEnvironmentRequest Request object for a describe environment request.
//
// swagger:model DescribeEnvironmentRequest
type DescribeEnvironmentRequest struct {

	// The names or CRN of the environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`
}

// Validate validates this describe environment request
func (m *DescribeEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeEnvironmentRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res DescribeEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

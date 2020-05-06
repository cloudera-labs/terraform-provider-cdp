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

// GetTruststorePasswordRequest Request object for Datalake truststore password.
//
// swagger:model GetTruststorePasswordRequest
type GetTruststorePasswordRequest struct {

	// The CRN of the Datalake
	// Required: true
	DatalakeCrn *string `json:"datalakeCrn"`
}

// Validate validates this get truststore password request
func (m *GetTruststorePasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTruststorePasswordRequest) validateDatalakeCrn(formats strfmt.Registry) error {

	if err := validate.Required("datalakeCrn", "body", m.DatalakeCrn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTruststorePasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTruststorePasswordRequest) UnmarshalBinary(b []byte) error {
	var res GetTruststorePasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

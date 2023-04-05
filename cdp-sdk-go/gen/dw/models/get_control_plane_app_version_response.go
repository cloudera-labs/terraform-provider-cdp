// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetControlPlaneAppVersionResponse Response object for the GetControlPlaneAppVersion method.
//
// swagger:model GetControlPlaneAppVersionResponse
type GetControlPlaneAppVersionResponse struct {

	// The version of the Cloudera Data Warehouse control plane app.
	// Required: true
	ControlPlaneAppVersion *string `json:"controlPlaneAppVersion"`
}

// Validate validates this get control plane app version response
func (m *GetControlPlaneAppVersionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlPlaneAppVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetControlPlaneAppVersionResponse) validateControlPlaneAppVersion(formats strfmt.Registry) error {

	if err := validate.Required("controlPlaneAppVersion", "body", m.ControlPlaneAppVersion); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get control plane app version response based on context it is used
func (m *GetControlPlaneAppVersionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetControlPlaneAppVersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetControlPlaneAppVersionResponse) UnmarshalBinary(b []byte) error {
	var res GetControlPlaneAppVersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
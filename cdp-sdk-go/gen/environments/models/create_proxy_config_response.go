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

// CreateProxyConfigResponse Response object for a create proxy config request.
//
// swagger:model CreateProxyConfigResponse
type CreateProxyConfigResponse struct {

	// The proxy config object.
	// Required: true
	ProxyConfig *ProxyConfig `json:"proxyConfig"`
}

// Validate validates this create proxy config response
func (m *CreateProxyConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxyConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProxyConfigResponse) validateProxyConfig(formats strfmt.Registry) error {

	if err := validate.Required("proxyConfig", "body", m.ProxyConfig); err != nil {
		return err
	}

	if m.ProxyConfig != nil {
		if err := m.ProxyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxyConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateProxyConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProxyConfigResponse) UnmarshalBinary(b []byte) error {
	var res CreateProxyConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

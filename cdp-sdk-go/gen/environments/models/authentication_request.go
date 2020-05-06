// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthenticationRequest Additional SSH key authentication configuration for accessing cluster node.
//
// swagger:model AuthenticationRequest
type AuthenticationRequest struct {

	// Public SSH key string. Mutually exclusive with publicKeyId.
	PublicKey string `json:"publicKey,omitempty"`

	// Public SSH key ID already registered in the cloud provider. Mutually exclusive with publicKey.
	PublicKeyID string `json:"publicKeyId,omitempty"`
}

// Validate validates this authentication request
func (m *AuthenticationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationRequest) UnmarshalBinary(b []byte) error {
	var res AuthenticationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

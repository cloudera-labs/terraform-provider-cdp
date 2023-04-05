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

// DeleteAuditCredentialRequest Request object for a delete audit credential request.
//
// swagger:model DeleteAuditCredentialRequest
type DeleteAuditCredentialRequest struct {

	// The name or CRN of the credential.
	// Required: true
	CredentialName *string `json:"credentialName"`
}

// Validate validates this delete audit credential request
func (m *DeleteAuditCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteAuditCredentialRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete audit credential request based on context it is used
func (m *DeleteAuditCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAuditCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAuditCredentialRequest) UnmarshalBinary(b []byte) error {
	var res DeleteAuditCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
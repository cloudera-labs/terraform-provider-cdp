// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentMetadata Environment metadata response object for the workspace backup.
//
// swagger:model EnvironmentMetadata
type EnvironmentMetadata struct {

	// The environment CRN.
	Crn string `json:"crn,omitempty"`

	// The environment name.
	Name string `json:"name,omitempty"`

	// The environment status. Possible status are `unavailable` and `available`.
	Status string `json:"status,omitempty"`
}

// Validate validates this environment metadata
func (m *EnvironmentMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this environment metadata based on context it is used
func (m *EnvironmentMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentMetadata) UnmarshalBinary(b []byte) error {
	var res EnvironmentMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
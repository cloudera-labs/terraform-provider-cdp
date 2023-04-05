// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentTags Environment tags object containing the tag values defined for the environment.
//
// swagger:model EnvironmentTags
type EnvironmentTags struct {

	// Map of tag names to values, for default tags.
	Defaults map[string]string `json:"defaults,omitempty"`

	// Map of tag names to values, for user-defined tags.
	UserDefined map[string]string `json:"userDefined,omitempty"`
}

// Validate validates this environment tags
func (m *EnvironmentTags) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this environment tags based on context it is used
func (m *EnvironmentTags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentTags) UnmarshalBinary(b []byte) error {
	var res EnvironmentTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
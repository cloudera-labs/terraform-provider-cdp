// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageReferenceSet Detailed set of cloud providers region to image mappings.
//
// swagger:model ImageReferenceSet
type ImageReferenceSet struct {

	// AWS-related region-to-image mappings.
	Aws map[string]string `json:"aws,omitempty"`

	// Azure-related region-to-image mappings.
	Azure map[string]string `json:"azure,omitempty"`

	// GCP-related region-to-image mappings.
	Gcp map[string]string `json:"gcp,omitempty"`
}

// Validate validates this image reference set
func (m *ImageReferenceSet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image reference set based on context it is used
func (m *ImageReferenceSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageReferenceSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageReferenceSet) UnmarshalBinary(b []byte) error {
	var res ImageReferenceSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
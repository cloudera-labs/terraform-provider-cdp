// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetConfigFilesResponse Response object for a Datalake client configs response.
//
// swagger:model GetConfigFilesResponse
type GetConfigFilesResponse struct {

	// config files for the data lake including root cert and krb.conf
	File string `json:"file,omitempty"`

	// an opaque string (sha or similar) to detect config changes.
	VersionHash string `json:"versionHash,omitempty"`
}

// Validate validates this get config files response
func (m *GetConfigFilesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetConfigFilesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConfigFilesResponse) UnmarshalBinary(b []byte) error {
	var res GetConfigFilesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

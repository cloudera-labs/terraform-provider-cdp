// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FreeipaDiagnosticsCollectionDetailsResponse Response object for diagnostic collection details.
//
// swagger:model FreeipaDiagnosticsCollectionDetailsResponse
type FreeipaDiagnosticsCollectionDetailsResponse struct {

	// Account Id that was used for the diagnostics collection.
	AccountID string `json:"accountId,omitempty"`

	// Case number for the diagnostics collection.
	Case string `json:"case,omitempty"`

	// Version of the cluster that was used for the diagnostics collection.
	ClusterVersion string `json:"clusterVersion,omitempty"`

	// Description of the diagnostics collection.
	Description string `json:"description,omitempty"`

	// Destination type of the diagnostics collection.
	Destination string `json:"destination,omitempty"`

	// Output destination of the diagnostics collection.
	Output string `json:"output,omitempty"`

	// Crn of the Freeipa cluster.
	ResourceCrn string `json:"resourceCrn,omitempty"`

	// Status message of a failed diagnostics.
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this freeipa diagnostics collection details response
func (m *FreeipaDiagnosticsCollectionDetailsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this freeipa diagnostics collection details response based on context it is used
func (m *FreeipaDiagnosticsCollectionDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FreeipaDiagnosticsCollectionDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeipaDiagnosticsCollectionDetailsResponse) UnmarshalBinary(b []byte) error {
	var res FreeipaDiagnosticsCollectionDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
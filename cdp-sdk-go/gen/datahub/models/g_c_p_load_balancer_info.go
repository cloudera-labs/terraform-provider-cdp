// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPLoadBalancerInfo The GCP load balancer information.
//
// swagger:model GCPLoadBalancerInfo
type GCPLoadBalancerInfo struct {

	// The GCP Load Balancer resource ID.
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this g c p load balancer info
func (m *GCPLoadBalancerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g c p load balancer info based on context it is used
func (m *GCPLoadBalancerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPLoadBalancerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPLoadBalancerInfo) UnmarshalBinary(b []byte) error {
	var res GCPLoadBalancerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
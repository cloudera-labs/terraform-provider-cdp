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

// RepairFreeipaRequest Request object for repairing the FreeIPA servers.
//
// swagger:model RepairFreeipaRequest
type RepairFreeipaRequest struct {

	// The environment name or CRN of the FreeIPA to repair
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Force the repair even if the status if the FreeIPA nodes are good.
	Force *bool `json:"force,omitempty"`

	// The instance Ids to repair. If not provided then all instances are looked at for repair.
	Instances []string `json:"instances"`
}

// Validate validates this repair freeipa request
func (m *RepairFreeipaRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepairFreeipaRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepairFreeipaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepairFreeipaRequest) UnmarshalBinary(b []byte) error {
	var res RepairFreeipaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

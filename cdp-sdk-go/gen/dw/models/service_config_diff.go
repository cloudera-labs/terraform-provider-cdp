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

// ServiceConfigDiff Differences between two service configurations.
//
// swagger:model ServiceConfigDiff
type ServiceConfigDiff struct {

	// Differences between the application specific configurations.
	ApplicationConfigsDiffs map[string]ApplicationConfigDiff `json:"applicationConfigsDiffs,omitempty"`

	// Differences between the common configurations.
	CommonConfigsDiff *ApplicationConfigDiff `json:"commonConfigsDiff,omitempty"`

	// Has any of the configurations changed?
	IsChanged bool `json:"isChanged,omitempty"`
}

// Validate validates this service config diff
func (m *ServiceConfigDiff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationConfigsDiffs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonConfigsDiff(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceConfigDiff) validateApplicationConfigsDiffs(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationConfigsDiffs) { // not required
		return nil
	}

	for k := range m.ApplicationConfigsDiffs {

		if err := validate.Required("applicationConfigsDiffs"+"."+k, "body", m.ApplicationConfigsDiffs[k]); err != nil {
			return err
		}
		if val, ok := m.ApplicationConfigsDiffs[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applicationConfigsDiffs" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applicationConfigsDiffs" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceConfigDiff) validateCommonConfigsDiff(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonConfigsDiff) { // not required
		return nil
	}

	if m.CommonConfigsDiff != nil {
		if err := m.CommonConfigsDiff.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commonConfigsDiff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commonConfigsDiff")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service config diff based on the context it is used
func (m *ServiceConfigDiff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationConfigsDiffs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommonConfigsDiff(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceConfigDiff) contextValidateApplicationConfigsDiffs(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.ApplicationConfigsDiffs {

		if val, ok := m.ApplicationConfigsDiffs[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ServiceConfigDiff) contextValidateCommonConfigsDiff(ctx context.Context, formats strfmt.Registry) error {

	if m.CommonConfigsDiff != nil {
		if err := m.CommonConfigsDiff.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commonConfigsDiff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commonConfigsDiff")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceConfigDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceConfigDiff) UnmarshalBinary(b []byte) error {
	var res ServiceConfigDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

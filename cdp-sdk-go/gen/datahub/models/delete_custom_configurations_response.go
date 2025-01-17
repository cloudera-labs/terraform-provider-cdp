// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteCustomConfigurationsResponse The response object for the delete custom configurations request.
//
// swagger:model DeleteCustomConfigurationsResponse
type DeleteCustomConfigurationsResponse struct {

	// The deleted custom configurations.
	// Required: true
	CustomConfigurations []*CustomConfigurations `json:"customConfigurations"`
}

// Validate validates this delete custom configurations response
func (m *DeleteCustomConfigurationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteCustomConfigurationsResponse) validateCustomConfigurations(formats strfmt.Registry) error {

	if err := validate.Required("customConfigurations", "body", m.CustomConfigurations); err != nil {
		return err
	}

	for i := 0; i < len(m.CustomConfigurations); i++ {
		if swag.IsZero(m.CustomConfigurations[i]) { // not required
			continue
		}

		if m.CustomConfigurations[i] != nil {
			if err := m.CustomConfigurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete custom configurations response based on the context it is used
func (m *DeleteCustomConfigurationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteCustomConfigurationsResponse) contextValidateCustomConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomConfigurations); i++ {

		if m.CustomConfigurations[i] != nil {
			if err := m.CustomConfigurations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customConfigurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteCustomConfigurationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteCustomConfigurationsResponse) UnmarshalBinary(b []byte) error {
	var res DeleteCustomConfigurationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

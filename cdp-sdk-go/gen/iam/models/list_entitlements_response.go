// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListEntitlementsResponse Response object for a list entitlements request.
//
// swagger:model ListEntitlementsResponse
type ListEntitlementsResponse struct {

	// The granted entitlements.
	// Required: true
	Entitlements []*Entitlement `json:"entitlements"`

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list entitlements response
func (m *ListEntitlementsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEntitlementsResponse) validateEntitlements(formats strfmt.Registry) error {

	if err := validate.Required("entitlements", "body", m.Entitlements); err != nil {
		return err
	}

	for i := 0; i < len(m.Entitlements); i++ {
		if swag.IsZero(m.Entitlements[i]) { // not required
			continue
		}

		if m.Entitlements[i] != nil {
			if err := m.Entitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListEntitlementsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEntitlementsResponse) UnmarshalBinary(b []byte) error {
	var res ListEntitlementsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

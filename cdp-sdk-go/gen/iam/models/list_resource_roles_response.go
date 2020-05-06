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

// ListResourceRolesResponse Response object for a list resource roles request.
//
// swagger:model ListResourceRolesResponse
type ListResourceRolesResponse struct {

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`

	// The list of resource roles. Cannot be empty.
	// Required: true
	ResourceRoles []*ResourceRole `json:"resourceRoles"`
}

// Validate validates this list resource roles response
func (m *ListResourceRolesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListResourceRolesResponse) validateResourceRoles(formats strfmt.Registry) error {

	if err := validate.Required("resourceRoles", "body", m.ResourceRoles); err != nil {
		return err
	}

	for i := 0; i < len(m.ResourceRoles); i++ {
		if swag.IsZero(m.ResourceRoles[i]) { // not required
			continue
		}

		if m.ResourceRoles[i] != nil {
			if err := m.ResourceRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListResourceRolesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListResourceRolesResponse) UnmarshalBinary(b []byte) error {
	var res ListResourceRolesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

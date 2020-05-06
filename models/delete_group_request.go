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

// DeleteGroupRequest Request object for delete group request.
//
// swagger:model DeleteGroupRequest
type DeleteGroupRequest struct {

	// The name or CRN of the group to delete.
	// Required: true
	GroupName *string `json:"groupName"`
}

// Validate validates this delete group request
func (m *DeleteGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteGroupRequest) validateGroupName(formats strfmt.Registry) error {

	if err := validate.Required("groupName", "body", m.GroupName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteGroupRequest) UnmarshalBinary(b []byte) error {
	var res DeleteGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

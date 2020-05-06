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

// ListGroupsForUserResponse Response object for a list groups for user request.
//
// swagger:model ListGroupsForUserResponse
type ListGroupsForUserResponse struct {

	// The list of groups.
	// Required: true
	GroupCrns []string `json:"groupCrns"`

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list groups for user response
func (m *ListGroupsForUserResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupCrns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListGroupsForUserResponse) validateGroupCrns(formats strfmt.Registry) error {

	if err := validate.Required("groupCrns", "body", m.GroupCrns); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListGroupsForUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListGroupsForUserResponse) UnmarshalBinary(b []byte) error {
	var res ListGroupsForUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

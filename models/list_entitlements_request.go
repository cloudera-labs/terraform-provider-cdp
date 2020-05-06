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

// ListEntitlementsRequest Request object for a list entitlements request.
//
// swagger:model ListEntitlementsRequest
type ListEntitlementsRequest struct {

	// The size of each page.
	// Maximum: 100
	// Minimum: 1
	PageSize int32 `json:"pageSize,omitempty"`

	// A token to specify where to start paginating. This is the nextToken from a previously truncated response.
	StartingToken string `json:"startingToken,omitempty"`
}

// Validate validates this list entitlements request
func (m *ListEntitlementsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListEntitlementsRequest) validatePageSize(formats strfmt.Registry) error {

	if swag.IsZero(m.PageSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("pageSize", "body", int64(m.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "body", int64(m.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListEntitlementsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEntitlementsRequest) UnmarshalBinary(b []byte) error {
	var res ListEntitlementsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ProgressResponse Response object for the progress updates.
//
// swagger:model ProgressResponse
type ProgressResponse struct {

	// The latest logs of the activity.
	Logs string `json:"logs,omitempty"`

	// The name of the activity.
	// Required: true
	Name *string `json:"name"`

	// The list of remaining progress responses.
	ProgressResponses []*ProgressResponse `json:"progressResponses"`

	// The latest result of the activity.
	// Required: true
	Result *string `json:"result"`

	// The latest status of the activity.
	// Required: true
	Status *string `json:"status"`

	// The latest summary of the activity.
	Summary string `json:"summary,omitempty"`
}

// Validate validates this progress response
func (m *ProgressResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgressResponses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgressResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProgressResponse) validateProgressResponses(formats strfmt.Registry) error {
	if swag.IsZero(m.ProgressResponses) { // not required
		return nil
	}

	for i := 0; i < len(m.ProgressResponses); i++ {
		if swag.IsZero(m.ProgressResponses[i]) { // not required
			continue
		}

		if m.ProgressResponses[i] != nil {
			if err := m.ProgressResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("progressResponses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("progressResponses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProgressResponse) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

func (m *ProgressResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this progress response based on the context it is used
func (m *ProgressResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProgressResponses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgressResponse) contextValidateProgressResponses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProgressResponses); i++ {

		if m.ProgressResponses[i] != nil {
			if err := m.ProgressResponses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("progressResponses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("progressResponses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProgressResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressResponse) UnmarshalBinary(b []byte) error {
	var res ProgressResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
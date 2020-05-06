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

// MachineUser Information about a Cloudera Altus machine user.
//
// swagger:model MachineUser
type MachineUser struct {

	// The date when this machine user record was created.
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The CRN of the user.
	// Required: true
	Crn *string `json:"crn"`

	// The machine user name.
	// Required: true
	MachineUserName *string `json:"machineUserName"`

	// The username used in all the workload clusters of the machine user.
	WorkloadUsername string `json:"workloadUsername,omitempty"`
}

// Validate validates this machine user
func (m *MachineUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineUser) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MachineUser) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *MachineUser) validateMachineUserName(formats strfmt.Registry) error {

	if err := validate.Required("machineUserName", "body", m.MachineUserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineUser) UnmarshalBinary(b []byte) error {
	var res MachineUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

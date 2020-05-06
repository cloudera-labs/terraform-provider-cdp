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

// IDBrokerSyncStatus Status of an ID Broker mappings sync operation.
//
// swagger:model IdBrokerSyncStatus
type IDBrokerSyncStatus struct {

	// The date when the mappings sync completed or was terminated. Omitted if status is NEVER_RUN or RUNNING.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The detail of the error. Omitted if status is not FAILED.
	ErrorDetail string `json:"errorDetail,omitempty"`

	// The date when the mappings sync started executing. Omitted if status is NEVER_RUN.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// The mappings sync summary status.
	// Required: true
	Status SyncStatus `json:"status"`
}

// Validate validates this Id broker sync status
func (m *IDBrokerSyncStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
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

func (m *IDBrokerSyncStatus) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IDBrokerSyncStatus) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IDBrokerSyncStatus) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IDBrokerSyncStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDBrokerSyncStatus) UnmarshalBinary(b []byte) error {
	var res IDBrokerSyncStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

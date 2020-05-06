// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateAccessKeyRequest Request object for an update access key request.
//
// swagger:model UpdateAccessKeyRequest
type UpdateAccessKeyRequest struct {

	// The ID of the access key to update.
	// Required: true
	AccessKeyID *string `json:"accessKeyId"`

	// The status to assign to the access key.
	// Required: true
	// Enum: [ACTIVE INACTIVE]
	Status *string `json:"status"`
}

// Validate validates this update access key request
func (m *UpdateAccessKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyID(formats); err != nil {
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

func (m *UpdateAccessKeyRequest) validateAccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("accessKeyId", "body", m.AccessKeyID); err != nil {
		return err
	}

	return nil
}

var updateAccessKeyRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateAccessKeyRequestTypeStatusPropEnum = append(updateAccessKeyRequestTypeStatusPropEnum, v)
	}
}

const (

	// UpdateAccessKeyRequestStatusACTIVE captures enum value "ACTIVE"
	UpdateAccessKeyRequestStatusACTIVE string = "ACTIVE"

	// UpdateAccessKeyRequestStatusINACTIVE captures enum value "INACTIVE"
	UpdateAccessKeyRequestStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *UpdateAccessKeyRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, updateAccessKeyRequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateAccessKeyRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAccessKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAccessKeyRequest) UnmarshalBinary(b []byte) error {
	var res UpdateAccessKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

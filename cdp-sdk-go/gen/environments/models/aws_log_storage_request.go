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

// AwsLogStorageRequest AWS storage configuration for cluster and audit logs.
//
// swagger:model AwsLogStorageRequest
type AwsLogStorageRequest struct {

	// The AWS instance profile that which contains the necessary permissions to access the S3 storage location.
	// Required: true
	InstanceProfile *string `json:"instanceProfile"`

	// The base location to store logs in S3. This should be an s3a:// url.
	// Required: true
	StorageLocationBase *string `json:"storageLocationBase"`
}

// Validate validates this aws log storage request
func (m *AwsLogStorageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageLocationBase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsLogStorageRequest) validateInstanceProfile(formats strfmt.Registry) error {

	if err := validate.Required("instanceProfile", "body", m.InstanceProfile); err != nil {
		return err
	}

	return nil
}

func (m *AwsLogStorageRequest) validateStorageLocationBase(formats strfmt.Registry) error {

	if err := validate.Required("storageLocationBase", "body", m.StorageLocationBase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsLogStorageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsLogStorageRequest) UnmarshalBinary(b []byte) error {
	var res AwsLogStorageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

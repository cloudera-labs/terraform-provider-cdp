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

// AzureLogStorageRequest Azure storage configuration for cluster and audit logs.
//
// swagger:model AzureLogStorageRequest
type AzureLogStorageRequest struct {

	// The managed identity associated with the logger. This identity should have Storage Blob Data Contributor role on the given storage account.
	// Required: true
	ManagedIdentity *string `json:"managedIdentity"`

	// The storage location to use. The location has to be in the following format abfs://filesystem@storage-account-name.dfs.core.windows.net. The filesystem must already exist and the storage account must be StorageV2.
	// Required: true
	StorageLocationBase *string `json:"storageLocationBase"`
}

// Validate validates this azure log storage request
func (m *AzureLogStorageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateManagedIdentity(formats); err != nil {
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

func (m *AzureLogStorageRequest) validateManagedIdentity(formats strfmt.Registry) error {

	if err := validate.Required("managedIdentity", "body", m.ManagedIdentity); err != nil {
		return err
	}

	return nil
}

func (m *AzureLogStorageRequest) validateStorageLocationBase(formats strfmt.Registry) error {

	if err := validate.Required("storageLocationBase", "body", m.StorageLocationBase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureLogStorageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureLogStorageRequest) UnmarshalBinary(b []byte) error {
	var res AzureLogStorageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

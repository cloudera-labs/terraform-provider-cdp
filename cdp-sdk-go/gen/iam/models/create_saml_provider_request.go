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

// CreateSamlProviderRequest Request object for creating SAML provider request.
//
// swagger:model CreateSamlProviderRequest
type CreateSamlProviderRequest struct {

	// SAML metadata document XML file. Length of meta data document cannot be more than 200 KB (200,000 bytes).
	// Required: true
	// Max Length: 200000
	SamlMetadataDocument *string `json:"samlMetadataDocument"`

	// The name of SAML provider. The name must be unique, must have a maximum of 128 characters, and must contain only alphanumeric characters, "-" and "_". Names are are not case-sensitive.
	// Required: true
	SamlProviderName *string `json:"samlProviderName"`

	// Whether to sync group information for users federated with this SAML provider. Group membership can be passed using the https://altus.cloudera.com/SAML/Attributes/groups SAML assertion. The default is to synchronize group membership.
	SyncGroupsOnLogin bool `json:"syncGroupsOnLogin,omitempty"`
}

// Validate validates this create saml provider request
func (m *CreateSamlProviderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamlMetadataDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamlProviderName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSamlProviderRequest) validateSamlMetadataDocument(formats strfmt.Registry) error {

	if err := validate.Required("samlMetadataDocument", "body", m.SamlMetadataDocument); err != nil {
		return err
	}

	if err := validate.MaxLength("samlMetadataDocument", "body", string(*m.SamlMetadataDocument), 200000); err != nil {
		return err
	}

	return nil
}

func (m *CreateSamlProviderRequest) validateSamlProviderName(formats strfmt.Registry) error {

	if err := validate.Required("samlProviderName", "body", m.SamlProviderName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSamlProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSamlProviderRequest) UnmarshalBinary(b []byte) error {
	var res CreateSamlProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

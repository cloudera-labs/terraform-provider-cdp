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

// CreateAzureCredentialRequest Request object for a create Azure credential request.
//
// swagger:model CreateAzureCredentialRequest
type CreateAzureCredentialRequest struct {

	// app based
	// Required: true
	AppBased *CreateAzureCredentialRequestAppBased `json:"appBased"`

	// The name of the credential.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// A description for the credential.
	Description string `json:"description,omitempty"`

	// The Azure subscription ID.
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`

	// The Azure AD tenant ID for the Azure subscription.
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this create azure credential request
func (m *CreateAzureCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppBased(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequest) validateAppBased(formats strfmt.Registry) error {

	if err := validate.Required("appBased", "body", m.AppBased); err != nil {
		return err
	}

	if m.AppBased != nil {
		if err := m.AppBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureCredentialRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureCredentialRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureCredentialRequest) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureCredentialRequest) UnmarshalBinary(b []byte) error {
	var res CreateAzureCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateAzureCredentialRequestAppBased Additional configurations needed for app-based authentication.
//
// swagger:model CreateAzureCredentialRequestAppBased
type CreateAzureCredentialRequestAppBased struct {

	// The id of the application registered in Azure.
	// Required: true
	ApplicationID *string `json:"applicationId"`

	// The client secret key (also referred to as application password) for the registered application.
	// Required: true
	SecretKey *string `json:"secretKey"`
}

// Validate validates this create azure credential request app based
func (m *CreateAzureCredentialRequestAppBased) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequestAppBased) validateApplicationID(formats strfmt.Registry) error {

	if err := validate.Required("appBased"+"."+"applicationId", "body", m.ApplicationID); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureCredentialRequestAppBased) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("appBased"+"."+"secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureCredentialRequestAppBased) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureCredentialRequestAppBased) UnmarshalBinary(b []byte) error {
	var res CreateAzureCredentialRequestAppBased
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

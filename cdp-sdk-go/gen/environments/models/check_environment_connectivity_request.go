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

// CheckEnvironmentConnectivityRequest Request object to check connectivity to private cloud environment.
//
// swagger:model CheckEnvironmentConnectivityRequest
type CheckEnvironmentConnectivityRequest struct {

	// The address of the Cloudera Manager managing the Datalake cluster.
	// Required: true
	Address *string `json:"address"`

	// A string (text or json) used to authenticate to the Cloudera Manager.
	// Required: true
	AuthenticationToken *string `json:"authenticationToken"`

	// How to interpret the authenticationToken field. Defaults to CLEARTEXT_PASSWORD.
	// Enum: [CLEARTEXT_PASSWORD]
	AuthenticationTokenType string `json:"authenticationTokenType,omitempty"`

	// The name of the cluster(s) to use as a Datalake for the environment.
	ClusterNames []string `json:"clusterNames"`

	// User name for accessing the Cloudera Manager.
	// Required: true
	User *string `json:"user"`
}

// Validate validates this check environment connectivity request
func (m *CheckEnvironmentConnectivityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationTokenType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckEnvironmentConnectivityRequest) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *CheckEnvironmentConnectivityRequest) validateAuthenticationToken(formats strfmt.Registry) error {

	if err := validate.Required("authenticationToken", "body", m.AuthenticationToken); err != nil {
		return err
	}

	return nil
}

var checkEnvironmentConnectivityRequestTypeAuthenticationTokenTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLEARTEXT_PASSWORD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checkEnvironmentConnectivityRequestTypeAuthenticationTokenTypePropEnum = append(checkEnvironmentConnectivityRequestTypeAuthenticationTokenTypePropEnum, v)
	}
}

const (

	// CheckEnvironmentConnectivityRequestAuthenticationTokenTypeCLEARTEXTPASSWORD captures enum value "CLEARTEXT_PASSWORD"
	CheckEnvironmentConnectivityRequestAuthenticationTokenTypeCLEARTEXTPASSWORD string = "CLEARTEXT_PASSWORD"
)

// prop value enum
func (m *CheckEnvironmentConnectivityRequest) validateAuthenticationTokenTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, checkEnvironmentConnectivityRequestTypeAuthenticationTokenTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CheckEnvironmentConnectivityRequest) validateAuthenticationTokenType(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationTokenType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTokenTypeEnum("authenticationTokenType", "body", m.AuthenticationTokenType); err != nil {
		return err
	}

	return nil
}

func (m *CheckEnvironmentConnectivityRequest) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckEnvironmentConnectivityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckEnvironmentConnectivityRequest) UnmarshalBinary(b []byte) error {
	var res CheckEnvironmentConnectivityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

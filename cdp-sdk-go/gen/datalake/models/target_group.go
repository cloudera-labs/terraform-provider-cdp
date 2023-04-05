// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TargetGroup The target instance group the load balancer routes traffic to.
//
// swagger:model TargetGroup
type TargetGroup struct {

	// The AWS-specific target group configuration for the load balancer.
	AwsTargetGroupConfiguration *AWSTargetGroupConfiguration `json:"awsTargetGroupConfiguration,omitempty"`

	// The Azure-specific target group configuration for the load balancer.
	AzureTargetGroupConfiguration *AzureTargetGroupConfiguration `json:"azureTargetGroupConfiguration,omitempty"`

	// The GCP-specific target group configuration for the load balancer.
	GcpTargetGroupConfiguration *GCPTargetGroupConfiguration `json:"gcpTargetGroupConfiguration,omitempty"`

	// The port through which the load balancer receives and forwards traffic to the associated targets.
	Port int32 `json:"port,omitempty"`

	// The IDs for the target instances receiving traffic from the load balancer on the defined port.
	// Unique: true
	TargetInstances []string `json:"targetInstances"`
}

// Validate validates this target group
func (m *TargetGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsTargetGroupConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureTargetGroupConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpTargetGroupConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetGroup) validateAwsTargetGroupConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsTargetGroupConfiguration) { // not required
		return nil
	}

	if m.AwsTargetGroupConfiguration != nil {
		if err := m.AwsTargetGroupConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *TargetGroup) validateAzureTargetGroupConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureTargetGroupConfiguration) { // not required
		return nil
	}

	if m.AzureTargetGroupConfiguration != nil {
		if err := m.AzureTargetGroupConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *TargetGroup) validateGcpTargetGroupConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpTargetGroupConfiguration) { // not required
		return nil
	}

	if m.GcpTargetGroupConfiguration != nil {
		if err := m.GcpTargetGroupConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *TargetGroup) validateTargetInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetInstances) { // not required
		return nil
	}

	if err := validate.UniqueItems("targetInstances", "body", m.TargetInstances); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this target group based on the context it is used
func (m *TargetGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsTargetGroupConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureTargetGroupConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpTargetGroupConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetGroup) contextValidateAwsTargetGroupConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsTargetGroupConfiguration != nil {
		if err := m.AwsTargetGroupConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *TargetGroup) contextValidateAzureTargetGroupConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureTargetGroupConfiguration != nil {
		if err := m.AzureTargetGroupConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *TargetGroup) contextValidateGcpTargetGroupConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpTargetGroupConfiguration != nil {
		if err := m.GcpTargetGroupConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpTargetGroupConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpTargetGroupConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetGroup) UnmarshalBinary(b []byte) error {
	var res TargetGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
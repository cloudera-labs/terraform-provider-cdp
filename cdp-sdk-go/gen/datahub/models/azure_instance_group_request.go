// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureInstanceGroupRequest Configurations for instance group
//
// swagger:model AzureInstanceGroupRequest
type AzureInstanceGroupRequest struct {

	// The attached volume configuration. This does not include root volume.
	// Required: true
	AttachedVolumeConfiguration []*AttachedVolumeRequest `json:"attachedVolumeConfiguration"`

	// The instance group name.
	// Required: true
	InstanceGroupName *string `json:"instanceGroupName"`

	// The instance group type.
	// Required: true
	InstanceGroupType *string `json:"instanceGroupType"`

	// The cloud provider specific instance type to be used.
	// Required: true
	InstanceType *string `json:"instanceType"`

	// Number of instances in the instance group
	// Required: true
	NodeCount *int32 `json:"nodeCount"`

	// The names or CRNs of the recipes that would be applied to the instance group.
	RecipeNames []string `json:"recipeNames"`

	// Recovery mode for the instance group.
	RecoveryMode string `json:"recoveryMode,omitempty"`

	// The root volume size.
	// Required: true
	RootVolumeSize *int32 `json:"rootVolumeSize"`

	// Percentage of spot instances in instance group.
	// Maximum: 100
	// Minimum: 0
	SpotPercentage *int32 `json:"spotPercentage,omitempty"`
}

// Validate validates this azure instance group request
func (m *AzureInstanceGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedVolumeConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootVolumeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpotPercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureInstanceGroupRequest) validateAttachedVolumeConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("attachedVolumeConfiguration", "body", m.AttachedVolumeConfiguration); err != nil {
		return err
	}

	for i := 0; i < len(m.AttachedVolumeConfiguration); i++ {
		if swag.IsZero(m.AttachedVolumeConfiguration[i]) { // not required
			continue
		}

		if m.AttachedVolumeConfiguration[i] != nil {
			if err := m.AttachedVolumeConfiguration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachedVolumeConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateInstanceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupName", "body", m.InstanceGroupName); err != nil {
		return err
	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateInstanceGroupType(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupType", "body", m.InstanceGroupType); err != nil {
		return err
	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateRootVolumeSize(formats strfmt.Registry) error {

	if err := validate.Required("rootVolumeSize", "body", m.RootVolumeSize); err != nil {
		return err
	}

	return nil
}

func (m *AzureInstanceGroupRequest) validateSpotPercentage(formats strfmt.Registry) error {

	if swag.IsZero(m.SpotPercentage) { // not required
		return nil
	}

	if err := validate.MinimumInt("spotPercentage", "body", int64(*m.SpotPercentage), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("spotPercentage", "body", int64(*m.SpotPercentage), 100, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureInstanceGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureInstanceGroupRequest) UnmarshalBinary(b []byte) error {
	var res AzureInstanceGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

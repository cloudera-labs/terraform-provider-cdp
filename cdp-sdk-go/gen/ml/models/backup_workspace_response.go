// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupWorkspaceResponse The response object for workspace backup.
//
// swagger:model BackupWorkspaceResponse
type BackupWorkspaceResponse struct {

	// CRN of the backup generated.
	BackupCrn string `json:"backupCrn,omitempty"`

	// Name of the vault where backup related volumes are stored.
	VaultName string `json:"vaultName,omitempty"`

	// The metadata of the workflow trigerred for the backup request.
	WorkflowMetadata *WorkflowMetadata `json:"workflowMetadata,omitempty"`
}

// Validate validates this backup workspace response
func (m *BackupWorkspaceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkflowMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupWorkspaceResponse) validateWorkflowMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkflowMetadata) { // not required
		return nil
	}

	if m.WorkflowMetadata != nil {
		if err := m.WorkflowMetadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflowMetadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup workspace response based on the context it is used
func (m *BackupWorkspaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkflowMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupWorkspaceResponse) contextValidateWorkflowMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkflowMetadata != nil {
		if err := m.WorkflowMetadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflowMetadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workflowMetadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res BackupWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
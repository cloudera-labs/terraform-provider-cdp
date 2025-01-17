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

// HbaseBackupRestoreState The state of each HBase backup/restore operation.
//
// swagger:model HbaseBackupRestoreState
type HbaseBackupRestoreState struct {

	// The status of the ATLAS_ENTITY_AUDIT_EVENTS_TABLE backup/restore.
	// Required: true
	AtlasEntityAuditEventTable *BackupRestoreOperationStatus `json:"atlasEntityAuditEventTable"`

	// The status of the ATLAS_JANUS_TABLE backup/restore.
	// Required: true
	AtlasJanusTable *BackupRestoreOperationStatus `json:"atlasJanusTable"`
}

// Validate validates this hbase backup restore state
func (m *HbaseBackupRestoreState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtlasEntityAuditEventTable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAtlasJanusTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HbaseBackupRestoreState) validateAtlasEntityAuditEventTable(formats strfmt.Registry) error {

	if err := validate.Required("atlasEntityAuditEventTable", "body", m.AtlasEntityAuditEventTable); err != nil {
		return err
	}

	if m.AtlasEntityAuditEventTable != nil {
		if err := m.AtlasEntityAuditEventTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atlasEntityAuditEventTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atlasEntityAuditEventTable")
			}
			return err
		}
	}

	return nil
}

func (m *HbaseBackupRestoreState) validateAtlasJanusTable(formats strfmt.Registry) error {

	if err := validate.Required("atlasJanusTable", "body", m.AtlasJanusTable); err != nil {
		return err
	}

	if m.AtlasJanusTable != nil {
		if err := m.AtlasJanusTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atlasJanusTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atlasJanusTable")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hbase backup restore state based on the context it is used
func (m *HbaseBackupRestoreState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAtlasEntityAuditEventTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAtlasJanusTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HbaseBackupRestoreState) contextValidateAtlasEntityAuditEventTable(ctx context.Context, formats strfmt.Registry) error {

	if m.AtlasEntityAuditEventTable != nil {
		if err := m.AtlasEntityAuditEventTable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atlasEntityAuditEventTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atlasEntityAuditEventTable")
			}
			return err
		}
	}

	return nil
}

func (m *HbaseBackupRestoreState) contextValidateAtlasJanusTable(ctx context.Context, formats strfmt.Registry) error {

	if m.AtlasJanusTable != nil {
		if err := m.AtlasJanusTable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("atlasJanusTable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("atlasJanusTable")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HbaseBackupRestoreState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HbaseBackupRestoreState) UnmarshalBinary(b []byte) error {
	var res HbaseBackupRestoreState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

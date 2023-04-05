// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectCmDiagnosticsRequest Request object for collecting DataLake diagnostics.
//
// swagger:model CollectCmDiagnosticsRequest
type CollectCmDiagnosticsRequest struct {

	// Diagnostics bundle size limit in MB.
	BundleSizeLimit int32 `json:"bundleSizeLimit,omitempty"`

	// Optional support case number in case of SUPPORT destination, otherwise only act as additional data.
	CaseNumber string `json:"caseNumber,omitempty"`

	// CRN of the Datalake cluster.
	// Required: true
	Crn *string `json:"crn"`

	// Additional information / title for the diagnostics collection.
	Description string `json:"description,omitempty"`

	// Destination of the diagnostics collection (Support, Own cloud storage, Engineering or collect only on the nodes)
	// Required: true
	// Enum: [SUPPORT CLOUD_STORAGE LOCAL]
	Destination *string `json:"destination"`

	// Restrict collected logs and metrics (until the provided date timestamp).
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Flag to enable collection of metrics for chart display in CM based diagnostics collection.
	MonitorMetricsCollection *bool `json:"monitorMetricsCollection,omitempty"`

	// Array of roles for which to get logs and metrics. If set, this restricts the roles for log and metrics collection.
	// Unique: true
	Roles []string `json:"roles"`

	// Restrict collected logs and metrics (from the provided date timestamp).
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Enable/disable node level storage validation (can be disabled for example, if you have too many hosts and do not want to do too much parallel writes to s3/abfs)
	StorageValidation *bool `json:"storageValidation,omitempty"`

	// If enabled, required package (cdp-telemetry) will be upgraded or installed on the nodes. (useful if package is not installed or needs to be upgraded) Network is required for this operation.
	UpdatePackage *bool `json:"updatePackage,omitempty"`
}

// Validate validates this collect cm diagnostics request
func (m *CollectCmDiagnosticsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectCmDiagnosticsRequest) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

var collectCmDiagnosticsRequestTypeDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUPPORT","CLOUD_STORAGE","LOCAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		collectCmDiagnosticsRequestTypeDestinationPropEnum = append(collectCmDiagnosticsRequestTypeDestinationPropEnum, v)
	}
}

const (

	// CollectCmDiagnosticsRequestDestinationSUPPORT captures enum value "SUPPORT"
	CollectCmDiagnosticsRequestDestinationSUPPORT string = "SUPPORT"

	// CollectCmDiagnosticsRequestDestinationCLOUDSTORAGE captures enum value "CLOUD_STORAGE"
	CollectCmDiagnosticsRequestDestinationCLOUDSTORAGE string = "CLOUD_STORAGE"

	// CollectCmDiagnosticsRequestDestinationLOCAL captures enum value "LOCAL"
	CollectCmDiagnosticsRequestDestinationLOCAL string = "LOCAL"
)

// prop value enum
func (m *CollectCmDiagnosticsRequest) validateDestinationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, collectCmDiagnosticsRequestTypeDestinationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CollectCmDiagnosticsRequest) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	// value enum
	if err := m.validateDestinationEnum("destination", "body", *m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *CollectCmDiagnosticsRequest) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CollectCmDiagnosticsRequest) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *CollectCmDiagnosticsRequest) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this collect cm diagnostics request based on context it is used
func (m *CollectCmDiagnosticsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectCmDiagnosticsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectCmDiagnosticsRequest) UnmarshalBinary(b []byte) error {
	var res CollectCmDiagnosticsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
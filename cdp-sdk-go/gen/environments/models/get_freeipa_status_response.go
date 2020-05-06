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

// GetFreeipaStatusResponse The overall status of the FreeIPA cluster.
//
// swagger:model GetFreeipaStatusResponse
type GetFreeipaStatusResponse struct {

	// The CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// The name of the Environment
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Status of individual nodes in the FreeIPA cluster
	// Required: true
	Instances map[string]FreeIPANodeStatus `json:"instances"`

	// The overall status of the FreeIPA cluster
	// Required: true
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE STACK_AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED UNREACHABLE UNHEALTHY DELETED_ON_PROVIDER_SIDE UNKNOWN]
	Status *string `json:"status"`
}

// Validate validates this get freeipa status response
func (m *GetFreeipaStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFreeipaStatusResponse) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *GetFreeipaStatusResponse) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *GetFreeipaStatusResponse) validateInstances(formats strfmt.Registry) error {

	for k := range m.Instances {

		if err := validate.Required("instances"+"."+k, "body", m.Instances[k]); err != nil {
			return err
		}
		if val, ok := m.Instances[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var getFreeipaStatusResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","STACK_AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","UNREACHABLE","UNHEALTHY","DELETED_ON_PROVIDER_SIDE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getFreeipaStatusResponseTypeStatusPropEnum = append(getFreeipaStatusResponseTypeStatusPropEnum, v)
	}
}

const (

	// GetFreeipaStatusResponseStatusREQUESTED captures enum value "REQUESTED"
	GetFreeipaStatusResponseStatusREQUESTED string = "REQUESTED"

	// GetFreeipaStatusResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	GetFreeipaStatusResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// GetFreeipaStatusResponseStatusAVAILABLE captures enum value "AVAILABLE"
	GetFreeipaStatusResponseStatusAVAILABLE string = "AVAILABLE"

	// GetFreeipaStatusResponseStatusSTACKAVAILABLE captures enum value "STACK_AVAILABLE"
	GetFreeipaStatusResponseStatusSTACKAVAILABLE string = "STACK_AVAILABLE"

	// GetFreeipaStatusResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	GetFreeipaStatusResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// GetFreeipaStatusResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	GetFreeipaStatusResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// GetFreeipaStatusResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	GetFreeipaStatusResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// GetFreeipaStatusResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	GetFreeipaStatusResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// GetFreeipaStatusResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	GetFreeipaStatusResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// GetFreeipaStatusResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	GetFreeipaStatusResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// GetFreeipaStatusResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	GetFreeipaStatusResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// GetFreeipaStatusResponseStatusSTOPPED captures enum value "STOPPED"
	GetFreeipaStatusResponseStatusSTOPPED string = "STOPPED"

	// GetFreeipaStatusResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	GetFreeipaStatusResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// GetFreeipaStatusResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	GetFreeipaStatusResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// GetFreeipaStatusResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	GetFreeipaStatusResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// GetFreeipaStatusResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	GetFreeipaStatusResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// GetFreeipaStatusResponseStatusSTARTFAILED captures enum value "START_FAILED"
	GetFreeipaStatusResponseStatusSTARTFAILED string = "START_FAILED"

	// GetFreeipaStatusResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	GetFreeipaStatusResponseStatusSTOPFAILED string = "STOP_FAILED"

	// GetFreeipaStatusResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	GetFreeipaStatusResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// GetFreeipaStatusResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	GetFreeipaStatusResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// GetFreeipaStatusResponseStatusUNREACHABLE captures enum value "UNREACHABLE"
	GetFreeipaStatusResponseStatusUNREACHABLE string = "UNREACHABLE"

	// GetFreeipaStatusResponseStatusUNHEALTHY captures enum value "UNHEALTHY"
	GetFreeipaStatusResponseStatusUNHEALTHY string = "UNHEALTHY"

	// GetFreeipaStatusResponseStatusDELETEDONPROVIDERSIDE captures enum value "DELETED_ON_PROVIDER_SIDE"
	GetFreeipaStatusResponseStatusDELETEDONPROVIDERSIDE string = "DELETED_ON_PROVIDER_SIDE"

	// GetFreeipaStatusResponseStatusUNKNOWN captures enum value "UNKNOWN"
	GetFreeipaStatusResponseStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *GetFreeipaStatusResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getFreeipaStatusResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetFreeipaStatusResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetFreeipaStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFreeipaStatusResponse) UnmarshalBinary(b []byte) error {
	var res GetFreeipaStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

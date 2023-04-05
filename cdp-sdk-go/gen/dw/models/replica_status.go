// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicaStatus Status information on the current state of replicas in the virtual warehouse.
//
// swagger:model ReplicaStatus
type ReplicaStatus struct {

	// Total number of ready coordinator replicas in the virtual warehouse. This number only contains the healthy executor replicas that have already started up and are ready to accept requests. If this number is less than the totalCoordinatorReplicas, then the virtual warehouse might still be starting up or there might be a problem scheduling these replicas.
	ReadyCoordinatorReplicas int32 `json:"readyCoordinatorReplicas"`

	// Total number of ready executor replicas in the virtual warehouse. This number only contains the healthy executor replicas that have already started up and are ready to accept requests. If this number is less than the totalExecutorReplicas, then the virtual warehouse might still be starting up or there might be a problem scheduling these replicas.
	ReadyExecutorReplicas int32 `json:"readyExecutorReplicas"`

	// Total number of coordinator replicas scheduled for the virtual warehouse. This number contains the number of coordinator replicas in pending state, as well as the replicas that are already running. If this number is zero, then the coordinator functionality is stopped.
	TotalCoordinatorReplicas int32 `json:"totalCoordinatorReplicas"`

	// Total number of executor replicas scheduled for the virtual warehouse. This number contains the number of executor replicas in pending state, as well as the replicas that are already running. If this number is zero, then the executor functionality is stopped.
	TotalExecutorReplicas int32 `json:"totalExecutorReplicas"`
}

// Validate validates this replica status
func (m *ReplicaStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this replica status based on context it is used
func (m *ReplicaStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicaStatus) UnmarshalBinary(b []byte) error {
	var res ReplicaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
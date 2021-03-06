// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyExecutionPoint policy execution point
//
// swagger:model PolicyExecutionPoint
type PolicyExecutionPoint struct {

	// optional policy id
	PolicyID string `json:"policy_id,omitempty"`

	// server id
	ServerID string `json:"server_id,omitempty"`

	// target id
	TargetID string `json:"target_fk,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// policy execution point type
	Type string `json:"type,omitempty"`
}

// Validate validates this policy execution point
func (m *PolicyExecutionPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyExecutionPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyExecutionPoint) UnmarshalBinary(b []byte) error {
	var res PolicyExecutionPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateGatewayRequest create gateway request
//
// swagger:model CreateGatewayRequest
type CreateGatewayRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// gateway name
	Name string `json:"name,omitempty"`

	// ServerID that this gateway should protect
	ServerID string `json:"server_id,omitempty"`

	// gateway type, one of: pyron, aws, azure
	Type string `json:"type,omitempty"`
}

// Validate validates this create gateway request
func (m *CreateGatewayRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGatewayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGatewayRequest) UnmarshalBinary(b []byte) error {
	var res CreateGatewayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

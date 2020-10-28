// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateGatewayRequest update gateway request
//
// swagger:model UpdateGatewayRequest
type UpdateGatewayRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// gateway name
	Name string `json:"name,omitempty"`
}

// Validate validates this update gateway request
func (m *UpdateGatewayRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateGatewayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateGatewayRequest) UnmarshalBinary(b []byte) error {
	var res UpdateGatewayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
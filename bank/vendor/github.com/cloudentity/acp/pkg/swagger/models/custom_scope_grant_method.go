// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomScopeGrantMethod custom scope grant method
//
// swagger:model CustomScopeGrantMethod
type CustomScopeGrantMethod struct {

	// scope grant URL
	ScopeGrantURL string `json:"scope_grant_url,omitempty"`
}

// Validate validates this custom scope grant method
func (m *CustomScopeGrantMethod) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomScopeGrantMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomScopeGrantMethod) UnmarshalBinary(b []byte) error {
	var res CustomScopeGrantMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
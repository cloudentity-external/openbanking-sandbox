// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsentGrantContext consent grant context
//
// swagger:model ConsentGrantContext
type ConsentGrantContext struct {

	// json object - device print of the End User's device
	Device map[string]interface{} `json:"device,omitempty"`

	// string in the form of a valid IP v 4 address, represents the current IP of the End User
	IP string `json:"ip,omitempty"`

	// 2 element array of floats - current geolocation of the end-user, the format is [lat, long]
	Location []float64 `json:"location"`
}

// Validate validates this consent grant context
func (m *ConsentGrantContext) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsentGrantContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentGrantContext) UnmarshalBinary(b []byte) error {
	var res ConsentGrantContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

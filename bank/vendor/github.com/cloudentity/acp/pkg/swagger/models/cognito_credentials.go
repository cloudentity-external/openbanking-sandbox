// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CognitoCredentials cognito credentials
//
// swagger:model CognitoCredentials
type CognitoCredentials struct {

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this cognito credentials
func (m *CognitoCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CognitoCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CognitoCredentials) UnmarshalBinary(b []byte) error {
	var res CognitoCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

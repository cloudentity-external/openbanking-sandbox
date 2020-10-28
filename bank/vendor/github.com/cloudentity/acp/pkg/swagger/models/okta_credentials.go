// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OktaCredentials okta credentials
//
// swagger:model OktaCredentials
type OktaCredentials struct {

	// supervisor client
	SupervisorClient *OktaSupervisorClient `json:"supervisor_client,omitempty"`
}

// Validate validates this okta credentials
func (m *OktaCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupervisorClient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OktaCredentials) validateSupervisorClient(formats strfmt.Registry) error {

	if swag.IsZero(m.SupervisorClient) { // not required
		return nil
	}

	if m.SupervisorClient != nil {
		if err := m.SupervisorClient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supervisor_client")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OktaCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OktaCredentials) UnmarshalBinary(b []byte) error {
	var res OktaCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

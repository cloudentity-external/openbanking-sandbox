// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Claim claim
//
// swagger:model Claim
type Claim struct {

	// unique claim id
	ID string `json:"id,omitempty"`

	// claim mapping - path to attribute in identity context from where claim value should be picked
	Mapping string `json:"mapping,omitempty"`

	// claim name in outgoing id / access token
	Name string `json:"name,omitempty"`

	// list of scopes - when at least one of listed scopes has been granted to a client, then claim will be added to id / access token. In case of empty array claim is always added.
	Scopes []string `json:"scopes"`

	// authorization server id
	ServerID string `json:"authorization_server_id,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// type
	Type ClaimType `json:"type,omitempty"`
}

// Validate validates this claim
func (m *Claim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Claim) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Claim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Claim) UnmarshalBinary(b []byte) error {
	var res Claim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

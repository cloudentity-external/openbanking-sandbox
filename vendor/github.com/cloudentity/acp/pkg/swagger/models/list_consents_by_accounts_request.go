// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListConsentsByAccountsRequest list consents by accounts request
//
// swagger:model ListConsentsByAccountsRequest
type ListConsentsByAccountsRequest struct {

	// list of account ids
	AccountIDs []string `json:"accounts"`
}

// Validate validates this list consents by accounts request
func (m *ListConsentsByAccountsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListConsentsByAccountsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListConsentsByAccountsRequest) UnmarshalBinary(b []byte) error {
	var res ListConsentsByAccountsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

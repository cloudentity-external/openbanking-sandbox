// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// ClaimType claim type, one of: id_token, access_token
// example: id_token
//
// swagger:model ClaimType
type ClaimType string

// Validate validates this claim type
func (m ClaimType) Validate(formats strfmt.Registry) error {
	return nil
}

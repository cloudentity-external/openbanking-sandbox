// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIs a p is
//
// swagger:model APIs
type APIs struct {

	// a p is
	APIs []*API `json:"apis"`
}

// Validate validates this a p is
func (m *APIs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIs) validateAPIs(formats strfmt.Registry) error {

	if swag.IsZero(m.APIs) { // not required
		return nil
	}

	for i := 0; i < len(m.APIs); i++ {
		if swag.IsZero(m.APIs[i]) { // not required
			continue
		}

		if m.APIs[i] != nil {
			if err := m.APIs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIs) UnmarshalBinary(b []byte) error {
	var res APIs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
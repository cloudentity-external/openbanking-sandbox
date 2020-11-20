// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBOtherCodeType16 Other application frequencies not covered in the standard code list
//
// swagger:model OB_OtherCodeType1_6
type OBOtherCodeType16 struct {

	// code
	Code OBCodeMnemonic `json:"Code,omitempty"`

	// description
	// Required: true
	Description Description3 `json:"Description"`

	// name
	// Required: true
	Name Name4 `json:"Name"`
}

// Validate validates this o b other code type1 6
func (m *OBOtherCodeType16) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBOtherCodeType16) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := m.Code.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Code")
		}
		return err
	}

	return nil
}

func (m *OBOtherCodeType16) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Description")
		}
		return err
	}

	return nil
}

func (m *OBOtherCodeType16) validateName(formats strfmt.Registry) error {

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBOtherCodeType16) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBOtherCodeType16) UnmarshalBinary(b []byte) error {
	var res OBOtherCodeType16
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
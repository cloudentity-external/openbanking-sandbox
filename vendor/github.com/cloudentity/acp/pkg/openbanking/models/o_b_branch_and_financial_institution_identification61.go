// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBBranchAndFinancialInstitutionIdentification61 Financial institution servicing an account for the creditor.
//
// swagger:model OBBranchAndFinancialInstitutionIdentification6_1
type OBBranchAndFinancialInstitutionIdentification61 struct {

	// identification
	Identification Identification2 `json:"Identification,omitempty"`

	// name
	Name Name1 `json:"Name,omitempty"`

	// postal address
	PostalAddress *OBPostalAddress6 `json:"PostalAddress,omitempty"`

	// scheme name
	SchemeName OBExternalFinancialInstitutionIdentification4Code `json:"SchemeName,omitempty"`
}

// Validate validates this o b branch and financial institution identification6 1
func (m *OBBranchAndFinancialInstitutionIdentification61) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBBranchAndFinancialInstitutionIdentification61) validateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if err := m.Identification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Identification")
		}
		return err
	}

	return nil
}

func (m *OBBranchAndFinancialInstitutionIdentification61) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

func (m *OBBranchAndFinancialInstitutionIdentification61) validatePostalAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PostalAddress) { // not required
		return nil
	}

	if m.PostalAddress != nil {
		if err := m.PostalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PostalAddress")
			}
			return err
		}
	}

	return nil
}

func (m *OBBranchAndFinancialInstitutionIdentification61) validateSchemeName(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeName) { // not required
		return nil
	}

	if err := m.SchemeName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SchemeName")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBBranchAndFinancialInstitutionIdentification61) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBBranchAndFinancialInstitutionIdentification61) UnmarshalBinary(b []byte) error {
	var res OBBranchAndFinancialInstitutionIdentification61
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

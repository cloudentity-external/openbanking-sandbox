// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBReadScheduledPayment3 o b read scheduled payment3
//
// swagger:model OBReadScheduledPayment3
type OBReadScheduledPayment3 struct {

	// data
	// Required: true
	Data *OBReadScheduledPayment3Data `json:"Data"`

	// links
	Links *Links `json:"Links,omitempty"`

	// meta
	Meta *Meta `json:"Meta,omitempty"`
}

// Validate validates this o b read scheduled payment3
func (m *OBReadScheduledPayment3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBReadScheduledPayment3) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

func (m *OBReadScheduledPayment3) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Links")
			}
			return err
		}
	}

	return nil
}

func (m *OBReadScheduledPayment3) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBReadScheduledPayment3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBReadScheduledPayment3) UnmarshalBinary(b []byte) error {
	var res OBReadScheduledPayment3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBReadScheduledPayment3Data o b read scheduled payment3 data
//
// swagger:model OBReadScheduledPayment3Data
type OBReadScheduledPayment3Data struct {

	// scheduled payment
	ScheduledPayment []*OBScheduledPayment3 `json:"ScheduledPayment"`
}

// Validate validates this o b read scheduled payment3 data
func (m *OBReadScheduledPayment3Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScheduledPayment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBReadScheduledPayment3Data) validateScheduledPayment(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledPayment) { // not required
		return nil
	}

	for i := 0; i < len(m.ScheduledPayment); i++ {
		if swag.IsZero(m.ScheduledPayment[i]) { // not required
			continue
		}

		if m.ScheduledPayment[i] != nil {
			if err := m.ScheduledPayment[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Data" + "." + "ScheduledPayment" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBReadScheduledPayment3Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBReadScheduledPayment3Data) UnmarshalBinary(b []byte) error {
	var res OBReadScheduledPayment3Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

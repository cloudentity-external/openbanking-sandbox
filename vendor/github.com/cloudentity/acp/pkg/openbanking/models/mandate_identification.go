// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MandateIdentification Direct Debit reference. For AUDDIS service users provide Core Reference. For non AUDDIS service users provide Core reference if possible or last used reference.
//
// swagger:model MandateIdentification
type MandateIdentification string

// Validate validates this mandate identification
func (m MandateIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 35); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

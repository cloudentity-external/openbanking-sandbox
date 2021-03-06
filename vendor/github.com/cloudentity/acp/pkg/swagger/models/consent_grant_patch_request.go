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

// ConsentGrantPatchRequest consent grant patch request
//
// swagger:model ConsentGrantPatchRequest
type ConsentGrantPatchRequest struct {

	// time when the grant occurred
	CollectionTimestamp int64 `json:"collection_timestamp,omitempty"`

	// an array of consent objects, consisting of consentId and granted - boolean flag marking if the user granted or revoked the consent
	Consents []*ConsentGrantPatch `json:"consents"`

	// language in which the consent was obtained [ISO 639]
	Language string `json:"language,omitempty"`

	// optional string with action_id - can be set if the consent grant/withdraw request was caused when an app asked the user for consent required for a specific action
	TriggeredByAction string `json:"triggered_by_action,omitempty"`

	// context
	Context *ConsentGrantContext `json:"context,omitempty"`
}

// Validate validates this consent grant patch request
func (m *ConsentGrantPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentGrantPatchRequest) validateConsents(formats strfmt.Registry) error {

	if swag.IsZero(m.Consents) { // not required
		return nil
	}

	for i := 0; i < len(m.Consents); i++ {
		if swag.IsZero(m.Consents[i]) { // not required
			continue
		}

		if m.Consents[i] != nil {
			if err := m.Consents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConsentGrantPatchRequest) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentGrantPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentGrantPatchRequest) UnmarshalBinary(b []byte) error {
	var res ConsentGrantPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

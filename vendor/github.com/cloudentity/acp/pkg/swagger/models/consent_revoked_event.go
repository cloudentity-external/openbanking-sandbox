// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsentRevokedEvent consent revoked event
//
// swagger:model ConsentRevokedEvent
type ConsentRevokedEvent struct {

	// time when the grant occurred
	CollectionTimestamp int64 `json:"collection_timestamp,omitempty"`

	// consent grant id
	ConsentGrantActID string `json:"consent_grant_act_id,omitempty"`

	// consent id
	ConsentID string `json:"consent_id,omitempty"`

	// given at timestamp
	// Format: date-time
	GivenAt strfmt.DateTime `json:"given_at,omitempty"`

	// grant type, one of: implicit, explicit
	GrantType string `json:"grant_type,omitempty"`

	// language in which the consent was obtained [ISO 639]
	Language string `json:"language,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// optional string with action_id - can be set if the consent grant/withdraw request was caused when an app asked the user for consent required for a specific action
	TriggeredByAction string `json:"triggered_by_action,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// consent
	Consent *Consent `json:"consent,omitempty"`

	// context
	Context *ConsentGrantContext `json:"context,omitempty"`
}

// Validate validates this consent revoked event
func (m *ConsentRevokedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGivenAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsent(formats); err != nil {
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

func (m *ConsentRevokedEvent) validateGivenAt(formats strfmt.Registry) error {

	if swag.IsZero(m.GivenAt) { // not required
		return nil
	}

	if err := validate.FormatOf("given_at", "body", "date-time", m.GivenAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsentRevokedEvent) validateConsent(formats strfmt.Registry) error {

	if swag.IsZero(m.Consent) { // not required
		return nil
	}

	if m.Consent != nil {
		if err := m.Consent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consent")
			}
			return err
		}
	}

	return nil
}

func (m *ConsentRevokedEvent) validateContext(formats strfmt.Registry) error {

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
func (m *ConsentRevokedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentRevokedEvent) UnmarshalBinary(b []byte) error {
	var res ConsentRevokedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

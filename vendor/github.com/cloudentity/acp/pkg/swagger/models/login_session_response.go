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

// LoginSessionResponse login session response
//
// swagger:model LoginSessionResponse
type LoginSessionResponse struct {

	// authentication context class reference
	Acr string `json:"acr,omitempty"`

	// scopes that passed policy validation
	AllowedScopes map[string]bool `json:"allowed_scopes,omitempty"`

	// authentication methods references
	Amr []string `json:"amr"`

	// time when user authenticated
	// Format: date-time
	AuthTime strfmt.DateTime `json:"auth_time,omitempty"`

	// OAuth client identifier
	ClientID string `json:"client_id,omitempty"`

	// list of granted audience
	GrantedAudience []string `json:"granted_audience"`

	// list of granted scopes
	GrantedScopes []string `json:"granted_scopes"`

	// unique id of login session
	ID string `json:"id,omitempty"`

	// idp identifier
	IDPID string `json:"idp_id,omitempty"`

	// is login approved
	LoginApproved bool `json:"login_approved,omitempty"`

	// is login rejected
	LoginRejected bool `json:"login_rejected,omitempty"`

	// max age for a session to live
	// Format: duration
	MaxAge strfmt.Duration `json:"max_age,omitempty"`

	// original url requested by oauth client
	RequestURL string `json:"request_url,omitempty"`

	// time when oauth client made a request
	// Format: date-time
	RequestedAt strfmt.DateTime `json:"requested_at,omitempty"`

	// list of requested audiences
	RequestedAudience []string `json:"requested_audience"`

	// requested grant type
	RequestedGrantType string `json:"requested_grant_type,omitempty"`

	// list of requested scopes
	RequestedScopes []*RequestedScope `json:"requested_scopes"`

	// is scope grant approved
	ScopeGrantApproved bool `json:"scope_grant_approved,omitempty"`

	// is scope grant rejected
	ScopeGrantRejected bool `json:"scope_grant_rejected,omitempty"`

	// authorization server identifier
	ServerID string `json:"server_id,omitempty"`

	// user identifier
	Subject string `json:"subject,omitempty"`

	// tenant identifier
	TenantID string `json:"tenant_id,omitempty"`

	// authentication context
	AuthenticationContext AuthenticationContext `json:"authentication_context,omitempty"`

	// error
	Error *RFC6749Error `json:"error,omitempty"`

	// request query params
	RequestQueryParams Values `json:"request_query_params,omitempty"`

	// requested claims
	RequestedClaims *ClaimsRequests `json:"requested_claims,omitempty"`
}

// Validate validates this login session response
func (m *LoginSessionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestQueryParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedClaims(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginSessionResponse) validateAuthTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthTime) { // not required
		return nil
	}

	if err := validate.FormatOf("auth_time", "body", "date-time", m.AuthTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoginSessionResponse) validateMaxAge(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxAge) { // not required
		return nil
	}

	if err := validate.FormatOf("max_age", "body", "duration", m.MaxAge.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoginSessionResponse) validateRequestedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_at", "body", "date-time", m.RequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoginSessionResponse) validateRequestedScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedScopes) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestedScopes); i++ {
		if swag.IsZero(m.RequestedScopes[i]) { // not required
			continue
		}

		if m.RequestedScopes[i] != nil {
			if err := m.RequestedScopes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requested_scopes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoginSessionResponse) validateAuthenticationContext(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationContext) { // not required
		return nil
	}

	if err := m.AuthenticationContext.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authentication_context")
		}
		return err
	}

	return nil
}

func (m *LoginSessionResponse) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *LoginSessionResponse) validateRequestQueryParams(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestQueryParams) { // not required
		return nil
	}

	if err := m.RequestQueryParams.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("request_query_params")
		}
		return err
	}

	return nil
}

func (m *LoginSessionResponse) validateRequestedClaims(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedClaims) { // not required
		return nil
	}

	if m.RequestedClaims != nil {
		if err := m.RequestedClaims.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requested_claims")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginSessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginSessionResponse) UnmarshalBinary(b []byte) error {
	var res LoginSessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
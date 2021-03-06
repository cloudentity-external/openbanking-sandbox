// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureB2CSettings Azure AD B2C authentication settings.
//
// Provide OAuth client details here.
//
// swagger:model AzureB2CSettings
type AzureB2CSettings struct {

	// OAuth client identifier
	ClientID string `json:"client_id,omitempty"`

	// a flag to fetch user groups
	FetchGroups bool `json:"fetch_groups,omitempty"`

	// flag to fetch additional user data from graph endpoint
	GetUser bool `json:"get_user,omitempty"`

	// list of user attributes to be fetched from graph
	GraphUserAttributes []string `json:"graph_user_attributes"`

	// user groups format: id or name
	GroupNameFormat string `json:"group_name_format,omitempty"`

	// should only security groups be fetched
	OnlySecurityGroups bool `json:"only_security_groups,omitempty"`

	// The user flow to be run.
	// Specify the name of a user flow you've created in your Azure AD B2C tenant.
	Policy string `json:"policy,omitempty"`

	// OAuth redirect URL
	RedirectURL string `json:"redirect_url,omitempty"`

	// OAuth scopes which client will be requesting
	Scopes []string `json:"scopes"`

	// azure tenant id
	Tenant string `json:"tenant,omitempty"`
}

// Validate validates this azure b2 c settings
func (m *AzureB2CSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureB2CSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureB2CSettings) UnmarshalBinary(b []byte) error {
	var res AzureB2CSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

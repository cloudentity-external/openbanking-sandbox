// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Dashboards dashboards
//
// swagger:model Dashboards
type Dashboards struct {

	// are dashboards enabled
	Enabled bool `json:"enabled,omitempty"`

	// sample API dashboard
	SampleAPIURL string `json:"sample_api_url,omitempty"`

	// sample business dashboard url
	SampleBusinessURL string `json:"sample_business_url,omitempty"`

	// sample threat dashboard url
	SampleThreatURL string `json:"sample_threat_url,omitempty"`

	// current tenant API dashboard
	TenantAPIURL string `json:"tenant_api_url,omitempty"`

	// current tenant business dashboard url
	TenantBusinessURL string `json:"tenant_business_url,omitempty"`

	// current tenant threat dashboard url
	TenantThreatURL string `json:"tenant_threat_url,omitempty"`
}

// Validate validates this dashboards
func (m *Dashboards) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Dashboards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dashboards) UnmarshalBinary(b []byte) error {
	var res Dashboards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

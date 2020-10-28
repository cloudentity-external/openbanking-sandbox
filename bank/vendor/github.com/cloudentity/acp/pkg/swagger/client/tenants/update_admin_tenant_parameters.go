// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// NewUpdateAdminTenantParams creates a new UpdateAdminTenantParams object
// with the default values initialized.
func NewUpdateAdminTenantParams() *UpdateAdminTenantParams {
	var (
		tidDefault = string("default")
	)
	return &UpdateAdminTenantParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdminTenantParamsWithTimeout creates a new UpdateAdminTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAdminTenantParamsWithTimeout(timeout time.Duration) *UpdateAdminTenantParams {
	var (
		tidDefault = string("default")
	)
	return &UpdateAdminTenantParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUpdateAdminTenantParamsWithContext creates a new UpdateAdminTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAdminTenantParamsWithContext(ctx context.Context) *UpdateAdminTenantParams {
	var (
		tidDefault = string("default")
	)
	return &UpdateAdminTenantParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUpdateAdminTenantParamsWithHTTPClient creates a new UpdateAdminTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAdminTenantParamsWithHTTPClient(client *http.Client) *UpdateAdminTenantParams {
	var (
		tidDefault = string("default")
	)
	return &UpdateAdminTenantParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UpdateAdminTenantParams contains all the parameters to send to the API endpoint
for the update admin tenant operation typically these are written to a http.Request
*/
type UpdateAdminTenantParams struct {

	/*Tenant*/
	Tenant *models.Tenant
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update admin tenant params
func (o *UpdateAdminTenantParams) WithTimeout(timeout time.Duration) *UpdateAdminTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update admin tenant params
func (o *UpdateAdminTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update admin tenant params
func (o *UpdateAdminTenantParams) WithContext(ctx context.Context) *UpdateAdminTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update admin tenant params
func (o *UpdateAdminTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update admin tenant params
func (o *UpdateAdminTenantParams) WithHTTPClient(client *http.Client) *UpdateAdminTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update admin tenant params
func (o *UpdateAdminTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenant adds the tenant to the update admin tenant params
func (o *UpdateAdminTenantParams) WithTenant(tenant *models.Tenant) *UpdateAdminTenantParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the update admin tenant params
func (o *UpdateAdminTenantParams) SetTenant(tenant *models.Tenant) {
	o.Tenant = tenant
}

// WithTid adds the tid to the update admin tenant params
func (o *UpdateAdminTenantParams) WithTid(tid string) *UpdateAdminTenantParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the update admin tenant params
func (o *UpdateAdminTenantParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdminTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Tenant != nil {
		if err := r.SetBodyParam(o.Tenant); err != nil {
			return err
		}
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// NewCreateTenantParams creates a new CreateTenantParams object
// with the default values initialized.
func NewCreateTenantParams() *CreateTenantParams {
	var ()
	return &CreateTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTenantParamsWithTimeout creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTenantParamsWithTimeout(timeout time.Duration) *CreateTenantParams {
	var ()
	return &CreateTenantParams{

		timeout: timeout,
	}
}

// NewCreateTenantParamsWithContext creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTenantParamsWithContext(ctx context.Context) *CreateTenantParams {
	var ()
	return &CreateTenantParams{

		Context: ctx,
	}
}

// NewCreateTenantParamsWithHTTPClient creates a new CreateTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTenantParamsWithHTTPClient(client *http.Client) *CreateTenantParams {
	var ()
	return &CreateTenantParams{
		HTTPClient: client,
	}
}

/*CreateTenantParams contains all the parameters to send to the API endpoint
for the create tenant operation typically these are written to a http.Request
*/
type CreateTenantParams struct {

	/*Tenant*/
	Tenant *models.Tenant
	/*WithDefaultWorkspaces
	  Create tenant with default workspaces

	*/
	WithDefaultWorkspaces *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create tenant params
func (o *CreateTenantParams) WithTimeout(timeout time.Duration) *CreateTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tenant params
func (o *CreateTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tenant params
func (o *CreateTenantParams) WithContext(ctx context.Context) *CreateTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tenant params
func (o *CreateTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tenant params
func (o *CreateTenantParams) WithHTTPClient(client *http.Client) *CreateTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tenant params
func (o *CreateTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenant adds the tenant to the create tenant params
func (o *CreateTenantParams) WithTenant(tenant *models.Tenant) *CreateTenantParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the create tenant params
func (o *CreateTenantParams) SetTenant(tenant *models.Tenant) {
	o.Tenant = tenant
}

// WithWithDefaultWorkspaces adds the withDefaultWorkspaces to the create tenant params
func (o *CreateTenantParams) WithWithDefaultWorkspaces(withDefaultWorkspaces *bool) *CreateTenantParams {
	o.SetWithDefaultWorkspaces(withDefaultWorkspaces)
	return o
}

// SetWithDefaultWorkspaces adds the withDefaultWorkspaces to the create tenant params
func (o *CreateTenantParams) SetWithDefaultWorkspaces(withDefaultWorkspaces *bool) {
	o.WithDefaultWorkspaces = withDefaultWorkspaces
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Tenant != nil {
		if err := r.SetBodyParam(o.Tenant); err != nil {
			return err
		}
	}

	if o.WithDefaultWorkspaces != nil {

		// query param with_default_workspaces
		var qrWithDefaultWorkspaces bool
		if o.WithDefaultWorkspaces != nil {
			qrWithDefaultWorkspaces = *o.WithDefaultWorkspaces
		}
		qWithDefaultWorkspaces := swag.FormatBool(qrWithDefaultWorkspaces)
		if qWithDefaultWorkspaces != "" {
			if err := r.SetQueryParam("with_default_workspaces", qWithDefaultWorkspaces); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewBindGroupToServiceParams creates a new BindGroupToServiceParams object
// with the default values initialized.
func NewBindGroupToServiceParams() *BindGroupToServiceParams {
	var (
		tidDefault = string("default")
	)
	return &BindGroupToServiceParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewBindGroupToServiceParamsWithTimeout creates a new BindGroupToServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBindGroupToServiceParamsWithTimeout(timeout time.Duration) *BindGroupToServiceParams {
	var (
		tidDefault = string("default")
	)
	return &BindGroupToServiceParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewBindGroupToServiceParamsWithContext creates a new BindGroupToServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewBindGroupToServiceParamsWithContext(ctx context.Context) *BindGroupToServiceParams {
	var (
		tidDefault = string("default")
	)
	return &BindGroupToServiceParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewBindGroupToServiceParamsWithHTTPClient creates a new BindGroupToServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBindGroupToServiceParamsWithHTTPClient(client *http.Client) *BindGroupToServiceParams {
	var (
		tidDefault = string("default")
	)
	return &BindGroupToServiceParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*BindGroupToServiceParams contains all the parameters to send to the API endpoint
for the bind group to service operation typically these are written to a http.Request
*/
type BindGroupToServiceParams struct {

	/*BindGroupToServiceRequest*/
	BindGroupToServiceRequest *models.BindGroupToServiceRequest
	/*APIGroup*/
	APIGroupID string
	/*Gw*/
	Gw string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bind group to service params
func (o *BindGroupToServiceParams) WithTimeout(timeout time.Duration) *BindGroupToServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind group to service params
func (o *BindGroupToServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind group to service params
func (o *BindGroupToServiceParams) WithContext(ctx context.Context) *BindGroupToServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind group to service params
func (o *BindGroupToServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind group to service params
func (o *BindGroupToServiceParams) WithHTTPClient(client *http.Client) *BindGroupToServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind group to service params
func (o *BindGroupToServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBindGroupToServiceRequest adds the bindGroupToServiceRequest to the bind group to service params
func (o *BindGroupToServiceParams) WithBindGroupToServiceRequest(bindGroupToServiceRequest *models.BindGroupToServiceRequest) *BindGroupToServiceParams {
	o.SetBindGroupToServiceRequest(bindGroupToServiceRequest)
	return o
}

// SetBindGroupToServiceRequest adds the bindGroupToServiceRequest to the bind group to service params
func (o *BindGroupToServiceParams) SetBindGroupToServiceRequest(bindGroupToServiceRequest *models.BindGroupToServiceRequest) {
	o.BindGroupToServiceRequest = bindGroupToServiceRequest
}

// WithAPIGroupID adds the aPIGroup to the bind group to service params
func (o *BindGroupToServiceParams) WithAPIGroupID(aPIGroup string) *BindGroupToServiceParams {
	o.SetAPIGroupID(aPIGroup)
	return o
}

// SetAPIGroupID adds the apiGroup to the bind group to service params
func (o *BindGroupToServiceParams) SetAPIGroupID(aPIGroup string) {
	o.APIGroupID = aPIGroup
}

// WithGw adds the gw to the bind group to service params
func (o *BindGroupToServiceParams) WithGw(gw string) *BindGroupToServiceParams {
	o.SetGw(gw)
	return o
}

// SetGw adds the gw to the bind group to service params
func (o *BindGroupToServiceParams) SetGw(gw string) {
	o.Gw = gw
}

// WithTid adds the tid to the bind group to service params
func (o *BindGroupToServiceParams) WithTid(tid string) *BindGroupToServiceParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the bind group to service params
func (o *BindGroupToServiceParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *BindGroupToServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BindGroupToServiceRequest != nil {
		if err := r.SetBodyParam(o.BindGroupToServiceRequest); err != nil {
			return err
		}
	}

	// path param apiGroup
	if err := r.SetPathParam("apiGroup", o.APIGroupID); err != nil {
		return err
	}

	// path param gw
	if err := r.SetPathParam("gw", o.Gw); err != nil {
		return err
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

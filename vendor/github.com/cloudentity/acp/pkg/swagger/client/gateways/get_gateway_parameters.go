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
)

// NewGetGatewayParams creates a new GetGatewayParams object
// with the default values initialized.
func NewGetGatewayParams() *GetGatewayParams {
	var (
		gwDefault  = string("default")
		tidDefault = string("default")
	)
	return &GetGatewayParams{
		Gw:  gwDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewayParamsWithTimeout creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewayParamsWithTimeout(timeout time.Duration) *GetGatewayParams {
	var (
		gwDefault  = string("default")
		tidDefault = string("default")
	)
	return &GetGatewayParams{
		Gw:  gwDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetGatewayParamsWithContext creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewayParamsWithContext(ctx context.Context) *GetGatewayParams {
	var (
		gwDefault  = string("default")
		tidDefault = string("default")
	)
	return &GetGatewayParams{
		Gw:  gwDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetGatewayParamsWithHTTPClient creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewayParamsWithHTTPClient(client *http.Client) *GetGatewayParams {
	var (
		gwDefault  = string("default")
		tidDefault = string("default")
	)
	return &GetGatewayParams{
		Gw:         gwDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetGatewayParams contains all the parameters to send to the API endpoint
for the get gateway operation typically these are written to a http.Request
*/
type GetGatewayParams struct {

	/*Gw
	  Gateway id

	*/
	Gw string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gateway params
func (o *GetGatewayParams) WithTimeout(timeout time.Duration) *GetGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateway params
func (o *GetGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateway params
func (o *GetGatewayParams) WithContext(ctx context.Context) *GetGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateway params
func (o *GetGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateway params
func (o *GetGatewayParams) WithHTTPClient(client *http.Client) *GetGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateway params
func (o *GetGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGw adds the gw to the get gateway params
func (o *GetGatewayParams) WithGw(gw string) *GetGatewayParams {
	o.SetGw(gw)
	return o
}

// SetGw adds the gw to the get gateway params
func (o *GetGatewayParams) SetGw(gw string) {
	o.Gw = gw
}

// WithTid adds the tid to the get gateway params
func (o *GetGatewayParams) WithTid(tid string) *GetGatewayParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get gateway params
func (o *GetGatewayParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

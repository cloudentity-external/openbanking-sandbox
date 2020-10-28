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

// NewDeleteGatewayParams creates a new DeleteGatewayParams object
// with the default values initialized.
func NewDeleteGatewayParams() *DeleteGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteGatewayParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGatewayParamsWithTimeout creates a new DeleteGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGatewayParamsWithTimeout(timeout time.Duration) *DeleteGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteGatewayParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDeleteGatewayParamsWithContext creates a new DeleteGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGatewayParamsWithContext(ctx context.Context) *DeleteGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteGatewayParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDeleteGatewayParamsWithHTTPClient creates a new DeleteGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGatewayParamsWithHTTPClient(client *http.Client) *DeleteGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteGatewayParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DeleteGatewayParams contains all the parameters to send to the API endpoint
for the delete gateway operation typically these are written to a http.Request
*/
type DeleteGatewayParams struct {

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

// WithTimeout adds the timeout to the delete gateway params
func (o *DeleteGatewayParams) WithTimeout(timeout time.Duration) *DeleteGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete gateway params
func (o *DeleteGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete gateway params
func (o *DeleteGatewayParams) WithContext(ctx context.Context) *DeleteGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete gateway params
func (o *DeleteGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete gateway params
func (o *DeleteGatewayParams) WithHTTPClient(client *http.Client) *DeleteGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete gateway params
func (o *DeleteGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGw adds the gw to the delete gateway params
func (o *DeleteGatewayParams) WithGw(gw string) *DeleteGatewayParams {
	o.SetGw(gw)
	return o
}

// SetGw adds the gw to the delete gateway params
func (o *DeleteGatewayParams) SetGw(gw string) {
	o.Gw = gw
}

// WithTid adds the tid to the delete gateway params
func (o *DeleteGatewayParams) WithTid(tid string) *DeleteGatewayParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the delete gateway params
func (o *DeleteGatewayParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

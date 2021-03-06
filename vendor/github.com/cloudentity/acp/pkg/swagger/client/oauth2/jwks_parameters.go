// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// NewJwksParams creates a new JwksParams object
// with the default values initialized.
func NewJwksParams() *JwksParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &JwksParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewJwksParamsWithTimeout creates a new JwksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJwksParamsWithTimeout(timeout time.Duration) *JwksParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &JwksParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewJwksParamsWithContext creates a new JwksParams object
// with the default values initialized, and the ability to set a context for a request
func NewJwksParamsWithContext(ctx context.Context) *JwksParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &JwksParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewJwksParamsWithHTTPClient creates a new JwksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJwksParamsWithHTTPClient(client *http.Client) *JwksParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &JwksParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*JwksParams contains all the parameters to send to the API endpoint
for the jwks operation typically these are written to a http.Request
*/
type JwksParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the jwks params
func (o *JwksParams) WithTimeout(timeout time.Duration) *JwksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jwks params
func (o *JwksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jwks params
func (o *JwksParams) WithContext(ctx context.Context) *JwksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jwks params
func (o *JwksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jwks params
func (o *JwksParams) WithHTTPClient(client *http.Client) *JwksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jwks params
func (o *JwksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the jwks params
func (o *JwksParams) WithAid(aid string) *JwksParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the jwks params
func (o *JwksParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the jwks params
func (o *JwksParams) WithTid(tid string) *JwksParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the jwks params
func (o *JwksParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *JwksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
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

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

// NewUserinfoParams creates a new UserinfoParams object
// with the default values initialized.
func NewUserinfoParams() *UserinfoParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &UserinfoParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUserinfoParamsWithTimeout creates a new UserinfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserinfoParamsWithTimeout(timeout time.Duration) *UserinfoParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &UserinfoParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUserinfoParamsWithContext creates a new UserinfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserinfoParamsWithContext(ctx context.Context) *UserinfoParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &UserinfoParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUserinfoParamsWithHTTPClient creates a new UserinfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserinfoParamsWithHTTPClient(client *http.Client) *UserinfoParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &UserinfoParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UserinfoParams contains all the parameters to send to the API endpoint
for the userinfo operation typically these are written to a http.Request
*/
type UserinfoParams struct {

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

// WithTimeout adds the timeout to the userinfo params
func (o *UserinfoParams) WithTimeout(timeout time.Duration) *UserinfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the userinfo params
func (o *UserinfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the userinfo params
func (o *UserinfoParams) WithContext(ctx context.Context) *UserinfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the userinfo params
func (o *UserinfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the userinfo params
func (o *UserinfoParams) WithHTTPClient(client *http.Client) *UserinfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the userinfo params
func (o *UserinfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the userinfo params
func (o *UserinfoParams) WithAid(aid string) *UserinfoParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the userinfo params
func (o *UserinfoParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the userinfo params
func (o *UserinfoParams) WithTid(tid string) *UserinfoParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the userinfo params
func (o *UserinfoParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UserinfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

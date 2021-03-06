// Code generated by go-swagger; DO NOT EDIT.

package servers

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

// NewGetAuthorizationServerParams creates a new GetAuthorizationServerParams object
// with the default values initialized.
func NewGetAuthorizationServerParams() *GetAuthorizationServerParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetAuthorizationServerParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthorizationServerParamsWithTimeout creates a new GetAuthorizationServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthorizationServerParamsWithTimeout(timeout time.Duration) *GetAuthorizationServerParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetAuthorizationServerParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetAuthorizationServerParamsWithContext creates a new GetAuthorizationServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthorizationServerParamsWithContext(ctx context.Context) *GetAuthorizationServerParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetAuthorizationServerParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetAuthorizationServerParamsWithHTTPClient creates a new GetAuthorizationServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthorizationServerParamsWithHTTPClient(client *http.Client) *GetAuthorizationServerParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetAuthorizationServerParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetAuthorizationServerParams contains all the parameters to send to the API endpoint
for the get authorization server operation typically these are written to a http.Request
*/
type GetAuthorizationServerParams struct {

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

// WithTimeout adds the timeout to the get authorization server params
func (o *GetAuthorizationServerParams) WithTimeout(timeout time.Duration) *GetAuthorizationServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get authorization server params
func (o *GetAuthorizationServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get authorization server params
func (o *GetAuthorizationServerParams) WithContext(ctx context.Context) *GetAuthorizationServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get authorization server params
func (o *GetAuthorizationServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get authorization server params
func (o *GetAuthorizationServerParams) WithHTTPClient(client *http.Client) *GetAuthorizationServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get authorization server params
func (o *GetAuthorizationServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the get authorization server params
func (o *GetAuthorizationServerParams) WithAid(aid string) *GetAuthorizationServerParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the get authorization server params
func (o *GetAuthorizationServerParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the get authorization server params
func (o *GetAuthorizationServerParams) WithTid(tid string) *GetAuthorizationServerParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get authorization server params
func (o *GetAuthorizationServerParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthorizationServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

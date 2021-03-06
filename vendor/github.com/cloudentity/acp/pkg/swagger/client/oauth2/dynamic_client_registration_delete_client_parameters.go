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

// NewDynamicClientRegistrationDeleteClientParams creates a new DynamicClientRegistrationDeleteClientParams object
// with the default values initialized.
func NewDynamicClientRegistrationDeleteClientParams() *DynamicClientRegistrationDeleteClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationDeleteClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDynamicClientRegistrationDeleteClientParamsWithTimeout creates a new DynamicClientRegistrationDeleteClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDynamicClientRegistrationDeleteClientParamsWithTimeout(timeout time.Duration) *DynamicClientRegistrationDeleteClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationDeleteClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDynamicClientRegistrationDeleteClientParamsWithContext creates a new DynamicClientRegistrationDeleteClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewDynamicClientRegistrationDeleteClientParamsWithContext(ctx context.Context) *DynamicClientRegistrationDeleteClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationDeleteClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDynamicClientRegistrationDeleteClientParamsWithHTTPClient creates a new DynamicClientRegistrationDeleteClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDynamicClientRegistrationDeleteClientParamsWithHTTPClient(client *http.Client) *DynamicClientRegistrationDeleteClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationDeleteClientParams{
		Aid:        aidDefault,
		Cid:        cidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DynamicClientRegistrationDeleteClientParams contains all the parameters to send to the API endpoint
for the dynamic client registration delete client operation typically these are written to a http.Request
*/
type DynamicClientRegistrationDeleteClientParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*Cid
	  Client id

	*/
	Cid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithTimeout(timeout time.Duration) *DynamicClientRegistrationDeleteClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithContext(ctx context.Context) *DynamicClientRegistrationDeleteClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithHTTPClient(client *http.Client) *DynamicClientRegistrationDeleteClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithAid(aid string) *DynamicClientRegistrationDeleteClientParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetAid(aid string) {
	o.Aid = aid
}

// WithCid adds the cid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithCid(cid string) *DynamicClientRegistrationDeleteClientParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetCid(cid string) {
	o.Cid = cid
}

// WithTid adds the tid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) WithTid(tid string) *DynamicClientRegistrationDeleteClientParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the dynamic client registration delete client params
func (o *DynamicClientRegistrationDeleteClientParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DynamicClientRegistrationDeleteClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
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

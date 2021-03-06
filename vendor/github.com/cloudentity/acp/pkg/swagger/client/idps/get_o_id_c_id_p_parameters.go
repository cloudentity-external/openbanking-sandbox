// Code generated by go-swagger; DO NOT EDIT.

package idps

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

// NewGetOIDCIDPParams creates a new GetOIDCIDPParams object
// with the default values initialized.
func NewGetOIDCIDPParams() *GetOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOIDCIDPParamsWithTimeout creates a new GetOIDCIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOIDCIDPParamsWithTimeout(timeout time.Duration) *GetOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetOIDCIDPParamsWithContext creates a new GetOIDCIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOIDCIDPParamsWithContext(ctx context.Context) *GetOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetOIDCIDPParamsWithHTTPClient creates a new GetOIDCIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOIDCIDPParamsWithHTTPClient(client *http.Client) *GetOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOIDCIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetOIDCIDPParams contains all the parameters to send to the API endpoint
for the get o ID c ID p operation typically these are written to a http.Request
*/
type GetOIDCIDPParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*Iid
	  IDP id

	*/
	Iid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithTimeout(timeout time.Duration) *GetOIDCIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithContext(ctx context.Context) *GetOIDCIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithHTTPClient(client *http.Client) *GetOIDCIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithAid(aid string) *GetOIDCIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithIid(iid string) *GetOIDCIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the get o ID c ID p params
func (o *GetOIDCIDPParams) WithTid(tid string) *GetOIDCIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get o ID c ID p params
func (o *GetOIDCIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOIDCIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	// path param iid
	if err := r.SetPathParam("iid", o.Iid); err != nil {
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

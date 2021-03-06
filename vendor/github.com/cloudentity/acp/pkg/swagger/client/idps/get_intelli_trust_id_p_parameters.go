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

// NewGetIntelliTrustIDPParams creates a new GetIntelliTrustIDPParams object
// with the default values initialized.
func NewGetIntelliTrustIDPParams() *GetIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetIntelliTrustIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntelliTrustIDPParamsWithTimeout creates a new GetIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIntelliTrustIDPParamsWithTimeout(timeout time.Duration) *GetIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetIntelliTrustIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetIntelliTrustIDPParamsWithContext creates a new GetIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIntelliTrustIDPParamsWithContext(ctx context.Context) *GetIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetIntelliTrustIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetIntelliTrustIDPParamsWithHTTPClient creates a new GetIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIntelliTrustIDPParamsWithHTTPClient(client *http.Client) *GetIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetIntelliTrustIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetIntelliTrustIDPParams contains all the parameters to send to the API endpoint
for the get intelli trust ID p operation typically these are written to a http.Request
*/
type GetIntelliTrustIDPParams struct {

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

// WithTimeout adds the timeout to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithTimeout(timeout time.Duration) *GetIntelliTrustIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithContext(ctx context.Context) *GetIntelliTrustIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithHTTPClient(client *http.Client) *GetIntelliTrustIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithAid(aid string) *GetIntelliTrustIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithIid(iid string) *GetIntelliTrustIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) WithTid(tid string) *GetIntelliTrustIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get intelli trust ID p params
func (o *GetIntelliTrustIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntelliTrustIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

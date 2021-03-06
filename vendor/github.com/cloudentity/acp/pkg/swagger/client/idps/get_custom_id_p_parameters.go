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

// NewGetCustomIDPParams creates a new GetCustomIDPParams object
// with the default values initialized.
func NewGetCustomIDPParams() *GetCustomIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetCustomIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomIDPParamsWithTimeout creates a new GetCustomIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomIDPParamsWithTimeout(timeout time.Duration) *GetCustomIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetCustomIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetCustomIDPParamsWithContext creates a new GetCustomIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCustomIDPParamsWithContext(ctx context.Context) *GetCustomIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetCustomIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetCustomIDPParamsWithHTTPClient creates a new GetCustomIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCustomIDPParamsWithHTTPClient(client *http.Client) *GetCustomIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetCustomIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetCustomIDPParams contains all the parameters to send to the API endpoint
for the get custom ID p operation typically these are written to a http.Request
*/
type GetCustomIDPParams struct {

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

// WithTimeout adds the timeout to the get custom ID p params
func (o *GetCustomIDPParams) WithTimeout(timeout time.Duration) *GetCustomIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get custom ID p params
func (o *GetCustomIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get custom ID p params
func (o *GetCustomIDPParams) WithContext(ctx context.Context) *GetCustomIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get custom ID p params
func (o *GetCustomIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get custom ID p params
func (o *GetCustomIDPParams) WithHTTPClient(client *http.Client) *GetCustomIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get custom ID p params
func (o *GetCustomIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the get custom ID p params
func (o *GetCustomIDPParams) WithAid(aid string) *GetCustomIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the get custom ID p params
func (o *GetCustomIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the get custom ID p params
func (o *GetCustomIDPParams) WithIid(iid string) *GetCustomIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get custom ID p params
func (o *GetCustomIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the get custom ID p params
func (o *GetCustomIDPParams) WithTid(tid string) *GetCustomIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get custom ID p params
func (o *GetCustomIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

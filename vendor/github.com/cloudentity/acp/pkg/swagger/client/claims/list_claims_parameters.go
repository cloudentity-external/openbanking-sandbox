// Code generated by go-swagger; DO NOT EDIT.

package claims

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

// NewListClaimsParams creates a new ListClaimsParams object
// with the default values initialized.
func NewListClaimsParams() *ListClaimsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListClaimsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListClaimsParamsWithTimeout creates a new ListClaimsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClaimsParamsWithTimeout(timeout time.Duration) *ListClaimsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListClaimsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewListClaimsParamsWithContext creates a new ListClaimsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClaimsParamsWithContext(ctx context.Context) *ListClaimsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListClaimsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewListClaimsParamsWithHTTPClient creates a new ListClaimsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClaimsParamsWithHTTPClient(client *http.Client) *ListClaimsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListClaimsParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListClaimsParams contains all the parameters to send to the API endpoint
for the list claims operation typically these are written to a http.Request
*/
type ListClaimsParams struct {

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

// WithTimeout adds the timeout to the list claims params
func (o *ListClaimsParams) WithTimeout(timeout time.Duration) *ListClaimsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list claims params
func (o *ListClaimsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list claims params
func (o *ListClaimsParams) WithContext(ctx context.Context) *ListClaimsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list claims params
func (o *ListClaimsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list claims params
func (o *ListClaimsParams) WithHTTPClient(client *http.Client) *ListClaimsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list claims params
func (o *ListClaimsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the list claims params
func (o *ListClaimsParams) WithAid(aid string) *ListClaimsParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the list claims params
func (o *ListClaimsParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the list claims params
func (o *ListClaimsParams) WithTid(tid string) *ListClaimsParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list claims params
func (o *ListClaimsParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ListClaimsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

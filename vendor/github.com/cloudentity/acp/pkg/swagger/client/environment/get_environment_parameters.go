// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// NewGetEnvironmentParams creates a new GetEnvironmentParams object
// with the default values initialized.
func NewGetEnvironmentParams() *GetEnvironmentParams {
	var (
		tidDefault = string("default")
	)
	return &GetEnvironmentParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentParamsWithTimeout creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentParamsWithTimeout(timeout time.Duration) *GetEnvironmentParams {
	var (
		tidDefault = string("default")
	)
	return &GetEnvironmentParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetEnvironmentParamsWithContext creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentParamsWithContext(ctx context.Context) *GetEnvironmentParams {
	var (
		tidDefault = string("default")
	)
	return &GetEnvironmentParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetEnvironmentParamsWithHTTPClient creates a new GetEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentParamsWithHTTPClient(client *http.Client) *GetEnvironmentParams {
	var (
		tidDefault = string("default")
	)
	return &GetEnvironmentParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetEnvironmentParams contains all the parameters to send to the API endpoint
for the get environment operation typically these are written to a http.Request
*/
type GetEnvironmentParams struct {

	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) WithTimeout(timeout time.Duration) *GetEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment params
func (o *GetEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment params
func (o *GetEnvironmentParams) WithContext(ctx context.Context) *GetEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment params
func (o *GetEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) WithHTTPClient(client *http.Client) *GetEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment params
func (o *GetEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTid adds the tid to the get environment params
func (o *GetEnvironmentParams) WithTid(tid string) *GetEnvironmentParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get environment params
func (o *GetEnvironmentParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

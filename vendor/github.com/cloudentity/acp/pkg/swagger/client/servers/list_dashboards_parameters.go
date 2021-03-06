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

// NewListDashboardsParams creates a new ListDashboardsParams object
// with the default values initialized.
func NewListDashboardsParams() *ListDashboardsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListDashboardsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListDashboardsParamsWithTimeout creates a new ListDashboardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDashboardsParamsWithTimeout(timeout time.Duration) *ListDashboardsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListDashboardsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewListDashboardsParamsWithContext creates a new ListDashboardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDashboardsParamsWithContext(ctx context.Context) *ListDashboardsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListDashboardsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewListDashboardsParamsWithHTTPClient creates a new ListDashboardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDashboardsParamsWithHTTPClient(client *http.Client) *ListDashboardsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListDashboardsParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListDashboardsParams contains all the parameters to send to the API endpoint
for the list dashboards operation typically these are written to a http.Request
*/
type ListDashboardsParams struct {

	/*Aid
	  Server id

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

// WithTimeout adds the timeout to the list dashboards params
func (o *ListDashboardsParams) WithTimeout(timeout time.Duration) *ListDashboardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dashboards params
func (o *ListDashboardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dashboards params
func (o *ListDashboardsParams) WithContext(ctx context.Context) *ListDashboardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dashboards params
func (o *ListDashboardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dashboards params
func (o *ListDashboardsParams) WithHTTPClient(client *http.Client) *ListDashboardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dashboards params
func (o *ListDashboardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the list dashboards params
func (o *ListDashboardsParams) WithAid(aid string) *ListDashboardsParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the list dashboards params
func (o *ListDashboardsParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the list dashboards params
func (o *ListDashboardsParams) WithTid(tid string) *ListDashboardsParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list dashboards params
func (o *ListDashboardsParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ListDashboardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

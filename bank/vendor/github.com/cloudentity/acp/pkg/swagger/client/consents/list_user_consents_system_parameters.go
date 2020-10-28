// Code generated by go-swagger; DO NOT EDIT.

package consents

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
	"github.com/go-openapi/swag"
)

// NewListUserConsentsSystemParams creates a new ListUserConsentsSystemParams object
// with the default values initialized.
func NewListUserConsentsSystemParams() *ListUserConsentsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &ListUserConsentsSystemParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListUserConsentsSystemParamsWithTimeout creates a new ListUserConsentsSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUserConsentsSystemParamsWithTimeout(timeout time.Duration) *ListUserConsentsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &ListUserConsentsSystemParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewListUserConsentsSystemParamsWithContext creates a new ListUserConsentsSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUserConsentsSystemParamsWithContext(ctx context.Context) *ListUserConsentsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &ListUserConsentsSystemParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewListUserConsentsSystemParamsWithHTTPClient creates a new ListUserConsentsSystemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUserConsentsSystemParamsWithHTTPClient(client *http.Client) *ListUserConsentsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &ListUserConsentsSystemParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListUserConsentsSystemParams contains all the parameters to send to the API endpoint
for the list user consents system operation typically these are written to a http.Request
*/
type ListUserConsentsSystemParams struct {

	/*ConsentID*/
	ConsentIDs []string
	/*Tid
	  Tenant id

	*/
	Tid string
	/*XSubject
	  user identifier

	*/
	Subject *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list user consents system params
func (o *ListUserConsentsSystemParams) WithTimeout(timeout time.Duration) *ListUserConsentsSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user consents system params
func (o *ListUserConsentsSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user consents system params
func (o *ListUserConsentsSystemParams) WithContext(ctx context.Context) *ListUserConsentsSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user consents system params
func (o *ListUserConsentsSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user consents system params
func (o *ListUserConsentsSystemParams) WithHTTPClient(client *http.Client) *ListUserConsentsSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user consents system params
func (o *ListUserConsentsSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsentIDs adds the consentID to the list user consents system params
func (o *ListUserConsentsSystemParams) WithConsentIDs(consentID []string) *ListUserConsentsSystemParams {
	o.SetConsentIDs(consentID)
	return o
}

// SetConsentIDs adds the consentId to the list user consents system params
func (o *ListUserConsentsSystemParams) SetConsentIDs(consentID []string) {
	o.ConsentIDs = consentID
}

// WithTid adds the tid to the list user consents system params
func (o *ListUserConsentsSystemParams) WithTid(tid string) *ListUserConsentsSystemParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list user consents system params
func (o *ListUserConsentsSystemParams) SetTid(tid string) {
	o.Tid = tid
}

// WithSubject adds the xSubject to the list user consents system params
func (o *ListUserConsentsSystemParams) WithSubject(xSubject *string) *ListUserConsentsSystemParams {
	o.SetSubject(xSubject)
	return o
}

// SetSubject adds the xSubject to the list user consents system params
func (o *ListUserConsentsSystemParams) SetSubject(xSubject *string) {
	o.Subject = xSubject
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserConsentsSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesConsentID := o.ConsentIDs

	joinedConsentID := swag.JoinByFormat(valuesConsentID, "")
	// query array param consent_id
	if err := r.SetQueryParam("consent_id", joinedConsentID...); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if o.Subject != nil {

		// header param x-subject
		if err := r.SetHeaderParam("x-subject", *o.Subject); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
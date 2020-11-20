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
)

// NewDeleteConsentActionParams creates a new DeleteConsentActionParams object
// with the default values initialized.
func NewDeleteConsentActionParams() *DeleteConsentActionParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteConsentActionParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConsentActionParamsWithTimeout creates a new DeleteConsentActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConsentActionParamsWithTimeout(timeout time.Duration) *DeleteConsentActionParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteConsentActionParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDeleteConsentActionParamsWithContext creates a new DeleteConsentActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConsentActionParamsWithContext(ctx context.Context) *DeleteConsentActionParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteConsentActionParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDeleteConsentActionParamsWithHTTPClient creates a new DeleteConsentActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConsentActionParamsWithHTTPClient(client *http.Client) *DeleteConsentActionParams {
	var (
		tidDefault = string("default")
	)
	return &DeleteConsentActionParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DeleteConsentActionParams contains all the parameters to send to the API endpoint
for the delete consent action operation typically these are written to a http.Request
*/
type DeleteConsentActionParams struct {

	/*Action*/
	Action string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete consent action params
func (o *DeleteConsentActionParams) WithTimeout(timeout time.Duration) *DeleteConsentActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete consent action params
func (o *DeleteConsentActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete consent action params
func (o *DeleteConsentActionParams) WithContext(ctx context.Context) *DeleteConsentActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete consent action params
func (o *DeleteConsentActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete consent action params
func (o *DeleteConsentActionParams) WithHTTPClient(client *http.Client) *DeleteConsentActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete consent action params
func (o *DeleteConsentActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the delete consent action params
func (o *DeleteConsentActionParams) WithAction(action string) *DeleteConsentActionParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the delete consent action params
func (o *DeleteConsentActionParams) SetAction(action string) {
	o.Action = action
}

// WithTid adds the tid to the delete consent action params
func (o *DeleteConsentActionParams) WithTid(tid string) *DeleteConsentActionParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the delete consent action params
func (o *DeleteConsentActionParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConsentActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param action
	if err := r.SetPathParam("action", o.Action); err != nil {
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
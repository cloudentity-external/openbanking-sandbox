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

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// NewPatchConsentGrantsSystemParams creates a new PatchConsentGrantsSystemParams object
// with the default values initialized.
func NewPatchConsentGrantsSystemParams() *PatchConsentGrantsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &PatchConsentGrantsSystemParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConsentGrantsSystemParamsWithTimeout creates a new PatchConsentGrantsSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConsentGrantsSystemParamsWithTimeout(timeout time.Duration) *PatchConsentGrantsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &PatchConsentGrantsSystemParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewPatchConsentGrantsSystemParamsWithContext creates a new PatchConsentGrantsSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchConsentGrantsSystemParamsWithContext(ctx context.Context) *PatchConsentGrantsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &PatchConsentGrantsSystemParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewPatchConsentGrantsSystemParamsWithHTTPClient creates a new PatchConsentGrantsSystemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchConsentGrantsSystemParamsWithHTTPClient(client *http.Client) *PatchConsentGrantsSystemParams {
	var (
		tidDefault = string("default")
	)
	return &PatchConsentGrantsSystemParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*PatchConsentGrantsSystemParams contains all the parameters to send to the API endpoint
for the patch consent grants system operation typically these are written to a http.Request
*/
type PatchConsentGrantsSystemParams struct {

	/*ConsentGrantPatchRequest*/
	ConsentGrantPatchRequest *models.ConsentGrantPatchRequest
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

// WithTimeout adds the timeout to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithTimeout(timeout time.Duration) *PatchConsentGrantsSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithContext(ctx context.Context) *PatchConsentGrantsSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithHTTPClient(client *http.Client) *PatchConsentGrantsSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsentGrantPatchRequest adds the consentGrantPatchRequest to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithConsentGrantPatchRequest(consentGrantPatchRequest *models.ConsentGrantPatchRequest) *PatchConsentGrantsSystemParams {
	o.SetConsentGrantPatchRequest(consentGrantPatchRequest)
	return o
}

// SetConsentGrantPatchRequest adds the consentGrantPatchRequest to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetConsentGrantPatchRequest(consentGrantPatchRequest *models.ConsentGrantPatchRequest) {
	o.ConsentGrantPatchRequest = consentGrantPatchRequest
}

// WithTid adds the tid to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithTid(tid string) *PatchConsentGrantsSystemParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetTid(tid string) {
	o.Tid = tid
}

// WithSubject adds the xSubject to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) WithSubject(xSubject *string) *PatchConsentGrantsSystemParams {
	o.SetSubject(xSubject)
	return o
}

// SetSubject adds the xSubject to the patch consent grants system params
func (o *PatchConsentGrantsSystemParams) SetSubject(xSubject *string) {
	o.Subject = xSubject
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConsentGrantsSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConsentGrantPatchRequest != nil {
		if err := r.SetBodyParam(o.ConsentGrantPatchRequest); err != nil {
			return err
		}
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
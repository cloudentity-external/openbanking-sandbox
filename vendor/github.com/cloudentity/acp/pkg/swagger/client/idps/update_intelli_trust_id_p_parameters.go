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

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// NewUpdateIntelliTrustIDPParams creates a new UpdateIntelliTrustIDPParams object
// with the default values initialized.
func NewUpdateIntelliTrustIDPParams() *UpdateIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateIntelliTrustIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntelliTrustIDPParamsWithTimeout creates a new UpdateIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateIntelliTrustIDPParamsWithTimeout(timeout time.Duration) *UpdateIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateIntelliTrustIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUpdateIntelliTrustIDPParamsWithContext creates a new UpdateIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateIntelliTrustIDPParamsWithContext(ctx context.Context) *UpdateIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateIntelliTrustIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUpdateIntelliTrustIDPParamsWithHTTPClient creates a new UpdateIntelliTrustIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateIntelliTrustIDPParamsWithHTTPClient(client *http.Client) *UpdateIntelliTrustIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateIntelliTrustIDPParams{
		Aid:        aidDefault,
		Iid:        iidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UpdateIntelliTrustIDPParams contains all the parameters to send to the API endpoint
for the update intelli trust ID p operation typically these are written to a http.Request
*/
type UpdateIntelliTrustIDPParams struct {

	/*IntelliTrustIDP
	  IntelliTrustIDP

	*/
	IntelliTrustIDP *models.IntelliTrustIDP
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

// WithTimeout adds the timeout to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithTimeout(timeout time.Duration) *UpdateIntelliTrustIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithContext(ctx context.Context) *UpdateIntelliTrustIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithHTTPClient(client *http.Client) *UpdateIntelliTrustIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntelliTrustIDP adds the intelliTrustIDP to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) *UpdateIntelliTrustIDPParams {
	o.SetIntelliTrustIDP(intelliTrustIDP)
	return o
}

// SetIntelliTrustIDP adds the intelliTrustIdP to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetIntelliTrustIDP(intelliTrustIDP *models.IntelliTrustIDP) {
	o.IntelliTrustIDP = intelliTrustIDP
}

// WithAid adds the aid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithAid(aid string) *UpdateIntelliTrustIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithIid(iid string) *UpdateIntelliTrustIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) WithTid(tid string) *UpdateIntelliTrustIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the update intelli trust ID p params
func (o *UpdateIntelliTrustIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntelliTrustIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IntelliTrustIDP != nil {
		if err := r.SetBodyParam(o.IntelliTrustIDP); err != nil {
			return err
		}
	}

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

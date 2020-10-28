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

// NewUpdateSAMLIDPParams creates a new UpdateSAMLIDPParams object
// with the default values initialized.
func NewUpdateSAMLIDPParams() *UpdateSAMLIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateSAMLIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSAMLIDPParamsWithTimeout creates a new UpdateSAMLIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSAMLIDPParamsWithTimeout(timeout time.Duration) *UpdateSAMLIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateSAMLIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUpdateSAMLIDPParamsWithContext creates a new UpdateSAMLIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSAMLIDPParamsWithContext(ctx context.Context) *UpdateSAMLIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateSAMLIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUpdateSAMLIDPParamsWithHTTPClient creates a new UpdateSAMLIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSAMLIDPParamsWithHTTPClient(client *http.Client) *UpdateSAMLIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateSAMLIDPParams{
		Aid:        aidDefault,
		Iid:        iidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UpdateSAMLIDPParams contains all the parameters to send to the API endpoint
for the update s a m l ID p operation typically these are written to a http.Request
*/
type UpdateSAMLIDPParams struct {

	/*SAMLIDP
	  SAMLIDP

	*/
	SAMLIDP *models.SAMLIDP
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

// WithTimeout adds the timeout to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithTimeout(timeout time.Duration) *UpdateSAMLIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithContext(ctx context.Context) *UpdateSAMLIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithHTTPClient(client *http.Client) *UpdateSAMLIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSAMLIDP adds the sAMLIDP to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithSAMLIDP(sAMLIDP *models.SAMLIDP) *UpdateSAMLIDPParams {
	o.SetSAMLIDP(sAMLIDP)
	return o
}

// SetSAMLIDP adds the sAMLIdP to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetSAMLIDP(sAMLIDP *models.SAMLIDP) {
	o.SAMLIDP = sAMLIDP
}

// WithAid adds the aid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithAid(aid string) *UpdateSAMLIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithIid(iid string) *UpdateSAMLIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) WithTid(tid string) *UpdateSAMLIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the update s a m l ID p params
func (o *UpdateSAMLIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSAMLIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SAMLIDP != nil {
		if err := r.SetBodyParam(o.SAMLIDP); err != nil {
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

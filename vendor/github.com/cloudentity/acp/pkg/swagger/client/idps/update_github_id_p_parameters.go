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

// NewUpdateGithubIDPParams creates a new UpdateGithubIDPParams object
// with the default values initialized.
func NewUpdateGithubIDPParams() *UpdateGithubIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateGithubIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGithubIDPParamsWithTimeout creates a new UpdateGithubIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGithubIDPParamsWithTimeout(timeout time.Duration) *UpdateGithubIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateGithubIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUpdateGithubIDPParamsWithContext creates a new UpdateGithubIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGithubIDPParamsWithContext(ctx context.Context) *UpdateGithubIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateGithubIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUpdateGithubIDPParamsWithHTTPClient creates a new UpdateGithubIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGithubIDPParamsWithHTTPClient(client *http.Client) *UpdateGithubIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateGithubIDPParams{
		Aid:        aidDefault,
		Iid:        iidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UpdateGithubIDPParams contains all the parameters to send to the API endpoint
for the update github ID p operation typically these are written to a http.Request
*/
type UpdateGithubIDPParams struct {

	/*GithubIDP
	  GithubIDP

	*/
	GithubIDP *models.GithubIDP
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

// WithTimeout adds the timeout to the update github ID p params
func (o *UpdateGithubIDPParams) WithTimeout(timeout time.Duration) *UpdateGithubIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update github ID p params
func (o *UpdateGithubIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update github ID p params
func (o *UpdateGithubIDPParams) WithContext(ctx context.Context) *UpdateGithubIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update github ID p params
func (o *UpdateGithubIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update github ID p params
func (o *UpdateGithubIDPParams) WithHTTPClient(client *http.Client) *UpdateGithubIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update github ID p params
func (o *UpdateGithubIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGithubIDP adds the githubIDP to the update github ID p params
func (o *UpdateGithubIDPParams) WithGithubIDP(githubIDP *models.GithubIDP) *UpdateGithubIDPParams {
	o.SetGithubIDP(githubIDP)
	return o
}

// SetGithubIDP adds the githubIdP to the update github ID p params
func (o *UpdateGithubIDPParams) SetGithubIDP(githubIDP *models.GithubIDP) {
	o.GithubIDP = githubIDP
}

// WithAid adds the aid to the update github ID p params
func (o *UpdateGithubIDPParams) WithAid(aid string) *UpdateGithubIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the update github ID p params
func (o *UpdateGithubIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the update github ID p params
func (o *UpdateGithubIDPParams) WithIid(iid string) *UpdateGithubIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update github ID p params
func (o *UpdateGithubIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the update github ID p params
func (o *UpdateGithubIDPParams) WithTid(tid string) *UpdateGithubIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the update github ID p params
func (o *UpdateGithubIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGithubIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GithubIDP != nil {
		if err := r.SetBodyParam(o.GithubIDP); err != nil {
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

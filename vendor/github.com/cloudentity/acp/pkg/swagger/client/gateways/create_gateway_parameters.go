// Code generated by go-swagger; DO NOT EDIT.

package gateways

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

// NewCreateGatewayParams creates a new CreateGatewayParams object
// with the default values initialized.
func NewCreateGatewayParams() *CreateGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &CreateGatewayParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGatewayParamsWithTimeout creates a new CreateGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGatewayParamsWithTimeout(timeout time.Duration) *CreateGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &CreateGatewayParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewCreateGatewayParamsWithContext creates a new CreateGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGatewayParamsWithContext(ctx context.Context) *CreateGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &CreateGatewayParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewCreateGatewayParamsWithHTTPClient creates a new CreateGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGatewayParamsWithHTTPClient(client *http.Client) *CreateGatewayParams {
	var (
		tidDefault = string("default")
	)
	return &CreateGatewayParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*CreateGatewayParams contains all the parameters to send to the API endpoint
for the create gateway operation typically these are written to a http.Request
*/
type CreateGatewayParams struct {

	/*Gateway*/
	Gateway *models.CreateGatewayRequest
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create gateway params
func (o *CreateGatewayParams) WithTimeout(timeout time.Duration) *CreateGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gateway params
func (o *CreateGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gateway params
func (o *CreateGatewayParams) WithContext(ctx context.Context) *CreateGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gateway params
func (o *CreateGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gateway params
func (o *CreateGatewayParams) WithHTTPClient(client *http.Client) *CreateGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gateway params
func (o *CreateGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGateway adds the gateway to the create gateway params
func (o *CreateGatewayParams) WithGateway(gateway *models.CreateGatewayRequest) *CreateGatewayParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the create gateway params
func (o *CreateGatewayParams) SetGateway(gateway *models.CreateGatewayRequest) {
	o.Gateway = gateway
}

// WithTid adds the tid to the create gateway params
func (o *CreateGatewayParams) WithTid(tid string) *CreateGatewayParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the create gateway params
func (o *CreateGatewayParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Gateway != nil {
		if err := r.SetBodyParam(o.Gateway); err != nil {
			return err
		}
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

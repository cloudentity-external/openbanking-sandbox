// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewListPoliciesParams creates a new ListPoliciesParams object
// with the default values initialized.
func NewListPoliciesParams() *ListPoliciesParams {
	var (
		aidDefault        = string("default")
		policyTypeDefault = string("api")
		tidDefault        = string("default")
	)
	return &ListPoliciesParams{
		Aid:        aidDefault,
		PolicyType: &policyTypeDefault,
		Tid:        tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListPoliciesParamsWithTimeout creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPoliciesParamsWithTimeout(timeout time.Duration) *ListPoliciesParams {
	var (
		aidDefault        = string("default")
		policyTypeDefault = string("api")
		tidDefault        = string("default")
	)
	return &ListPoliciesParams{
		Aid:        aidDefault,
		PolicyType: &policyTypeDefault,
		Tid:        tidDefault,

		timeout: timeout,
	}
}

// NewListPoliciesParamsWithContext creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPoliciesParamsWithContext(ctx context.Context) *ListPoliciesParams {
	var (
		aidDefault        = string("default")
		policyTypeDefault = string("api")
		tidDefault        = string("default")
	)
	return &ListPoliciesParams{
		Aid:        aidDefault,
		PolicyType: &policyTypeDefault,
		Tid:        tidDefault,

		Context: ctx,
	}
}

// NewListPoliciesParamsWithHTTPClient creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPoliciesParamsWithHTTPClient(client *http.Client) *ListPoliciesParams {
	var (
		aidDefault        = string("default")
		policyTypeDefault = string("api")
		tidDefault        = string("default")
	)
	return &ListPoliciesParams{
		Aid:        aidDefault,
		PolicyType: &policyTypeDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListPoliciesParams contains all the parameters to send to the API endpoint
for the list policies operation typically these are written to a http.Request
*/
type ListPoliciesParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*PolicyType
	  Policy type

	*/
	PolicyType *string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) WithTimeout(timeout time.Duration) *ListPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list policies params
func (o *ListPoliciesParams) WithContext(ctx context.Context) *ListPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list policies params
func (o *ListPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) WithHTTPClient(client *http.Client) *ListPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the list policies params
func (o *ListPoliciesParams) WithAid(aid string) *ListPoliciesParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the list policies params
func (o *ListPoliciesParams) SetAid(aid string) {
	o.Aid = aid
}

// WithPolicyType adds the policyType to the list policies params
func (o *ListPoliciesParams) WithPolicyType(policyType *string) *ListPoliciesParams {
	o.SetPolicyType(policyType)
	return o
}

// SetPolicyType adds the policyType to the list policies params
func (o *ListPoliciesParams) SetPolicyType(policyType *string) {
	o.PolicyType = policyType
}

// WithTid adds the tid to the list policies params
func (o *ListPoliciesParams) WithTid(tid string) *ListPoliciesParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list policies params
func (o *ListPoliciesParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ListPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	if o.PolicyType != nil {

		// query param policy_type
		var qrPolicyType string
		if o.PolicyType != nil {
			qrPolicyType = *o.PolicyType
		}
		qPolicyType := qrPolicyType
		if qPolicyType != "" {
			if err := r.SetQueryParam("policy_type", qPolicyType); err != nil {
				return err
			}
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

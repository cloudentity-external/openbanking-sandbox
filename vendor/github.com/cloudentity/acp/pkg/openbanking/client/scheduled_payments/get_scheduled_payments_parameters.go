// Code generated by go-swagger; DO NOT EDIT.

package scheduled_payments

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

// NewGetScheduledPaymentsParams creates a new GetScheduledPaymentsParams object
// with the default values initialized.
func NewGetScheduledPaymentsParams() *GetScheduledPaymentsParams {
	var ()
	return &GetScheduledPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduledPaymentsParamsWithTimeout creates a new GetScheduledPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScheduledPaymentsParamsWithTimeout(timeout time.Duration) *GetScheduledPaymentsParams {
	var ()
	return &GetScheduledPaymentsParams{

		timeout: timeout,
	}
}

// NewGetScheduledPaymentsParamsWithContext creates a new GetScheduledPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScheduledPaymentsParamsWithContext(ctx context.Context) *GetScheduledPaymentsParams {
	var ()
	return &GetScheduledPaymentsParams{

		Context: ctx,
	}
}

// NewGetScheduledPaymentsParamsWithHTTPClient creates a new GetScheduledPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScheduledPaymentsParamsWithHTTPClient(client *http.Client) *GetScheduledPaymentsParams {
	var ()
	return &GetScheduledPaymentsParams{
		HTTPClient: client,
	}
}

/*GetScheduledPaymentsParams contains all the parameters to send to the API endpoint
for the get scheduled payments operation typically these are written to a http.Request
*/
type GetScheduledPaymentsParams struct {

	/*Authorization
	  An Authorisation Token as per https://tools.ietf.org/html/rfc6750

	*/
	Authorization string
	/*XCustomerUserAgent
	  Indicates the user-agent that the PSU is using.

	*/
	XCustomerUserAgent *string
	/*XFapiAuthDate
	  The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC

	*/
	XFapiAuthDate *string
	/*XFapiCustomerIPAddress
	  The PSU's IP address if the PSU is currently logged in with the TPP.

	*/
	XFapiCustomerIPAddress *string
	/*XFapiInteractionID
	  An RFC4122 UID used as a correlation id.

	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithTimeout(timeout time.Duration) *GetScheduledPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithContext(ctx context.Context) *GetScheduledPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithHTTPClient(client *http.Client) *GetScheduledPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithAuthorization(authorization string) *GetScheduledPaymentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetScheduledPaymentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetScheduledPaymentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetScheduledPaymentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get scheduled payments params
func (o *GetScheduledPaymentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetScheduledPaymentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get scheduled payments params
func (o *GetScheduledPaymentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduledPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}

	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}

	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}

	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

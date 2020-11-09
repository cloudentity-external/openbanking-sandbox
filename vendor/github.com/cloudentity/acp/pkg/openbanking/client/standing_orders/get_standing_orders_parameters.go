// Code generated by go-swagger; DO NOT EDIT.

package standing_orders

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

// NewGetStandingOrdersParams creates a new GetStandingOrdersParams object
// with the default values initialized.
func NewGetStandingOrdersParams() *GetStandingOrdersParams {
	var ()
	return &GetStandingOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStandingOrdersParamsWithTimeout creates a new GetStandingOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStandingOrdersParamsWithTimeout(timeout time.Duration) *GetStandingOrdersParams {
	var ()
	return &GetStandingOrdersParams{

		timeout: timeout,
	}
}

// NewGetStandingOrdersParamsWithContext creates a new GetStandingOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStandingOrdersParamsWithContext(ctx context.Context) *GetStandingOrdersParams {
	var ()
	return &GetStandingOrdersParams{

		Context: ctx,
	}
}

// NewGetStandingOrdersParamsWithHTTPClient creates a new GetStandingOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStandingOrdersParamsWithHTTPClient(client *http.Client) *GetStandingOrdersParams {
	var ()
	return &GetStandingOrdersParams{
		HTTPClient: client,
	}
}

/*GetStandingOrdersParams contains all the parameters to send to the API endpoint
for the get standing orders operation typically these are written to a http.Request
*/
type GetStandingOrdersParams struct {

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

// WithTimeout adds the timeout to the get standing orders params
func (o *GetStandingOrdersParams) WithTimeout(timeout time.Duration) *GetStandingOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get standing orders params
func (o *GetStandingOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get standing orders params
func (o *GetStandingOrdersParams) WithContext(ctx context.Context) *GetStandingOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get standing orders params
func (o *GetStandingOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get standing orders params
func (o *GetStandingOrdersParams) WithHTTPClient(client *http.Client) *GetStandingOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get standing orders params
func (o *GetStandingOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get standing orders params
func (o *GetStandingOrdersParams) WithAuthorization(authorization string) *GetStandingOrdersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get standing orders params
func (o *GetStandingOrdersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get standing orders params
func (o *GetStandingOrdersParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetStandingOrdersParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get standing orders params
func (o *GetStandingOrdersParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get standing orders params
func (o *GetStandingOrdersParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetStandingOrdersParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get standing orders params
func (o *GetStandingOrdersParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get standing orders params
func (o *GetStandingOrdersParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetStandingOrdersParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get standing orders params
func (o *GetStandingOrdersParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get standing orders params
func (o *GetStandingOrdersParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetStandingOrdersParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get standing orders params
func (o *GetStandingOrdersParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStandingOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

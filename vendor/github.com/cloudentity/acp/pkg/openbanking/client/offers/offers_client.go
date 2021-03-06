// Code generated by go-swagger; DO NOT EDIT.

package offers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new offers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for offers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountsAccountIDOffers(params *GetAccountsAccountIDOffersParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDOffersOK, error)

	GetOffers(params *GetOffersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOffersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountsAccountIDOffers gets offers
*/
func (a *Client) GetAccountsAccountIDOffers(params *GetAccountsAccountIDOffersParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDOffersOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetAccountsAccountIDOffersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountsAccountIdOffers",
		Method:             "GET",
		PathPattern:        "/accounts/{AccountId}/offers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsAccountIDOffersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsAccountIDOffersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountsAccountIdOffers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOffers gets offers
*/
func (a *Client) GetOffers(params *GetOffersParams, authInfo runtime.ClientAuthInfoWriter) (*GetOffersOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetOffersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOffers",
		Method:             "GET",
		PathPattern:        "/offers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOffersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOffersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOffers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

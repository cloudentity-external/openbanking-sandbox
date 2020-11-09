// Code generated by go-swagger; DO NOT EDIT.

package account_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new account access API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account access API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAccountAccessConsents(params *CreateAccountAccessConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAccountAccessConsentsCreated, error)

	DeleteAccountAccessConsentsConsentID(params *DeleteAccountAccessConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAccountAccessConsentsConsentIDNoContent, error)

	GetAccountAccessConsentsConsentID(params *GetAccountAccessConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountAccessConsentsConsentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAccountAccessConsents creates account access consents
*/
func (a *Client) CreateAccountAccessConsents(params *CreateAccountAccessConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAccountAccessConsentsCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewCreateAccountAccessConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateAccountAccessConsents",
		Method:             "POST",
		PathPattern:        "/account-access-consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAccountAccessConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAccountAccessConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateAccountAccessConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAccountAccessConsentsConsentID deletes account access consents
*/
func (a *Client) DeleteAccountAccessConsentsConsentID(params *DeleteAccountAccessConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAccountAccessConsentsConsentIDNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDeleteAccountAccessConsentsConsentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAccountAccessConsentsConsentId",
		Method:             "DELETE",
		PathPattern:        "/account-access-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccountAccessConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAccountAccessConsentsConsentIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccountAccessConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccountAccessConsentsConsentID gets account access consents
*/
func (a *Client) GetAccountAccessConsentsConsentID(params *GetAccountAccessConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountAccessConsentsConsentIDOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetAccountAccessConsentsConsentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountAccessConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/account-access-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountAccessConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountAccessConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountAccessConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new oauth2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for oauth2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Authorize(params *AuthorizeParams) error

	DynamicClientRegistration(params *DynamicClientRegistrationParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationCreated, error)

	DynamicClientRegistrationDeleteClient(params *DynamicClientRegistrationDeleteClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationDeleteClientNoContent, error)

	DynamicClientRegistrationGetClient(params *DynamicClientRegistrationGetClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationGetClientOK, error)

	DynamicClientRegistrationUpdateClient(params *DynamicClientRegistrationUpdateClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationUpdateClientOK, error)

	GatewayIntrospect(params *GatewayIntrospectParams, authInfo runtime.ClientAuthInfoWriter) (*GatewayIntrospectOK, error)

	Introspect(params *IntrospectParams, authInfo runtime.ClientAuthInfoWriter) (*IntrospectOK, error)

	Jwks(params *JwksParams) (*JwksOK, error)

	Revoke(params *RevokeParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeOK, error)

	Token(params *TokenParams) (*TokenOK, error)

	Userinfo(params *UserinfoParams, authInfo runtime.ClientAuthInfoWriter) (*UserinfoOK, error)

	WellKnown(params *WellKnownParams) (*WellKnownOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Authorize thes o auth 2 0 authorize endpoint

  The authorization endpoint is used to interact with the resource
owner and obtain an authorization grant.
*/
func (a *Client) Authorize(params *AuthorizeParams) error {
	// : Validate the params before sending
	if params == nil {
		params = NewAuthorizeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "authorize",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/oauth2/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthorizeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  DynamicClientRegistration The OAuth 2.0 Dynamic Client Registration endpoint
*/
func (a *Client) DynamicClientRegistration(params *DynamicClientRegistrationParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dynamicClientRegistration",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/oauth2/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dynamicClientRegistration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DynamicClientRegistrationDeleteClient The OAuth 2.0 Dynamic Client Registration Delete Client endpoint
*/
func (a *Client) DynamicClientRegistrationDeleteClient(params *DynamicClientRegistrationDeleteClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationDeleteClientNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationDeleteClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationDeleteClient",
		Method:             "DELETE",
		PathPattern:        "/{tid}/{aid}/oauth2/register/{cid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationDeleteClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationDeleteClientNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dynamicClientRegistrationDeleteClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DynamicClientRegistrationGetClient The OAuth 2.0 Dynamic Client Registration Get Client endpoint
*/
func (a *Client) DynamicClientRegistrationGetClient(params *DynamicClientRegistrationGetClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationGetClientOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationGetClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationGetClient",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/oauth2/register/{cid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationGetClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationGetClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dynamicClientRegistrationGetClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DynamicClientRegistrationUpdateClient The OAuth 2.0 Dynamic Client Registration Update Client endpoint
*/
func (a *Client) DynamicClientRegistrationUpdateClient(params *DynamicClientRegistrationUpdateClientParams, authInfo runtime.ClientAuthInfoWriter) (*DynamicClientRegistrationUpdateClientOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationUpdateClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationUpdateClient",
		Method:             "PUT",
		PathPattern:        "/{tid}/{aid}/oauth2/register/{cid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationUpdateClientReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationUpdateClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dynamicClientRegistrationUpdateClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GatewayIntrospect Introspect access token endpoint as a gateway
*/
func (a *Client) GatewayIntrospect(params *GatewayIntrospectParams, authInfo runtime.ClientAuthInfoWriter) (*GatewayIntrospectOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGatewayIntrospectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "gatewayIntrospect",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/gateways/introspect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GatewayIntrospectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GatewayIntrospectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for gatewayIntrospect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Introspect thes o auth 2 0 introspection endpoint

  The introspection endpoint is an OAuth 2.0 endpoint that takes a
parameter representing an OAuth 2.0 token and returns a JSON
document representing the meta information surrounding the
token, including whether this token is currently active.
*/
func (a *Client) Introspect(params *IntrospectParams, authInfo runtime.ClientAuthInfoWriter) (*IntrospectOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewIntrospectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "introspect",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/oauth2/introspect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IntrospectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IntrospectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for introspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Jwks JSONs web keys discovery endpoint

  This endpoint returns the signing key(s) the client uses to validate
signatures from the authorization server.
*/
func (a *Client) Jwks(params *JwksParams) (*JwksOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewJwksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "jwks",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/.well-known/jwks.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &JwksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*JwksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for jwks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Revoke thes o auth 2 0 revocation endpoint

  Supports revocation of access and refresh tokens.
*/
func (a *Client) Revoke(params *RevokeParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRevokeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revoke",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/oauth2/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revoke: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Token thes o auth 2 0 token endpoint

  The token endpoint is used by the client to obtain an access token by
presenting its authorization grant or refresh token.
*/
func (a *Client) Token(params *TokenParams) (*TokenOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "token",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/oauth2/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Userinfo opens ID connect userinfo endpoint

  The UserInfo Endpoint is an OAuth 2.0 Protected Resource that
returns Claims about the authenticated End-User. To obtain the requested
Claims about the End-User, the Client makes a request to the UserInfo Endpoint
using an Access Token obtained through OpenID Connect Authentication. These Claims
are represented by a JSON object that contains a collection of name and value
pairs for the Claims.
*/
func (a *Client) Userinfo(params *UserinfoParams, authInfo runtime.ClientAuthInfoWriter) (*UserinfoOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userinfo",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserinfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserinfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userinfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WellKnown opens ID connect discovery endpoint

  Returns OpenID configuration.
*/
func (a *Client) WellKnown(params *WellKnownParams) (*WellKnownOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewWellKnownParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "wellKnown",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/.well-known/openid-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WellKnownReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WellKnownOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for wellKnown: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

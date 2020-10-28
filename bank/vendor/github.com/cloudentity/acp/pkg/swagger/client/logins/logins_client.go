// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new logins API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for logins API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AcceptGlobalLoginRequest(params *AcceptGlobalLoginRequestParams) (*AcceptGlobalLoginRequestOK, error)

	AcceptLoginRequest(params *AcceptLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptLoginRequestOK, error)

	AcceptScopeGrantRequest(params *AcceptScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptScopeGrantRequestOK, error)

	GetGlobalLoginRequest(params *GetGlobalLoginRequestParams) (*GetGlobalLoginRequestOK, error)

	GetLoginRequest(params *GetLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoginRequestOK, error)

	GetScopeGrantRequest(params *GetScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*GetScopeGrantRequestOK, error)

	RejectGlobalLoginRequest(params *RejectGlobalLoginRequestParams) (*RejectGlobalLoginRequestOK, error)

	RejectLoginRequest(params *RejectLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*RejectLoginRequestOK, error)

	RejectScopeGrantRequest(params *RejectScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*RejectScopeGrantRequestOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcceptGlobalLoginRequest accepts login request for any server

  This API can be globally used to notify ACP that user has successfully authenticated.
*/
func (a *Client) AcceptGlobalLoginRequest(params *AcceptGlobalLoginRequestParams) (*AcceptGlobalLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewAcceptGlobalLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acceptGlobalLoginRequest",
		Method:             "POST",
		PathPattern:        "/api/system/logins/{login}/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptGlobalLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptGlobalLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acceptGlobalLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AcceptLoginRequest accepts login request

  This API is used by a login page to notify ACP that user has successfully authenticated.
*/
func (a *Client) AcceptLoginRequest(params *AcceptLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewAcceptLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acceptLoginRequest",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/logins/{login}/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptLoginRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acceptLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AcceptScopeGrantRequest accepts login request

  This API is used by a consent page to notify ACP that user granted consent.
*/
func (a *Client) AcceptScopeGrantRequest(params *AcceptScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*AcceptScopeGrantRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewAcceptScopeGrantRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acceptScopeGrantRequest",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/scope-grants/{login}/accept",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcceptScopeGrantRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AcceptScopeGrantRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for acceptScopeGrantRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGlobalLoginRequest gets login request for any server

  This API can be globally used to make a decision if user should authenticate.
*/
func (a *Client) GetGlobalLoginRequest(params *GetGlobalLoginRequestParams) (*GetGlobalLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetGlobalLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGlobalLoginRequest",
		Method:             "GET",
		PathPattern:        "/api/system/logins/{login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGlobalLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLoginRequest gets login request

  This API is used by a login page to make a decision if user should authenticate.
*/
func (a *Client) GetLoginRequest(params *GetLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoginRequest",
		Method:             "GET",
		PathPattern:        "/api/system/{tid}/logins/{login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoginRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetScopeGrantRequest gets consent request

  This API is used by a consent page to display requested scopes.
*/
func (a *Client) GetScopeGrantRequest(params *GetScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*GetScopeGrantRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetScopeGrantRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScopeGrantRequest",
		Method:             "GET",
		PathPattern:        "/api/system/{tid}/scope-grants/{login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScopeGrantRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScopeGrantRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScopeGrantRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RejectGlobalLoginRequest rejects login request for any server

  This API can be globally used to notify ACP that login has been rejected.
*/
func (a *Client) RejectGlobalLoginRequest(params *RejectGlobalLoginRequestParams) (*RejectGlobalLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRejectGlobalLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rejectGlobalLoginRequest",
		Method:             "POST",
		PathPattern:        "/api/system/logins/{login}/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectGlobalLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RejectGlobalLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rejectGlobalLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RejectLoginRequest rejects login request

  This API is used by a login page to notify ACP that login has been rejected.
*/
func (a *Client) RejectLoginRequest(params *RejectLoginRequestParams, authInfo runtime.ClientAuthInfoWriter) (*RejectLoginRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRejectLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rejectLoginRequest",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/logins/{login}/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectLoginRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RejectLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rejectLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RejectScopeGrantRequest rejects login request

  This API is used by a consent page to notify ACP that scope grant has been rejected.
*/
func (a *Client) RejectScopeGrantRequest(params *RejectScopeGrantRequestParams, authInfo runtime.ClientAuthInfoWriter) (*RejectScopeGrantRequestOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRejectScopeGrantRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rejectScopeGrantRequest",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/scope-grants/{login}/reject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RejectScopeGrantRequestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RejectScopeGrantRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rejectScopeGrantRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

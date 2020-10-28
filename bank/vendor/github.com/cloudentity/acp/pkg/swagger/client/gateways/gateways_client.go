// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gateways API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gateways API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BindGroupToService(params *BindGroupToServiceParams, authInfo runtime.ClientAuthInfoWriter) (*BindGroupToServiceOK, error)

	CreateGateway(params *CreateGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*CreateGatewayCreated, error)

	DeleteGateway(params *DeleteGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteGatewayNoContent, error)

	GetGateway(params *GetGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewayOK, error)

	GetGatewayConfiguration(params *GetGatewayConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewayConfigurationOK, error)

	GetGatewayPackage(params *GetGatewayPackageParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetGatewayPackageOK, error)

	ListGatewayAPIGroups(params *ListGatewayAPIGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ListGatewayAPIGroupsOK, error)

	ListGateways(params *ListGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*ListGatewaysOK, error)

	PushGatewayRequests(params *PushGatewayRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*PushGatewayRequestsOK, error)

	SetGatewayConfiguration(params *SetGatewayConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*SetGatewayConfigurationOK, error)

	UnbindGroupFromService(params *UnbindGroupFromServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindGroupFromServiceOK, error)

	UpdateGateway(params *UpdateGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateGatewayOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BindGroupToService binds a group to a service

  It unbinds service first, if it is already connected to an api group.

It removes specification from a service, if it has specification imported already.

It removes any APIs added manually to a service, if it contains any.
*/
func (a *Client) BindGroupToService(params *BindGroupToServiceParams, authInfo runtime.ClientAuthInfoWriter) (*BindGroupToServiceOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewBindGroupToServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bindGroupToService",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}/groups/{apiGroup}/bind",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BindGroupToServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BindGroupToServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bindGroupToService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateGateway creates gateway

  Create gateway.
*/
func (a *Client) CreateGateway(params *CreateGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*CreateGatewayCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewCreateGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createGateway",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGatewayCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteGateway deletes gateway

  This removes configuration for all services connected to this gateway.
*/
func (a *Client) DeleteGateway(params *DeleteGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteGatewayNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDeleteGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGateway",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGatewayNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGateway gets gateway

  Get gateway.
*/
func (a *Client) GetGateway(params *GetGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewayOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGateway",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGatewayConfiguration gets configuration for a gateway

  This endpoint is used by api authorizers to fetch rules, policies and services to protect.
*/
func (a *Client) GetGatewayConfiguration(params *GetGatewayConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewayConfigurationOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetGatewayConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGatewayConfiguration",
		Method:             "GET",
		PathPattern:        "/api/system/{tid}/gateways/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGatewayConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewayConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGatewayConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGatewayPackage gets package for a gateway

  Get package for a gateway.
*/
func (a *Client) GetGatewayPackage(params *GetGatewayPackageParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetGatewayPackageOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetGatewayPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGatewayPackage",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}/package",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGatewayPackageReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewayPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGatewayPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListGatewayAPIGroups lists gateway api groups

  List gateway api groups.
*/
func (a *Client) ListGatewayAPIGroups(params *ListGatewayAPIGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ListGatewayAPIGroupsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListGatewayAPIGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listGatewayAPIGroups",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGatewayAPIGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGatewayAPIGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listGatewayAPIGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListGateways lists gateways

  List gateways.
*/
func (a *Client) ListGateways(params *ListGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*ListGatewaysOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listGateways",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listGateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PushGatewayRequests pushes gateway request

  Push gateway request.
*/
func (a *Client) PushGatewayRequests(params *PushGatewayRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*PushGatewayRequestsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewPushGatewayRequestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pushGatewayRequests",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/gateways/requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PushGatewayRequestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushGatewayRequestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushGatewayRequests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetGatewayConfiguration sets configuration for a gateway

  This endpoint is used to push apis protected by api gateway to acp.
*/
func (a *Client) SetGatewayConfiguration(params *SetGatewayConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*SetGatewayConfigurationOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewSetGatewayConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setGatewayConfiguration",
		Method:             "POST",
		PathPattern:        "/api/system/{tid}/gateways/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetGatewayConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetGatewayConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setGatewayConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnbindGroupFromService unbinds a group from a service

  Remove all apis and policies.

If policy is used by another service it will not be removed.

If service is connected to a gateway, it will be disconnected.

If a gateway api group is connected to this service, it will be disconnected.
*/
func (a *Client) UnbindGroupFromService(params *UnbindGroupFromServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindGroupFromServiceOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUnbindGroupFromServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unbindGroupFromService",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}/groups/{apiGroup}/unbind",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnbindGroupFromServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnbindGroupFromServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unbindGroupFromService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateGateway updates gateway

  Update gateway.
*/
func (a *Client) UpdateGateway(params *UpdateGatewayParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateGatewayOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUpdateGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGateway",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/gateways/{gw}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateGatewayReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

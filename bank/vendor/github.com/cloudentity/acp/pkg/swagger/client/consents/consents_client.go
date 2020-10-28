// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new consents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for consents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateConsent(params *CreateConsentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConsentCreated, error)

	CreateConsentAction(params *CreateConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConsentActionCreated, error)

	DeleteConsent(params *DeleteConsentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConsentNoContent, error)

	DeleteConsentAction(params *DeleteConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConsentActionNoContent, error)

	GetConsent(params *GetConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GetConsentOK, error)

	GetConsentAction(params *GetConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConsentActionOK, error)

	GrantConsent(params *GrantConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GrantConsentCreated, error)

	ListConsentActions(params *ListConsentActionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConsentActionsOK, error)

	ListConsents(params *ListConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConsentsOK, error)

	ListPrivacyLedgerEvents(params *ListPrivacyLedgerEventsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPrivacyLedgerEventsOK, error)

	ListPrivacyLedgerEventsBySubject(params *ListPrivacyLedgerEventsBySubjectParams, authInfo runtime.ClientAuthInfoWriter) (*ListPrivacyLedgerEventsBySubjectOK, error)

	ListUserConsents(params *ListUserConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsOK, error)

	ListUserConsentsByAction(params *ListUserConsentsByActionParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsByActionOK, error)

	ListUserConsentsByActionSystem(params *ListUserConsentsByActionSystemParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsByActionSystemOK, error)

	ListUserConsentsSystem(params *ListUserConsentsSystemParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsSystemOK, error)

	PatchConsentGrants(params *PatchConsentGrantsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConsentGrantsCreated, error)

	PatchConsentGrantsSystem(params *PatchConsentGrantsSystemParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConsentGrantsSystemCreated, error)

	RevokeConsent(params *RevokeConsentParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeConsentOK, error)

	UpdateConsent(params *UpdateConsentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConsentCreated, error)

	UpdateConsentAction(params *UpdateConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConsentActionCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateConsent creates consent

  Consents are created per tenant.

ID, Name and UpdateExistingGrants strategy are required fields when creating a new consent.

UpdateExistingGrants has the following options:

explicitAll - all the existing grants should not be updated. It is required that the user grants the consent explicitly.

implicitAll - all previously existing consent grants should be updated, to the new version of the consent, but
all of those grants should be implicit from the moment on.

keepCurrent - if a previously existing consent grant was set implicitly, it is automatically updated and a new
consent grant is produced, which is also implicit. if a previously existing consent grant was set
explicitly, it should not be updated. It is required that the user grants the consent explicitly.
*/
func (a *Client) CreateConsent(params *CreateConsentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConsentCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewCreateConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createConsent",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateConsentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateConsentAction creates consent action

  Consent action allows to group list of consents. A Tenant Application asking for the consent it can
ask for +by the action name+ instead of asking for a list of consents directly. It allows for cleaner
abstraction and adds the ability to dynamically configure the consents required for particular applications.

ID and Name are required fields.

For each consent provide id and specify if it is required.
*/
func (a *Client) CreateConsentAction(params *CreateConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*CreateConsentActionCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewCreateConsentActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createConsentAction",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConsentActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateConsentActionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createConsentAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteConsent deletes consent

  Delete consent.
*/
func (a *Client) DeleteConsent(params *DeleteConsentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConsentNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDeleteConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConsent",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/consents/{consent}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteConsentNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteConsentAction deletes consent action

  Delete consent action.
*/
func (a *Client) DeleteConsentAction(params *DeleteConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConsentActionNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDeleteConsentActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConsentAction",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/actions/{action}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConsentActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteConsentActionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteConsentAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConsent gets consent details

  Get consent details.
*/
func (a *Client) GetConsent(params *GetConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GetConsentOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConsent",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/consents/{consent}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConsentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConsentAction gets consent action with consents

  Get consent action with consents.
*/
func (a *Client) GetConsentAction(params *GetConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConsentActionOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetConsentActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConsentAction",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/actions/{action}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConsentActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConsentActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConsentAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GrantConsent grants privacy consent

  Consent id must be provided in the request body.

When a user grants consent which was already granted, it does not result in an error but it silently skipped instead.
*/
func (a *Client) GrantConsent(params *GrantConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GrantConsentCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGrantConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "grantConsent",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/privacy/consents/grant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GrantConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GrantConsentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for grantConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListConsentActions lists consent actions

  List consent actions.
*/
func (a *Client) ListConsentActions(params *ListConsentActionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConsentActionsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListConsentActionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConsentActions",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListConsentActionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListConsentActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listConsentActions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListConsents lists consents

  List consents.
*/
func (a *Client) ListConsents(params *ListConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListConsentsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listConsents",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListConsentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPrivacyLedgerEvents lists privacy ledger events

  It is possible to provide time constraints using from and to query params.
*/
func (a *Client) ListPrivacyLedgerEvents(params *ListPrivacyLedgerEventsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPrivacyLedgerEventsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListPrivacyLedgerEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPrivacyLedgerEvents",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/privacy/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPrivacyLedgerEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPrivacyLedgerEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPrivacyLedgerEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPrivacyLedgerEventsBySubject lists privacy ledger events

  List privacy ledger events.
*/
func (a *Client) ListPrivacyLedgerEventsBySubject(params *ListPrivacyLedgerEventsBySubjectParams, authInfo runtime.ClientAuthInfoWriter) (*ListPrivacyLedgerEventsBySubjectOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListPrivacyLedgerEventsBySubjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPrivacyLedgerEventsBySubject",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPrivacyLedgerEventsBySubjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPrivacyLedgerEventsBySubjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPrivacyLedgerEventsBySubject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserConsents lists consents

  If you want to list only specific consents, provide consent identifiers in query params.
*/
func (a *Client) ListUserConsents(params *ListUserConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListUserConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserConsents",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/privacy/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserConsentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserConsentsByAction lists consents by action

  Returns any possible required consents that the app should ask the User about.

The response includes a list of consents (including the ones user already agreed to).
Inclusion of the consents which the user already agreed to can be used to inform the user what he already agreed to.
*/
func (a *Client) ListUserConsentsByAction(params *ListUserConsentsByActionParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsByActionOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListUserConsentsByActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserConsentsByAction",
		Method:             "GET",
		PathPattern:        "/{tid}/{aid}/privacy/consents/{action}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserConsentsByActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserConsentsByActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserConsentsByAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserConsentsByActionSystem lists user consents by action

  User identifier must be provided in the header.
*/
func (a *Client) ListUserConsentsByActionSystem(params *ListUserConsentsByActionSystemParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsByActionSystemOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListUserConsentsByActionSystemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserConsentsByActionSystem",
		Method:             "GET",
		PathPattern:        "/api/system/{tid}/consents/{action}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserConsentsByActionSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserConsentsByActionSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserConsentsByActionSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUserConsentsSystem lists user consents

  User identifier must be provided in the header.

To limit list of consents, you can provide consent identifiers in query param.
*/
func (a *Client) ListUserConsentsSystem(params *ListUserConsentsSystemParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserConsentsSystemOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListUserConsentsSystemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserConsentsSystem",
		Method:             "GET",
		PathPattern:        "/api/system/{tid}/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserConsentsSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUserConsentsSystemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserConsentsSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchConsentGrants patches consent grants

  This is a non-standardized PATCH request.

Allows to update multiple consents approval in one API call.

See ConsentGrantPatchRequest object for more information.
*/
func (a *Client) PatchConsentGrants(params *PatchConsentGrantsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConsentGrantsCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewPatchConsentGrantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchConsentGrants",
		Method:             "PATCH",
		PathPattern:        "/{tid}/{aid}/privacy/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchConsentGrantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConsentGrantsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchConsentGrants: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchConsentGrantsSystem patches consent grants

  This is a non-standardized PATCH request.

Allows to update multiple consents approval in one API call.

User identifier must be provided in the header.

See ConsentGrantPatchRequest object for more information.
*/
func (a *Client) PatchConsentGrantsSystem(params *PatchConsentGrantsSystemParams, authInfo runtime.ClientAuthInfoWriter) (*PatchConsentGrantsSystemCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewPatchConsentGrantsSystemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchConsentGrantsSystem",
		Method:             "PATCH",
		PathPattern:        "/api/system/{tid}/consents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchConsentGrantsSystemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConsentGrantsSystemCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchConsentGrantsSystem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevokeConsent revokes privacy consent

  This API can be used to withdraw a consent which user previously gave.

Consent id must be provided in the request body.

When consent has the can_be_withdrawn flag set to false the API fails with an error saying that the consent cannot be revoked.
This flag is useful for scenarios in which the application cannot function without the consent from a User.
*/
func (a *Client) RevokeConsent(params *RevokeConsentParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeConsentOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRevokeConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeConsent",
		Method:             "POST",
		PathPattern:        "/{tid}/{aid}/privacy/consents/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeConsentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revokeConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateConsent updates consent

  Every time consent is updated, its version is incremented.

If ValidFrom attribute is not provided it will be set to current time.
*/
func (a *Client) UpdateConsent(params *UpdateConsentParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConsentCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUpdateConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConsent",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/consents/{consent}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateConsentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateConsentAction updates consent action

  Update consent action.
*/
func (a *Client) UpdateConsentAction(params *UpdateConsentActionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConsentActionCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUpdateConsentActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConsentAction",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/actions/{action}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateConsentActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateConsentActionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateConsentAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

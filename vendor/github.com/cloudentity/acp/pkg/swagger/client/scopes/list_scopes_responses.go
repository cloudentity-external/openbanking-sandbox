// Code generated by go-swagger; DO NOT EDIT.

package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// ListScopesReader is a Reader for the ListScopes structure.
type ListScopesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListScopesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListScopesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListScopesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListScopesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListScopesOK creates a ListScopesOK with default headers values
func NewListScopesOK() *ListScopesOK {
	return &ListScopesOK{}
}

/*ListScopesOK handles this case with default header values.

ScopesWithServices
*/
type ListScopesOK struct {
	Payload *models.ScopesWithServices
}

func (o *ListScopesOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/scopes][%d] listScopesOK  %+v", 200, o.Payload)
}

func (o *ListScopesOK) GetPayload() *models.ScopesWithServices {
	return o.Payload
}

func (o *ListScopesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopesWithServices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScopesUnauthorized creates a ListScopesUnauthorized with default headers values
func NewListScopesUnauthorized() *ListScopesUnauthorized {
	return &ListScopesUnauthorized{}
}

/*ListScopesUnauthorized handles this case with default header values.

HttpError
*/
type ListScopesUnauthorized struct {
	Payload *models.Error
}

func (o *ListScopesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/scopes][%d] listScopesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListScopesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListScopesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListScopesForbidden creates a ListScopesForbidden with default headers values
func NewListScopesForbidden() *ListScopesForbidden {
	return &ListScopesForbidden{}
}

/*ListScopesForbidden handles this case with default header values.

HttpError
*/
type ListScopesForbidden struct {
	Payload *models.Error
}

func (o *ListScopesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/scopes][%d] listScopesForbidden  %+v", 403, o.Payload)
}

func (o *ListScopesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListScopesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

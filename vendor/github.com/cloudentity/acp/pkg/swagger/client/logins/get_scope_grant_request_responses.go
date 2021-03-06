// Code generated by go-swagger; DO NOT EDIT.

package logins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetScopeGrantRequestReader is a Reader for the GetScopeGrantRequest structure.
type GetScopeGrantRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScopeGrantRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScopeGrantRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScopeGrantRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScopeGrantRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScopeGrantRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScopeGrantRequestOK creates a GetScopeGrantRequestOK with default headers values
func NewGetScopeGrantRequestOK() *GetScopeGrantRequestOK {
	return &GetScopeGrantRequestOK{}
}

/*GetScopeGrantRequestOK handles this case with default header values.

ScopeGrantSessionResponse
*/
type GetScopeGrantRequestOK struct {
	Payload *models.ScopeGrantSessionResponse
}

func (o *GetScopeGrantRequestOK) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/scope-grants/{login}][%d] getScopeGrantRequestOK  %+v", 200, o.Payload)
}

func (o *GetScopeGrantRequestOK) GetPayload() *models.ScopeGrantSessionResponse {
	return o.Payload
}

func (o *GetScopeGrantRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScopeGrantSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeGrantRequestUnauthorized creates a GetScopeGrantRequestUnauthorized with default headers values
func NewGetScopeGrantRequestUnauthorized() *GetScopeGrantRequestUnauthorized {
	return &GetScopeGrantRequestUnauthorized{}
}

/*GetScopeGrantRequestUnauthorized handles this case with default header values.

HttpError
*/
type GetScopeGrantRequestUnauthorized struct {
	Payload *models.Error
}

func (o *GetScopeGrantRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/scope-grants/{login}][%d] getScopeGrantRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScopeGrantRequestUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeGrantRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeGrantRequestForbidden creates a GetScopeGrantRequestForbidden with default headers values
func NewGetScopeGrantRequestForbidden() *GetScopeGrantRequestForbidden {
	return &GetScopeGrantRequestForbidden{}
}

/*GetScopeGrantRequestForbidden handles this case with default header values.

HttpError
*/
type GetScopeGrantRequestForbidden struct {
	Payload *models.Error
}

func (o *GetScopeGrantRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/scope-grants/{login}][%d] getScopeGrantRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetScopeGrantRequestForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeGrantRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeGrantRequestNotFound creates a GetScopeGrantRequestNotFound with default headers values
func NewGetScopeGrantRequestNotFound() *GetScopeGrantRequestNotFound {
	return &GetScopeGrantRequestNotFound{}
}

/*GetScopeGrantRequestNotFound handles this case with default header values.

HttpError
*/
type GetScopeGrantRequestNotFound struct {
	Payload *models.Error
}

func (o *GetScopeGrantRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/scope-grants/{login}][%d] getScopeGrantRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetScopeGrantRequestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScopeGrantRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

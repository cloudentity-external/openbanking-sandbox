// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetAuthorizationServerReader is a Reader for the GetAuthorizationServer structure.
type GetAuthorizationServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAuthorizationServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthorizationServerOK creates a GetAuthorizationServerOK with default headers values
func NewGetAuthorizationServerOK() *GetAuthorizationServerOK {
	return &GetAuthorizationServerOK{}
}

/*GetAuthorizationServerOK handles this case with default header values.

Server
*/
type GetAuthorizationServerOK struct {
	Payload *models.Server
}

func (o *GetAuthorizationServerOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}][%d] getAuthorizationServerOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationServerOK) GetPayload() *models.Server {
	return o.Payload
}

func (o *GetAuthorizationServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationServerUnauthorized creates a GetAuthorizationServerUnauthorized with default headers values
func NewGetAuthorizationServerUnauthorized() *GetAuthorizationServerUnauthorized {
	return &GetAuthorizationServerUnauthorized{}
}

/*GetAuthorizationServerUnauthorized handles this case with default header values.

HttpError
*/
type GetAuthorizationServerUnauthorized struct {
	Payload *models.Error
}

func (o *GetAuthorizationServerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}][%d] getAuthorizationServerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationServerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuthorizationServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationServerForbidden creates a GetAuthorizationServerForbidden with default headers values
func NewGetAuthorizationServerForbidden() *GetAuthorizationServerForbidden {
	return &GetAuthorizationServerForbidden{}
}

/*GetAuthorizationServerForbidden handles this case with default header values.

HttpError
*/
type GetAuthorizationServerForbidden struct {
	Payload *models.Error
}

func (o *GetAuthorizationServerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}][%d] getAuthorizationServerForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationServerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuthorizationServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationServerNotFound creates a GetAuthorizationServerNotFound with default headers values
func NewGetAuthorizationServerNotFound() *GetAuthorizationServerNotFound {
	return &GetAuthorizationServerNotFound{}
}

/*GetAuthorizationServerNotFound handles this case with default header values.

HttpError
*/
type GetAuthorizationServerNotFound struct {
	Payload *models.Error
}

func (o *GetAuthorizationServerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}][%d] getAuthorizationServerNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationServerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuthorizationServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

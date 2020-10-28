// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetClientReader is a Reader for the GetClient structure.
type GetClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientOK creates a GetClientOK with default headers values
func NewGetClientOK() *GetClientOK {
	return &GetClientOK{}
}

/*GetClientOK handles this case with default header values.

ClientAdminResponse
*/
type GetClientOK struct {
	Payload *models.ClientAdminResponse
}

func (o *GetClientOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/clients/{cid}][%d] getClientOK  %+v", 200, o.Payload)
}

func (o *GetClientOK) GetPayload() *models.ClientAdminResponse {
	return o.Payload
}

func (o *GetClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientAdminResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientUnauthorized creates a GetClientUnauthorized with default headers values
func NewGetClientUnauthorized() *GetClientUnauthorized {
	return &GetClientUnauthorized{}
}

/*GetClientUnauthorized handles this case with default header values.

HttpError
*/
type GetClientUnauthorized struct {
	Payload *models.Error
}

func (o *GetClientUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/clients/{cid}][%d] getClientUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClientUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientForbidden creates a GetClientForbidden with default headers values
func NewGetClientForbidden() *GetClientForbidden {
	return &GetClientForbidden{}
}

/*GetClientForbidden handles this case with default header values.

HttpError
*/
type GetClientForbidden struct {
	Payload *models.Error
}

func (o *GetClientForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/clients/{cid}][%d] getClientForbidden  %+v", 403, o.Payload)
}

func (o *GetClientForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientNotFound creates a GetClientNotFound with default headers values
func NewGetClientNotFound() *GetClientNotFound {
	return &GetClientNotFound{}
}

/*GetClientNotFound handles this case with default header values.

HttpError
*/
type GetClientNotFound struct {
	Payload *models.Error
}

func (o *GetClientNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/clients/{cid}][%d] getClientNotFound  %+v", 404, o.Payload)
}

func (o *GetClientNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// RevokeClientAccessReader is a Reader for the RevokeClientAccess structure.
type RevokeClientAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClientAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRevokeClientAccessNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClientAccessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClientAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevokeClientAccessNoContent creates a RevokeClientAccessNoContent with default headers values
func NewRevokeClientAccessNoContent() *RevokeClientAccessNoContent {
	return &RevokeClientAccessNoContent{}
}

/*RevokeClientAccessNoContent handles this case with default header values.

Client access has been revoked
*/
type RevokeClientAccessNoContent struct {
}

func (o *RevokeClientAccessNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{tid}/{aid}/clients/{cid}][%d] revokeClientAccessNoContent ", 204)
}

func (o *RevokeClientAccessNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClientAccessUnauthorized creates a RevokeClientAccessUnauthorized with default headers values
func NewRevokeClientAccessUnauthorized() *RevokeClientAccessUnauthorized {
	return &RevokeClientAccessUnauthorized{}
}

/*RevokeClientAccessUnauthorized handles this case with default header values.

HttpError
*/
type RevokeClientAccessUnauthorized struct {
	Payload *models.Error
}

func (o *RevokeClientAccessUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{tid}/{aid}/clients/{cid}][%d] revokeClientAccessUnauthorized  %+v", 401, o.Payload)
}

func (o *RevokeClientAccessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeClientAccessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeClientAccessForbidden creates a RevokeClientAccessForbidden with default headers values
func NewRevokeClientAccessForbidden() *RevokeClientAccessForbidden {
	return &RevokeClientAccessForbidden{}
}

/*RevokeClientAccessForbidden handles this case with default header values.

HttpError
*/
type RevokeClientAccessForbidden struct {
	Payload *models.Error
}

func (o *RevokeClientAccessForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{tid}/{aid}/clients/{cid}][%d] revokeClientAccessForbidden  %+v", 403, o.Payload)
}

func (o *RevokeClientAccessForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeClientAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

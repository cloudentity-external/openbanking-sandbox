// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// CreateGatewayReader is a Reader for the CreateGateway structure.
type CreateGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGatewayCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGatewayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateGatewayConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateGatewayUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGatewayCreated creates a CreateGatewayCreated with default headers values
func NewCreateGatewayCreated() *CreateGatewayCreated {
	return &CreateGatewayCreated{}
}

/*CreateGatewayCreated handles this case with default header values.

GatewayWithClient
*/
type CreateGatewayCreated struct {
	Payload *models.GatewayWithClient
}

func (o *CreateGatewayCreated) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayCreated  %+v", 201, o.Payload)
}

func (o *CreateGatewayCreated) GetPayload() *models.GatewayWithClient {
	return o.Payload
}

func (o *CreateGatewayCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayWithClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayUnauthorized creates a CreateGatewayUnauthorized with default headers values
func NewCreateGatewayUnauthorized() *CreateGatewayUnauthorized {
	return &CreateGatewayUnauthorized{}
}

/*CreateGatewayUnauthorized handles this case with default header values.

HttpError
*/
type CreateGatewayUnauthorized struct {
	Payload *models.Error
}

func (o *CreateGatewayUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateGatewayUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGatewayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayForbidden creates a CreateGatewayForbidden with default headers values
func NewCreateGatewayForbidden() *CreateGatewayForbidden {
	return &CreateGatewayForbidden{}
}

/*CreateGatewayForbidden handles this case with default header values.

HttpError
*/
type CreateGatewayForbidden struct {
	Payload *models.Error
}

func (o *CreateGatewayForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayForbidden  %+v", 403, o.Payload)
}

func (o *CreateGatewayForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayNotFound creates a CreateGatewayNotFound with default headers values
func NewCreateGatewayNotFound() *CreateGatewayNotFound {
	return &CreateGatewayNotFound{}
}

/*CreateGatewayNotFound handles this case with default header values.

HttpError
*/
type CreateGatewayNotFound struct {
	Payload *models.Error
}

func (o *CreateGatewayNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayNotFound  %+v", 404, o.Payload)
}

func (o *CreateGatewayNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayConflict creates a CreateGatewayConflict with default headers values
func NewCreateGatewayConflict() *CreateGatewayConflict {
	return &CreateGatewayConflict{}
}

/*CreateGatewayConflict handles this case with default header values.

HttpError
*/
type CreateGatewayConflict struct {
	Payload *models.Error
}

func (o *CreateGatewayConflict) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayConflict  %+v", 409, o.Payload)
}

func (o *CreateGatewayConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGatewayConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGatewayUnprocessableEntity creates a CreateGatewayUnprocessableEntity with default headers values
func NewCreateGatewayUnprocessableEntity() *CreateGatewayUnprocessableEntity {
	return &CreateGatewayUnprocessableEntity{}
}

/*CreateGatewayUnprocessableEntity handles this case with default header values.

HttpError
*/
type CreateGatewayUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateGatewayUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/gateways][%d] createGatewayUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateGatewayUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGatewayUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

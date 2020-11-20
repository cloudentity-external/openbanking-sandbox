// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// CreateConsentActionReader is a Reader for the CreateConsentAction structure.
type CreateConsentActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConsentActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConsentActionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateConsentActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateConsentActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateConsentActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateConsentActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateConsentActionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateConsentActionCreated creates a CreateConsentActionCreated with default headers values
func NewCreateConsentActionCreated() *CreateConsentActionCreated {
	return &CreateConsentActionCreated{}
}

/*CreateConsentActionCreated handles this case with default header values.

ConsentActionWithConsents
*/
type CreateConsentActionCreated struct {
	Payload *models.ConsentActionWithConsents
}

func (o *CreateConsentActionCreated) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionCreated  %+v", 201, o.Payload)
}

func (o *CreateConsentActionCreated) GetPayload() *models.ConsentActionWithConsents {
	return o.Payload
}

func (o *CreateConsentActionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentActionWithConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionUnauthorized creates a CreateConsentActionUnauthorized with default headers values
func NewCreateConsentActionUnauthorized() *CreateConsentActionUnauthorized {
	return &CreateConsentActionUnauthorized{}
}

/*CreateConsentActionUnauthorized handles this case with default header values.

HttpError
*/
type CreateConsentActionUnauthorized struct {
	Payload *models.Error
}

func (o *CreateConsentActionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateConsentActionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionForbidden creates a CreateConsentActionForbidden with default headers values
func NewCreateConsentActionForbidden() *CreateConsentActionForbidden {
	return &CreateConsentActionForbidden{}
}

/*CreateConsentActionForbidden handles this case with default header values.

HttpError
*/
type CreateConsentActionForbidden struct {
	Payload *models.Error
}

func (o *CreateConsentActionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionForbidden  %+v", 403, o.Payload)
}

func (o *CreateConsentActionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionNotFound creates a CreateConsentActionNotFound with default headers values
func NewCreateConsentActionNotFound() *CreateConsentActionNotFound {
	return &CreateConsentActionNotFound{}
}

/*CreateConsentActionNotFound handles this case with default header values.

HttpError
*/
type CreateConsentActionNotFound struct {
	Payload *models.Error
}

func (o *CreateConsentActionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionNotFound  %+v", 404, o.Payload)
}

func (o *CreateConsentActionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionConflict creates a CreateConsentActionConflict with default headers values
func NewCreateConsentActionConflict() *CreateConsentActionConflict {
	return &CreateConsentActionConflict{}
}

/*CreateConsentActionConflict handles this case with default header values.

HttpError
*/
type CreateConsentActionConflict struct {
	Payload *models.Error
}

func (o *CreateConsentActionConflict) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionConflict  %+v", 409, o.Payload)
}

func (o *CreateConsentActionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConsentActionUnprocessableEntity creates a CreateConsentActionUnprocessableEntity with default headers values
func NewCreateConsentActionUnprocessableEntity() *CreateConsentActionUnprocessableEntity {
	return &CreateConsentActionUnprocessableEntity{}
}

/*CreateConsentActionUnprocessableEntity handles this case with default header values.

HttpError
*/
type CreateConsentActionUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateConsentActionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/actions][%d] createConsentActionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateConsentActionUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateConsentActionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
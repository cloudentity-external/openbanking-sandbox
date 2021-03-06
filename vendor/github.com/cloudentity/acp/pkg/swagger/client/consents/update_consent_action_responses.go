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

// UpdateConsentActionReader is a Reader for the UpdateConsentAction structure.
type UpdateConsentActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConsentActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateConsentActionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateConsentActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateConsentActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConsentActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateConsentActionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateConsentActionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateConsentActionCreated creates a UpdateConsentActionCreated with default headers values
func NewUpdateConsentActionCreated() *UpdateConsentActionCreated {
	return &UpdateConsentActionCreated{}
}

/*UpdateConsentActionCreated handles this case with default header values.

ConsentActionWithConsents
*/
type UpdateConsentActionCreated struct {
	Payload *models.ConsentActionWithConsents
}

func (o *UpdateConsentActionCreated) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionCreated  %+v", 201, o.Payload)
}

func (o *UpdateConsentActionCreated) GetPayload() *models.ConsentActionWithConsents {
	return o.Payload
}

func (o *UpdateConsentActionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentActionWithConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentActionUnauthorized creates a UpdateConsentActionUnauthorized with default headers values
func NewUpdateConsentActionUnauthorized() *UpdateConsentActionUnauthorized {
	return &UpdateConsentActionUnauthorized{}
}

/*UpdateConsentActionUnauthorized handles this case with default header values.

HttpError
*/
type UpdateConsentActionUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateConsentActionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateConsentActionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentActionForbidden creates a UpdateConsentActionForbidden with default headers values
func NewUpdateConsentActionForbidden() *UpdateConsentActionForbidden {
	return &UpdateConsentActionForbidden{}
}

/*UpdateConsentActionForbidden handles this case with default header values.

HttpError
*/
type UpdateConsentActionForbidden struct {
	Payload *models.Error
}

func (o *UpdateConsentActionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConsentActionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentActionNotFound creates a UpdateConsentActionNotFound with default headers values
func NewUpdateConsentActionNotFound() *UpdateConsentActionNotFound {
	return &UpdateConsentActionNotFound{}
}

/*UpdateConsentActionNotFound handles this case with default header values.

HttpError
*/
type UpdateConsentActionNotFound struct {
	Payload *models.Error
}

func (o *UpdateConsentActionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConsentActionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentActionConflict creates a UpdateConsentActionConflict with default headers values
func NewUpdateConsentActionConflict() *UpdateConsentActionConflict {
	return &UpdateConsentActionConflict{}
}

/*UpdateConsentActionConflict handles this case with default header values.

HttpError
*/
type UpdateConsentActionConflict struct {
	Payload *models.Error
}

func (o *UpdateConsentActionConflict) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionConflict  %+v", 409, o.Payload)
}

func (o *UpdateConsentActionConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentActionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConsentActionUnprocessableEntity creates a UpdateConsentActionUnprocessableEntity with default headers values
func NewUpdateConsentActionUnprocessableEntity() *UpdateConsentActionUnprocessableEntity {
	return &UpdateConsentActionUnprocessableEntity{}
}

/*UpdateConsentActionUnprocessableEntity handles this case with default header values.

HttpError
*/
type UpdateConsentActionUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateConsentActionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/actions/{action}][%d] updateConsentActionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateConsentActionUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConsentActionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

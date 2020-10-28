// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// UpdateAPIReader is a Reader for the UpdateAPI structure.
type UpdateAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateAPIUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAPIOK creates a UpdateAPIOK with default headers values
func NewUpdateAPIOK() *UpdateAPIOK {
	return &UpdateAPIOK{}
}

/*UpdateAPIOK handles this case with default header values.

API
*/
type UpdateAPIOK struct {
	Payload *models.API
}

func (o *UpdateAPIOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiOK  %+v", 200, o.Payload)
}

func (o *UpdateAPIOK) GetPayload() *models.API {
	return o.Payload
}

func (o *UpdateAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.API)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIBadRequest creates a UpdateAPIBadRequest with default headers values
func NewUpdateAPIBadRequest() *UpdateAPIBadRequest {
	return &UpdateAPIBadRequest{}
}

/*UpdateAPIBadRequest handles this case with default header values.

HttpError
*/
type UpdateAPIBadRequest struct {
	Payload *models.Error
}

func (o *UpdateAPIBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAPIBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIUnauthorized creates a UpdateAPIUnauthorized with default headers values
func NewUpdateAPIUnauthorized() *UpdateAPIUnauthorized {
	return &UpdateAPIUnauthorized{}
}

/*UpdateAPIUnauthorized handles this case with default header values.

HttpError
*/
type UpdateAPIUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateAPIUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAPIUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIForbidden creates a UpdateAPIForbidden with default headers values
func NewUpdateAPIForbidden() *UpdateAPIForbidden {
	return &UpdateAPIForbidden{}
}

/*UpdateAPIForbidden handles this case with default header values.

HttpError
*/
type UpdateAPIForbidden struct {
	Payload *models.Error
}

func (o *UpdateAPIForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAPIForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPINotFound creates a UpdateAPINotFound with default headers values
func NewUpdateAPINotFound() *UpdateAPINotFound {
	return &UpdateAPINotFound{}
}

/*UpdateAPINotFound handles this case with default header values.

HttpError
*/
type UpdateAPINotFound struct {
	Payload *models.Error
}

func (o *UpdateAPINotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAPINotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAPIUnprocessableEntity creates a UpdateAPIUnprocessableEntity with default headers values
func NewUpdateAPIUnprocessableEntity() *UpdateAPIUnprocessableEntity {
	return &UpdateAPIUnprocessableEntity{}
}

/*UpdateAPIUnprocessableEntity handles this case with default header values.

HttpError
*/
type UpdateAPIUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateAPIUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/apis/{api}][%d] updateApiUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateAPIUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAPIUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

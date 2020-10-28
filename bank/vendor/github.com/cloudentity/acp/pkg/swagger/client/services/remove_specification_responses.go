// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// RemoveSpecificationReader is a Reader for the RemoveSpecification structure.
type RemoveSpecificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveSpecificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveSpecificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveSpecificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveSpecificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveSpecificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveSpecificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRemoveSpecificationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRemoveSpecificationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveSpecificationOK creates a RemoveSpecificationOK with default headers values
func NewRemoveSpecificationOK() *RemoveSpecificationOK {
	return &RemoveSpecificationOK{}
}

/*RemoveSpecificationOK handles this case with default header values.

RemoveServiceConfigurationResult
*/
type RemoveSpecificationOK struct {
	Payload *models.RemoveServiceConfigurationResult
}

func (o *RemoveSpecificationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationOK  %+v", 200, o.Payload)
}

func (o *RemoveSpecificationOK) GetPayload() *models.RemoveServiceConfigurationResult {
	return o.Payload
}

func (o *RemoveSpecificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoveServiceConfigurationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationBadRequest creates a RemoveSpecificationBadRequest with default headers values
func NewRemoveSpecificationBadRequest() *RemoveSpecificationBadRequest {
	return &RemoveSpecificationBadRequest{}
}

/*RemoveSpecificationBadRequest handles this case with default header values.

HttpError
*/
type RemoveSpecificationBadRequest struct {
	Payload *models.Error
}

func (o *RemoveSpecificationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveSpecificationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationUnauthorized creates a RemoveSpecificationUnauthorized with default headers values
func NewRemoveSpecificationUnauthorized() *RemoveSpecificationUnauthorized {
	return &RemoveSpecificationUnauthorized{}
}

/*RemoveSpecificationUnauthorized handles this case with default header values.

HttpError
*/
type RemoveSpecificationUnauthorized struct {
	Payload *models.Error
}

func (o *RemoveSpecificationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveSpecificationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationForbidden creates a RemoveSpecificationForbidden with default headers values
func NewRemoveSpecificationForbidden() *RemoveSpecificationForbidden {
	return &RemoveSpecificationForbidden{}
}

/*RemoveSpecificationForbidden handles this case with default header values.

HttpError
*/
type RemoveSpecificationForbidden struct {
	Payload *models.Error
}

func (o *RemoveSpecificationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationForbidden  %+v", 403, o.Payload)
}

func (o *RemoveSpecificationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationNotFound creates a RemoveSpecificationNotFound with default headers values
func NewRemoveSpecificationNotFound() *RemoveSpecificationNotFound {
	return &RemoveSpecificationNotFound{}
}

/*RemoveSpecificationNotFound handles this case with default header values.

HttpError
*/
type RemoveSpecificationNotFound struct {
	Payload *models.Error
}

func (o *RemoveSpecificationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationNotFound  %+v", 404, o.Payload)
}

func (o *RemoveSpecificationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationConflict creates a RemoveSpecificationConflict with default headers values
func NewRemoveSpecificationConflict() *RemoveSpecificationConflict {
	return &RemoveSpecificationConflict{}
}

/*RemoveSpecificationConflict handles this case with default header values.

HttpError
*/
type RemoveSpecificationConflict struct {
	Payload *models.Error
}

func (o *RemoveSpecificationConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationConflict  %+v", 409, o.Payload)
}

func (o *RemoveSpecificationConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveSpecificationUnprocessableEntity creates a RemoveSpecificationUnprocessableEntity with default headers values
func NewRemoveSpecificationUnprocessableEntity() *RemoveSpecificationUnprocessableEntity {
	return &RemoveSpecificationUnprocessableEntity{}
}

/*RemoveSpecificationUnprocessableEntity handles this case with default header values.

HttpError
*/
type RemoveSpecificationUnprocessableEntity struct {
	Payload *models.Error
}

func (o *RemoveSpecificationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/services/{sid}/apis][%d] removeSpecificationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RemoveSpecificationUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveSpecificationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

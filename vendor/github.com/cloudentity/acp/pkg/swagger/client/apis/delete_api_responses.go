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

// DeleteAPIReader is a Reader for the DeleteAPI structure.
type DeleteAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPINoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPINoContent creates a DeleteAPINoContent with default headers values
func NewDeleteAPINoContent() *DeleteAPINoContent {
	return &DeleteAPINoContent{}
}

/*DeleteAPINoContent handles this case with default header values.

API has been deleted
*/
type DeleteAPINoContent struct {
}

func (o *DeleteAPINoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/apis/{api}][%d] deleteApiNoContent ", 204)
}

func (o *DeleteAPINoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIUnauthorized creates a DeleteAPIUnauthorized with default headers values
func NewDeleteAPIUnauthorized() *DeleteAPIUnauthorized {
	return &DeleteAPIUnauthorized{}
}

/*DeleteAPIUnauthorized handles this case with default header values.

HttpError
*/
type DeleteAPIUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteAPIUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/apis/{api}][%d] deleteApiUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAPIUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPIForbidden creates a DeleteAPIForbidden with default headers values
func NewDeleteAPIForbidden() *DeleteAPIForbidden {
	return &DeleteAPIForbidden{}
}

/*DeleteAPIForbidden handles this case with default header values.

HttpError
*/
type DeleteAPIForbidden struct {
	Payload *models.Error
}

func (o *DeleteAPIForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/apis/{api}][%d] deleteApiForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAPIForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPINotFound creates a DeleteAPINotFound with default headers values
func NewDeleteAPINotFound() *DeleteAPINotFound {
	return &DeleteAPINotFound{}
}

/*DeleteAPINotFound handles this case with default header values.

HttpError
*/
type DeleteAPINotFound struct {
	Payload *models.Error
}

func (o *DeleteAPINotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/apis/{api}][%d] deleteApiNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAPINotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

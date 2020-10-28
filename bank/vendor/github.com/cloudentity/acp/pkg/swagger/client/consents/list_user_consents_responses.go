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

// ListUserConsentsReader is a Reader for the ListUserConsents structure.
type ListUserConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUserConsentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUserConsentsOK creates a ListUserConsentsOK with default headers values
func NewListUserConsentsOK() *ListUserConsentsOK {
	return &ListUserConsentsOK{}
}

/*ListUserConsentsOK handles this case with default header values.

ConsentsWithGrants
*/
type ListUserConsentsOK struct {
	Payload *models.ConsentsWithGrants
}

func (o *ListUserConsentsOK) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/privacy/consents][%d] listUserConsentsOK  %+v", 200, o.Payload)
}

func (o *ListUserConsentsOK) GetPayload() *models.ConsentsWithGrants {
	return o.Payload
}

func (o *ListUserConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentsWithGrants)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserConsentsUnauthorized creates a ListUserConsentsUnauthorized with default headers values
func NewListUserConsentsUnauthorized() *ListUserConsentsUnauthorized {
	return &ListUserConsentsUnauthorized{}
}

/*ListUserConsentsUnauthorized handles this case with default header values.

HttpError
*/
type ListUserConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *ListUserConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/privacy/consents][%d] listUserConsentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUserConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListUserConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserConsentsForbidden creates a ListUserConsentsForbidden with default headers values
func NewListUserConsentsForbidden() *ListUserConsentsForbidden {
	return &ListUserConsentsForbidden{}
}

/*ListUserConsentsForbidden handles this case with default header values.

HttpError
*/
type ListUserConsentsForbidden struct {
	Payload *models.Error
}

func (o *ListUserConsentsForbidden) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/privacy/consents][%d] listUserConsentsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListUserConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserConsentsNotFound creates a ListUserConsentsNotFound with default headers values
func NewListUserConsentsNotFound() *ListUserConsentsNotFound {
	return &ListUserConsentsNotFound{}
}

/*ListUserConsentsNotFound handles this case with default header values.

HttpError
*/
type ListUserConsentsNotFound struct {
	Payload *models.Error
}

func (o *ListUserConsentsNotFound) Error() string {
	return fmt.Sprintf("[GET /{tid}/{aid}/privacy/consents][%d] listUserConsentsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserConsentsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListUserConsentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

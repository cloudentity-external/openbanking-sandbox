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

// GetConsentActionReader is a Reader for the GetConsentAction structure.
type GetConsentActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsentActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsentActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConsentActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConsentActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConsentActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConsentActionOK creates a GetConsentActionOK with default headers values
func NewGetConsentActionOK() *GetConsentActionOK {
	return &GetConsentActionOK{}
}

/*GetConsentActionOK handles this case with default header values.

ConsentActionWithConsents
*/
type GetConsentActionOK struct {
	Payload *models.ConsentActionWithConsents
}

func (o *GetConsentActionOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/actions/{action}][%d] getConsentActionOK  %+v", 200, o.Payload)
}

func (o *GetConsentActionOK) GetPayload() *models.ConsentActionWithConsents {
	return o.Payload
}

func (o *GetConsentActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentActionWithConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentActionUnauthorized creates a GetConsentActionUnauthorized with default headers values
func NewGetConsentActionUnauthorized() *GetConsentActionUnauthorized {
	return &GetConsentActionUnauthorized{}
}

/*GetConsentActionUnauthorized handles this case with default header values.

HttpError
*/
type GetConsentActionUnauthorized struct {
	Payload *models.Error
}

func (o *GetConsentActionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/actions/{action}][%d] getConsentActionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConsentActionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentActionForbidden creates a GetConsentActionForbidden with default headers values
func NewGetConsentActionForbidden() *GetConsentActionForbidden {
	return &GetConsentActionForbidden{}
}

/*GetConsentActionForbidden handles this case with default header values.

HttpError
*/
type GetConsentActionForbidden struct {
	Payload *models.Error
}

func (o *GetConsentActionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/actions/{action}][%d] getConsentActionForbidden  %+v", 403, o.Payload)
}

func (o *GetConsentActionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentActionNotFound creates a GetConsentActionNotFound with default headers values
func NewGetConsentActionNotFound() *GetConsentActionNotFound {
	return &GetConsentActionNotFound{}
}

/*GetConsentActionNotFound handles this case with default header values.

HttpError
*/
type GetConsentActionNotFound struct {
	Payload *models.Error
}

func (o *GetConsentActionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/actions/{action}][%d] getConsentActionNotFound  %+v", 404, o.Payload)
}

func (o *GetConsentActionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsentActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

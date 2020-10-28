// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// DynamicClientRegistrationReader is a Reader for the DynamicClientRegistration structure.
type DynamicClientRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DynamicClientRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDynamicClientRegistrationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDynamicClientRegistrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDynamicClientRegistrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDynamicClientRegistrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDynamicClientRegistrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDynamicClientRegistrationCreated creates a DynamicClientRegistrationCreated with default headers values
func NewDynamicClientRegistrationCreated() *DynamicClientRegistrationCreated {
	return &DynamicClientRegistrationCreated{}
}

/*DynamicClientRegistrationCreated handles this case with default header values.

DynamicClientRegistrationResponse
*/
type DynamicClientRegistrationCreated struct {
	Payload *models.DynamicClientRegistrationResponse
}

func (o *DynamicClientRegistrationCreated) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/register][%d] dynamicClientRegistrationCreated  %+v", 201, o.Payload)
}

func (o *DynamicClientRegistrationCreated) GetPayload() *models.DynamicClientRegistrationResponse {
	return o.Payload
}

func (o *DynamicClientRegistrationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DynamicClientRegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDynamicClientRegistrationBadRequest creates a DynamicClientRegistrationBadRequest with default headers values
func NewDynamicClientRegistrationBadRequest() *DynamicClientRegistrationBadRequest {
	return &DynamicClientRegistrationBadRequest{}
}

/*DynamicClientRegistrationBadRequest handles this case with default header values.

RFC6749Error
*/
type DynamicClientRegistrationBadRequest struct {
	Payload *models.RFC6749Error
}

func (o *DynamicClientRegistrationBadRequest) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/register][%d] dynamicClientRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *DynamicClientRegistrationBadRequest) GetPayload() *models.RFC6749Error {
	return o.Payload
}

func (o *DynamicClientRegistrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RFC6749Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDynamicClientRegistrationUnauthorized creates a DynamicClientRegistrationUnauthorized with default headers values
func NewDynamicClientRegistrationUnauthorized() *DynamicClientRegistrationUnauthorized {
	return &DynamicClientRegistrationUnauthorized{}
}

/*DynamicClientRegistrationUnauthorized handles this case with default header values.

RFC6749Error
*/
type DynamicClientRegistrationUnauthorized struct {
	Payload *models.RFC6749Error
}

func (o *DynamicClientRegistrationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/register][%d] dynamicClientRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *DynamicClientRegistrationUnauthorized) GetPayload() *models.RFC6749Error {
	return o.Payload
}

func (o *DynamicClientRegistrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RFC6749Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDynamicClientRegistrationForbidden creates a DynamicClientRegistrationForbidden with default headers values
func NewDynamicClientRegistrationForbidden() *DynamicClientRegistrationForbidden {
	return &DynamicClientRegistrationForbidden{}
}

/*DynamicClientRegistrationForbidden handles this case with default header values.

RFC6749Error
*/
type DynamicClientRegistrationForbidden struct {
	Payload *models.RFC6749Error
}

func (o *DynamicClientRegistrationForbidden) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/register][%d] dynamicClientRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *DynamicClientRegistrationForbidden) GetPayload() *models.RFC6749Error {
	return o.Payload
}

func (o *DynamicClientRegistrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RFC6749Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDynamicClientRegistrationNotFound creates a DynamicClientRegistrationNotFound with default headers values
func NewDynamicClientRegistrationNotFound() *DynamicClientRegistrationNotFound {
	return &DynamicClientRegistrationNotFound{}
}

/*DynamicClientRegistrationNotFound handles this case with default header values.

genericError
*/
type DynamicClientRegistrationNotFound struct {
	Payload *models.GenericError
}

func (o *DynamicClientRegistrationNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/register][%d] dynamicClientRegistrationNotFound  %+v", 404, o.Payload)
}

func (o *DynamicClientRegistrationNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DynamicClientRegistrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

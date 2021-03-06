// Code generated by go-swagger; DO NOT EDIT.

package openbanking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// OpenbankingAccountAccessConsentIntrospectReader is a Reader for the OpenbankingAccountAccessConsentIntrospect structure.
type OpenbankingAccountAccessConsentIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenbankingAccountAccessConsentIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenbankingAccountAccessConsentIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOpenbankingAccountAccessConsentIntrospectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenbankingAccountAccessConsentIntrospectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOpenbankingAccountAccessConsentIntrospectOK creates a OpenbankingAccountAccessConsentIntrospectOK with default headers values
func NewOpenbankingAccountAccessConsentIntrospectOK() *OpenbankingAccountAccessConsentIntrospectOK {
	return &OpenbankingAccountAccessConsentIntrospectOK{}
}

/*OpenbankingAccountAccessConsentIntrospectOK handles this case with default header values.

IntrospectOpenbankingAccountAccessConsentResponse
*/
type OpenbankingAccountAccessConsentIntrospectOK struct {
	Payload *models.IntrospectOpenbankingAccountAccessConsentResponse
}

func (o *OpenbankingAccountAccessConsentIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents/introspect][%d] openbankingAccountAccessConsentIntrospectOK  %+v", 200, o.Payload)
}

func (o *OpenbankingAccountAccessConsentIntrospectOK) GetPayload() *models.IntrospectOpenbankingAccountAccessConsentResponse {
	return o.Payload
}

func (o *OpenbankingAccountAccessConsentIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntrospectOpenbankingAccountAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingAccountAccessConsentIntrospectUnauthorized creates a OpenbankingAccountAccessConsentIntrospectUnauthorized with default headers values
func NewOpenbankingAccountAccessConsentIntrospectUnauthorized() *OpenbankingAccountAccessConsentIntrospectUnauthorized {
	return &OpenbankingAccountAccessConsentIntrospectUnauthorized{}
}

/*OpenbankingAccountAccessConsentIntrospectUnauthorized handles this case with default header values.

genericError
*/
type OpenbankingAccountAccessConsentIntrospectUnauthorized struct {
	Payload *models.GenericError
}

func (o *OpenbankingAccountAccessConsentIntrospectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents/introspect][%d] openbankingAccountAccessConsentIntrospectUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenbankingAccountAccessConsentIntrospectUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingAccountAccessConsentIntrospectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenbankingAccountAccessConsentIntrospectNotFound creates a OpenbankingAccountAccessConsentIntrospectNotFound with default headers values
func NewOpenbankingAccountAccessConsentIntrospectNotFound() *OpenbankingAccountAccessConsentIntrospectNotFound {
	return &OpenbankingAccountAccessConsentIntrospectNotFound{}
}

/*OpenbankingAccountAccessConsentIntrospectNotFound handles this case with default header values.

genericError
*/
type OpenbankingAccountAccessConsentIntrospectNotFound struct {
	Payload *models.GenericError
}

func (o *OpenbankingAccountAccessConsentIntrospectNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents/introspect][%d] openbankingAccountAccessConsentIntrospectNotFound  %+v", 404, o.Payload)
}

func (o *OpenbankingAccountAccessConsentIntrospectNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *OpenbankingAccountAccessConsentIntrospectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

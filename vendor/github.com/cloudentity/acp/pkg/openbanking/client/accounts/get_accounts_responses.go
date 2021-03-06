// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp/pkg/openbanking/models"
)

// GetAccountsReader is a Reader for the GetAccounts structure.
type GetAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountsOK creates a GetAccountsOK with default headers values
func NewGetAccountsOK() *GetAccountsOK {
	return &GetAccountsOK{}
}

/*GetAccountsOK handles this case with default header values.

Accounts Read
*/
type GetAccountsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadAccount6
}

func (o *GetAccountsOK) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsOK) GetPayload() *models.OBReadAccount6 {
	return o.Payload
}

func (o *GetAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadAccount6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsBadRequest creates a GetAccountsBadRequest with default headers values
func NewGetAccountsBadRequest() *GetAccountsBadRequest {
	return &GetAccountsBadRequest{}
}

/*GetAccountsBadRequest handles this case with default header values.

Bad request
*/
type GetAccountsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsUnauthorized creates a GetAccountsUnauthorized with default headers values
func NewGetAccountsUnauthorized() *GetAccountsUnauthorized {
	return &GetAccountsUnauthorized{}
}

/*GetAccountsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsUnauthorized ", 401)
}

func (o *GetAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsForbidden creates a GetAccountsForbidden with default headers values
func NewGetAccountsForbidden() *GetAccountsForbidden {
	return &GetAccountsForbidden{}
}

/*GetAccountsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsMethodNotAllowed creates a GetAccountsMethodNotAllowed with default headers values
func NewGetAccountsMethodNotAllowed() *GetAccountsMethodNotAllowed {
	return &GetAccountsMethodNotAllowed{}
}

/*GetAccountsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAccountsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsMethodNotAllowed ", 405)
}

func (o *GetAccountsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsNotAcceptable creates a GetAccountsNotAcceptable with default headers values
func NewGetAccountsNotAcceptable() *GetAccountsNotAcceptable {
	return &GetAccountsNotAcceptable{}
}

/*GetAccountsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetAccountsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsNotAcceptable ", 406)
}

func (o *GetAccountsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsTooManyRequests creates a GetAccountsTooManyRequests with default headers values
func NewGetAccountsTooManyRequests() *GetAccountsTooManyRequests {
	return &GetAccountsTooManyRequests{}
}

/*GetAccountsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetAccountsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsTooManyRequests ", 429)
}

func (o *GetAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Retry-After
	retryAfter, err := swag.ConvertInt64(response.GetHeader("Retry-After"))
	if err != nil {
		return errors.InvalidType("Retry-After", "header", "int64", response.GetHeader("Retry-After"))
	}
	o.RetryAfter = retryAfter

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsInternalServerError creates a GetAccountsInternalServerError with default headers values
func NewGetAccountsInternalServerError() *GetAccountsInternalServerError {
	return &GetAccountsInternalServerError{}
}

/*GetAccountsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts][%d] getAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

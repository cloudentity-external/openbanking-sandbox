// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// GetTransactionsReader is a Reader for the GetTransactions structure.
type GetTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetTransactionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetTransactionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTransactionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionsOK creates a GetTransactionsOK with default headers values
func NewGetTransactionsOK() *GetTransactionsOK {
	return &GetTransactionsOK{}
}

/*GetTransactionsOK handles this case with default header values.

Transactions Read
*/
type GetTransactionsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadTransaction6
}

func (o *GetTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionsOK) GetPayload() *models.OBReadTransaction6 {
	return o.Payload
}

func (o *GetTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadTransaction6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsBadRequest creates a GetTransactionsBadRequest with default headers values
func NewGetTransactionsBadRequest() *GetTransactionsBadRequest {
	return &GetTransactionsBadRequest{}
}

/*GetTransactionsBadRequest handles this case with default header values.

Bad request
*/
type GetTransactionsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsUnauthorized creates a GetTransactionsUnauthorized with default headers values
func NewGetTransactionsUnauthorized() *GetTransactionsUnauthorized {
	return &GetTransactionsUnauthorized{}
}

/*GetTransactionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTransactionsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsUnauthorized ", 401)
}

func (o *GetTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetTransactionsForbidden creates a GetTransactionsForbidden with default headers values
func NewGetTransactionsForbidden() *GetTransactionsForbidden {
	return &GetTransactionsForbidden{}
}

/*GetTransactionsForbidden handles this case with default header values.

Forbidden
*/
type GetTransactionsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsNotFound creates a GetTransactionsNotFound with default headers values
func NewGetTransactionsNotFound() *GetTransactionsNotFound {
	return &GetTransactionsNotFound{}
}

/*GetTransactionsNotFound handles this case with default header values.

Not found
*/
type GetTransactionsNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetTransactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsNotFound ", 404)
}

func (o *GetTransactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetTransactionsMethodNotAllowed creates a GetTransactionsMethodNotAllowed with default headers values
func NewGetTransactionsMethodNotAllowed() *GetTransactionsMethodNotAllowed {
	return &GetTransactionsMethodNotAllowed{}
}

/*GetTransactionsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetTransactionsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetTransactionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsMethodNotAllowed ", 405)
}

func (o *GetTransactionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetTransactionsNotAcceptable creates a GetTransactionsNotAcceptable with default headers values
func NewGetTransactionsNotAcceptable() *GetTransactionsNotAcceptable {
	return &GetTransactionsNotAcceptable{}
}

/*GetTransactionsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetTransactionsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetTransactionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsNotAcceptable ", 406)
}

func (o *GetTransactionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetTransactionsTooManyRequests creates a GetTransactionsTooManyRequests with default headers values
func NewGetTransactionsTooManyRequests() *GetTransactionsTooManyRequests {
	return &GetTransactionsTooManyRequests{}
}

/*GetTransactionsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetTransactionsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetTransactionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsTooManyRequests ", 429)
}

func (o *GetTransactionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTransactionsInternalServerError creates a GetTransactionsInternalServerError with default headers values
func NewGetTransactionsInternalServerError() *GetTransactionsInternalServerError {
	return &GetTransactionsInternalServerError{}
}

/*GetTransactionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetTransactionsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /transactions][%d] getTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTransactionsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

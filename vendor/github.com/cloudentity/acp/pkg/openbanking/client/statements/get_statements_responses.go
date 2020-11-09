// Code generated by go-swagger; DO NOT EDIT.

package statements

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

// GetStatementsReader is a Reader for the GetStatements structure.
type GetStatementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStatementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStatementsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStatementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStatementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetStatementsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetStatementsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStatementsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStatementsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStatementsOK creates a GetStatementsOK with default headers values
func NewGetStatementsOK() *GetStatementsOK {
	return &GetStatementsOK{}
}

/*GetStatementsOK handles this case with default header values.

Statements Read
*/
type GetStatementsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadStatement2
}

func (o *GetStatementsOK) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsOK  %+v", 200, o.Payload)
}

func (o *GetStatementsOK) GetPayload() *models.OBReadStatement2 {
	return o.Payload
}

func (o *GetStatementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadStatement2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatementsBadRequest creates a GetStatementsBadRequest with default headers values
func NewGetStatementsBadRequest() *GetStatementsBadRequest {
	return &GetStatementsBadRequest{}
}

/*GetStatementsBadRequest handles this case with default header values.

Bad request
*/
type GetStatementsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStatementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsBadRequest  %+v", 400, o.Payload)
}

func (o *GetStatementsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStatementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatementsUnauthorized creates a GetStatementsUnauthorized with default headers values
func NewGetStatementsUnauthorized() *GetStatementsUnauthorized {
	return &GetStatementsUnauthorized{}
}

/*GetStatementsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetStatementsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStatementsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsUnauthorized ", 401)
}

func (o *GetStatementsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetStatementsForbidden creates a GetStatementsForbidden with default headers values
func NewGetStatementsForbidden() *GetStatementsForbidden {
	return &GetStatementsForbidden{}
}

/*GetStatementsForbidden handles this case with default header values.

Forbidden
*/
type GetStatementsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStatementsForbidden) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsForbidden  %+v", 403, o.Payload)
}

func (o *GetStatementsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStatementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatementsNotFound creates a GetStatementsNotFound with default headers values
func NewGetStatementsNotFound() *GetStatementsNotFound {
	return &GetStatementsNotFound{}
}

/*GetStatementsNotFound handles this case with default header values.

Not found
*/
type GetStatementsNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStatementsNotFound) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsNotFound ", 404)
}

func (o *GetStatementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetStatementsMethodNotAllowed creates a GetStatementsMethodNotAllowed with default headers values
func NewGetStatementsMethodNotAllowed() *GetStatementsMethodNotAllowed {
	return &GetStatementsMethodNotAllowed{}
}

/*GetStatementsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetStatementsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStatementsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsMethodNotAllowed ", 405)
}

func (o *GetStatementsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetStatementsNotAcceptable creates a GetStatementsNotAcceptable with default headers values
func NewGetStatementsNotAcceptable() *GetStatementsNotAcceptable {
	return &GetStatementsNotAcceptable{}
}

/*GetStatementsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetStatementsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStatementsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsNotAcceptable ", 406)
}

func (o *GetStatementsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetStatementsTooManyRequests creates a GetStatementsTooManyRequests with default headers values
func NewGetStatementsTooManyRequests() *GetStatementsTooManyRequests {
	return &GetStatementsTooManyRequests{}
}

/*GetStatementsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetStatementsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetStatementsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsTooManyRequests ", 429)
}

func (o *GetStatementsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetStatementsInternalServerError creates a GetStatementsInternalServerError with default headers values
func NewGetStatementsInternalServerError() *GetStatementsInternalServerError {
	return &GetStatementsInternalServerError{}
}

/*GetStatementsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStatementsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetStatementsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /statements][%d] getStatementsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStatementsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetStatementsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

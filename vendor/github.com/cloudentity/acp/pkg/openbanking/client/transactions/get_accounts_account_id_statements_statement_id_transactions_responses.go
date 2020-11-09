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

// GetAccountsAccountIDStatementsStatementIDTransactionsReader is a Reader for the GetAccountsAccountIDStatementsStatementIDTransactions structure.
type GetAccountsAccountIDStatementsStatementIDTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsOK creates a GetAccountsAccountIDStatementsStatementIDTransactionsOK with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsOK() *GetAccountsAccountIDStatementsStatementIDTransactionsOK {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsOK{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsOK handles this case with default header values.

Transactions Read
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadTransaction6
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsOK) GetPayload() *models.OBReadTransaction6 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadTransaction6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsBadRequest creates a GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsBadRequest() *GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest handles this case with default header values.

Bad request
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized creates a GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized() *GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsUnauthorized ", 401)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsForbidden creates a GetAccountsAccountIDStatementsStatementIDTransactionsForbidden with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsForbidden() *GetAccountsAccountIDStatementsStatementIDTransactionsForbidden {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsForbidden{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsNotFound creates a GetAccountsAccountIDStatementsStatementIDTransactionsNotFound with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsNotFound() *GetAccountsAccountIDStatementsStatementIDTransactionsNotFound {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsNotFound{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsNotFound handles this case with default header values.

Not found
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsNotFound ", 404)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed creates a GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed() *GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable creates a GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable() *GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests creates a GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests() *GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError creates a GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError() *GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError{}
}

/*GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}/transactions][%d] getAccountsAccountIdStatementsStatementIdTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

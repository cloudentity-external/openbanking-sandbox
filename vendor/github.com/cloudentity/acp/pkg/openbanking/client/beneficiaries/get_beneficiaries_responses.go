// Code generated by go-swagger; DO NOT EDIT.

package beneficiaries

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

// GetBeneficiariesReader is a Reader for the GetBeneficiaries structure.
type GetBeneficiariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBeneficiariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBeneficiariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBeneficiariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBeneficiariesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBeneficiariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBeneficiariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetBeneficiariesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetBeneficiariesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBeneficiariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBeneficiariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBeneficiariesOK creates a GetBeneficiariesOK with default headers values
func NewGetBeneficiariesOK() *GetBeneficiariesOK {
	return &GetBeneficiariesOK{}
}

/*GetBeneficiariesOK handles this case with default header values.

Beneficiaries Read
*/
type GetBeneficiariesOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadBeneficiary5
}

func (o *GetBeneficiariesOK) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesOK  %+v", 200, o.Payload)
}

func (o *GetBeneficiariesOK) GetPayload() *models.OBReadBeneficiary5 {
	return o.Payload
}

func (o *GetBeneficiariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadBeneficiary5)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBeneficiariesBadRequest creates a GetBeneficiariesBadRequest with default headers values
func NewGetBeneficiariesBadRequest() *GetBeneficiariesBadRequest {
	return &GetBeneficiariesBadRequest{}
}

/*GetBeneficiariesBadRequest handles this case with default header values.

Bad request
*/
type GetBeneficiariesBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBeneficiariesBadRequest) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetBeneficiariesBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBeneficiariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBeneficiariesUnauthorized creates a GetBeneficiariesUnauthorized with default headers values
func NewGetBeneficiariesUnauthorized() *GetBeneficiariesUnauthorized {
	return &GetBeneficiariesUnauthorized{}
}

/*GetBeneficiariesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBeneficiariesUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBeneficiariesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesUnauthorized ", 401)
}

func (o *GetBeneficiariesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetBeneficiariesForbidden creates a GetBeneficiariesForbidden with default headers values
func NewGetBeneficiariesForbidden() *GetBeneficiariesForbidden {
	return &GetBeneficiariesForbidden{}
}

/*GetBeneficiariesForbidden handles this case with default header values.

Forbidden
*/
type GetBeneficiariesForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBeneficiariesForbidden) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesForbidden  %+v", 403, o.Payload)
}

func (o *GetBeneficiariesForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBeneficiariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBeneficiariesNotFound creates a GetBeneficiariesNotFound with default headers values
func NewGetBeneficiariesNotFound() *GetBeneficiariesNotFound {
	return &GetBeneficiariesNotFound{}
}

/*GetBeneficiariesNotFound handles this case with default header values.

Not found
*/
type GetBeneficiariesNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBeneficiariesNotFound) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesNotFound ", 404)
}

func (o *GetBeneficiariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetBeneficiariesMethodNotAllowed creates a GetBeneficiariesMethodNotAllowed with default headers values
func NewGetBeneficiariesMethodNotAllowed() *GetBeneficiariesMethodNotAllowed {
	return &GetBeneficiariesMethodNotAllowed{}
}

/*GetBeneficiariesMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetBeneficiariesMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBeneficiariesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesMethodNotAllowed ", 405)
}

func (o *GetBeneficiariesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetBeneficiariesNotAcceptable creates a GetBeneficiariesNotAcceptable with default headers values
func NewGetBeneficiariesNotAcceptable() *GetBeneficiariesNotAcceptable {
	return &GetBeneficiariesNotAcceptable{}
}

/*GetBeneficiariesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetBeneficiariesNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBeneficiariesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesNotAcceptable ", 406)
}

func (o *GetBeneficiariesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetBeneficiariesTooManyRequests creates a GetBeneficiariesTooManyRequests with default headers values
func NewGetBeneficiariesTooManyRequests() *GetBeneficiariesTooManyRequests {
	return &GetBeneficiariesTooManyRequests{}
}

/*GetBeneficiariesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetBeneficiariesTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetBeneficiariesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesTooManyRequests ", 429)
}

func (o *GetBeneficiariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBeneficiariesInternalServerError creates a GetBeneficiariesInternalServerError with default headers values
func NewGetBeneficiariesInternalServerError() *GetBeneficiariesInternalServerError {
	return &GetBeneficiariesInternalServerError{}
}

/*GetBeneficiariesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetBeneficiariesInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetBeneficiariesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /beneficiaries][%d] getBeneficiariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBeneficiariesInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetBeneficiariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

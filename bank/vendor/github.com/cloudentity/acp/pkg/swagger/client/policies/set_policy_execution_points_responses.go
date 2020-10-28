// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// SetPolicyExecutionPointsReader is a Reader for the SetPolicyExecutionPoints structure.
type SetPolicyExecutionPointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPolicyExecutionPointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetPolicyExecutionPointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetPolicyExecutionPointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetPolicyExecutionPointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetPolicyExecutionPointsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetPolicyExecutionPointsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetPolicyExecutionPointsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSetPolicyExecutionPointsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetPolicyExecutionPointsOK creates a SetPolicyExecutionPointsOK with default headers values
func NewSetPolicyExecutionPointsOK() *SetPolicyExecutionPointsOK {
	return &SetPolicyExecutionPointsOK{}
}

/*SetPolicyExecutionPointsOK handles this case with default header values.

PolicyExecutionPoints
*/
type SetPolicyExecutionPointsOK struct {
	Payload *models.PolicyExecutionPoints
}

func (o *SetPolicyExecutionPointsOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsOK  %+v", 200, o.Payload)
}

func (o *SetPolicyExecutionPointsOK) GetPayload() *models.PolicyExecutionPoints {
	return o.Payload
}

func (o *SetPolicyExecutionPointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyExecutionPoints)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsBadRequest creates a SetPolicyExecutionPointsBadRequest with default headers values
func NewSetPolicyExecutionPointsBadRequest() *SetPolicyExecutionPointsBadRequest {
	return &SetPolicyExecutionPointsBadRequest{}
}

/*SetPolicyExecutionPointsBadRequest handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsBadRequest struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsBadRequest  %+v", 400, o.Payload)
}

func (o *SetPolicyExecutionPointsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsUnauthorized creates a SetPolicyExecutionPointsUnauthorized with default headers values
func NewSetPolicyExecutionPointsUnauthorized() *SetPolicyExecutionPointsUnauthorized {
	return &SetPolicyExecutionPointsUnauthorized{}
}

/*SetPolicyExecutionPointsUnauthorized handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsUnauthorized struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsUnauthorized  %+v", 401, o.Payload)
}

func (o *SetPolicyExecutionPointsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsForbidden creates a SetPolicyExecutionPointsForbidden with default headers values
func NewSetPolicyExecutionPointsForbidden() *SetPolicyExecutionPointsForbidden {
	return &SetPolicyExecutionPointsForbidden{}
}

/*SetPolicyExecutionPointsForbidden handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsForbidden struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsForbidden  %+v", 403, o.Payload)
}

func (o *SetPolicyExecutionPointsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsNotFound creates a SetPolicyExecutionPointsNotFound with default headers values
func NewSetPolicyExecutionPointsNotFound() *SetPolicyExecutionPointsNotFound {
	return &SetPolicyExecutionPointsNotFound{}
}

/*SetPolicyExecutionPointsNotFound handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsNotFound struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsNotFound  %+v", 404, o.Payload)
}

func (o *SetPolicyExecutionPointsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsConflict creates a SetPolicyExecutionPointsConflict with default headers values
func NewSetPolicyExecutionPointsConflict() *SetPolicyExecutionPointsConflict {
	return &SetPolicyExecutionPointsConflict{}
}

/*SetPolicyExecutionPointsConflict handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsConflict struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsConflict) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsConflict  %+v", 409, o.Payload)
}

func (o *SetPolicyExecutionPointsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPolicyExecutionPointsUnprocessableEntity creates a SetPolicyExecutionPointsUnprocessableEntity with default headers values
func NewSetPolicyExecutionPointsUnprocessableEntity() *SetPolicyExecutionPointsUnprocessableEntity {
	return &SetPolicyExecutionPointsUnprocessableEntity{}
}

/*SetPolicyExecutionPointsUnprocessableEntity handles this case with default header values.

HttpError
*/
type SetPolicyExecutionPointsUnprocessableEntity struct {
	Payload *models.Error
}

func (o *SetPolicyExecutionPointsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/policy-execution-points][%d] setPolicyExecutionPointsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SetPolicyExecutionPointsUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPolicyExecutionPointsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetPolicyReader is a Reader for the GetPolicy structure.
type GetPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPolicyOK creates a GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {
	return &GetPolicyOK{}
}

/*GetPolicyOK handles this case with default header values.

Policy
*/
type GetPolicyOK struct {
	Payload *models.Policy
}

func (o *GetPolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/policies/{pid}][%d] getPolicyOK  %+v", 200, o.Payload)
}

func (o *GetPolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *GetPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyUnauthorized creates a GetPolicyUnauthorized with default headers values
func NewGetPolicyUnauthorized() *GetPolicyUnauthorized {
	return &GetPolicyUnauthorized{}
}

/*GetPolicyUnauthorized handles this case with default header values.

HttpError
*/
type GetPolicyUnauthorized struct {
	Payload *models.Error
}

func (o *GetPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/policies/{pid}][%d] getPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPolicyUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyForbidden creates a GetPolicyForbidden with default headers values
func NewGetPolicyForbidden() *GetPolicyForbidden {
	return &GetPolicyForbidden{}
}

/*GetPolicyForbidden handles this case with default header values.

HttpError
*/
type GetPolicyForbidden struct {
	Payload *models.Error
}

func (o *GetPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/policies/{pid}][%d] getPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetPolicyForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyNotFound creates a GetPolicyNotFound with default headers values
func NewGetPolicyNotFound() *GetPolicyNotFound {
	return &GetPolicyNotFound{}
}

/*GetPolicyNotFound handles this case with default header values.

HttpError
*/
type GetPolicyNotFound struct {
	Payload *models.Error
}

func (o *GetPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/policies/{pid}][%d] getPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

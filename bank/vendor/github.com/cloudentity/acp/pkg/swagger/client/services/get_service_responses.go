// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

ServiceWithScopesResponse
*/
type GetServiceOK struct {
	Payload *models.ServiceWithScopesResponse
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/services/{sid}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) GetPayload() *models.ServiceWithScopesResponse {
	return o.Payload
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceWithScopesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceUnauthorized creates a GetServiceUnauthorized with default headers values
func NewGetServiceUnauthorized() *GetServiceUnauthorized {
	return &GetServiceUnauthorized{}
}

/*GetServiceUnauthorized handles this case with default header values.

HttpError
*/
type GetServiceUnauthorized struct {
	Payload *models.Error
}

func (o *GetServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/services/{sid}][%d] getServiceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetServiceUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceForbidden creates a GetServiceForbidden with default headers values
func NewGetServiceForbidden() *GetServiceForbidden {
	return &GetServiceForbidden{}
}

/*GetServiceForbidden handles this case with default header values.

HttpError
*/
type GetServiceForbidden struct {
	Payload *models.Error
}

func (o *GetServiceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/services/{sid}][%d] getServiceForbidden  %+v", 403, o.Payload)
}

func (o *GetServiceForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceNotFound creates a GetServiceNotFound with default headers values
func NewGetServiceNotFound() *GetServiceNotFound {
	return &GetServiceNotFound{}
}

/*GetServiceNotFound handles this case with default header values.

HttpError
*/
type GetServiceNotFound struct {
	Payload *models.Error
}

func (o *GetServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/services/{sid}][%d] getServiceNotFound  %+v", 404, o.Payload)
}

func (o *GetServiceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

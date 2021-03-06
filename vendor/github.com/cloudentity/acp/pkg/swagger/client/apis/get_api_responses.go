// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetAPIReader is a Reader for the GetAPI structure.
type GetAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIOK creates a GetAPIOK with default headers values
func NewGetAPIOK() *GetAPIOK {
	return &GetAPIOK{}
}

/*GetAPIOK handles this case with default header values.

API
*/
type GetAPIOK struct {
	Payload *models.API
}

func (o *GetAPIOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/apis/{api}][%d] getApiOK  %+v", 200, o.Payload)
}

func (o *GetAPIOK) GetPayload() *models.API {
	return o.Payload
}

func (o *GetAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.API)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIUnauthorized creates a GetAPIUnauthorized with default headers values
func NewGetAPIUnauthorized() *GetAPIUnauthorized {
	return &GetAPIUnauthorized{}
}

/*GetAPIUnauthorized handles this case with default header values.

HttpError
*/
type GetAPIUnauthorized struct {
	Payload *models.Error
}

func (o *GetAPIUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/apis/{api}][%d] getApiUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAPIUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIForbidden creates a GetAPIForbidden with default headers values
func NewGetAPIForbidden() *GetAPIForbidden {
	return &GetAPIForbidden{}
}

/*GetAPIForbidden handles this case with default header values.

HttpError
*/
type GetAPIForbidden struct {
	Payload *models.Error
}

func (o *GetAPIForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/apis/{api}][%d] getApiForbidden  %+v", 403, o.Payload)
}

func (o *GetAPIForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPINotFound creates a GetAPINotFound with default headers values
func NewGetAPINotFound() *GetAPINotFound {
	return &GetAPINotFound{}
}

/*GetAPINotFound handles this case with default header values.

HttpError
*/
type GetAPINotFound struct {
	Payload *models.Error
}

func (o *GetAPINotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/apis/{api}][%d] getApiNotFound  %+v", 404, o.Payload)
}

func (o *GetAPINotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package idps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GetCustomIDPReader is a Reader for the GetCustomIDP structure.
type GetCustomIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCustomIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomIDPOK creates a GetCustomIDPOK with default headers values
func NewGetCustomIDPOK() *GetCustomIDPOK {
	return &GetCustomIDPOK{}
}

/*GetCustomIDPOK handles this case with default header values.

CustomIDP
*/
type GetCustomIDPOK struct {
	Payload *models.CustomIDP
}

func (o *GetCustomIDPOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/custom/{iid}][%d] getCustomIdPOK  %+v", 200, o.Payload)
}

func (o *GetCustomIDPOK) GetPayload() *models.CustomIDP {
	return o.Payload
}

func (o *GetCustomIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomIDPUnauthorized creates a GetCustomIDPUnauthorized with default headers values
func NewGetCustomIDPUnauthorized() *GetCustomIDPUnauthorized {
	return &GetCustomIDPUnauthorized{}
}

/*GetCustomIDPUnauthorized handles this case with default header values.

HttpError
*/
type GetCustomIDPUnauthorized struct {
	Payload *models.Error
}

func (o *GetCustomIDPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/custom/{iid}][%d] getCustomIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCustomIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCustomIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomIDPForbidden creates a GetCustomIDPForbidden with default headers values
func NewGetCustomIDPForbidden() *GetCustomIDPForbidden {
	return &GetCustomIDPForbidden{}
}

/*GetCustomIDPForbidden handles this case with default header values.

HttpError
*/
type GetCustomIDPForbidden struct {
	Payload *models.Error
}

func (o *GetCustomIDPForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/custom/{iid}][%d] getCustomIdPForbidden  %+v", 403, o.Payload)
}

func (o *GetCustomIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCustomIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomIDPNotFound creates a GetCustomIDPNotFound with default headers values
func NewGetCustomIDPNotFound() *GetCustomIDPNotFound {
	return &GetCustomIDPNotFound{}
}

/*GetCustomIDPNotFound handles this case with default header values.

HttpError
*/
type GetCustomIDPNotFound struct {
	Payload *models.Error
}

func (o *GetCustomIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/custom/{iid}][%d] getCustomIdPNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCustomIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
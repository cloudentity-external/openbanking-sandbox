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

// GetSAMLIDPReader is a Reader for the GetSAMLIDP structure.
type GetSAMLIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSAMLIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSAMLIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSAMLIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSAMLIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSAMLIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSAMLIDPOK creates a GetSAMLIDPOK with default headers values
func NewGetSAMLIDPOK() *GetSAMLIDPOK {
	return &GetSAMLIDPOK{}
}

/*GetSAMLIDPOK handles this case with default header values.

SAMLIDP
*/
type GetSAMLIDPOK struct {
	Payload *models.SAMLIDP
}

func (o *GetSAMLIDPOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/saml/{iid}][%d] getSAMLIdPOK  %+v", 200, o.Payload)
}

func (o *GetSAMLIDPOK) GetPayload() *models.SAMLIDP {
	return o.Payload
}

func (o *GetSAMLIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAMLIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSAMLIDPUnauthorized creates a GetSAMLIDPUnauthorized with default headers values
func NewGetSAMLIDPUnauthorized() *GetSAMLIDPUnauthorized {
	return &GetSAMLIDPUnauthorized{}
}

/*GetSAMLIDPUnauthorized handles this case with default header values.

HttpError
*/
type GetSAMLIDPUnauthorized struct {
	Payload *models.Error
}

func (o *GetSAMLIDPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/saml/{iid}][%d] getSAMLIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSAMLIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSAMLIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSAMLIDPForbidden creates a GetSAMLIDPForbidden with default headers values
func NewGetSAMLIDPForbidden() *GetSAMLIDPForbidden {
	return &GetSAMLIDPForbidden{}
}

/*GetSAMLIDPForbidden handles this case with default header values.

HttpError
*/
type GetSAMLIDPForbidden struct {
	Payload *models.Error
}

func (o *GetSAMLIDPForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/saml/{iid}][%d] getSAMLIdPForbidden  %+v", 403, o.Payload)
}

func (o *GetSAMLIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSAMLIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSAMLIDPNotFound creates a GetSAMLIDPNotFound with default headers values
func NewGetSAMLIDPNotFound() *GetSAMLIDPNotFound {
	return &GetSAMLIDPNotFound{}
}

/*GetSAMLIDPNotFound handles this case with default header values.

HttpError
*/
type GetSAMLIDPNotFound struct {
	Payload *models.Error
}

func (o *GetSAMLIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/saml/{iid}][%d] getSAMLIdPNotFound  %+v", 404, o.Payload)
}

func (o *GetSAMLIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSAMLIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

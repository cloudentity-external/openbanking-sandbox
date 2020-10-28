// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// GlobalIntrospectReader is a Reader for the GlobalIntrospect structure.
type GlobalIntrospectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GlobalIntrospectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGlobalIntrospectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGlobalIntrospectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGlobalIntrospectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGlobalIntrospectOK creates a GlobalIntrospectOK with default headers values
func NewGlobalIntrospectOK() *GlobalIntrospectOK {
	return &GlobalIntrospectOK{}
}

/*GlobalIntrospectOK handles this case with default header values.

IntrospectResponse
*/
type GlobalIntrospectOK struct {
	Payload *models.IntrospectResponse
}

func (o *GlobalIntrospectOK) Error() string {
	return fmt.Sprintf("[POST /api/system/introspect][%d] globalIntrospectOK  %+v", 200, o.Payload)
}

func (o *GlobalIntrospectOK) GetPayload() *models.IntrospectResponse {
	return o.Payload
}

func (o *GlobalIntrospectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntrospectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalIntrospectBadRequest creates a GlobalIntrospectBadRequest with default headers values
func NewGlobalIntrospectBadRequest() *GlobalIntrospectBadRequest {
	return &GlobalIntrospectBadRequest{}
}

/*GlobalIntrospectBadRequest handles this case with default header values.

HttpError
*/
type GlobalIntrospectBadRequest struct {
	Payload *models.Error
}

func (o *GlobalIntrospectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/system/introspect][%d] globalIntrospectBadRequest  %+v", 400, o.Payload)
}

func (o *GlobalIntrospectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GlobalIntrospectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGlobalIntrospectUnprocessableEntity creates a GlobalIntrospectUnprocessableEntity with default headers values
func NewGlobalIntrospectUnprocessableEntity() *GlobalIntrospectUnprocessableEntity {
	return &GlobalIntrospectUnprocessableEntity{}
}

/*GlobalIntrospectUnprocessableEntity handles this case with default header values.

HttpError
*/
type GlobalIntrospectUnprocessableEntity struct {
	Payload *models.Error
}

func (o *GlobalIntrospectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/system/introspect][%d] globalIntrospectUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GlobalIntrospectUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GlobalIntrospectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

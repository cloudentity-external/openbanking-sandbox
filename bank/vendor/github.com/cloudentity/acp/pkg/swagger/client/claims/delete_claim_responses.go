// Code generated by go-swagger; DO NOT EDIT.

package claims

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// DeleteClaimReader is a Reader for the DeleteClaim structure.
type DeleteClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteClaimNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClaimUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClaimForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClaimNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClaimNoContent creates a DeleteClaimNoContent with default headers values
func NewDeleteClaimNoContent() *DeleteClaimNoContent {
	return &DeleteClaimNoContent{}
}

/*DeleteClaimNoContent handles this case with default header values.

Claim has been deleted
*/
type DeleteClaimNoContent struct {
}

func (o *DeleteClaimNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/claims/{claim}][%d] deleteClaimNoContent ", 204)
}

func (o *DeleteClaimNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClaimUnauthorized creates a DeleteClaimUnauthorized with default headers values
func NewDeleteClaimUnauthorized() *DeleteClaimUnauthorized {
	return &DeleteClaimUnauthorized{}
}

/*DeleteClaimUnauthorized handles this case with default header values.

HttpError
*/
type DeleteClaimUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteClaimUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/claims/{claim}][%d] deleteClaimUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClaimUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClaimUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClaimForbidden creates a DeleteClaimForbidden with default headers values
func NewDeleteClaimForbidden() *DeleteClaimForbidden {
	return &DeleteClaimForbidden{}
}

/*DeleteClaimForbidden handles this case with default header values.

HttpError
*/
type DeleteClaimForbidden struct {
	Payload *models.Error
}

func (o *DeleteClaimForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/claims/{claim}][%d] deleteClaimForbidden  %+v", 403, o.Payload)
}

func (o *DeleteClaimForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClaimForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClaimNotFound creates a DeleteClaimNotFound with default headers values
func NewDeleteClaimNotFound() *DeleteClaimNotFound {
	return &DeleteClaimNotFound{}
}

/*DeleteClaimNotFound handles this case with default header values.

HttpError
*/
type DeleteClaimNotFound struct {
	Payload *models.Error
}

func (o *DeleteClaimNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/claims/{claim}][%d] deleteClaimNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClaimNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClaimNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

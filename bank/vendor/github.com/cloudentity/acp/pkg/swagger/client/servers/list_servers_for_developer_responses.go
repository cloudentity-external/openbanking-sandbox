// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// ListServersForDeveloperReader is a Reader for the ListServersForDeveloper structure.
type ListServersForDeveloperReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServersForDeveloperReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServersForDeveloperOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServersForDeveloperUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServersForDeveloperForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListServersForDeveloperNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListServersForDeveloperOK creates a ListServersForDeveloperOK with default headers values
func NewListServersForDeveloperOK() *ListServersForDeveloperOK {
	return &ListServersForDeveloperOK{}
}

/*ListServersForDeveloperOK handles this case with default header values.

ListServersDeveloperResponse
*/
type ListServersForDeveloperOK struct {
	Payload *models.ListServersDeveloperResponse
}

func (o *ListServersForDeveloperOK) Error() string {
	return fmt.Sprintf("[GET /api/developer/{tid}/{aid}/servers][%d] listServersForDeveloperOK  %+v", 200, o.Payload)
}

func (o *ListServersForDeveloperOK) GetPayload() *models.ListServersDeveloperResponse {
	return o.Payload
}

func (o *ListServersForDeveloperOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListServersDeveloperResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersForDeveloperUnauthorized creates a ListServersForDeveloperUnauthorized with default headers values
func NewListServersForDeveloperUnauthorized() *ListServersForDeveloperUnauthorized {
	return &ListServersForDeveloperUnauthorized{}
}

/*ListServersForDeveloperUnauthorized handles this case with default header values.

HttpError
*/
type ListServersForDeveloperUnauthorized struct {
	Payload *models.Error
}

func (o *ListServersForDeveloperUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/developer/{tid}/{aid}/servers][%d] listServersForDeveloperUnauthorized  %+v", 401, o.Payload)
}

func (o *ListServersForDeveloperUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersForDeveloperUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersForDeveloperForbidden creates a ListServersForDeveloperForbidden with default headers values
func NewListServersForDeveloperForbidden() *ListServersForDeveloperForbidden {
	return &ListServersForDeveloperForbidden{}
}

/*ListServersForDeveloperForbidden handles this case with default header values.

HttpError
*/
type ListServersForDeveloperForbidden struct {
	Payload *models.Error
}

func (o *ListServersForDeveloperForbidden) Error() string {
	return fmt.Sprintf("[GET /api/developer/{tid}/{aid}/servers][%d] listServersForDeveloperForbidden  %+v", 403, o.Payload)
}

func (o *ListServersForDeveloperForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersForDeveloperForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServersForDeveloperNotFound creates a ListServersForDeveloperNotFound with default headers values
func NewListServersForDeveloperNotFound() *ListServersForDeveloperNotFound {
	return &ListServersForDeveloperNotFound{}
}

/*ListServersForDeveloperNotFound handles this case with default header values.

HttpError
*/
type ListServersForDeveloperNotFound struct {
	Payload *models.Error
}

func (o *ListServersForDeveloperNotFound) Error() string {
	return fmt.Sprintf("[GET /api/developer/{tid}/{aid}/servers][%d] listServersForDeveloperNotFound  %+v", 404, o.Payload)
}

func (o *ListServersForDeveloperNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListServersForDeveloperNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

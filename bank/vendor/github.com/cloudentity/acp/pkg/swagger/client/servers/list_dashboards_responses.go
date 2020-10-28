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

// ListDashboardsReader is a Reader for the ListDashboards structure.
type ListDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListDashboardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListDashboardsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListDashboardsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDashboardsOK creates a ListDashboardsOK with default headers values
func NewListDashboardsOK() *ListDashboardsOK {
	return &ListDashboardsOK{}
}

/*ListDashboardsOK handles this case with default header values.

Dashboards
*/
type ListDashboardsOK struct {
	Payload *models.Dashboards
}

func (o *ListDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/dashboards][%d] listDashboardsOK  %+v", 200, o.Payload)
}

func (o *ListDashboardsOK) GetPayload() *models.Dashboards {
	return o.Payload
}

func (o *ListDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Dashboards)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardsBadRequest creates a ListDashboardsBadRequest with default headers values
func NewListDashboardsBadRequest() *ListDashboardsBadRequest {
	return &ListDashboardsBadRequest{}
}

/*ListDashboardsBadRequest handles this case with default header values.

HttpError
*/
type ListDashboardsBadRequest struct {
	Payload *models.Error
}

func (o *ListDashboardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/dashboards][%d] listDashboardsBadRequest  %+v", 400, o.Payload)
}

func (o *ListDashboardsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDashboardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardsUnauthorized creates a ListDashboardsUnauthorized with default headers values
func NewListDashboardsUnauthorized() *ListDashboardsUnauthorized {
	return &ListDashboardsUnauthorized{}
}

/*ListDashboardsUnauthorized handles this case with default header values.

HttpError
*/
type ListDashboardsUnauthorized struct {
	Payload *models.Error
}

func (o *ListDashboardsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/dashboards][%d] listDashboardsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListDashboardsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDashboardsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardsForbidden creates a ListDashboardsForbidden with default headers values
func NewListDashboardsForbidden() *ListDashboardsForbidden {
	return &ListDashboardsForbidden{}
}

/*ListDashboardsForbidden handles this case with default header values.

HttpError
*/
type ListDashboardsForbidden struct {
	Payload *models.Error
}

func (o *ListDashboardsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/dashboards][%d] listDashboardsForbidden  %+v", 403, o.Payload)
}

func (o *ListDashboardsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDashboardsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardsNotFound creates a ListDashboardsNotFound with default headers values
func NewListDashboardsNotFound() *ListDashboardsNotFound {
	return &ListDashboardsNotFound{}
}

/*ListDashboardsNotFound handles this case with default header values.

HttpError
*/
type ListDashboardsNotFound struct {
	Payload *models.Error
}

func (o *ListDashboardsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/dashboards][%d] listDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *ListDashboardsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

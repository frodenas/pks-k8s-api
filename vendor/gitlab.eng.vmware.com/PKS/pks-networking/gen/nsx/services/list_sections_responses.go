// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// ListSectionsReader is a Reader for the ListSections structure.
type ListSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListSectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListSectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewListSectionsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListSectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewListSectionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSectionsOK creates a ListSectionsOK with default headers values
func NewListSectionsOK() *ListSectionsOK {
	return &ListSectionsOK{}
}

/*ListSectionsOK handles this case with default header values.

Success
*/
type ListSectionsOK struct {
	Payload *models.FirewallSectionListResult
}

func (o *ListSectionsOK) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsOK  %+v", 200, o.Payload)
}

func (o *ListSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FirewallSectionListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsBadRequest creates a ListSectionsBadRequest with default headers values
func NewListSectionsBadRequest() *ListSectionsBadRequest {
	return &ListSectionsBadRequest{}
}

/*ListSectionsBadRequest handles this case with default header values.

Bad request
*/
type ListSectionsBadRequest struct {
	Payload *models.APIError
}

func (o *ListSectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListSectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsForbidden creates a ListSectionsForbidden with default headers values
func NewListSectionsForbidden() *ListSectionsForbidden {
	return &ListSectionsForbidden{}
}

/*ListSectionsForbidden handles this case with default header values.

Operation forbidden
*/
type ListSectionsForbidden struct {
	Payload *models.APIError
}

func (o *ListSectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsForbidden  %+v", 403, o.Payload)
}

func (o *ListSectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsNotFound creates a ListSectionsNotFound with default headers values
func NewListSectionsNotFound() *ListSectionsNotFound {
	return &ListSectionsNotFound{}
}

/*ListSectionsNotFound handles this case with default header values.

Resource not found
*/
type ListSectionsNotFound struct {
	Payload *models.APIError
}

func (o *ListSectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsNotFound  %+v", 404, o.Payload)
}

func (o *ListSectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsPreconditionFailed creates a ListSectionsPreconditionFailed with default headers values
func NewListSectionsPreconditionFailed() *ListSectionsPreconditionFailed {
	return &ListSectionsPreconditionFailed{}
}

/*ListSectionsPreconditionFailed handles this case with default header values.

Precondition failed
*/
type ListSectionsPreconditionFailed struct {
	Payload *models.APIError
}

func (o *ListSectionsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ListSectionsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsInternalServerError creates a ListSectionsInternalServerError with default headers values
func NewListSectionsInternalServerError() *ListSectionsInternalServerError {
	return &ListSectionsInternalServerError{}
}

/*ListSectionsInternalServerError handles this case with default header values.

Internal server error
*/
type ListSectionsInternalServerError struct {
	Payload *models.APIError
}

func (o *ListSectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSectionsServiceUnavailable creates a ListSectionsServiceUnavailable with default headers values
func NewListSectionsServiceUnavailable() *ListSectionsServiceUnavailable {
	return &ListSectionsServiceUnavailable{}
}

/*ListSectionsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type ListSectionsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *ListSectionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /firewall/sections][%d] listSectionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListSectionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

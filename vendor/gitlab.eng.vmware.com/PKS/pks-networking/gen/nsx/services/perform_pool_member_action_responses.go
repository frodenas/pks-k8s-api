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

// PerformPoolMemberActionReader is a Reader for the PerformPoolMemberAction structure.
type PerformPoolMemberActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformPoolMemberActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPerformPoolMemberActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPerformPoolMemberActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPerformPoolMemberActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPerformPoolMemberActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPerformPoolMemberActionPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPerformPoolMemberActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPerformPoolMemberActionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPerformPoolMemberActionOK creates a PerformPoolMemberActionOK with default headers values
func NewPerformPoolMemberActionOK() *PerformPoolMemberActionOK {
	return &PerformPoolMemberActionOK{}
}

/*PerformPoolMemberActionOK handles this case with default header values.

OK
*/
type PerformPoolMemberActionOK struct {
	Payload *models.LbPool
}

func (o *PerformPoolMemberActionOK) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionOK  %+v", 200, o.Payload)
}

func (o *PerformPoolMemberActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LbPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionBadRequest creates a PerformPoolMemberActionBadRequest with default headers values
func NewPerformPoolMemberActionBadRequest() *PerformPoolMemberActionBadRequest {
	return &PerformPoolMemberActionBadRequest{}
}

/*PerformPoolMemberActionBadRequest handles this case with default header values.

Bad request
*/
type PerformPoolMemberActionBadRequest struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformPoolMemberActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionForbidden creates a PerformPoolMemberActionForbidden with default headers values
func NewPerformPoolMemberActionForbidden() *PerformPoolMemberActionForbidden {
	return &PerformPoolMemberActionForbidden{}
}

/*PerformPoolMemberActionForbidden handles this case with default header values.

Operation forbidden
*/
type PerformPoolMemberActionForbidden struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionForbidden) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformPoolMemberActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionNotFound creates a PerformPoolMemberActionNotFound with default headers values
func NewPerformPoolMemberActionNotFound() *PerformPoolMemberActionNotFound {
	return &PerformPoolMemberActionNotFound{}
}

/*PerformPoolMemberActionNotFound handles this case with default header values.

Resource not found
*/
type PerformPoolMemberActionNotFound struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionNotFound) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformPoolMemberActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionPreconditionFailed creates a PerformPoolMemberActionPreconditionFailed with default headers values
func NewPerformPoolMemberActionPreconditionFailed() *PerformPoolMemberActionPreconditionFailed {
	return &PerformPoolMemberActionPreconditionFailed{}
}

/*PerformPoolMemberActionPreconditionFailed handles this case with default header values.

Precondition failed
*/
type PerformPoolMemberActionPreconditionFailed struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PerformPoolMemberActionPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionInternalServerError creates a PerformPoolMemberActionInternalServerError with default headers values
func NewPerformPoolMemberActionInternalServerError() *PerformPoolMemberActionInternalServerError {
	return &PerformPoolMemberActionInternalServerError{}
}

/*PerformPoolMemberActionInternalServerError handles this case with default header values.

Internal server error
*/
type PerformPoolMemberActionInternalServerError struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformPoolMemberActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformPoolMemberActionServiceUnavailable creates a PerformPoolMemberActionServiceUnavailable with default headers values
func NewPerformPoolMemberActionServiceUnavailable() *PerformPoolMemberActionServiceUnavailable {
	return &PerformPoolMemberActionServiceUnavailable{}
}

/*PerformPoolMemberActionServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PerformPoolMemberActionServiceUnavailable struct {
	Payload *models.APIError
}

func (o *PerformPoolMemberActionServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /loadbalancer/pools/{pool-id}][%d] performPoolMemberActionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PerformPoolMemberActionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package pool_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// DeleteIPBlockSubnetReader is a Reader for the DeleteIPBlockSubnet structure.
type DeleteIPBlockSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPBlockSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteIPBlockSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteIPBlockSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteIPBlockSubnetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteIPBlockSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIPBlockSubnetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteIPBlockSubnetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteIPBlockSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteIPBlockSubnetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIPBlockSubnetOK creates a DeleteIPBlockSubnetOK with default headers values
func NewDeleteIPBlockSubnetOK() *DeleteIPBlockSubnetOK {
	return &DeleteIPBlockSubnetOK{}
}

/*DeleteIPBlockSubnetOK handles this case with default header values.

Ip Block subnet successfully deleted
*/
type DeleteIPBlockSubnetOK struct {
}

func (o *DeleteIPBlockSubnetOK) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetOK ", 200)
}

func (o *DeleteIPBlockSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPBlockSubnetBadRequest creates a DeleteIPBlockSubnetBadRequest with default headers values
func NewDeleteIPBlockSubnetBadRequest() *DeleteIPBlockSubnetBadRequest {
	return &DeleteIPBlockSubnetBadRequest{}
}

/*DeleteIPBlockSubnetBadRequest handles this case with default header values.

Bad request
*/
type DeleteIPBlockSubnetBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIPBlockSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetForbidden creates a DeleteIPBlockSubnetForbidden with default headers values
func NewDeleteIPBlockSubnetForbidden() *DeleteIPBlockSubnetForbidden {
	return &DeleteIPBlockSubnetForbidden{}
}

/*DeleteIPBlockSubnetForbidden handles this case with default header values.

Operation forbidden
*/
type DeleteIPBlockSubnetForbidden struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIPBlockSubnetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetNotFound creates a DeleteIPBlockSubnetNotFound with default headers values
func NewDeleteIPBlockSubnetNotFound() *DeleteIPBlockSubnetNotFound {
	return &DeleteIPBlockSubnetNotFound{}
}

/*DeleteIPBlockSubnetNotFound handles this case with default header values.

Resource not found
*/
type DeleteIPBlockSubnetNotFound struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIPBlockSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetConflict creates a DeleteIPBlockSubnetConflict with default headers values
func NewDeleteIPBlockSubnetConflict() *DeleteIPBlockSubnetConflict {
	return &DeleteIPBlockSubnetConflict{}
}

/*DeleteIPBlockSubnetConflict handles this case with default header values.

Resource conflict
*/
type DeleteIPBlockSubnetConflict struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetConflict) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetConflict  %+v", 409, o.Payload)
}

func (o *DeleteIPBlockSubnetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetPreconditionFailed creates a DeleteIPBlockSubnetPreconditionFailed with default headers values
func NewDeleteIPBlockSubnetPreconditionFailed() *DeleteIPBlockSubnetPreconditionFailed {
	return &DeleteIPBlockSubnetPreconditionFailed{}
}

/*DeleteIPBlockSubnetPreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteIPBlockSubnetPreconditionFailed struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteIPBlockSubnetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetInternalServerError creates a DeleteIPBlockSubnetInternalServerError with default headers values
func NewDeleteIPBlockSubnetInternalServerError() *DeleteIPBlockSubnetInternalServerError {
	return &DeleteIPBlockSubnetInternalServerError{}
}

/*DeleteIPBlockSubnetInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteIPBlockSubnetInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPBlockSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPBlockSubnetServiceUnavailable creates a DeleteIPBlockSubnetServiceUnavailable with default headers values
func NewDeleteIPBlockSubnetServiceUnavailable() *DeleteIPBlockSubnetServiceUnavailable {
	return &DeleteIPBlockSubnetServiceUnavailable{}
}

/*DeleteIPBlockSubnetServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteIPBlockSubnetServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteIPBlockSubnetServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /pools/ip-subnets/{subnet-id}][%d] deleteIpBlockSubnetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIPBlockSubnetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

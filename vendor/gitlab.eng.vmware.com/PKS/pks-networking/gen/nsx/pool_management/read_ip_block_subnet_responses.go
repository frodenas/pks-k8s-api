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

// ReadIPBlockSubnetReader is a Reader for the ReadIPBlockSubnet structure.
type ReadIPBlockSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadIPBlockSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReadIPBlockSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReadIPBlockSubnetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReadIPBlockSubnetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReadIPBlockSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewReadIPBlockSubnetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewReadIPBlockSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewReadIPBlockSubnetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadIPBlockSubnetOK creates a ReadIPBlockSubnetOK with default headers values
func NewReadIPBlockSubnetOK() *ReadIPBlockSubnetOK {
	return &ReadIPBlockSubnetOK{}
}

/*ReadIPBlockSubnetOK handles this case with default header values.

Success
*/
type ReadIPBlockSubnetOK struct {
	Payload *models.IPBlockSubnet
}

func (o *ReadIPBlockSubnetOK) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetOK  %+v", 200, o.Payload)
}

func (o *ReadIPBlockSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPBlockSubnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetBadRequest creates a ReadIPBlockSubnetBadRequest with default headers values
func NewReadIPBlockSubnetBadRequest() *ReadIPBlockSubnetBadRequest {
	return &ReadIPBlockSubnetBadRequest{}
}

/*ReadIPBlockSubnetBadRequest handles this case with default header values.

Bad request
*/
type ReadIPBlockSubnetBadRequest struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetBadRequest  %+v", 400, o.Payload)
}

func (o *ReadIPBlockSubnetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetForbidden creates a ReadIPBlockSubnetForbidden with default headers values
func NewReadIPBlockSubnetForbidden() *ReadIPBlockSubnetForbidden {
	return &ReadIPBlockSubnetForbidden{}
}

/*ReadIPBlockSubnetForbidden handles this case with default header values.

Operation forbidden
*/
type ReadIPBlockSubnetForbidden struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetForbidden) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetForbidden  %+v", 403, o.Payload)
}

func (o *ReadIPBlockSubnetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetNotFound creates a ReadIPBlockSubnetNotFound with default headers values
func NewReadIPBlockSubnetNotFound() *ReadIPBlockSubnetNotFound {
	return &ReadIPBlockSubnetNotFound{}
}

/*ReadIPBlockSubnetNotFound handles this case with default header values.

Resource not found
*/
type ReadIPBlockSubnetNotFound struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetNotFound) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetNotFound  %+v", 404, o.Payload)
}

func (o *ReadIPBlockSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetPreconditionFailed creates a ReadIPBlockSubnetPreconditionFailed with default headers values
func NewReadIPBlockSubnetPreconditionFailed() *ReadIPBlockSubnetPreconditionFailed {
	return &ReadIPBlockSubnetPreconditionFailed{}
}

/*ReadIPBlockSubnetPreconditionFailed handles this case with default header values.

Precondition failed
*/
type ReadIPBlockSubnetPreconditionFailed struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ReadIPBlockSubnetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetInternalServerError creates a ReadIPBlockSubnetInternalServerError with default headers values
func NewReadIPBlockSubnetInternalServerError() *ReadIPBlockSubnetInternalServerError {
	return &ReadIPBlockSubnetInternalServerError{}
}

/*ReadIPBlockSubnetInternalServerError handles this case with default header values.

Internal server error
*/
type ReadIPBlockSubnetInternalServerError struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadIPBlockSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIPBlockSubnetServiceUnavailable creates a ReadIPBlockSubnetServiceUnavailable with default headers values
func NewReadIPBlockSubnetServiceUnavailable() *ReadIPBlockSubnetServiceUnavailable {
	return &ReadIPBlockSubnetServiceUnavailable{}
}

/*ReadIPBlockSubnetServiceUnavailable handles this case with default header values.

Service unavailable
*/
type ReadIPBlockSubnetServiceUnavailable struct {
	Payload *models.APIError
}

func (o *ReadIPBlockSubnetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets/{subnet-id}][%d] readIpBlockSubnetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ReadIPBlockSubnetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

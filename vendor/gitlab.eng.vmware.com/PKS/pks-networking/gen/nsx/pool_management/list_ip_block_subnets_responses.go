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

// ListIPBlockSubnetsReader is a Reader for the ListIPBlockSubnets structure.
type ListIPBlockSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIPBlockSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListIPBlockSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListIPBlockSubnetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListIPBlockSubnetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListIPBlockSubnetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewListIPBlockSubnetsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListIPBlockSubnetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewListIPBlockSubnetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIPBlockSubnetsOK creates a ListIPBlockSubnetsOK with default headers values
func NewListIPBlockSubnetsOK() *ListIPBlockSubnetsOK {
	return &ListIPBlockSubnetsOK{}
}

/*ListIPBlockSubnetsOK handles this case with default header values.

Success
*/
type ListIPBlockSubnetsOK struct {
	Payload *models.IPBlockSubnetListResult
}

func (o *ListIPBlockSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsOK  %+v", 200, o.Payload)
}

func (o *ListIPBlockSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPBlockSubnetListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsBadRequest creates a ListIPBlockSubnetsBadRequest with default headers values
func NewListIPBlockSubnetsBadRequest() *ListIPBlockSubnetsBadRequest {
	return &ListIPBlockSubnetsBadRequest{}
}

/*ListIPBlockSubnetsBadRequest handles this case with default header values.

Bad request
*/
type ListIPBlockSubnetsBadRequest struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsBadRequest  %+v", 400, o.Payload)
}

func (o *ListIPBlockSubnetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsForbidden creates a ListIPBlockSubnetsForbidden with default headers values
func NewListIPBlockSubnetsForbidden() *ListIPBlockSubnetsForbidden {
	return &ListIPBlockSubnetsForbidden{}
}

/*ListIPBlockSubnetsForbidden handles this case with default header values.

Operation forbidden
*/
type ListIPBlockSubnetsForbidden struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsForbidden) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsForbidden  %+v", 403, o.Payload)
}

func (o *ListIPBlockSubnetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsNotFound creates a ListIPBlockSubnetsNotFound with default headers values
func NewListIPBlockSubnetsNotFound() *ListIPBlockSubnetsNotFound {
	return &ListIPBlockSubnetsNotFound{}
}

/*ListIPBlockSubnetsNotFound handles this case with default header values.

Resource not found
*/
type ListIPBlockSubnetsNotFound struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsNotFound) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsNotFound  %+v", 404, o.Payload)
}

func (o *ListIPBlockSubnetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsPreconditionFailed creates a ListIPBlockSubnetsPreconditionFailed with default headers values
func NewListIPBlockSubnetsPreconditionFailed() *ListIPBlockSubnetsPreconditionFailed {
	return &ListIPBlockSubnetsPreconditionFailed{}
}

/*ListIPBlockSubnetsPreconditionFailed handles this case with default header values.

Precondition failed
*/
type ListIPBlockSubnetsPreconditionFailed struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ListIPBlockSubnetsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsInternalServerError creates a ListIPBlockSubnetsInternalServerError with default headers values
func NewListIPBlockSubnetsInternalServerError() *ListIPBlockSubnetsInternalServerError {
	return &ListIPBlockSubnetsInternalServerError{}
}

/*ListIPBlockSubnetsInternalServerError handles this case with default header values.

Internal server error
*/
type ListIPBlockSubnetsInternalServerError struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIPBlockSubnetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPBlockSubnetsServiceUnavailable creates a ListIPBlockSubnetsServiceUnavailable with default headers values
func NewListIPBlockSubnetsServiceUnavailable() *ListIPBlockSubnetsServiceUnavailable {
	return &ListIPBlockSubnetsServiceUnavailable{}
}

/*ListIPBlockSubnetsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type ListIPBlockSubnetsServiceUnavailable struct {
	Payload *models.APIError
}

func (o *ListIPBlockSubnetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /pools/ip-subnets][%d] listIpBlockSubnetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListIPBlockSubnetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

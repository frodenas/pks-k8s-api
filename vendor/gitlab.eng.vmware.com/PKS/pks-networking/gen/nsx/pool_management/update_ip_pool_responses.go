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

// UpdateIPPoolReader is a Reader for the UpdateIPPool structure.
type UpdateIPPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIPPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateIPPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateIPPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateIPPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateIPPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateIPPoolPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateIPPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewUpdateIPPoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateIPPoolOK creates a UpdateIPPoolOK with default headers values
func NewUpdateIPPoolOK() *UpdateIPPoolOK {
	return &UpdateIPPoolOK{}
}

/*UpdateIPPoolOK handles this case with default header values.

Success
*/
type UpdateIPPoolOK struct {
	Payload *models.IPPool
}

func (o *UpdateIPPoolOK) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolOK  %+v", 200, o.Payload)
}

func (o *UpdateIPPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolBadRequest creates a UpdateIPPoolBadRequest with default headers values
func NewUpdateIPPoolBadRequest() *UpdateIPPoolBadRequest {
	return &UpdateIPPoolBadRequest{}
}

/*UpdateIPPoolBadRequest handles this case with default header values.

Bad request
*/
type UpdateIPPoolBadRequest struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIPPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolForbidden creates a UpdateIPPoolForbidden with default headers values
func NewUpdateIPPoolForbidden() *UpdateIPPoolForbidden {
	return &UpdateIPPoolForbidden{}
}

/*UpdateIPPoolForbidden handles this case with default header values.

Operation forbidden
*/
type UpdateIPPoolForbidden struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolForbidden) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolForbidden  %+v", 403, o.Payload)
}

func (o *UpdateIPPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolNotFound creates a UpdateIPPoolNotFound with default headers values
func NewUpdateIPPoolNotFound() *UpdateIPPoolNotFound {
	return &UpdateIPPoolNotFound{}
}

/*UpdateIPPoolNotFound handles this case with default header values.

Resource not found
*/
type UpdateIPPoolNotFound struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolNotFound) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIPPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolPreconditionFailed creates a UpdateIPPoolPreconditionFailed with default headers values
func NewUpdateIPPoolPreconditionFailed() *UpdateIPPoolPreconditionFailed {
	return &UpdateIPPoolPreconditionFailed{}
}

/*UpdateIPPoolPreconditionFailed handles this case with default header values.

Precondition failed
*/
type UpdateIPPoolPreconditionFailed struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateIPPoolPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolInternalServerError creates a UpdateIPPoolInternalServerError with default headers values
func NewUpdateIPPoolInternalServerError() *UpdateIPPoolInternalServerError {
	return &UpdateIPPoolInternalServerError{}
}

/*UpdateIPPoolInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateIPPoolInternalServerError struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIPPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPPoolServiceUnavailable creates a UpdateIPPoolServiceUnavailable with default headers values
func NewUpdateIPPoolServiceUnavailable() *UpdateIPPoolServiceUnavailable {
	return &UpdateIPPoolServiceUnavailable{}
}

/*UpdateIPPoolServiceUnavailable handles this case with default header values.

Service unavailable
*/
type UpdateIPPoolServiceUnavailable struct {
	Payload *models.APIError
}

func (o *UpdateIPPoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-pools/{pool-id}][%d] updateIpPoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateIPPoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// AllocateOrReleaseFromIPPoolReader is a Reader for the AllocateOrReleaseFromIPPool structure.
type AllocateOrReleaseFromIPPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateOrReleaseFromIPPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllocateOrReleaseFromIPPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllocateOrReleaseFromIPPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAllocateOrReleaseFromIPPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllocateOrReleaseFromIPPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAllocateOrReleaseFromIPPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewAllocateOrReleaseFromIPPoolPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAllocateOrReleaseFromIPPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewAllocateOrReleaseFromIPPoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllocateOrReleaseFromIPPoolOK creates a AllocateOrReleaseFromIPPoolOK with default headers values
func NewAllocateOrReleaseFromIPPoolOK() *AllocateOrReleaseFromIPPoolOK {
	return &AllocateOrReleaseFromIPPoolOK{}
}

/*AllocateOrReleaseFromIPPoolOK handles this case with default header values.

Success
*/
type AllocateOrReleaseFromIPPoolOK struct {
	Payload *models.AllocationIPAddress
}

func (o *AllocateOrReleaseFromIPPoolOK) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolOK  %+v", 200, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllocationIPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolBadRequest creates a AllocateOrReleaseFromIPPoolBadRequest with default headers values
func NewAllocateOrReleaseFromIPPoolBadRequest() *AllocateOrReleaseFromIPPoolBadRequest {
	return &AllocateOrReleaseFromIPPoolBadRequest{}
}

/*AllocateOrReleaseFromIPPoolBadRequest handles this case with default header values.

Bad request
*/
type AllocateOrReleaseFromIPPoolBadRequest struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolBadRequest) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolBadRequest  %+v", 400, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolForbidden creates a AllocateOrReleaseFromIPPoolForbidden with default headers values
func NewAllocateOrReleaseFromIPPoolForbidden() *AllocateOrReleaseFromIPPoolForbidden {
	return &AllocateOrReleaseFromIPPoolForbidden{}
}

/*AllocateOrReleaseFromIPPoolForbidden handles this case with default header values.

Operation forbidden
*/
type AllocateOrReleaseFromIPPoolForbidden struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolForbidden) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolForbidden  %+v", 403, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolNotFound creates a AllocateOrReleaseFromIPPoolNotFound with default headers values
func NewAllocateOrReleaseFromIPPoolNotFound() *AllocateOrReleaseFromIPPoolNotFound {
	return &AllocateOrReleaseFromIPPoolNotFound{}
}

/*AllocateOrReleaseFromIPPoolNotFound handles this case with default header values.

Resource not found
*/
type AllocateOrReleaseFromIPPoolNotFound struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolNotFound) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolNotFound  %+v", 404, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolConflict creates a AllocateOrReleaseFromIPPoolConflict with default headers values
func NewAllocateOrReleaseFromIPPoolConflict() *AllocateOrReleaseFromIPPoolConflict {
	return &AllocateOrReleaseFromIPPoolConflict{}
}

/*AllocateOrReleaseFromIPPoolConflict handles this case with default header values.

Resource conflict
*/
type AllocateOrReleaseFromIPPoolConflict struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolConflict) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolConflict  %+v", 409, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolPreconditionFailed creates a AllocateOrReleaseFromIPPoolPreconditionFailed with default headers values
func NewAllocateOrReleaseFromIPPoolPreconditionFailed() *AllocateOrReleaseFromIPPoolPreconditionFailed {
	return &AllocateOrReleaseFromIPPoolPreconditionFailed{}
}

/*AllocateOrReleaseFromIPPoolPreconditionFailed handles this case with default header values.

Precondition failed
*/
type AllocateOrReleaseFromIPPoolPreconditionFailed struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolInternalServerError creates a AllocateOrReleaseFromIPPoolInternalServerError with default headers values
func NewAllocateOrReleaseFromIPPoolInternalServerError() *AllocateOrReleaseFromIPPoolInternalServerError {
	return &AllocateOrReleaseFromIPPoolInternalServerError{}
}

/*AllocateOrReleaseFromIPPoolInternalServerError handles this case with default header values.

Internal server error
*/
type AllocateOrReleaseFromIPPoolInternalServerError struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateOrReleaseFromIPPoolServiceUnavailable creates a AllocateOrReleaseFromIPPoolServiceUnavailable with default headers values
func NewAllocateOrReleaseFromIPPoolServiceUnavailable() *AllocateOrReleaseFromIPPoolServiceUnavailable {
	return &AllocateOrReleaseFromIPPoolServiceUnavailable{}
}

/*AllocateOrReleaseFromIPPoolServiceUnavailable handles this case with default header values.

Service unavailable
*/
type AllocateOrReleaseFromIPPoolServiceUnavailable struct {
	Payload *models.APIError
}

func (o *AllocateOrReleaseFromIPPoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /pools/ip-pools/{pool-id}][%d] allocateOrReleaseFromIpPoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *AllocateOrReleaseFromIPPoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

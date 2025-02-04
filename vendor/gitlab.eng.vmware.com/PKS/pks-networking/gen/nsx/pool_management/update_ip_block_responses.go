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

// UpdateIPBlockReader is a Reader for the UpdateIPBlock structure.
type UpdateIPBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIPBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateIPBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateIPBlockBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateIPBlockForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateIPBlockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateIPBlockConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateIPBlockPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateIPBlockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewUpdateIPBlockServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateIPBlockOK creates a UpdateIPBlockOK with default headers values
func NewUpdateIPBlockOK() *UpdateIPBlockOK {
	return &UpdateIPBlockOK{}
}

/*UpdateIPBlockOK handles this case with default header values.

Success
*/
type UpdateIPBlockOK struct {
	Payload *models.IPBlock
}

func (o *UpdateIPBlockOK) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockOK  %+v", 200, o.Payload)
}

func (o *UpdateIPBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPBlock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockBadRequest creates a UpdateIPBlockBadRequest with default headers values
func NewUpdateIPBlockBadRequest() *UpdateIPBlockBadRequest {
	return &UpdateIPBlockBadRequest{}
}

/*UpdateIPBlockBadRequest handles this case with default header values.

Bad request
*/
type UpdateIPBlockBadRequest struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIPBlockBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockForbidden creates a UpdateIPBlockForbidden with default headers values
func NewUpdateIPBlockForbidden() *UpdateIPBlockForbidden {
	return &UpdateIPBlockForbidden{}
}

/*UpdateIPBlockForbidden handles this case with default header values.

Operation forbidden
*/
type UpdateIPBlockForbidden struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockForbidden) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockForbidden  %+v", 403, o.Payload)
}

func (o *UpdateIPBlockForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockNotFound creates a UpdateIPBlockNotFound with default headers values
func NewUpdateIPBlockNotFound() *UpdateIPBlockNotFound {
	return &UpdateIPBlockNotFound{}
}

/*UpdateIPBlockNotFound handles this case with default header values.

Resource not found
*/
type UpdateIPBlockNotFound struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockNotFound) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIPBlockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockConflict creates a UpdateIPBlockConflict with default headers values
func NewUpdateIPBlockConflict() *UpdateIPBlockConflict {
	return &UpdateIPBlockConflict{}
}

/*UpdateIPBlockConflict handles this case with default header values.

UpdateIPBlockConflict update Ip block conflict
*/
type UpdateIPBlockConflict struct {
}

func (o *UpdateIPBlockConflict) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockConflict ", 409)
}

func (o *UpdateIPBlockConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateIPBlockPreconditionFailed creates a UpdateIPBlockPreconditionFailed with default headers values
func NewUpdateIPBlockPreconditionFailed() *UpdateIPBlockPreconditionFailed {
	return &UpdateIPBlockPreconditionFailed{}
}

/*UpdateIPBlockPreconditionFailed handles this case with default header values.

Precondition failed
*/
type UpdateIPBlockPreconditionFailed struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateIPBlockPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockInternalServerError creates a UpdateIPBlockInternalServerError with default headers values
func NewUpdateIPBlockInternalServerError() *UpdateIPBlockInternalServerError {
	return &UpdateIPBlockInternalServerError{}
}

/*UpdateIPBlockInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateIPBlockInternalServerError struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIPBlockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIPBlockServiceUnavailable creates a UpdateIPBlockServiceUnavailable with default headers values
func NewUpdateIPBlockServiceUnavailable() *UpdateIPBlockServiceUnavailable {
	return &UpdateIPBlockServiceUnavailable{}
}

/*UpdateIPBlockServiceUnavailable handles this case with default header values.

Service unavailable
*/
type UpdateIPBlockServiceUnavailable struct {
	Payload *models.APIError
}

func (o *UpdateIPBlockServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /pools/ip-blocks/{block-id}][%d] updateIpBlockServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateIPBlockServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

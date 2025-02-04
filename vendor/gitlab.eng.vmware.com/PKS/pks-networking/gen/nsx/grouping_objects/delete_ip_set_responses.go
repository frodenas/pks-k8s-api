// Code generated by go-swagger; DO NOT EDIT.

package grouping_objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// DeleteIPSetReader is a Reader for the DeleteIPSet structure.
type DeleteIPSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteIPSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteIPSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteIPSetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteIPSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIPSetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteIPSetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteIPSetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteIPSetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIPSetOK creates a DeleteIPSetOK with default headers values
func NewDeleteIPSetOK() *DeleteIPSetOK {
	return &DeleteIPSetOK{}
}

/*DeleteIPSetOK handles this case with default header values.

DeleteIPSetOK delete Ip set o k
*/
type DeleteIPSetOK struct {
}

func (o *DeleteIPSetOK) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetOK ", 200)
}

func (o *DeleteIPSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPSetBadRequest creates a DeleteIPSetBadRequest with default headers values
func NewDeleteIPSetBadRequest() *DeleteIPSetBadRequest {
	return &DeleteIPSetBadRequest{}
}

/*DeleteIPSetBadRequest handles this case with default header values.

Bad request
*/
type DeleteIPSetBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteIPSetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIPSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPSetForbidden creates a DeleteIPSetForbidden with default headers values
func NewDeleteIPSetForbidden() *DeleteIPSetForbidden {
	return &DeleteIPSetForbidden{}
}

/*DeleteIPSetForbidden handles this case with default header values.

Operation forbidden
*/
type DeleteIPSetForbidden struct {
	Payload *models.APIError
}

func (o *DeleteIPSetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIPSetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPSetNotFound creates a DeleteIPSetNotFound with default headers values
func NewDeleteIPSetNotFound() *DeleteIPSetNotFound {
	return &DeleteIPSetNotFound{}
}

/*DeleteIPSetNotFound handles this case with default header values.

Resource not found
*/
type DeleteIPSetNotFound struct {
	Payload *models.APIError
}

func (o *DeleteIPSetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIPSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPSetConflict creates a DeleteIPSetConflict with default headers values
func NewDeleteIPSetConflict() *DeleteIPSetConflict {
	return &DeleteIPSetConflict{}
}

/*DeleteIPSetConflict handles this case with default header values.

DeleteIPSetConflict delete Ip set conflict
*/
type DeleteIPSetConflict struct {
}

func (o *DeleteIPSetConflict) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetConflict ", 409)
}

func (o *DeleteIPSetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIPSetPreconditionFailed creates a DeleteIPSetPreconditionFailed with default headers values
func NewDeleteIPSetPreconditionFailed() *DeleteIPSetPreconditionFailed {
	return &DeleteIPSetPreconditionFailed{}
}

/*DeleteIPSetPreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteIPSetPreconditionFailed struct {
	Payload *models.APIError
}

func (o *DeleteIPSetPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteIPSetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPSetInternalServerError creates a DeleteIPSetInternalServerError with default headers values
func NewDeleteIPSetInternalServerError() *DeleteIPSetInternalServerError {
	return &DeleteIPSetInternalServerError{}
}

/*DeleteIPSetInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteIPSetInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteIPSetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPSetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPSetServiceUnavailable creates a DeleteIPSetServiceUnavailable with default headers values
func NewDeleteIPSetServiceUnavailable() *DeleteIPSetServiceUnavailable {
	return &DeleteIPSetServiceUnavailable{}
}

/*DeleteIPSetServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteIPSetServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteIPSetServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /ip-sets/{ip-set-id}][%d] deleteIpSetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIPSetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

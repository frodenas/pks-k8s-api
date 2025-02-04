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

// DeleteNSGroupReader is a Reader for the DeleteNSGroup structure.
type DeleteNSGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNSGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteNSGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteNSGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteNSGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteNSGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteNSGroupPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteNSGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteNSGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNSGroupOK creates a DeleteNSGroupOK with default headers values
func NewDeleteNSGroupOK() *DeleteNSGroupOK {
	return &DeleteNSGroupOK{}
}

/*DeleteNSGroupOK handles this case with default header values.

OK
*/
type DeleteNSGroupOK struct {
}

func (o *DeleteNSGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupOK ", 200)
}

func (o *DeleteNSGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNSGroupBadRequest creates a DeleteNSGroupBadRequest with default headers values
func NewDeleteNSGroupBadRequest() *DeleteNSGroupBadRequest {
	return &DeleteNSGroupBadRequest{}
}

/*DeleteNSGroupBadRequest handles this case with default header values.

Bad request
*/
type DeleteNSGroupBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNSGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNSGroupForbidden creates a DeleteNSGroupForbidden with default headers values
func NewDeleteNSGroupForbidden() *DeleteNSGroupForbidden {
	return &DeleteNSGroupForbidden{}
}

/*DeleteNSGroupForbidden handles this case with default header values.

Operation forbidden
*/
type DeleteNSGroupForbidden struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNSGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNSGroupNotFound creates a DeleteNSGroupNotFound with default headers values
func NewDeleteNSGroupNotFound() *DeleteNSGroupNotFound {
	return &DeleteNSGroupNotFound{}
}

/*DeleteNSGroupNotFound handles this case with default header values.

Resource not found
*/
type DeleteNSGroupNotFound struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNSGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNSGroupPreconditionFailed creates a DeleteNSGroupPreconditionFailed with default headers values
func NewDeleteNSGroupPreconditionFailed() *DeleteNSGroupPreconditionFailed {
	return &DeleteNSGroupPreconditionFailed{}
}

/*DeleteNSGroupPreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteNSGroupPreconditionFailed struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteNSGroupPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNSGroupInternalServerError creates a DeleteNSGroupInternalServerError with default headers values
func NewDeleteNSGroupInternalServerError() *DeleteNSGroupInternalServerError {
	return &DeleteNSGroupInternalServerError{}
}

/*DeleteNSGroupInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteNSGroupInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNSGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNSGroupServiceUnavailable creates a DeleteNSGroupServiceUnavailable with default headers values
func NewDeleteNSGroupServiceUnavailable() *DeleteNSGroupServiceUnavailable {
	return &DeleteNSGroupServiceUnavailable{}
}

/*DeleteNSGroupServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteNSGroupServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteNSGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /ns-groups/{ns-group-id}][%d] deleteNSGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteNSGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

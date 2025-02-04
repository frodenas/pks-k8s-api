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

// DeleteLoadBalancerPersistenceProfileReader is a Reader for the DeleteLoadBalancerPersistenceProfile structure.
type DeleteLoadBalancerPersistenceProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancerPersistenceProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteLoadBalancerPersistenceProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLoadBalancerPersistenceProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteLoadBalancerPersistenceProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLoadBalancerPersistenceProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteLoadBalancerPersistenceProfilePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteLoadBalancerPersistenceProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteLoadBalancerPersistenceProfileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancerPersistenceProfileOK creates a DeleteLoadBalancerPersistenceProfileOK with default headers values
func NewDeleteLoadBalancerPersistenceProfileOK() *DeleteLoadBalancerPersistenceProfileOK {
	return &DeleteLoadBalancerPersistenceProfileOK{}
}

/*DeleteLoadBalancerPersistenceProfileOK handles this case with default header values.

OK
*/
type DeleteLoadBalancerPersistenceProfileOK struct {
}

func (o *DeleteLoadBalancerPersistenceProfileOK) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileOK ", 200)
}

func (o *DeleteLoadBalancerPersistenceProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancerPersistenceProfileBadRequest creates a DeleteLoadBalancerPersistenceProfileBadRequest with default headers values
func NewDeleteLoadBalancerPersistenceProfileBadRequest() *DeleteLoadBalancerPersistenceProfileBadRequest {
	return &DeleteLoadBalancerPersistenceProfileBadRequest{}
}

/*DeleteLoadBalancerPersistenceProfileBadRequest handles this case with default header values.

Bad request
*/
type DeleteLoadBalancerPersistenceProfileBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerPersistenceProfileForbidden creates a DeleteLoadBalancerPersistenceProfileForbidden with default headers values
func NewDeleteLoadBalancerPersistenceProfileForbidden() *DeleteLoadBalancerPersistenceProfileForbidden {
	return &DeleteLoadBalancerPersistenceProfileForbidden{}
}

/*DeleteLoadBalancerPersistenceProfileForbidden handles this case with default header values.

Operation forbidden
*/
type DeleteLoadBalancerPersistenceProfileForbidden struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerPersistenceProfileNotFound creates a DeleteLoadBalancerPersistenceProfileNotFound with default headers values
func NewDeleteLoadBalancerPersistenceProfileNotFound() *DeleteLoadBalancerPersistenceProfileNotFound {
	return &DeleteLoadBalancerPersistenceProfileNotFound{}
}

/*DeleteLoadBalancerPersistenceProfileNotFound handles this case with default header values.

Resource not found
*/
type DeleteLoadBalancerPersistenceProfileNotFound struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerPersistenceProfilePreconditionFailed creates a DeleteLoadBalancerPersistenceProfilePreconditionFailed with default headers values
func NewDeleteLoadBalancerPersistenceProfilePreconditionFailed() *DeleteLoadBalancerPersistenceProfilePreconditionFailed {
	return &DeleteLoadBalancerPersistenceProfilePreconditionFailed{}
}

/*DeleteLoadBalancerPersistenceProfilePreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteLoadBalancerPersistenceProfilePreconditionFailed struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfilePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfilePreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfilePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerPersistenceProfileInternalServerError creates a DeleteLoadBalancerPersistenceProfileInternalServerError with default headers values
func NewDeleteLoadBalancerPersistenceProfileInternalServerError() *DeleteLoadBalancerPersistenceProfileInternalServerError {
	return &DeleteLoadBalancerPersistenceProfileInternalServerError{}
}

/*DeleteLoadBalancerPersistenceProfileInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteLoadBalancerPersistenceProfileInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfileInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerPersistenceProfileServiceUnavailable creates a DeleteLoadBalancerPersistenceProfileServiceUnavailable with default headers values
func NewDeleteLoadBalancerPersistenceProfileServiceUnavailable() *DeleteLoadBalancerPersistenceProfileServiceUnavailable {
	return &DeleteLoadBalancerPersistenceProfileServiceUnavailable{}
}

/*DeleteLoadBalancerPersistenceProfileServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteLoadBalancerPersistenceProfileServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerPersistenceProfileServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/persistence-profiles/{persistence-profile-id}][%d] deleteLoadBalancerPersistenceProfileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLoadBalancerPersistenceProfileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DeleteLoadBalancerMonitorReader is a Reader for the DeleteLoadBalancerMonitor structure.
type DeleteLoadBalancerMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancerMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteLoadBalancerMonitorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteLoadBalancerMonitorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteLoadBalancerMonitorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLoadBalancerMonitorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteLoadBalancerMonitorPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteLoadBalancerMonitorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewDeleteLoadBalancerMonitorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancerMonitorOK creates a DeleteLoadBalancerMonitorOK with default headers values
func NewDeleteLoadBalancerMonitorOK() *DeleteLoadBalancerMonitorOK {
	return &DeleteLoadBalancerMonitorOK{}
}

/*DeleteLoadBalancerMonitorOK handles this case with default header values.

OK
*/
type DeleteLoadBalancerMonitorOK struct {
}

func (o *DeleteLoadBalancerMonitorOK) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorOK ", 200)
}

func (o *DeleteLoadBalancerMonitorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancerMonitorBadRequest creates a DeleteLoadBalancerMonitorBadRequest with default headers values
func NewDeleteLoadBalancerMonitorBadRequest() *DeleteLoadBalancerMonitorBadRequest {
	return &DeleteLoadBalancerMonitorBadRequest{}
}

/*DeleteLoadBalancerMonitorBadRequest handles this case with default header values.

Bad request
*/
type DeleteLoadBalancerMonitorBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLoadBalancerMonitorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerMonitorForbidden creates a DeleteLoadBalancerMonitorForbidden with default headers values
func NewDeleteLoadBalancerMonitorForbidden() *DeleteLoadBalancerMonitorForbidden {
	return &DeleteLoadBalancerMonitorForbidden{}
}

/*DeleteLoadBalancerMonitorForbidden handles this case with default header values.

Operation forbidden
*/
type DeleteLoadBalancerMonitorForbidden struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLoadBalancerMonitorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerMonitorNotFound creates a DeleteLoadBalancerMonitorNotFound with default headers values
func NewDeleteLoadBalancerMonitorNotFound() *DeleteLoadBalancerMonitorNotFound {
	return &DeleteLoadBalancerMonitorNotFound{}
}

/*DeleteLoadBalancerMonitorNotFound handles this case with default header values.

Resource not found
*/
type DeleteLoadBalancerMonitorNotFound struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLoadBalancerMonitorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerMonitorPreconditionFailed creates a DeleteLoadBalancerMonitorPreconditionFailed with default headers values
func NewDeleteLoadBalancerMonitorPreconditionFailed() *DeleteLoadBalancerMonitorPreconditionFailed {
	return &DeleteLoadBalancerMonitorPreconditionFailed{}
}

/*DeleteLoadBalancerMonitorPreconditionFailed handles this case with default header values.

Precondition failed
*/
type DeleteLoadBalancerMonitorPreconditionFailed struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteLoadBalancerMonitorPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerMonitorInternalServerError creates a DeleteLoadBalancerMonitorInternalServerError with default headers values
func NewDeleteLoadBalancerMonitorInternalServerError() *DeleteLoadBalancerMonitorInternalServerError {
	return &DeleteLoadBalancerMonitorInternalServerError{}
}

/*DeleteLoadBalancerMonitorInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteLoadBalancerMonitorInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLoadBalancerMonitorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerMonitorServiceUnavailable creates a DeleteLoadBalancerMonitorServiceUnavailable with default headers values
func NewDeleteLoadBalancerMonitorServiceUnavailable() *DeleteLoadBalancerMonitorServiceUnavailable {
	return &DeleteLoadBalancerMonitorServiceUnavailable{}
}

/*DeleteLoadBalancerMonitorServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteLoadBalancerMonitorServiceUnavailable struct {
	Payload *models.APIError
}

func (o *DeleteLoadBalancerMonitorServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /loadbalancer/monitors/{monitor-id}][%d] deleteLoadBalancerMonitorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLoadBalancerMonitorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package logical_routing_and_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// CreateLogicalRouterPortReader is a Reader for the CreateLogicalRouterPort structure.
type CreateLogicalRouterPortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLogicalRouterPortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateLogicalRouterPortCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateLogicalRouterPortBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateLogicalRouterPortForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateLogicalRouterPortNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewCreateLogicalRouterPortPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateLogicalRouterPortInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCreateLogicalRouterPortServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLogicalRouterPortCreated creates a CreateLogicalRouterPortCreated with default headers values
func NewCreateLogicalRouterPortCreated() *CreateLogicalRouterPortCreated {
	return &CreateLogicalRouterPortCreated{}
}

/*CreateLogicalRouterPortCreated handles this case with default header values.

Resource created successfully
*/
type CreateLogicalRouterPortCreated struct {
	Payload *models.LogicalRouterPort
}

func (o *CreateLogicalRouterPortCreated) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortCreated  %+v", 201, o.Payload)
}

func (o *CreateLogicalRouterPortCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogicalRouterPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortBadRequest creates a CreateLogicalRouterPortBadRequest with default headers values
func NewCreateLogicalRouterPortBadRequest() *CreateLogicalRouterPortBadRequest {
	return &CreateLogicalRouterPortBadRequest{}
}

/*CreateLogicalRouterPortBadRequest handles this case with default header values.

Bad request
*/
type CreateLogicalRouterPortBadRequest struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortBadRequest) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLogicalRouterPortBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortForbidden creates a CreateLogicalRouterPortForbidden with default headers values
func NewCreateLogicalRouterPortForbidden() *CreateLogicalRouterPortForbidden {
	return &CreateLogicalRouterPortForbidden{}
}

/*CreateLogicalRouterPortForbidden handles this case with default header values.

Operation forbidden
*/
type CreateLogicalRouterPortForbidden struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortForbidden) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortForbidden  %+v", 403, o.Payload)
}

func (o *CreateLogicalRouterPortForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortNotFound creates a CreateLogicalRouterPortNotFound with default headers values
func NewCreateLogicalRouterPortNotFound() *CreateLogicalRouterPortNotFound {
	return &CreateLogicalRouterPortNotFound{}
}

/*CreateLogicalRouterPortNotFound handles this case with default header values.

Resource not found
*/
type CreateLogicalRouterPortNotFound struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortNotFound) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortNotFound  %+v", 404, o.Payload)
}

func (o *CreateLogicalRouterPortNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortPreconditionFailed creates a CreateLogicalRouterPortPreconditionFailed with default headers values
func NewCreateLogicalRouterPortPreconditionFailed() *CreateLogicalRouterPortPreconditionFailed {
	return &CreateLogicalRouterPortPreconditionFailed{}
}

/*CreateLogicalRouterPortPreconditionFailed handles this case with default header values.

Precondition failed
*/
type CreateLogicalRouterPortPreconditionFailed struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreateLogicalRouterPortPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortInternalServerError creates a CreateLogicalRouterPortInternalServerError with default headers values
func NewCreateLogicalRouterPortInternalServerError() *CreateLogicalRouterPortInternalServerError {
	return &CreateLogicalRouterPortInternalServerError{}
}

/*CreateLogicalRouterPortInternalServerError handles this case with default header values.

Internal server error
*/
type CreateLogicalRouterPortInternalServerError struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortInternalServerError) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLogicalRouterPortInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLogicalRouterPortServiceUnavailable creates a CreateLogicalRouterPortServiceUnavailable with default headers values
func NewCreateLogicalRouterPortServiceUnavailable() *CreateLogicalRouterPortServiceUnavailable {
	return &CreateLogicalRouterPortServiceUnavailable{}
}

/*CreateLogicalRouterPortServiceUnavailable handles this case with default header values.

Service unavailable
*/
type CreateLogicalRouterPortServiceUnavailable struct {
	Payload *models.APIError
}

func (o *CreateLogicalRouterPortServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /logical-router-ports][%d] createLogicalRouterPortServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateLogicalRouterPortServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

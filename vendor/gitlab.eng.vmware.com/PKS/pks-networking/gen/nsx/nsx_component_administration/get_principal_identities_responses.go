// Code generated by go-swagger; DO NOT EDIT.

package nsx_component_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// GetPrincipalIdentitiesReader is a Reader for the GetPrincipalIdentities structure.
type GetPrincipalIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrincipalIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrincipalIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPrincipalIdentitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPrincipalIdentitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPrincipalIdentitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewGetPrincipalIdentitiesPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPrincipalIdentitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetPrincipalIdentitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrincipalIdentitiesOK creates a GetPrincipalIdentitiesOK with default headers values
func NewGetPrincipalIdentitiesOK() *GetPrincipalIdentitiesOK {
	return &GetPrincipalIdentitiesOK{}
}

/*GetPrincipalIdentitiesOK handles this case with default header values.

OK
*/
type GetPrincipalIdentitiesOK struct {
	Payload *models.PrincipalIdentityList
}

func (o *GetPrincipalIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesOK  %+v", 200, o.Payload)
}

func (o *GetPrincipalIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrincipalIdentityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesBadRequest creates a GetPrincipalIdentitiesBadRequest with default headers values
func NewGetPrincipalIdentitiesBadRequest() *GetPrincipalIdentitiesBadRequest {
	return &GetPrincipalIdentitiesBadRequest{}
}

/*GetPrincipalIdentitiesBadRequest handles this case with default header values.

Bad request
*/
type GetPrincipalIdentitiesBadRequest struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPrincipalIdentitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesForbidden creates a GetPrincipalIdentitiesForbidden with default headers values
func NewGetPrincipalIdentitiesForbidden() *GetPrincipalIdentitiesForbidden {
	return &GetPrincipalIdentitiesForbidden{}
}

/*GetPrincipalIdentitiesForbidden handles this case with default header values.

Operation forbidden
*/
type GetPrincipalIdentitiesForbidden struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetPrincipalIdentitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesNotFound creates a GetPrincipalIdentitiesNotFound with default headers values
func NewGetPrincipalIdentitiesNotFound() *GetPrincipalIdentitiesNotFound {
	return &GetPrincipalIdentitiesNotFound{}
}

/*GetPrincipalIdentitiesNotFound handles this case with default header values.

Resource not found
*/
type GetPrincipalIdentitiesNotFound struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesNotFound  %+v", 404, o.Payload)
}

func (o *GetPrincipalIdentitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesPreconditionFailed creates a GetPrincipalIdentitiesPreconditionFailed with default headers values
func NewGetPrincipalIdentitiesPreconditionFailed() *GetPrincipalIdentitiesPreconditionFailed {
	return &GetPrincipalIdentitiesPreconditionFailed{}
}

/*GetPrincipalIdentitiesPreconditionFailed handles this case with default header values.

Precondition failed
*/
type GetPrincipalIdentitiesPreconditionFailed struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetPrincipalIdentitiesPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesInternalServerError creates a GetPrincipalIdentitiesInternalServerError with default headers values
func NewGetPrincipalIdentitiesInternalServerError() *GetPrincipalIdentitiesInternalServerError {
	return &GetPrincipalIdentitiesInternalServerError{}
}

/*GetPrincipalIdentitiesInternalServerError handles this case with default header values.

Internal server error
*/
type GetPrincipalIdentitiesInternalServerError struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPrincipalIdentitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrincipalIdentitiesServiceUnavailable creates a GetPrincipalIdentitiesServiceUnavailable with default headers values
func NewGetPrincipalIdentitiesServiceUnavailable() *GetPrincipalIdentitiesServiceUnavailable {
	return &GetPrincipalIdentitiesServiceUnavailable{}
}

/*GetPrincipalIdentitiesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetPrincipalIdentitiesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetPrincipalIdentitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /trust-management/principal-identities][%d] getPrincipalIdentitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPrincipalIdentitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

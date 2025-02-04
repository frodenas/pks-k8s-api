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

// GetCertificateReader is a Reader for the GetCertificate structure.
type GetCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewGetCertificatePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCertificateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCertificateOK creates a GetCertificateOK with default headers values
func NewGetCertificateOK() *GetCertificateOK {
	return &GetCertificateOK{}
}

/*GetCertificateOK handles this case with default header values.

OK
*/
type GetCertificateOK struct {
	Payload *models.Certificate
}

func (o *GetCertificateOK) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateOK  %+v", 200, o.Payload)
}

func (o *GetCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Certificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateBadRequest creates a GetCertificateBadRequest with default headers values
func NewGetCertificateBadRequest() *GetCertificateBadRequest {
	return &GetCertificateBadRequest{}
}

/*GetCertificateBadRequest handles this case with default header values.

Bad request
*/
type GetCertificateBadRequest struct {
	Payload *models.APIError
}

func (o *GetCertificateBadRequest) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *GetCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateForbidden creates a GetCertificateForbidden with default headers values
func NewGetCertificateForbidden() *GetCertificateForbidden {
	return &GetCertificateForbidden{}
}

/*GetCertificateForbidden handles this case with default header values.

Operation forbidden
*/
type GetCertificateForbidden struct {
	Payload *models.APIError
}

func (o *GetCertificateForbidden) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateForbidden  %+v", 403, o.Payload)
}

func (o *GetCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateNotFound creates a GetCertificateNotFound with default headers values
func NewGetCertificateNotFound() *GetCertificateNotFound {
	return &GetCertificateNotFound{}
}

/*GetCertificateNotFound handles this case with default header values.

Resource not found
*/
type GetCertificateNotFound struct {
	Payload *models.APIError
}

func (o *GetCertificateNotFound) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateNotFound  %+v", 404, o.Payload)
}

func (o *GetCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificatePreconditionFailed creates a GetCertificatePreconditionFailed with default headers values
func NewGetCertificatePreconditionFailed() *GetCertificatePreconditionFailed {
	return &GetCertificatePreconditionFailed{}
}

/*GetCertificatePreconditionFailed handles this case with default header values.

Precondition failed
*/
type GetCertificatePreconditionFailed struct {
	Payload *models.APIError
}

func (o *GetCertificatePreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificatePreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetCertificatePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateInternalServerError creates a GetCertificateInternalServerError with default headers values
func NewGetCertificateInternalServerError() *GetCertificateInternalServerError {
	return &GetCertificateInternalServerError{}
}

/*GetCertificateInternalServerError handles this case with default header values.

Internal server error
*/
type GetCertificateInternalServerError struct {
	Payload *models.APIError
}

func (o *GetCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateServiceUnavailable creates a GetCertificateServiceUnavailable with default headers values
func NewGetCertificateServiceUnavailable() *GetCertificateServiceUnavailable {
	return &GetCertificateServiceUnavailable{}
}

/*GetCertificateServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCertificateServiceUnavailable struct {
	Payload *models.APIError
}

func (o *GetCertificateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /trust-management/certificates/{cert-id}][%d] getCertificateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCertificateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

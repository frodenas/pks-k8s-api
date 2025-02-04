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

// ListLoadBalancerRulesReader is a Reader for the ListLoadBalancerRules structure.
type ListLoadBalancerRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLoadBalancerRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListLoadBalancerRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListLoadBalancerRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListLoadBalancerRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListLoadBalancerRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewListLoadBalancerRulesPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListLoadBalancerRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewListLoadBalancerRulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListLoadBalancerRulesOK creates a ListLoadBalancerRulesOK with default headers values
func NewListLoadBalancerRulesOK() *ListLoadBalancerRulesOK {
	return &ListLoadBalancerRulesOK{}
}

/*ListLoadBalancerRulesOK handles this case with default header values.

OK
*/
type ListLoadBalancerRulesOK struct {
	Payload *models.LbRuleListResult
}

func (o *ListLoadBalancerRulesOK) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesOK  %+v", 200, o.Payload)
}

func (o *ListLoadBalancerRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LbRuleListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesBadRequest creates a ListLoadBalancerRulesBadRequest with default headers values
func NewListLoadBalancerRulesBadRequest() *ListLoadBalancerRulesBadRequest {
	return &ListLoadBalancerRulesBadRequest{}
}

/*ListLoadBalancerRulesBadRequest handles this case with default header values.

Bad request
*/
type ListLoadBalancerRulesBadRequest struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesBadRequest  %+v", 400, o.Payload)
}

func (o *ListLoadBalancerRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesForbidden creates a ListLoadBalancerRulesForbidden with default headers values
func NewListLoadBalancerRulesForbidden() *ListLoadBalancerRulesForbidden {
	return &ListLoadBalancerRulesForbidden{}
}

/*ListLoadBalancerRulesForbidden handles this case with default header values.

Operation forbidden
*/
type ListLoadBalancerRulesForbidden struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesForbidden  %+v", 403, o.Payload)
}

func (o *ListLoadBalancerRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesNotFound creates a ListLoadBalancerRulesNotFound with default headers values
func NewListLoadBalancerRulesNotFound() *ListLoadBalancerRulesNotFound {
	return &ListLoadBalancerRulesNotFound{}
}

/*ListLoadBalancerRulesNotFound handles this case with default header values.

Resource not found
*/
type ListLoadBalancerRulesNotFound struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesNotFound  %+v", 404, o.Payload)
}

func (o *ListLoadBalancerRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesPreconditionFailed creates a ListLoadBalancerRulesPreconditionFailed with default headers values
func NewListLoadBalancerRulesPreconditionFailed() *ListLoadBalancerRulesPreconditionFailed {
	return &ListLoadBalancerRulesPreconditionFailed{}
}

/*ListLoadBalancerRulesPreconditionFailed handles this case with default header values.

Precondition failed
*/
type ListLoadBalancerRulesPreconditionFailed struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ListLoadBalancerRulesPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesInternalServerError creates a ListLoadBalancerRulesInternalServerError with default headers values
func NewListLoadBalancerRulesInternalServerError() *ListLoadBalancerRulesInternalServerError {
	return &ListLoadBalancerRulesInternalServerError{}
}

/*ListLoadBalancerRulesInternalServerError handles this case with default header values.

Internal server error
*/
type ListLoadBalancerRulesInternalServerError struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListLoadBalancerRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLoadBalancerRulesServiceUnavailable creates a ListLoadBalancerRulesServiceUnavailable with default headers values
func NewListLoadBalancerRulesServiceUnavailable() *ListLoadBalancerRulesServiceUnavailable {
	return &ListLoadBalancerRulesServiceUnavailable{}
}

/*ListLoadBalancerRulesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type ListLoadBalancerRulesServiceUnavailable struct {
	Payload *models.APIError
}

func (o *ListLoadBalancerRulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /loadbalancer/rules][%d] listLoadBalancerRulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListLoadBalancerRulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

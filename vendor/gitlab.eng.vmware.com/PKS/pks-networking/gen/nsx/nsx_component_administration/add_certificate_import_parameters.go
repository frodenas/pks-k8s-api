// Code generated by go-swagger; DO NOT EDIT.

package nsx_component_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// NewAddCertificateImportParams creates a new AddCertificateImportParams object
// with the default values initialized.
func NewAddCertificateImportParams() *AddCertificateImportParams {
	var ()
	return &AddCertificateImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCertificateImportParamsWithTimeout creates a new AddCertificateImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCertificateImportParamsWithTimeout(timeout time.Duration) *AddCertificateImportParams {
	var ()
	return &AddCertificateImportParams{

		timeout: timeout,
	}
}

// NewAddCertificateImportParamsWithContext creates a new AddCertificateImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCertificateImportParamsWithContext(ctx context.Context) *AddCertificateImportParams {
	var ()
	return &AddCertificateImportParams{

		Context: ctx,
	}
}

// NewAddCertificateImportParamsWithHTTPClient creates a new AddCertificateImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCertificateImportParamsWithHTTPClient(client *http.Client) *AddCertificateImportParams {
	var ()
	return &AddCertificateImportParams{
		HTTPClient: client,
	}
}

/*AddCertificateImportParams contains all the parameters to send to the API endpoint
for the add certificate import operation typically these are written to a http.Request
*/
type AddCertificateImportParams struct {

	/*TrustObjectData*/
	TrustObjectData *models.TrustObjectData
	/*Action
	  Specifies allocate or release action

	*/
	Action string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add certificate import params
func (o *AddCertificateImportParams) WithTimeout(timeout time.Duration) *AddCertificateImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add certificate import params
func (o *AddCertificateImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add certificate import params
func (o *AddCertificateImportParams) WithContext(ctx context.Context) *AddCertificateImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add certificate import params
func (o *AddCertificateImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add certificate import params
func (o *AddCertificateImportParams) WithHTTPClient(client *http.Client) *AddCertificateImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add certificate import params
func (o *AddCertificateImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrustObjectData adds the trustObjectData to the add certificate import params
func (o *AddCertificateImportParams) WithTrustObjectData(trustObjectData *models.TrustObjectData) *AddCertificateImportParams {
	o.SetTrustObjectData(trustObjectData)
	return o
}

// SetTrustObjectData adds the trustObjectData to the add certificate import params
func (o *AddCertificateImportParams) SetTrustObjectData(trustObjectData *models.TrustObjectData) {
	o.TrustObjectData = trustObjectData
}

// WithAction adds the action to the add certificate import params
func (o *AddCertificateImportParams) WithAction(action string) *AddCertificateImportParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the add certificate import params
func (o *AddCertificateImportParams) SetAction(action string) {
	o.Action = action
}

// WriteToRequest writes these params to a swagger request
func (o *AddCertificateImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TrustObjectData == nil {
		o.TrustObjectData = new(models.TrustObjectData)
	}

	if err := r.SetBodyParam(o.TrustObjectData); err != nil {
		return err
	}

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {
		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

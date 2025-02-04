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
)

// NewReadClusterStatusParams creates a new ReadClusterStatusParams object
// with the default values initialized.
func NewReadClusterStatusParams() *ReadClusterStatusParams {
	var ()
	return &ReadClusterStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadClusterStatusParamsWithTimeout creates a new ReadClusterStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadClusterStatusParamsWithTimeout(timeout time.Duration) *ReadClusterStatusParams {
	var ()
	return &ReadClusterStatusParams{

		timeout: timeout,
	}
}

// NewReadClusterStatusParamsWithContext creates a new ReadClusterStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadClusterStatusParamsWithContext(ctx context.Context) *ReadClusterStatusParams {
	var ()
	return &ReadClusterStatusParams{

		Context: ctx,
	}
}

// NewReadClusterStatusParamsWithHTTPClient creates a new ReadClusterStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadClusterStatusParamsWithHTTPClient(client *http.Client) *ReadClusterStatusParams {
	var ()
	return &ReadClusterStatusParams{
		HTTPClient: client,
	}
}

/*ReadClusterStatusParams contains all the parameters to send to the API endpoint
for the read cluster status operation typically these are written to a http.Request
*/
type ReadClusterStatusParams struct {

	/*Source
	  Data source type.

	*/
	Source *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read cluster status params
func (o *ReadClusterStatusParams) WithTimeout(timeout time.Duration) *ReadClusterStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read cluster status params
func (o *ReadClusterStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read cluster status params
func (o *ReadClusterStatusParams) WithContext(ctx context.Context) *ReadClusterStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read cluster status params
func (o *ReadClusterStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read cluster status params
func (o *ReadClusterStatusParams) WithHTTPClient(client *http.Client) *ReadClusterStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read cluster status params
func (o *ReadClusterStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSource adds the source to the read cluster status params
func (o *ReadClusterStatusParams) WithSource(source *string) *ReadClusterStatusParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the read cluster status params
func (o *ReadClusterStatusParams) SetSource(source *string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *ReadClusterStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Source != nil {

		// query param source
		var qrSource string
		if o.Source != nil {
			qrSource = *o.Source
		}
		qSource := qrSource
		if qSource != "" {
			if err := r.SetQueryParam("source", qSource); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

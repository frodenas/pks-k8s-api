// Code generated by go-swagger; DO NOT EDIT.

package logical_switching

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

// NewGetLogicalPortParams creates a new GetLogicalPortParams object
// with the default values initialized.
func NewGetLogicalPortParams() *GetLogicalPortParams {
	var ()
	return &GetLogicalPortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogicalPortParamsWithTimeout creates a new GetLogicalPortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLogicalPortParamsWithTimeout(timeout time.Duration) *GetLogicalPortParams {
	var ()
	return &GetLogicalPortParams{

		timeout: timeout,
	}
}

// NewGetLogicalPortParamsWithContext creates a new GetLogicalPortParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLogicalPortParamsWithContext(ctx context.Context) *GetLogicalPortParams {
	var ()
	return &GetLogicalPortParams{

		Context: ctx,
	}
}

// NewGetLogicalPortParamsWithHTTPClient creates a new GetLogicalPortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLogicalPortParamsWithHTTPClient(client *http.Client) *GetLogicalPortParams {
	var ()
	return &GetLogicalPortParams{
		HTTPClient: client,
	}
}

/*GetLogicalPortParams contains all the parameters to send to the API endpoint
for the get logical port operation typically these are written to a http.Request
*/
type GetLogicalPortParams struct {

	/*LportID*/
	LportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get logical port params
func (o *GetLogicalPortParams) WithTimeout(timeout time.Duration) *GetLogicalPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logical port params
func (o *GetLogicalPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logical port params
func (o *GetLogicalPortParams) WithContext(ctx context.Context) *GetLogicalPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logical port params
func (o *GetLogicalPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logical port params
func (o *GetLogicalPortParams) WithHTTPClient(client *http.Client) *GetLogicalPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logical port params
func (o *GetLogicalPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLportID adds the lportID to the get logical port params
func (o *GetLogicalPortParams) WithLportID(lportID string) *GetLogicalPortParams {
	o.SetLportID(lportID)
	return o
}

// SetLportID adds the lportId to the get logical port params
func (o *GetLogicalPortParams) SetLportID(lportID string) {
	o.LportID = lportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogicalPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lport-id
	if err := r.SetPathParam("lport-id", o.LportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

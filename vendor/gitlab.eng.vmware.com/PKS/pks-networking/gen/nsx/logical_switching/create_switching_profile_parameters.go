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

	"gitlab.eng.vmware.com/PKS/pks-networking/gen/models"
)

// NewCreateSwitchingProfileParams creates a new CreateSwitchingProfileParams object
// with the default values initialized.
func NewCreateSwitchingProfileParams() *CreateSwitchingProfileParams {
	var ()
	return &CreateSwitchingProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSwitchingProfileParamsWithTimeout creates a new CreateSwitchingProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSwitchingProfileParamsWithTimeout(timeout time.Duration) *CreateSwitchingProfileParams {
	var ()
	return &CreateSwitchingProfileParams{

		timeout: timeout,
	}
}

// NewCreateSwitchingProfileParamsWithContext creates a new CreateSwitchingProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSwitchingProfileParamsWithContext(ctx context.Context) *CreateSwitchingProfileParams {
	var ()
	return &CreateSwitchingProfileParams{

		Context: ctx,
	}
}

// NewCreateSwitchingProfileParamsWithHTTPClient creates a new CreateSwitchingProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSwitchingProfileParamsWithHTTPClient(client *http.Client) *CreateSwitchingProfileParams {
	var ()
	return &CreateSwitchingProfileParams{
		HTTPClient: client,
	}
}

/*CreateSwitchingProfileParams contains all the parameters to send to the API endpoint
for the create switching profile operation typically these are written to a http.Request
*/
type CreateSwitchingProfileParams struct {

	/*SpoofGuardSwitchingProfile*/
	SpoofGuardSwitchingProfile *models.SpoofGuardSwitchingProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create switching profile params
func (o *CreateSwitchingProfileParams) WithTimeout(timeout time.Duration) *CreateSwitchingProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create switching profile params
func (o *CreateSwitchingProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create switching profile params
func (o *CreateSwitchingProfileParams) WithContext(ctx context.Context) *CreateSwitchingProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create switching profile params
func (o *CreateSwitchingProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create switching profile params
func (o *CreateSwitchingProfileParams) WithHTTPClient(client *http.Client) *CreateSwitchingProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create switching profile params
func (o *CreateSwitchingProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSpoofGuardSwitchingProfile adds the spoofGuardSwitchingProfile to the create switching profile params
func (o *CreateSwitchingProfileParams) WithSpoofGuardSwitchingProfile(spoofGuardSwitchingProfile *models.SpoofGuardSwitchingProfile) *CreateSwitchingProfileParams {
	o.SetSpoofGuardSwitchingProfile(spoofGuardSwitchingProfile)
	return o
}

// SetSpoofGuardSwitchingProfile adds the spoofGuardSwitchingProfile to the create switching profile params
func (o *CreateSwitchingProfileParams) SetSpoofGuardSwitchingProfile(spoofGuardSwitchingProfile *models.SpoofGuardSwitchingProfile) {
	o.SpoofGuardSwitchingProfile = spoofGuardSwitchingProfile
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSwitchingProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SpoofGuardSwitchingProfile == nil {
		o.SpoofGuardSwitchingProfile = new(models.SpoofGuardSwitchingProfile)
	}

	if err := r.SetBodyParam(o.SpoofGuardSwitchingProfile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

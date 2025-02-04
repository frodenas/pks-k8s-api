// Code generated by go-swagger; DO NOT EDIT.

package logical_routing_and_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListLogicalRouterPortsParams creates a new ListLogicalRouterPortsParams object
// with the default values initialized.
func NewListLogicalRouterPortsParams() *ListLogicalRouterPortsParams {
	var (
		pageSizeDefault = int64(1000)
	)
	return &ListLogicalRouterPortsParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListLogicalRouterPortsParamsWithTimeout creates a new ListLogicalRouterPortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListLogicalRouterPortsParamsWithTimeout(timeout time.Duration) *ListLogicalRouterPortsParams {
	var (
		pageSizeDefault = int64(1000)
	)
	return &ListLogicalRouterPortsParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListLogicalRouterPortsParamsWithContext creates a new ListLogicalRouterPortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListLogicalRouterPortsParamsWithContext(ctx context.Context) *ListLogicalRouterPortsParams {
	var (
		pageSizeDefault = int64(1000)
	)
	return &ListLogicalRouterPortsParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewListLogicalRouterPortsParamsWithHTTPClient creates a new ListLogicalRouterPortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListLogicalRouterPortsParamsWithHTTPClient(client *http.Client) *ListLogicalRouterPortsParams {
	var (
		pageSizeDefault = int64(1000)
	)
	return &ListLogicalRouterPortsParams{
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*ListLogicalRouterPortsParams contains all the parameters to send to the API endpoint
for the list logical router ports operation typically these are written to a http.Request
*/
type ListLogicalRouterPortsParams struct {

	/*Cursor
	  Opaque cursor to be used for getting next page of records (supplied by current result page)

	*/
	Cursor *string
	/*IncludedFields
	  Comma separated list of fields that should be included to result of query

	*/
	IncludedFields *string
	/*LogicalRouterID
	  Logical Router identifier

	*/
	LogicalRouterID *string
	/*LogicalSwitchID
	  Logical Switch identifier

	*/
	LogicalSwitchID *string
	/*PageSize
	  Maximum number of results to return in this page (server may return fewer)

	*/
	PageSize *int64
	/*ResourceType
	  Resource types of logical router port

	*/
	ResourceType *string
	/*SortAscending*/
	SortAscending *bool
	/*SortBy
	  Field by which records are sorted

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithTimeout(timeout time.Duration) *ListLogicalRouterPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithContext(ctx context.Context) *ListLogicalRouterPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithHTTPClient(client *http.Client) *ListLogicalRouterPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithCursor(cursor *string) *ListLogicalRouterPortsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithIncludedFields adds the includedFields to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithIncludedFields(includedFields *string) *ListLogicalRouterPortsParams {
	o.SetIncludedFields(includedFields)
	return o
}

// SetIncludedFields adds the includedFields to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetIncludedFields(includedFields *string) {
	o.IncludedFields = includedFields
}

// WithLogicalRouterID adds the logicalRouterID to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithLogicalRouterID(logicalRouterID *string) *ListLogicalRouterPortsParams {
	o.SetLogicalRouterID(logicalRouterID)
	return o
}

// SetLogicalRouterID adds the logicalRouterId to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetLogicalRouterID(logicalRouterID *string) {
	o.LogicalRouterID = logicalRouterID
}

// WithLogicalSwitchID adds the logicalSwitchID to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithLogicalSwitchID(logicalSwitchID *string) *ListLogicalRouterPortsParams {
	o.SetLogicalSwitchID(logicalSwitchID)
	return o
}

// SetLogicalSwitchID adds the logicalSwitchId to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetLogicalSwitchID(logicalSwitchID *string) {
	o.LogicalSwitchID = logicalSwitchID
}

// WithPageSize adds the pageSize to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithPageSize(pageSize *int64) *ListLogicalRouterPortsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithResourceType adds the resourceType to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithResourceType(resourceType *string) *ListLogicalRouterPortsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithSortAscending adds the sortAscending to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithSortAscending(sortAscending *bool) *ListLogicalRouterPortsParams {
	o.SetSortAscending(sortAscending)
	return o
}

// SetSortAscending adds the sortAscending to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetSortAscending(sortAscending *bool) {
	o.SortAscending = sortAscending
}

// WithSortBy adds the sortBy to the list logical router ports params
func (o *ListLogicalRouterPortsParams) WithSortBy(sortBy *string) *ListLogicalRouterPortsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list logical router ports params
func (o *ListLogicalRouterPortsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *ListLogicalRouterPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.IncludedFields != nil {

		// query param included_fields
		var qrIncludedFields string
		if o.IncludedFields != nil {
			qrIncludedFields = *o.IncludedFields
		}
		qIncludedFields := qrIncludedFields
		if qIncludedFields != "" {
			if err := r.SetQueryParam("included_fields", qIncludedFields); err != nil {
				return err
			}
		}

	}

	if o.LogicalRouterID != nil {

		// query param logical_router_id
		var qrLogicalRouterID string
		if o.LogicalRouterID != nil {
			qrLogicalRouterID = *o.LogicalRouterID
		}
		qLogicalRouterID := qrLogicalRouterID
		if qLogicalRouterID != "" {
			if err := r.SetQueryParam("logical_router_id", qLogicalRouterID); err != nil {
				return err
			}
		}

	}

	if o.LogicalSwitchID != nil {

		// query param logical_switch_id
		var qrLogicalSwitchID string
		if o.LogicalSwitchID != nil {
			qrLogicalSwitchID = *o.LogicalSwitchID
		}
		qLogicalSwitchID := qrLogicalSwitchID
		if qLogicalSwitchID != "" {
			if err := r.SetQueryParam("logical_switch_id", qLogicalSwitchID); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.ResourceType != nil {

		// query param resource_type
		var qrResourceType string
		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {
			if err := r.SetQueryParam("resource_type", qResourceType); err != nil {
				return err
			}
		}

	}

	if o.SortAscending != nil {

		// query param sort_ascending
		var qrSortAscending bool
		if o.SortAscending != nil {
			qrSortAscending = *o.SortAscending
		}
		qSortAscending := swag.FormatBool(qrSortAscending)
		if qSortAscending != "" {
			if err := r.SetQueryParam("sort_ascending", qSortAscending); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewQueryFirewallFieldsParams creates a new QueryFirewallFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryFirewallFieldsParams() *QueryFirewallFieldsParams {
	return &QueryFirewallFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryFirewallFieldsParamsWithTimeout creates a new QueryFirewallFieldsParams object
// with the ability to set a timeout on a request.
func NewQueryFirewallFieldsParamsWithTimeout(timeout time.Duration) *QueryFirewallFieldsParams {
	return &QueryFirewallFieldsParams{
		timeout: timeout,
	}
}

// NewQueryFirewallFieldsParamsWithContext creates a new QueryFirewallFieldsParams object
// with the ability to set a context for a request.
func NewQueryFirewallFieldsParamsWithContext(ctx context.Context) *QueryFirewallFieldsParams {
	return &QueryFirewallFieldsParams{
		Context: ctx,
	}
}

// NewQueryFirewallFieldsParamsWithHTTPClient creates a new QueryFirewallFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryFirewallFieldsParamsWithHTTPClient(client *http.Client) *QueryFirewallFieldsParams {
	return &QueryFirewallFieldsParams{
		HTTPClient: client,
	}
}

/* QueryFirewallFieldsParams contains all the parameters to send to the API endpoint
   for the query firewall fields operation.

   Typically these are written to a http.Request.
*/
type QueryFirewallFieldsParams struct {

	/* Limit.

	   Number of ids to return.
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return ids.
	*/
	Offset *string

	/* PlatformID.

	   Get fields configuration for this platform
	*/
	PlatformID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query firewall fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryFirewallFieldsParams) WithDefaults() *QueryFirewallFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query firewall fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryFirewallFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithTimeout(timeout time.Duration) *QueryFirewallFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithContext(ctx context.Context) *QueryFirewallFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithHTTPClient(client *http.Client) *QueryFirewallFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithLimit(limit *int64) *QueryFirewallFieldsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithOffset(offset *string) *QueryFirewallFieldsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithPlatformID adds the platformID to the query firewall fields params
func (o *QueryFirewallFieldsParams) WithPlatformID(platformID *string) *QueryFirewallFieldsParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the query firewall fields params
func (o *QueryFirewallFieldsParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryFirewallFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string

		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {

			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

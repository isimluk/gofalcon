// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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
)

// NewRTRDeleteScriptsParams creates a new RTRDeleteScriptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRDeleteScriptsParams() *RTRDeleteScriptsParams {
	return &RTRDeleteScriptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRDeleteScriptsParamsWithTimeout creates a new RTRDeleteScriptsParams object
// with the ability to set a timeout on a request.
func NewRTRDeleteScriptsParamsWithTimeout(timeout time.Duration) *RTRDeleteScriptsParams {
	return &RTRDeleteScriptsParams{
		timeout: timeout,
	}
}

// NewRTRDeleteScriptsParamsWithContext creates a new RTRDeleteScriptsParams object
// with the ability to set a context for a request.
func NewRTRDeleteScriptsParamsWithContext(ctx context.Context) *RTRDeleteScriptsParams {
	return &RTRDeleteScriptsParams{
		Context: ctx,
	}
}

// NewRTRDeleteScriptsParamsWithHTTPClient creates a new RTRDeleteScriptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRDeleteScriptsParamsWithHTTPClient(client *http.Client) *RTRDeleteScriptsParams {
	return &RTRDeleteScriptsParams{
		HTTPClient: client,
	}
}

/* RTRDeleteScriptsParams contains all the parameters to send to the API endpoint
   for the r t r delete scripts operation.

   Typically these are written to a http.Request.
*/
type RTRDeleteScriptsParams struct {

	/* Ids.

	   File id
	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r delete scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRDeleteScriptsParams) WithDefaults() *RTRDeleteScriptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r delete scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRDeleteScriptsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) WithTimeout(timeout time.Duration) *RTRDeleteScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) WithContext(ctx context.Context) *RTRDeleteScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) WithHTTPClient(client *http.Client) *RTRDeleteScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) WithIds(ids string) *RTRDeleteScriptsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the r t r delete scripts params
func (o *RTRDeleteScriptsParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *RTRDeleteScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {

		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

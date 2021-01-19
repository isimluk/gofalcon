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
	"github.com/go-openapi/swag"
)

// NewRTRGetPutFilesParams creates a new RTRGetPutFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRGetPutFilesParams() *RTRGetPutFilesParams {
	return &RTRGetPutFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRGetPutFilesParamsWithTimeout creates a new RTRGetPutFilesParams object
// with the ability to set a timeout on a request.
func NewRTRGetPutFilesParamsWithTimeout(timeout time.Duration) *RTRGetPutFilesParams {
	return &RTRGetPutFilesParams{
		timeout: timeout,
	}
}

// NewRTRGetPutFilesParamsWithContext creates a new RTRGetPutFilesParams object
// with the ability to set a context for a request.
func NewRTRGetPutFilesParamsWithContext(ctx context.Context) *RTRGetPutFilesParams {
	return &RTRGetPutFilesParams{
		Context: ctx,
	}
}

// NewRTRGetPutFilesParamsWithHTTPClient creates a new RTRGetPutFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRGetPutFilesParamsWithHTTPClient(client *http.Client) *RTRGetPutFilesParams {
	return &RTRGetPutFilesParams{
		HTTPClient: client,
	}
}

/* RTRGetPutFilesParams contains all the parameters to send to the API endpoint
   for the r t r get put files operation.

   Typically these are written to a http.Request.
*/
type RTRGetPutFilesParams struct {

	/* Ids.

	   File IDs
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r get put files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRGetPutFilesParams) WithDefaults() *RTRGetPutFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r get put files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRGetPutFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r get put files params
func (o *RTRGetPutFilesParams) WithTimeout(timeout time.Duration) *RTRGetPutFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r get put files params
func (o *RTRGetPutFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r get put files params
func (o *RTRGetPutFilesParams) WithContext(ctx context.Context) *RTRGetPutFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r get put files params
func (o *RTRGetPutFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r get put files params
func (o *RTRGetPutFilesParams) WithHTTPClient(client *http.Client) *RTRGetPutFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r get put files params
func (o *RTRGetPutFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the r t r get put files params
func (o *RTRGetPutFilesParams) WithIds(ids []string) *RTRGetPutFilesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the r t r get put files params
func (o *RTRGetPutFilesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *RTRGetPutFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRTRGetPutFiles binds the parameter ids
func (o *RTRGetPutFilesParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}

// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

// NewQueryMLExclusionsV1Params creates a new QueryMLExclusionsV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryMLExclusionsV1Params() *QueryMLExclusionsV1Params {
	return &QueryMLExclusionsV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryMLExclusionsV1ParamsWithTimeout creates a new QueryMLExclusionsV1Params object
// with the ability to set a timeout on a request.
func NewQueryMLExclusionsV1ParamsWithTimeout(timeout time.Duration) *QueryMLExclusionsV1Params {
	return &QueryMLExclusionsV1Params{
		timeout: timeout,
	}
}

// NewQueryMLExclusionsV1ParamsWithContext creates a new QueryMLExclusionsV1Params object
// with the ability to set a context for a request.
func NewQueryMLExclusionsV1ParamsWithContext(ctx context.Context) *QueryMLExclusionsV1Params {
	return &QueryMLExclusionsV1Params{
		Context: ctx,
	}
}

// NewQueryMLExclusionsV1ParamsWithHTTPClient creates a new QueryMLExclusionsV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewQueryMLExclusionsV1ParamsWithHTTPClient(client *http.Client) *QueryMLExclusionsV1Params {
	return &QueryMLExclusionsV1Params{
		HTTPClient: client,
	}
}

/* QueryMLExclusionsV1Params contains all the parameters to send to the API endpoint
   for the query m l exclusions v1 operation.

   Typically these are written to a http.Request.
*/
type QueryMLExclusionsV1Params struct {

	/* Filter.

	   The filter expression that should be used to limit the results.
	*/
	Filter *string

	/* Limit.

	   The maximum records to return. [1-500]
	*/
	Limit *int64

	/* Offset.

	   The offset to start retrieving records from
	*/
	Offset *int64

	/* Sort.

	   The sort expression that should be used to sort the results.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMLExclusionsV1Params) WithDefaults() *QueryMLExclusionsV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query m l exclusions v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryMLExclusionsV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithTimeout(timeout time.Duration) *QueryMLExclusionsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithContext(ctx context.Context) *QueryMLExclusionsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithHTTPClient(client *http.Client) *QueryMLExclusionsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithFilter(filter *string) *QueryMLExclusionsV1Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithLimit(limit *int64) *QueryMLExclusionsV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithOffset(offset *int64) *QueryMLExclusionsV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) WithSort(sort *string) *QueryMLExclusionsV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query m l exclusions v1 params
func (o *QueryMLExclusionsV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryMLExclusionsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

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
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

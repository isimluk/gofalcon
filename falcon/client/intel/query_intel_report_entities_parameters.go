// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// NewQueryIntelReportEntitiesParams creates a new QueryIntelReportEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryIntelReportEntitiesParams() *QueryIntelReportEntitiesParams {
	return &QueryIntelReportEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryIntelReportEntitiesParamsWithTimeout creates a new QueryIntelReportEntitiesParams object
// with the ability to set a timeout on a request.
func NewQueryIntelReportEntitiesParamsWithTimeout(timeout time.Duration) *QueryIntelReportEntitiesParams {
	return &QueryIntelReportEntitiesParams{
		timeout: timeout,
	}
}

// NewQueryIntelReportEntitiesParamsWithContext creates a new QueryIntelReportEntitiesParams object
// with the ability to set a context for a request.
func NewQueryIntelReportEntitiesParamsWithContext(ctx context.Context) *QueryIntelReportEntitiesParams {
	return &QueryIntelReportEntitiesParams{
		Context: ctx,
	}
}

// NewQueryIntelReportEntitiesParamsWithHTTPClient creates a new QueryIntelReportEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryIntelReportEntitiesParamsWithHTTPClient(client *http.Client) *QueryIntelReportEntitiesParams {
	return &QueryIntelReportEntitiesParams{
		HTTPClient: client,
	}
}

/* QueryIntelReportEntitiesParams contains all the parameters to send to the API endpoint
   for the query intel report entities operation.

   Typically these are written to a http.Request.
*/
type QueryIntelReportEntitiesParams struct {

	/* Fields.

	     The fields to return, or a predefined set of fields in the form of the collection name surrounded by two underscores like:

	\_\_\<collection\>\_\_.

	Ex: slug \_\_full\_\_.

	Defaults to \_\_basic\_\_.
	*/
	Fields []string

	/* Filter.

	     Filter your query by specifying FQL filter parameters. Filter parameters include:

	actors, actors.id, actors.name, actors.slug, actors.url, created_date, description, id, last_modified_date, motivations, motivations.id, motivations.slug, motivations.value, name, name.raw, short_description, slug, sub_type, sub_type.id, sub_type.name, sub_type.slug, tags, tags.id, tags.slug, tags.value, target_countries, target_countries.id, target_countries.slug, target_countries.value, target_industries, target_industries.id, target_industries.slug, target_industries.value, type, type.id, type.name, type.slug, url.
	*/
	Filter *string

	/* Limit.

	   Set the number of reports to return. The value must be between 1 and 5000.
	*/
	Limit *int64

	/* Offset.

	   Set the starting row number to return reports from. Defaults to 0.
	*/
	Offset *int64

	/* Q.

	   Perform a generic substring search across all fields.
	*/
	Q *string

	/* Sort.

	   Order fields in ascending or descending order. Ex: created_date|asc.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query intel report entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelReportEntitiesParams) WithDefaults() *QueryIntelReportEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query intel report entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryIntelReportEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithTimeout(timeout time.Duration) *QueryIntelReportEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithContext(ctx context.Context) *QueryIntelReportEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithHTTPClient(client *http.Client) *QueryIntelReportEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithFields(fields []string) *QueryIntelReportEntitiesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilter adds the filter to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithFilter(filter *string) *QueryIntelReportEntitiesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithLimit(limit *int64) *QueryIntelReportEntitiesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithOffset(offset *int64) *QueryIntelReportEntitiesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithQ(q *string) *QueryIntelReportEntitiesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) WithSort(sort *string) *QueryIntelReportEntitiesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query intel report entities params
func (o *QueryIntelReportEntitiesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *QueryIntelReportEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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

// bindParamQueryIntelReportEntities binds the parameter fields
func (o *QueryIntelReportEntitiesParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}

// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewRTRAggregateSessionsParams creates a new RTRAggregateSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRAggregateSessionsParams() *RTRAggregateSessionsParams {
	return &RTRAggregateSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRAggregateSessionsParamsWithTimeout creates a new RTRAggregateSessionsParams object
// with the ability to set a timeout on a request.
func NewRTRAggregateSessionsParamsWithTimeout(timeout time.Duration) *RTRAggregateSessionsParams {
	return &RTRAggregateSessionsParams{
		timeout: timeout,
	}
}

// NewRTRAggregateSessionsParamsWithContext creates a new RTRAggregateSessionsParams object
// with the ability to set a context for a request.
func NewRTRAggregateSessionsParamsWithContext(ctx context.Context) *RTRAggregateSessionsParams {
	return &RTRAggregateSessionsParams{
		Context: ctx,
	}
}

// NewRTRAggregateSessionsParamsWithHTTPClient creates a new RTRAggregateSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRAggregateSessionsParamsWithHTTPClient(client *http.Client) *RTRAggregateSessionsParams {
	return &RTRAggregateSessionsParams{
		HTTPClient: client,
	}
}

/* RTRAggregateSessionsParams contains all the parameters to send to the API endpoint
   for the r t r aggregate sessions operation.

   Typically these are written to a http.Request.
*/
type RTRAggregateSessionsParams struct {

	/* Body.

	     Supported aggregations:
	- `term`
	- `date_range`

	Supported aggregation members:

	**`date_ranges`** If peforming a date range query specify the **`from`** and **`to`** date ranges.  These can be in common date formats like `2019-07-18` or `now`
	**`field`** Term you want to aggregate on.  If doing a `date_range` query, this is the date field you want to apply the date ranges to
	**`filter`** Optional filter criteria in the form of an FQL query. For more information about FQL queries, see our [FQL documentation in Falcon](https://falcon.crowdstrike.com/support/documentation/45/falcon-query-language-feature-guide).
	**`name`** Name of the aggregation
	**`size`** Size limit to apply to the queries.
	*/
	Body []*models.MsaAggregateQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r aggregate sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRAggregateSessionsParams) WithDefaults() *RTRAggregateSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r aggregate sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRAggregateSessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) WithTimeout(timeout time.Duration) *RTRAggregateSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) WithContext(ctx context.Context) *RTRAggregateSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) WithHTTPClient(client *http.Client) *RTRAggregateSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) WithBody(body []*models.MsaAggregateQueryRequest) *RTRAggregateSessionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the r t r aggregate sessions params
func (o *RTRAggregateSessionsParams) SetBody(body []*models.MsaAggregateQueryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RTRAggregateSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

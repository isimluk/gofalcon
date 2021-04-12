// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// NewGetChildrenParams creates a new GetChildrenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetChildrenParams() *GetChildrenParams {
	return &GetChildrenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildrenParamsWithTimeout creates a new GetChildrenParams object
// with the ability to set a timeout on a request.
func NewGetChildrenParamsWithTimeout(timeout time.Duration) *GetChildrenParams {
	return &GetChildrenParams{
		timeout: timeout,
	}
}

// NewGetChildrenParamsWithContext creates a new GetChildrenParams object
// with the ability to set a context for a request.
func NewGetChildrenParamsWithContext(ctx context.Context) *GetChildrenParams {
	return &GetChildrenParams{
		Context: ctx,
	}
}

// NewGetChildrenParamsWithHTTPClient creates a new GetChildrenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetChildrenParamsWithHTTPClient(client *http.Client) *GetChildrenParams {
	return &GetChildrenParams{
		HTTPClient: client,
	}
}

/* GetChildrenParams contains all the parameters to send to the API endpoint
   for the get children operation.

   Typically these are written to a http.Request.
*/
type GetChildrenParams struct {

	/* Ids.

	   CID of a child customer
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChildrenParams) WithDefaults() *GetChildrenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get children params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetChildrenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get children params
func (o *GetChildrenParams) WithTimeout(timeout time.Duration) *GetChildrenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get children params
func (o *GetChildrenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get children params
func (o *GetChildrenParams) WithContext(ctx context.Context) *GetChildrenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get children params
func (o *GetChildrenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get children params
func (o *GetChildrenParams) WithHTTPClient(client *http.Client) *GetChildrenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get children params
func (o *GetChildrenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get children params
func (o *GetChildrenParams) WithIds(ids []string) *GetChildrenParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get children params
func (o *GetChildrenParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildrenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetChildren binds the parameter ids
func (o *GetChildrenParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "csv"
	idsIS := swag.JoinByFormat(idsIC, "csv")

	return idsIS
}

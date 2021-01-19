// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// NewGetAvailableRoleIdsParams creates a new GetAvailableRoleIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAvailableRoleIdsParams() *GetAvailableRoleIdsParams {
	return &GetAvailableRoleIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableRoleIdsParamsWithTimeout creates a new GetAvailableRoleIdsParams object
// with the ability to set a timeout on a request.
func NewGetAvailableRoleIdsParamsWithTimeout(timeout time.Duration) *GetAvailableRoleIdsParams {
	return &GetAvailableRoleIdsParams{
		timeout: timeout,
	}
}

// NewGetAvailableRoleIdsParamsWithContext creates a new GetAvailableRoleIdsParams object
// with the ability to set a context for a request.
func NewGetAvailableRoleIdsParamsWithContext(ctx context.Context) *GetAvailableRoleIdsParams {
	return &GetAvailableRoleIdsParams{
		Context: ctx,
	}
}

// NewGetAvailableRoleIdsParamsWithHTTPClient creates a new GetAvailableRoleIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAvailableRoleIdsParamsWithHTTPClient(client *http.Client) *GetAvailableRoleIdsParams {
	return &GetAvailableRoleIdsParams{
		HTTPClient: client,
	}
}

/* GetAvailableRoleIdsParams contains all the parameters to send to the API endpoint
   for the get available role ids operation.

   Typically these are written to a http.Request.
*/
type GetAvailableRoleIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get available role ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableRoleIdsParams) WithDefaults() *GetAvailableRoleIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get available role ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAvailableRoleIdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get available role ids params
func (o *GetAvailableRoleIdsParams) WithTimeout(timeout time.Duration) *GetAvailableRoleIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available role ids params
func (o *GetAvailableRoleIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available role ids params
func (o *GetAvailableRoleIdsParams) WithContext(ctx context.Context) *GetAvailableRoleIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available role ids params
func (o *GetAvailableRoleIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available role ids params
func (o *GetAvailableRoleIdsParams) WithHTTPClient(client *http.Client) *GetAvailableRoleIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available role ids params
func (o *GetAvailableRoleIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableRoleIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

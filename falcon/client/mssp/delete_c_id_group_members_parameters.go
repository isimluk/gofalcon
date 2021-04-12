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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewDeleteCIDGroupMembersParams creates a new DeleteCIDGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCIDGroupMembersParams() *DeleteCIDGroupMembersParams {
	return &DeleteCIDGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCIDGroupMembersParamsWithTimeout creates a new DeleteCIDGroupMembersParams object
// with the ability to set a timeout on a request.
func NewDeleteCIDGroupMembersParamsWithTimeout(timeout time.Duration) *DeleteCIDGroupMembersParams {
	return &DeleteCIDGroupMembersParams{
		timeout: timeout,
	}
}

// NewDeleteCIDGroupMembersParamsWithContext creates a new DeleteCIDGroupMembersParams object
// with the ability to set a context for a request.
func NewDeleteCIDGroupMembersParamsWithContext(ctx context.Context) *DeleteCIDGroupMembersParams {
	return &DeleteCIDGroupMembersParams{
		Context: ctx,
	}
}

// NewDeleteCIDGroupMembersParamsWithHTTPClient creates a new DeleteCIDGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCIDGroupMembersParamsWithHTTPClient(client *http.Client) *DeleteCIDGroupMembersParams {
	return &DeleteCIDGroupMembersParams{
		HTTPClient: client,
	}
}

/* DeleteCIDGroupMembersParams contains all the parameters to send to the API endpoint
   for the delete c ID group members operation.

   Typically these are written to a http.Request.
*/
type DeleteCIDGroupMembersParams struct {

	// Body.
	Body *models.DomainCIDGroupMembersRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete c ID group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCIDGroupMembersParams) WithDefaults() *DeleteCIDGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete c ID group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCIDGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) WithTimeout(timeout time.Duration) *DeleteCIDGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) WithContext(ctx context.Context) *DeleteCIDGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) WithHTTPClient(client *http.Client) *DeleteCIDGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) WithBody(body *models.DomainCIDGroupMembersRequestV1) *DeleteCIDGroupMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete c ID group members params
func (o *DeleteCIDGroupMembersParams) SetBody(body *models.DomainCIDGroupMembersRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCIDGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

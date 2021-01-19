// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

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

// NewUpdateAWSAccountsParams creates a new UpdateAWSAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAWSAccountsParams() *UpdateAWSAccountsParams {
	return &UpdateAWSAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAWSAccountsParamsWithTimeout creates a new UpdateAWSAccountsParams object
// with the ability to set a timeout on a request.
func NewUpdateAWSAccountsParamsWithTimeout(timeout time.Duration) *UpdateAWSAccountsParams {
	return &UpdateAWSAccountsParams{
		timeout: timeout,
	}
}

// NewUpdateAWSAccountsParamsWithContext creates a new UpdateAWSAccountsParams object
// with the ability to set a context for a request.
func NewUpdateAWSAccountsParamsWithContext(ctx context.Context) *UpdateAWSAccountsParams {
	return &UpdateAWSAccountsParams{
		Context: ctx,
	}
}

// NewUpdateAWSAccountsParamsWithHTTPClient creates a new UpdateAWSAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAWSAccountsParamsWithHTTPClient(client *http.Client) *UpdateAWSAccountsParams {
	return &UpdateAWSAccountsParams{
		HTTPClient: client,
	}
}

/* UpdateAWSAccountsParams contains all the parameters to send to the API endpoint
   for the update a w s accounts operation.

   Typically these are written to a http.Request.
*/
type UpdateAWSAccountsParams struct {

	// Body.
	Body *models.ModelsUpdateAWSAccountsV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update a w s accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAWSAccountsParams) WithDefaults() *UpdateAWSAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update a w s accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAWSAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update a w s accounts params
func (o *UpdateAWSAccountsParams) WithTimeout(timeout time.Duration) *UpdateAWSAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update a w s accounts params
func (o *UpdateAWSAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update a w s accounts params
func (o *UpdateAWSAccountsParams) WithContext(ctx context.Context) *UpdateAWSAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update a w s accounts params
func (o *UpdateAWSAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update a w s accounts params
func (o *UpdateAWSAccountsParams) WithHTTPClient(client *http.Client) *UpdateAWSAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update a w s accounts params
func (o *UpdateAWSAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update a w s accounts params
func (o *UpdateAWSAccountsParams) WithBody(body *models.ModelsUpdateAWSAccountsV1) *UpdateAWSAccountsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update a w s accounts params
func (o *UpdateAWSAccountsParams) SetBody(body *models.ModelsUpdateAWSAccountsV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAWSAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

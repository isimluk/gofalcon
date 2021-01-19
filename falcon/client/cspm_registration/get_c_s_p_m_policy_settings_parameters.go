// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// NewGetCSPMPolicySettingsParams creates a new GetCSPMPolicySettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCSPMPolicySettingsParams() *GetCSPMPolicySettingsParams {
	return &GetCSPMPolicySettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCSPMPolicySettingsParamsWithTimeout creates a new GetCSPMPolicySettingsParams object
// with the ability to set a timeout on a request.
func NewGetCSPMPolicySettingsParamsWithTimeout(timeout time.Duration) *GetCSPMPolicySettingsParams {
	return &GetCSPMPolicySettingsParams{
		timeout: timeout,
	}
}

// NewGetCSPMPolicySettingsParamsWithContext creates a new GetCSPMPolicySettingsParams object
// with the ability to set a context for a request.
func NewGetCSPMPolicySettingsParamsWithContext(ctx context.Context) *GetCSPMPolicySettingsParams {
	return &GetCSPMPolicySettingsParams{
		Context: ctx,
	}
}

// NewGetCSPMPolicySettingsParamsWithHTTPClient creates a new GetCSPMPolicySettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCSPMPolicySettingsParamsWithHTTPClient(client *http.Client) *GetCSPMPolicySettingsParams {
	return &GetCSPMPolicySettingsParams{
		HTTPClient: client,
	}
}

/* GetCSPMPolicySettingsParams contains all the parameters to send to the API endpoint
   for the get c s p m policy settings operation.

   Typically these are written to a http.Request.
*/
type GetCSPMPolicySettingsParams struct {

	/* PolicyID.

	   Policy ID
	*/
	PolicyID *string

	/* Service.

	   Service type to filter policy settings by.
	*/
	Service string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get c s p m policy settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMPolicySettingsParams) WithDefaults() *GetCSPMPolicySettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get c s p m policy settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCSPMPolicySettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) WithTimeout(timeout time.Duration) *GetCSPMPolicySettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) WithContext(ctx context.Context) *GetCSPMPolicySettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) WithHTTPClient(client *http.Client) *GetCSPMPolicySettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyID adds the policyID to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) WithPolicyID(policyID *string) *GetCSPMPolicySettingsParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) SetPolicyID(policyID *string) {
	o.PolicyID = policyID
}

// WithService adds the service to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) WithService(service string) *GetCSPMPolicySettingsParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the get c s p m policy settings params
func (o *GetCSPMPolicySettingsParams) SetService(service string) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *GetCSPMPolicySettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PolicyID != nil {

		// query param policy-id
		var qrPolicyID string

		if o.PolicyID != nil {
			qrPolicyID = *o.PolicyID
		}
		qPolicyID := qrPolicyID
		if qPolicyID != "" {

			if err := r.SetQueryParam("policy-id", qPolicyID); err != nil {
				return err
			}
		}
	}

	// query param service
	qrService := o.Service
	qService := qrService
	if qService != "" {

		if err := r.SetQueryParam("service", qService); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

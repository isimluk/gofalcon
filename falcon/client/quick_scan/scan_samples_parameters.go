// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

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

// NewScanSamplesParams creates a new ScanSamplesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScanSamplesParams() *ScanSamplesParams {
	return &ScanSamplesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScanSamplesParamsWithTimeout creates a new ScanSamplesParams object
// with the ability to set a timeout on a request.
func NewScanSamplesParamsWithTimeout(timeout time.Duration) *ScanSamplesParams {
	return &ScanSamplesParams{
		timeout: timeout,
	}
}

// NewScanSamplesParamsWithContext creates a new ScanSamplesParams object
// with the ability to set a context for a request.
func NewScanSamplesParamsWithContext(ctx context.Context) *ScanSamplesParams {
	return &ScanSamplesParams{
		Context: ctx,
	}
}

// NewScanSamplesParamsWithHTTPClient creates a new ScanSamplesParams object
// with the ability to set a custom HTTPClient for a request.
func NewScanSamplesParamsWithHTTPClient(client *http.Client) *ScanSamplesParams {
	return &ScanSamplesParams{
		HTTPClient: client,
	}
}

/* ScanSamplesParams contains all the parameters to send to the API endpoint
   for the scan samples operation.

   Typically these are written to a http.Request.
*/
type ScanSamplesParams struct {

	/* Body.

	   Submit a batch of SHA256s for ml scanning. The samples must have been previously uploaded through `/samples/entities/samples/v3`
	*/
	Body *models.MlscannerSamplesScanParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the scan samples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanSamplesParams) WithDefaults() *ScanSamplesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the scan samples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScanSamplesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the scan samples params
func (o *ScanSamplesParams) WithTimeout(timeout time.Duration) *ScanSamplesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scan samples params
func (o *ScanSamplesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scan samples params
func (o *ScanSamplesParams) WithContext(ctx context.Context) *ScanSamplesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scan samples params
func (o *ScanSamplesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scan samples params
func (o *ScanSamplesParams) WithHTTPClient(client *http.Client) *ScanSamplesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scan samples params
func (o *ScanSamplesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the scan samples params
func (o *ScanSamplesParams) WithBody(body *models.MlscannerSamplesScanParameters) *ScanSamplesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the scan samples params
func (o *ScanSamplesParams) SetBody(body *models.MlscannerSamplesScanParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ScanSamplesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

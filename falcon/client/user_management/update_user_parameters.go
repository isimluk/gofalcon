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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateUserParams creates a new UpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserParams() *UpdateUserParams {
	return &UpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserParamsWithTimeout creates a new UpdateUserParams object
// with the ability to set a timeout on a request.
func NewUpdateUserParamsWithTimeout(timeout time.Duration) *UpdateUserParams {
	return &UpdateUserParams{
		timeout: timeout,
	}
}

// NewUpdateUserParamsWithContext creates a new UpdateUserParams object
// with the ability to set a context for a request.
func NewUpdateUserParamsWithContext(ctx context.Context) *UpdateUserParams {
	return &UpdateUserParams{
		Context: ctx,
	}
}

// NewUpdateUserParamsWithHTTPClient creates a new UpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserParamsWithHTTPClient(client *http.Client) *UpdateUserParams {
	return &UpdateUserParams{
		HTTPClient: client,
	}
}

/* UpdateUserParams contains all the parameters to send to the API endpoint
   for the update user operation.

   Typically these are written to a http.Request.
*/
type UpdateUserParams struct {

	/* Body.

	   Attributes for this user. All attributes (shown below) are optional.
	*/
	Body *models.DomainUpdateUserFields

	/* UserUUID.

	   ID of a user. Find a user's ID from `/users/entities/user/v1`.
	*/
	UserUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserParams) WithDefaults() *UpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user params
func (o *UpdateUserParams) WithTimeout(timeout time.Duration) *UpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user params
func (o *UpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user params
func (o *UpdateUserParams) WithContext(ctx context.Context) *UpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user params
func (o *UpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) WithHTTPClient(client *http.Client) *UpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user params
func (o *UpdateUserParams) WithBody(body *models.DomainUpdateUserFields) *UpdateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user params
func (o *UpdateUserParams) SetBody(body *models.DomainUpdateUserFields) {
	o.Body = body
}

// WithUserUUID adds the userUUID to the update user params
func (o *UpdateUserParams) WithUserUUID(userUUID string) *UpdateUserParams {
	o.SetUserUUID(userUUID)
	return o
}

// SetUserUUID adds the userUuid to the update user params
func (o *UpdateUserParams) SetUserUUID(userUUID string) {
	o.UserUUID = userUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param user_uuid
	qrUserUUID := o.UserUUID
	qUserUUID := qrUserUUID
	if qUserUUID != "" {

		if err := r.SetQueryParam("user_uuid", qUserUUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

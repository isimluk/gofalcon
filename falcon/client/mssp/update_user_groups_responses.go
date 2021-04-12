// Code generated by go-swagger; DO NOT EDIT.

package mssp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateUserGroupsReader is a Reader for the UpdateUserGroups structure.
type UpdateUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewUpdateUserGroupsMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateUserGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateUserGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserGroupsOK creates a UpdateUserGroupsOK with default headers values
func NewUpdateUserGroupsOK() *UpdateUserGroupsOK {
	return &UpdateUserGroupsOK{}
}

/* UpdateUserGroupsOK describes a response with status code 200, with default header values.

OK
*/
type UpdateUserGroupsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupsResponseV1
}

func (o *UpdateUserGroupsOK) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroupsOK  %+v", 200, o.Payload)
}
func (o *UpdateUserGroupsOK) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupsMultiStatus creates a UpdateUserGroupsMultiStatus with default headers values
func NewUpdateUserGroupsMultiStatus() *UpdateUserGroupsMultiStatus {
	return &UpdateUserGroupsMultiStatus{}
}

/* UpdateUserGroupsMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type UpdateUserGroupsMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserGroupsResponseV1
}

func (o *UpdateUserGroupsMultiStatus) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroupsMultiStatus  %+v", 207, o.Payload)
}
func (o *UpdateUserGroupsMultiStatus) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateUserGroupsMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupsBadRequest creates a UpdateUserGroupsBadRequest with default headers values
func NewUpdateUserGroupsBadRequest() *UpdateUserGroupsBadRequest {
	return &UpdateUserGroupsBadRequest{}
}

/* UpdateUserGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateUserGroupsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateUserGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateUserGroupsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateUserGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupsForbidden creates a UpdateUserGroupsForbidden with default headers values
func NewUpdateUserGroupsForbidden() *UpdateUserGroupsForbidden {
	return &UpdateUserGroupsForbidden{}
}

/* UpdateUserGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUserGroupsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateUserGroupsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroupsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserGroupsForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateUserGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupsTooManyRequests creates a UpdateUserGroupsTooManyRequests with default headers values
func NewUpdateUserGroupsTooManyRequests() *UpdateUserGroupsTooManyRequests {
	return &UpdateUserGroupsTooManyRequests{}
}

/* UpdateUserGroupsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateUserGroupsTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateUserGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroupsTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateUserGroupsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateUserGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserGroupsDefault creates a UpdateUserGroupsDefault with default headers values
func NewUpdateUserGroupsDefault(code int) *UpdateUserGroupsDefault {
	return &UpdateUserGroupsDefault{
		_statusCode: code,
	}
}

/* UpdateUserGroupsDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateUserGroupsDefault struct {
	_statusCode int

	Payload *models.DomainUserGroupsResponseV1
}

// Code gets the status code for the update user groups default response
func (o *UpdateUserGroupsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserGroupsDefault) Error() string {
	return fmt.Sprintf("[PATCH /mssp/entities/user-groups/v1][%d] updateUserGroups default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateUserGroupsDefault) GetPayload() *models.DomainUserGroupsResponseV1 {
	return o.Payload
}

func (o *UpdateUserGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainUserGroupsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

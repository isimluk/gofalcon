// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// DeletePreventionPoliciesReader is a Reader for the DeletePreventionPolicies structure.
type DeletePreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeletePreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePreventionPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePreventionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePreventionPoliciesOK creates a DeletePreventionPoliciesOK with default headers values
func NewDeletePreventionPoliciesOK() *DeletePreventionPoliciesOK {
	return &DeletePreventionPoliciesOK{}
}

/* DeletePreventionPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type DeletePreventionPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeletePreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPoliciesOK  %+v", 200, o.Payload)
}
func (o *DeletePreventionPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeletePreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePreventionPoliciesForbidden creates a DeletePreventionPoliciesForbidden with default headers values
func NewDeletePreventionPoliciesForbidden() *DeletePreventionPoliciesForbidden {
	return &DeletePreventionPoliciesForbidden{}
}

/* DeletePreventionPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletePreventionPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *DeletePreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPoliciesForbidden  %+v", 403, o.Payload)
}
func (o *DeletePreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeletePreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePreventionPoliciesNotFound creates a DeletePreventionPoliciesNotFound with default headers values
func NewDeletePreventionPoliciesNotFound() *DeletePreventionPoliciesNotFound {
	return &DeletePreventionPoliciesNotFound{}
}

/* DeletePreventionPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeletePreventionPoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeletePreventionPoliciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPoliciesNotFound  %+v", 404, o.Payload)
}
func (o *DeletePreventionPoliciesNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeletePreventionPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePreventionPoliciesTooManyRequests creates a DeletePreventionPoliciesTooManyRequests with default headers values
func NewDeletePreventionPoliciesTooManyRequests() *DeletePreventionPoliciesTooManyRequests {
	return &DeletePreventionPoliciesTooManyRequests{}
}

/* DeletePreventionPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeletePreventionPoliciesTooManyRequests struct {

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

func (o *DeletePreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeletePreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeletePreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePreventionPoliciesInternalServerError creates a DeletePreventionPoliciesInternalServerError with default headers values
func NewDeletePreventionPoliciesInternalServerError() *DeletePreventionPoliciesInternalServerError {
	return &DeletePreventionPoliciesInternalServerError{}
}

/* DeletePreventionPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeletePreventionPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeletePreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *DeletePreventionPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeletePreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePreventionPoliciesDefault creates a DeletePreventionPoliciesDefault with default headers values
func NewDeletePreventionPoliciesDefault(code int) *DeletePreventionPoliciesDefault {
	return &DeletePreventionPoliciesDefault{
		_statusCode: code,
	}
}

/* DeletePreventionPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type DeletePreventionPoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the delete prevention policies default response
func (o *DeletePreventionPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *DeletePreventionPoliciesDefault) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/prevention/v1][%d] deletePreventionPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *DeletePreventionPoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeletePreventionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package spotlight_vulnerabilities

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

// QueryVulnerabilitiesReader is a Reader for the QueryVulnerabilities structure.
type QueryVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryVulnerabilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryVulnerabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryVulnerabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryVulnerabilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryVulnerabilitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryVulnerabilitiesOK creates a QueryVulnerabilitiesOK with default headers values
func NewQueryVulnerabilitiesOK() *QueryVulnerabilitiesOK {
	return &QueryVulnerabilitiesOK{}
}

/* QueryVulnerabilitiesOK describes a response with status code 200, with default header values.

OK
*/
type QueryVulnerabilitiesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIQueryVulnerabilitiesResponse
}

func (o *QueryVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilitiesOK  %+v", 200, o.Payload)
}
func (o *QueryVulnerabilitiesOK) GetPayload() *models.DomainSPAPIQueryVulnerabilitiesResponse {
	return o.Payload
}

func (o *QueryVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIQueryVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVulnerabilitiesBadRequest creates a QueryVulnerabilitiesBadRequest with default headers values
func NewQueryVulnerabilitiesBadRequest() *QueryVulnerabilitiesBadRequest {
	return &QueryVulnerabilitiesBadRequest{}
}

/* QueryVulnerabilitiesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryVulnerabilitiesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIQueryVulnerabilitiesResponse
}

func (o *QueryVulnerabilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilitiesBadRequest  %+v", 400, o.Payload)
}
func (o *QueryVulnerabilitiesBadRequest) GetPayload() *models.DomainSPAPIQueryVulnerabilitiesResponse {
	return o.Payload
}

func (o *QueryVulnerabilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIQueryVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVulnerabilitiesForbidden creates a QueryVulnerabilitiesForbidden with default headers values
func NewQueryVulnerabilitiesForbidden() *QueryVulnerabilitiesForbidden {
	return &QueryVulnerabilitiesForbidden{}
}

/* QueryVulnerabilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryVulnerabilitiesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryVulnerabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilitiesForbidden  %+v", 403, o.Payload)
}
func (o *QueryVulnerabilitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryVulnerabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVulnerabilitiesTooManyRequests creates a QueryVulnerabilitiesTooManyRequests with default headers values
func NewQueryVulnerabilitiesTooManyRequests() *QueryVulnerabilitiesTooManyRequests {
	return &QueryVulnerabilitiesTooManyRequests{}
}

/* QueryVulnerabilitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryVulnerabilitiesTooManyRequests struct {

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

func (o *QueryVulnerabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilitiesTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryVulnerabilitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryVulnerabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryVulnerabilitiesInternalServerError creates a QueryVulnerabilitiesInternalServerError with default headers values
func NewQueryVulnerabilitiesInternalServerError() *QueryVulnerabilitiesInternalServerError {
	return &QueryVulnerabilitiesInternalServerError{}
}

/* QueryVulnerabilitiesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryVulnerabilitiesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainSPAPIQueryVulnerabilitiesResponse
}

func (o *QueryVulnerabilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilitiesInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryVulnerabilitiesInternalServerError) GetPayload() *models.DomainSPAPIQueryVulnerabilitiesResponse {
	return o.Payload
}

func (o *QueryVulnerabilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainSPAPIQueryVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVulnerabilitiesDefault creates a QueryVulnerabilitiesDefault with default headers values
func NewQueryVulnerabilitiesDefault(code int) *QueryVulnerabilitiesDefault {
	return &QueryVulnerabilitiesDefault{
		_statusCode: code,
	}
}

/* QueryVulnerabilitiesDefault describes a response with status code -1, with default header values.

OK
*/
type QueryVulnerabilitiesDefault struct {
	_statusCode int

	Payload *models.DomainSPAPIQueryVulnerabilitiesResponse
}

// Code gets the status code for the query vulnerabilities default response
func (o *QueryVulnerabilitiesDefault) Code() int {
	return o._statusCode
}

func (o *QueryVulnerabilitiesDefault) Error() string {
	return fmt.Sprintf("[GET /spotlight/queries/vulnerabilities/v1][%d] queryVulnerabilities default  %+v", o._statusCode, o.Payload)
}
func (o *QueryVulnerabilitiesDefault) GetPayload() *models.DomainSPAPIQueryVulnerabilitiesResponse {
	return o.Payload
}

func (o *QueryVulnerabilitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainSPAPIQueryVulnerabilitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// GetSensorInstallersCCIDByQueryReader is a Reader for the GetSensorInstallersCCIDByQuery structure.
type GetSensorInstallersCCIDByQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorInstallersCCIDByQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorInstallersCCIDByQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSensorInstallersCCIDByQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSensorInstallersCCIDByQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorInstallersCCIDByQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSensorInstallersCCIDByQueryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSensorInstallersCCIDByQueryOK creates a GetSensorInstallersCCIDByQueryOK with default headers values
func NewGetSensorInstallersCCIDByQueryOK() *GetSensorInstallersCCIDByQueryOK {
	return &GetSensorInstallersCCIDByQueryOK{}
}

/* GetSensorInstallersCCIDByQueryOK describes a response with status code 200, with default header values.

OK
*/
type GetSensorInstallersCCIDByQueryOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *GetSensorInstallersCCIDByQueryOK) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryOK  %+v", 200, o.Payload)
}
func (o *GetSensorInstallersCCIDByQueryOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorInstallersCCIDByQueryBadRequest creates a GetSensorInstallersCCIDByQueryBadRequest with default headers values
func NewGetSensorInstallersCCIDByQueryBadRequest() *GetSensorInstallersCCIDByQueryBadRequest {
	return &GetSensorInstallersCCIDByQueryBadRequest{}
}

/* GetSensorInstallersCCIDByQueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSensorInstallersCCIDByQueryBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryBadRequest  %+v", 400, o.Payload)
}
func (o *GetSensorInstallersCCIDByQueryBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorInstallersCCIDByQueryForbidden creates a GetSensorInstallersCCIDByQueryForbidden with default headers values
func NewGetSensorInstallersCCIDByQueryForbidden() *GetSensorInstallersCCIDByQueryForbidden {
	return &GetSensorInstallersCCIDByQueryForbidden{}
}

/* GetSensorInstallersCCIDByQueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorInstallersCCIDByQueryForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetSensorInstallersCCIDByQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryForbidden  %+v", 403, o.Payload)
}
func (o *GetSensorInstallersCCIDByQueryForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorInstallersCCIDByQueryTooManyRequests creates a GetSensorInstallersCCIDByQueryTooManyRequests with default headers values
func NewGetSensorInstallersCCIDByQueryTooManyRequests() *GetSensorInstallersCCIDByQueryTooManyRequests {
	return &GetSensorInstallersCCIDByQueryTooManyRequests{}
}

/* GetSensorInstallersCCIDByQueryTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorInstallersCCIDByQueryTooManyRequests struct {

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

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] getSensorInstallersCCIdByQueryTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetSensorInstallersCCIDByQueryTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorInstallersCCIDByQueryDefault creates a GetSensorInstallersCCIDByQueryDefault with default headers values
func NewGetSensorInstallersCCIDByQueryDefault(code int) *GetSensorInstallersCCIDByQueryDefault {
	return &GetSensorInstallersCCIDByQueryDefault{
		_statusCode: code,
	}
}

/* GetSensorInstallersCCIDByQueryDefault describes a response with status code -1, with default header values.

OK
*/
type GetSensorInstallersCCIDByQueryDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the get sensor installers c c ID by query default response
func (o *GetSensorInstallersCCIDByQueryDefault) Code() int {
	return o._statusCode
}

func (o *GetSensorInstallersCCIDByQueryDefault) Error() string {
	return fmt.Sprintf("[GET /sensors/queries/installers/ccid/v1][%d] GetSensorInstallersCCIDByQuery default  %+v", o._statusCode, o.Payload)
}
func (o *GetSensorInstallersCCIDByQueryDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetSensorInstallersCCIDByQueryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

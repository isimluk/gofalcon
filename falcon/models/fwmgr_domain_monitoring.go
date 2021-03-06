// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FwmgrDomainMonitoring fwmgr domain monitoring
//
// swagger:model fwmgr.domain.Monitoring
type FwmgrDomainMonitoring struct {

	// count
	// Required: true
	Count *string `json:"count"`

	// period ms
	// Required: true
	PeriodMs *string `json:"period_ms"`
}

// Validate validates this fwmgr domain monitoring
func (m *FwmgrDomainMonitoring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodMs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainMonitoring) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrDomainMonitoring) validatePeriodMs(formats strfmt.Registry) error {

	if err := validate.Required("period_ms", "body", m.PeriodMs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr domain monitoring based on context it is used
func (m *FwmgrDomainMonitoring) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrDomainMonitoring) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrDomainMonitoring) UnmarshalBinary(b []byte) error {
	var res FwmgrDomainMonitoring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

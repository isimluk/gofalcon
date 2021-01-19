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

// DomainBatchExecuteCommandRequest domain batch execute command request
//
// swagger:model domain.BatchExecuteCommandRequest
type DomainBatchExecuteCommandRequest struct {

	// base command
	// Required: true
	BaseCommand *string `json:"base_command"`

	// batch id
	// Required: true
	BatchID *string `json:"batch_id"`

	// command string
	// Required: true
	CommandString *string `json:"command_string"`

	// optional hosts
	// Required: true
	OptionalHosts []string `json:"optional_hosts"`

	// persist all
	// Required: true
	PersistAll *bool `json:"persist_all"`
}

// Validate validates this domain batch execute command request
func (m *DomainBatchExecuteCommandRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBatchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommandString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptionalHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistAll(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainBatchExecuteCommandRequest) validateBaseCommand(formats strfmt.Registry) error {

	if err := validate.Required("base_command", "body", m.BaseCommand); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchExecuteCommandRequest) validateBatchID(formats strfmt.Registry) error {

	if err := validate.Required("batch_id", "body", m.BatchID); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchExecuteCommandRequest) validateCommandString(formats strfmt.Registry) error {

	if err := validate.Required("command_string", "body", m.CommandString); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchExecuteCommandRequest) validateOptionalHosts(formats strfmt.Registry) error {

	if err := validate.Required("optional_hosts", "body", m.OptionalHosts); err != nil {
		return err
	}

	return nil
}

func (m *DomainBatchExecuteCommandRequest) validatePersistAll(formats strfmt.Registry) error {

	if err := validate.Required("persist_all", "body", m.PersistAll); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain batch execute command request based on context it is used
func (m *DomainBatchExecuteCommandRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainBatchExecuteCommandRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainBatchExecuteCommandRequest) UnmarshalBinary(b []byte) error {
	var res DomainBatchExecuteCommandRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

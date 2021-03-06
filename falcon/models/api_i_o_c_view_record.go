// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIIOCViewRecord api i o c view record
//
// swagger:model api.IOCViewRecord
type APIIOCViewRecord struct {

	// batch id
	BatchID string `json:"batch_id,omitempty"`

	// created by
	CreatedBy string `json:"created_by,omitempty"`

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// expiration days
	ExpirationDays int32 `json:"expiration_days,omitempty"`

	// expiration timestamp
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`

	// modified by
	ModifiedBy string `json:"modified_by,omitempty"`

	// modified timestamp
	ModifiedTimestamp string `json:"modified_timestamp,omitempty"`

	// policy
	Policy string `json:"policy,omitempty"`

	// share level
	ShareLevel string `json:"share_level,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this api i o c view record
func (m *APIIOCViewRecord) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api i o c view record based on context it is used
func (m *APIIOCViewRecord) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIIOCViewRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIIOCViewRecord) UnmarshalBinary(b []byte) error {
	var res APIIOCViewRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

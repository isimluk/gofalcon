// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxRelatedIndicator falconx related indicator
//
// swagger:model falconx.RelatedIndicator
type FalconxRelatedIndicator struct {

	// created timestamp
	CreatedTimestamp string `json:"created_timestamp,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated timestamp
	UpdatedTimestamp string `json:"updated_timestamp,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this falconx related indicator
func (m *FalconxRelatedIndicator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this falconx related indicator based on context it is used
func (m *FalconxRelatedIndicator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxRelatedIndicator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxRelatedIndicator) UnmarshalBinary(b []byte) error {
	var res FalconxRelatedIndicator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxExtractedInterestingString falconx extracted interesting string
//
// swagger:model falconx.ExtractedInterestingString
type FalconxExtractedInterestingString struct {

	// filename
	Filename string `json:"filename,omitempty"`

	// process
	Process string `json:"process,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this falconx extracted interesting string
func (m *FalconxExtractedInterestingString) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this falconx extracted interesting string based on context it is used
func (m *FalconxExtractedInterestingString) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxExtractedInterestingString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxExtractedInterestingString) UnmarshalBinary(b []byte) error {
	var res FalconxExtractedInterestingString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

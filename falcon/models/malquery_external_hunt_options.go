// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MalqueryExternalHuntOptions malquery external hunt options
//
// swagger:model malquery.ExternalHuntOptions
type MalqueryExternalHuntOptions struct {

	// Limit results to files of certain types such as EMAIL, PCAP, PDF, PE32. Full list can be found in the documentation
	FilterFiletypes []string `json:"filter_filetypes"`

	// Specify a subset of metadata fields to return in the results. Possible values: sha256, md5, type, size, first_seen, label, family
	FilterMeta []string `json:"filter_meta"`

	// Maximum number of results to be returned
	Limit int32 `json:"limit,omitempty"`

	// Limit results to files first seen before this date. The format is YYYY/MM/DD
	MaxDate string `json:"max_date,omitempty"`

	// Maximum file size. The value can be specified either in bytes or in multiples of KB/MB/GB. Examples: 128000, 1.3 KB, 8mb
	MaxSize string `json:"max_size,omitempty"`

	// Limit results to files first seen after this date. The format is YYYY/MM/DD
	MinDate string `json:"min_date,omitempty"`

	// Minimum file size. The value can be specified either in bytes or in multiples of KB/MB/GB. Examples: 128000, 1.3 KB, 8mb
	MinSize string `json:"min_size,omitempty"`
}

// Validate validates this malquery external hunt options
func (m *MalqueryExternalHuntOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this malquery external hunt options based on context it is used
func (m *MalqueryExternalHuntOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MalqueryExternalHuntOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MalqueryExternalHuntOptions) UnmarshalBinary(b []byte) error {
	var res MalqueryExternalHuntOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

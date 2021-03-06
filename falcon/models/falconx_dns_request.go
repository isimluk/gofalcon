// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FalconxDNSRequest falconx DNS request
//
// swagger:model falconx.DNSRequest
type FalconxDNSRequest struct {

	// address
	Address string `json:"address,omitempty"`

	// compromised
	Compromised bool `json:"compromised,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// registrar creation timestamp
	RegistrarCreationTimestamp string `json:"registrar_creation_timestamp,omitempty"`

	// registrar name
	RegistrarName string `json:"registrar_name,omitempty"`

	// registrar name servers
	RegistrarNameServers string `json:"registrar_name_servers,omitempty"`

	// registrar organization
	RegistrarOrganization string `json:"registrar_organization,omitempty"`
}

// Validate validates this falconx DNS request
func (m *FalconxDNSRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this falconx DNS request based on context it is used
func (m *FalconxDNSRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FalconxDNSRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FalconxDNSRequest) UnmarshalBinary(b []byte) error {
	var res FalconxDNSRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

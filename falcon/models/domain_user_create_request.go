// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainUserCreateRequest domain user create request
//
// swagger:model domain.UserCreateRequest
type DomainUserCreateRequest struct {

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this domain user create request
func (m *DomainUserCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain user create request based on context it is used
func (m *DomainUserCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainUserCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUserCreateRequest) UnmarshalBinary(b []byte) error {
	var res DomainUserCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

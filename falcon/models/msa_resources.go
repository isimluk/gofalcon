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

// MsaResources msa resources
//
// swagger:model msa.Resources
type MsaResources struct {

	// resources affected
	// Required: true
	ResourcesAffected *int32 `json:"resources_affected"`
}

// Validate validates this msa resources
func (m *MsaResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourcesAffected(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaResources) validateResourcesAffected(formats strfmt.Registry) error {

	if err := validate.Required("resources_affected", "body", m.ResourcesAffected); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this msa resources based on context it is used
func (m *MsaResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsaResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsaResources) UnmarshalBinary(b []byte) error {
	var res MsaResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

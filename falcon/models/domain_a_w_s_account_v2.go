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

// DomainAWSAccountV2 domain a w s account v2
//
// swagger:model domain.AWSAccountV2
type DomainAWSAccountV2 struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"CreatedAt"`

	// deleted at
	// Required: true
	// Format: date-time
	DeletedAt *strfmt.DateTime `json:"DeletedAt"`

	// ID
	// Required: true
	ID *int64 `json:"ID"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"UpdatedAt"`

	// 12 digit AWS provided unique identifier for the account.
	AccountID string `json:"account_id,omitempty"`

	// cid
	Cid string `json:"cid,omitempty"`

	// ID assigned for use with cross account IAM role access.
	ExternalID string `json:"external_id,omitempty"`

	// The full arn of the IAM role created in this account to control access.
	IamRoleArn string `json:"iam_role_arn,omitempty"`

	// is master
	// Required: true
	IsMaster *bool `json:"is_master"`

	// Up to 34 character AWS provided unique identifier for the organization.
	OrganizationID string `json:"organization_id,omitempty"`

	// Account registration status.
	Status string `json:"status,omitempty"`
}

// Validate validates this domain a w s account v2
func (m *DomainAWSAccountV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsMaster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAWSAccountV2) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("CreatedAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAWSAccountV2) validateDeletedAt(formats strfmt.Registry) error {

	if err := validate.Required("DeletedAt", "body", m.DeletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("DeletedAt", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAWSAccountV2) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAWSAccountV2) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("UpdatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainAWSAccountV2) validateIsMaster(formats strfmt.Registry) error {

	if err := validate.Required("is_master", "body", m.IsMaster); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain a w s account v2 based on context it is used
func (m *DomainAWSAccountV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAWSAccountV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAWSAccountV2) UnmarshalBinary(b []byte) error {
	var res DomainAWSAccountV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

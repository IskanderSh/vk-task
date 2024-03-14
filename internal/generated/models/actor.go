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

// Actor actor
//
// swagger:model Actor
type Actor struct {

	// birthday
	// Required: true
	// Format: date
	Birthday *strfmt.Date `json:"birthday"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sex
	// Required: true
	Sex *string `json:"sex"`
}

// Validate validates this actor
func (m *Actor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Actor) validateBirthday(formats strfmt.Registry) error {

	if err := validate.Required("birthday", "body", m.Birthday); err != nil {
		return err
	}

	if err := validate.FormatOf("birthday", "body", "date", m.Birthday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Actor) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Actor) validateSex(formats strfmt.Registry) error {

	if err := validate.Required("sex", "body", m.Sex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this actor based on context it is used
func (m *Actor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Actor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Actor) UnmarshalBinary(b []byte) error {
	var res Actor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

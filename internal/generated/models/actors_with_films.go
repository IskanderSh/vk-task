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

// ActorsWithFilms actors with films
//
// swagger:model ActorsWithFilms
type ActorsWithFilms struct {

	// birthday
	// Format: date
	Birthday strfmt.Date `json:"birthday,omitempty"`

	// films
	Films []string `json:"films"`

	// name
	Name string `json:"name,omitempty"`

	// sex
	Sex string `json:"sex,omitempty"`
}

// Validate validates this actors with films
func (m *ActorsWithFilms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActorsWithFilms) validateBirthday(formats strfmt.Registry) error {
	if swag.IsZero(m.Birthday) { // not required
		return nil
	}

	if err := validate.FormatOf("birthday", "body", "date", m.Birthday.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this actors with films based on context it is used
func (m *ActorsWithFilms) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActorsWithFilms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActorsWithFilms) UnmarshalBinary(b []byte) error {
	var res ActorsWithFilms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
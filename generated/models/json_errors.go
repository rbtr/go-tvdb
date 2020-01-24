// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JSONErrors JSON errors
// swagger:model JSONErrors
type JSONErrors struct {

	// Invalid filters passed to route
	InvalidFilters []string `json:"invalidFilters"`

	// Invalid language or translation missing
	InvalidLanguage string `json:"invalidLanguage,omitempty"`

	// Invalid query params passed to route
	InvalidQueryParams []string `json:"invalidQueryParams"`
}

// Validate validates this JSON errors
func (m *JSONErrors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JSONErrors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONErrors) UnmarshalBinary(b []byte) error {
	var res JSONErrors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

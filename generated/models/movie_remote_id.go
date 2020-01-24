// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MovieRemoteID movie remote Id
// swagger:model MovieRemoteId
type MovieRemoteID struct {

	// id
	ID string `json:"id,omitempty"`

	// source id
	SourceID int64 `json:"source_id,omitempty"`

	// source name
	SourceName string `json:"source_name,omitempty"`

	// source url
	SourceURL string `json:"source_url,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this movie remote Id
func (m *MovieRemoteID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MovieRemoteID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MovieRemoteID) UnmarshalBinary(b []byte) error {
	var res MovieRemoteID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
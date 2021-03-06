// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BlockTimeSummary block time summary
// swagger:model BlockTimeSummary
type BlockTimeSummary struct {

	// difficulty
	Difficulty float64 `json:"difficulty,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// time delta to parent
	TimeDeltaToParent int64 `json:"time_delta_to_parent,omitempty"`
}

// Validate validates this block time summary
func (m *BlockTimeSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockTimeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockTimeSummary) UnmarshalBinary(b []byte) error {
	var res BlockTimeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

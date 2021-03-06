// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelativeTTL relative TTL
// swagger:model RelativeTTL
type RelativeTTL struct {

	// type
	// Enum: [delta]
	Type string `json:"type,omitempty"`

	// value
	// Minimum: 1
	Value int64 `json:"value,omitempty"`
}

// Validate validates this relative TTL
func (m *RelativeTTL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var relativeTtlTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["delta"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relativeTtlTypeTypePropEnum = append(relativeTtlTypeTypePropEnum, v)
	}
}

const (

	// RelativeTTLTypeDelta captures enum value "delta"
	RelativeTTLTypeDelta string = "delta"
)

// prop value enum
func (m *RelativeTTL) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, relativeTtlTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RelativeTTL) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *RelativeTTL) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.MinimumInt("value", "body", int64(m.Value), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelativeTTL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelativeTTL) UnmarshalBinary(b []byte) error {
	var res RelativeTTL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NameClaimTx name claim tx
// swagger:model NameClaimTx
type NameClaimTx struct {

	// account
	Account EncodedHash `json:"account,omitempty"`

	// fee
	// Required: true
	Fee *int64 `json:"fee"`

	// name
	// Required: true
	Name *string `json:"name"`

	// name salt
	// Required: true
	NameSalt *int64 `json:"name_salt"`

	// nonce
	Nonce int64 `json:"nonce,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this name claim tx
func (m *NameClaimTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameSalt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameClaimTx) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if err := m.Account.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account")
		}
		return err
	}

	return nil
}

func (m *NameClaimTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

func (m *NameClaimTx) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NameClaimTx) validateNameSalt(formats strfmt.Registry) error {

	if err := validate.Required("name_salt", "body", m.NameSalt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameClaimTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameClaimTx) UnmarshalBinary(b []byte) error {
	var res NameClaimTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// NamePreclaimTx name preclaim tx
// swagger:model NamePreclaimTx
type NamePreclaimTx struct {

	// account
	Account EncodedHash `json:"account,omitempty"`

	// commitment
	// Required: true
	Commitment *string `json:"commitment"`

	// fee
	// Required: true
	Fee *int64 `json:"fee"`

	// nonce
	Nonce int64 `json:"nonce,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this name preclaim tx
func (m *NamePreclaimTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamePreclaimTx) validateAccount(formats strfmt.Registry) error {

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

func (m *NamePreclaimTx) validateCommitment(formats strfmt.Registry) error {

	if err := validate.Required("commitment", "body", m.Commitment); err != nil {
		return err
	}

	return nil
}

func (m *NamePreclaimTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamePreclaimTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamePreclaimTx) UnmarshalBinary(b []byte) error {
	var res NamePreclaimTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

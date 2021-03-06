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

// ChannelCloseMutualTx channel close mutual tx
// swagger:model ChannelCloseMutualTx
type ChannelCloseMutualTx struct {

	// channel id
	// Required: true
	ChannelID EncodedHash `json:"channel_id"`

	// fee
	// Required: true
	// Minimum: 0
	Fee *int64 `json:"fee"`

	// initiator amount final
	// Required: true
	// Minimum: 0
	InitiatorAmountFinal *int64 `json:"initiator_amount_final"`

	// nonce
	// Required: true
	// Minimum: 0
	Nonce *int64 `json:"nonce"`

	// responder amount final
	// Required: true
	// Minimum: 0
	ResponderAmountFinal *int64 `json:"responder_amount_final"`

	// ttl
	// Minimum: 0
	TTL *int64 `json:"ttl,omitempty"`
}

// Validate validates this channel close mutual tx
func (m *ChannelCloseMutualTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAmountFinal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmountFinal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelCloseMutualTx) validateChannelID(formats strfmt.Registry) error {

	if err := m.ChannelID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel_id")
		}
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateInitiatorAmountFinal(formats strfmt.Registry) error {

	if err := validate.Required("initiator_amount_final", "body", m.InitiatorAmountFinal); err != nil {
		return err
	}

	if err := validate.MinimumInt("initiator_amount_final", "body", int64(*m.InitiatorAmountFinal), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	if err := validate.MinimumInt("nonce", "body", int64(*m.Nonce), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateResponderAmountFinal(formats strfmt.Registry) error {

	if err := validate.Required("responder_amount_final", "body", m.ResponderAmountFinal); err != nil {
		return err
	}

	if err := validate.MinimumInt("responder_amount_final", "body", int64(*m.ResponderAmountFinal), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseMutualTx) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCloseMutualTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCloseMutualTx) UnmarshalBinary(b []byte) error {
	var res ChannelCloseMutualTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

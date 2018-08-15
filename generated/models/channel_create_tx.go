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

// ChannelCreateTx channel create tx
// swagger:model ChannelCreateTx
type ChannelCreateTx struct {

	// channel reserve
	// Required: true
	// Minimum: 0
	ChannelReserve *int64 `json:"channel_reserve"`

	// fee
	// Required: true
	// Minimum: 0
	Fee *int64 `json:"fee"`

	// initiator
	// Required: true
	Initiator EncodedHash `json:"initiator"`

	// initiator amount
	// Required: true
	// Minimum: 0
	InitiatorAmount *int64 `json:"initiator_amount"`

	// lock period
	// Required: true
	// Minimum: 0
	LockPeriod *int64 `json:"lock_period"`

	// nonce
	// Minimum: 0
	Nonce *int64 `json:"nonce,omitempty"`

	// push amount
	// Required: true
	// Minimum: 0
	PushAmount *int64 `json:"push_amount"`

	// responder
	// Required: true
	Responder EncodedHash `json:"responder"`

	// responder amount
	// Required: true
	// Minimum: 0
	ResponderAmount *int64 `json:"responder_amount"`

	// Root hash of the channel's internal state tree
	// Required: true
	StateHash EncodedHash `json:"state_hash"`

	// ttl
	// Minimum: 0
	TTL *int64 `json:"ttl,omitempty"`
}

// Validate validates this channel create tx
func (m *ChannelCreateTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelReserve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponderAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
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

func (m *ChannelCreateTx) validateChannelReserve(formats strfmt.Registry) error {

	if err := validate.Required("channel_reserve", "body", m.ChannelReserve); err != nil {
		return err
	}

	if err := validate.MinimumInt("channel_reserve", "body", int64(*m.ChannelReserve), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateInitiator(formats strfmt.Registry) error {

	if err := m.Initiator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initiator")
		}
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateInitiatorAmount(formats strfmt.Registry) error {

	if err := validate.Required("initiator_amount", "body", m.InitiatorAmount); err != nil {
		return err
	}

	if err := validate.MinimumInt("initiator_amount", "body", int64(*m.InitiatorAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateLockPeriod(formats strfmt.Registry) error {

	if err := validate.Required("lock_period", "body", m.LockPeriod); err != nil {
		return err
	}

	if err := validate.MinimumInt("lock_period", "body", int64(*m.LockPeriod), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateNonce(formats strfmt.Registry) error {

	if swag.IsZero(m.Nonce) { // not required
		return nil
	}

	if err := validate.MinimumInt("nonce", "body", int64(*m.Nonce), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validatePushAmount(formats strfmt.Registry) error {

	if err := validate.Required("push_amount", "body", m.PushAmount); err != nil {
		return err
	}

	if err := validate.MinimumInt("push_amount", "body", int64(*m.PushAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateResponder(formats strfmt.Registry) error {

	if err := m.Responder.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder")
		}
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateResponderAmount(formats strfmt.Registry) error {

	if err := validate.Required("responder_amount", "body", m.ResponderAmount); err != nil {
		return err
	}

	if err := validate.MinimumInt("responder_amount", "body", int64(*m.ResponderAmount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateStateHash(formats strfmt.Registry) error {

	if err := m.StateHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state_hash")
		}
		return err
	}

	return nil
}

func (m *ChannelCreateTx) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCreateTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCreateTx) UnmarshalBinary(b []byte) error {
	var res ChannelCreateTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
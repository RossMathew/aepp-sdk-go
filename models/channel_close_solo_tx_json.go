// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ChannelCloseSoloTxJSON channel close solo tx JSON
// swagger:model ChannelCloseSoloTxJSON
type ChannelCloseSoloTxJSON struct {
	vsnField int64

	ChannelCloseSoloTx
}

// DataSchema gets the data schema of this subtype
func (m *ChannelCloseSoloTxJSON) DataSchema() string {
	return "ChannelCloseSoloTxJSON"
}

// SetDataSchema sets the data schema of this subtype
func (m *ChannelCloseSoloTxJSON) SetDataSchema(val string) {

}

// Vsn gets the vsn of this subtype
func (m *ChannelCloseSoloTxJSON) Vsn() int64 {
	return m.vsnField
}

// SetVsn sets the vsn of this subtype
func (m *ChannelCloseSoloTxJSON) SetVsn(val int64) {
	m.vsnField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ChannelCloseSoloTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		ChannelCloseSoloTx
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ChannelCloseSoloTxJSON

	if base.DataSchema != result.DataSchema() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid data_schema value: %q", base.DataSchema)
	}

	result.vsnField = base.Vsn

	result.ChannelCloseSoloTx = data.ChannelCloseSoloTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ChannelCloseSoloTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ChannelCloseSoloTx
	}{

		ChannelCloseSoloTx: m.ChannelCloseSoloTx,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}{

		DataSchema: m.DataSchema(),

		Vsn: m.Vsn(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this channel close solo tx JSON
func (m *ChannelCloseSoloTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChannelCloseSoloTx
	if err := m.ChannelCloseSoloTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCloseSoloTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCloseSoloTxJSON) UnmarshalBinary(b []byte) error {
	var res ChannelCloseSoloTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

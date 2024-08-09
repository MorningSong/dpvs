// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IpsetInfoArray IpsetInfoArray contains an array of ipset.
//
// swagger:model IpsetInfoArray
type IpsetInfoArray struct {

	// count
	Count int32 `json:"Count,omitempty"`

	// infos
	Infos []*IpsetInfo `json:"Infos"`
}

// Validate validates this ipset info array
func (m *IpsetInfoArray) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpsetInfoArray) validateInfos(formats strfmt.Registry) error {
	if swag.IsZero(m.Infos) { // not required
		return nil
	}

	for i := 0; i < len(m.Infos); i++ {
		if swag.IsZero(m.Infos[i]) { // not required
			continue
		}

		if m.Infos[i] != nil {
			if err := m.Infos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipset info array based on the context it is used
func (m *IpsetInfoArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpsetInfoArray) contextValidateInfos(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Infos); i++ {

		if m.Infos[i] != nil {
			if err := m.Infos[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Infos" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Infos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IpsetInfoArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpsetInfoArray) UnmarshalBinary(b []byte) error {
	var res IpsetInfoArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

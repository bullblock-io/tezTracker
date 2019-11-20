// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FutureBakingRightsPerBlock future baking rights per block
// swagger:model FutureBakingRightsPerBlock
type FutureBakingRightsPerBlock struct {

	// level
	Level int64 `json:"level,omitempty"`

	// rights
	Rights []*BakingRightsRow `json:"rights"`
}

// Validate validates this future baking rights per block
func (m *FutureBakingRightsPerBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FutureBakingRightsPerBlock) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {
		if swag.IsZero(m.Rights[i]) { // not required
			continue
		}

		if m.Rights[i] != nil {
			if err := m.Rights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FutureBakingRightsPerBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FutureBakingRightsPerBlock) UnmarshalBinary(b []byte) error {
	var res FutureBakingRightsPerBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
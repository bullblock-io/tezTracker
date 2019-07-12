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

// AccountsRow accounts row
// swagger:model AccountsRow
type AccountsRow struct {

	// account Id
	// Required: true
	AccountID *string `json:"accountId"`

	// balance
	// Required: true
	Balance *float64 `json:"balance"`

	// block Id
	// Required: true
	BlockID *string `json:"blockId"`

	// block level
	// Required: true
	BlockLevel *float64 `json:"blockLevel"`

	// counter
	// Required: true
	Counter *int32 `json:"counter"`

	// delegate setable
	// Required: true
	DelegateSetable *bool `json:"delegateSetable"`

	// delegate value
	DelegateValue string `json:"delegateValue,omitempty"`

	// manager
	// Required: true
	Manager *string `json:"manager"`

	// script
	Script string `json:"script,omitempty"`

	// spendable
	// Required: true
	Spendable *bool `json:"spendable"`

	// storage
	Storage string `json:"storage,omitempty"`
}

// Validate validates this accounts row
func (m *AccountsRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegateSetable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpendable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsRow) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBlockID(formats strfmt.Registry) error {

	if err := validate.Required("blockId", "body", m.BlockID); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateBlockLevel(formats strfmt.Registry) error {

	if err := validate.Required("blockLevel", "body", m.BlockLevel); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateCounter(formats strfmt.Registry) error {

	if err := validate.Required("counter", "body", m.Counter); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateDelegateSetable(formats strfmt.Registry) error {

	if err := validate.Required("delegateSetable", "body", m.DelegateSetable); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateManager(formats strfmt.Registry) error {

	if err := validate.Required("manager", "body", m.Manager); err != nil {
		return err
	}

	return nil
}

func (m *AccountsRow) validateSpendable(formats strfmt.Registry) error {

	if err := validate.Required("spendable", "body", m.Spendable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsRow) UnmarshalBinary(b []byte) error {
	var res AccountsRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

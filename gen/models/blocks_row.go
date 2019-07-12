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

// BlocksRow blocks row
// swagger:model BlocksRow
type BlocksRow struct {

	// active proposal
	ActiveProposal string `json:"activeProposal,omitempty"`

	// baker
	Baker string `json:"baker,omitempty"`

	// chain Id
	ChainID string `json:"chainId,omitempty"`

	// consumed gas
	ConsumedGas float64 `json:"consumedGas,omitempty"`

	// context
	Context string `json:"context,omitempty"`

	// current expected quorum
	CurrentExpectedQuorum int32 `json:"currentExpectedQuorum,omitempty"`

	// expected commitment
	ExpectedCommitment bool `json:"expectedCommitment,omitempty"`

	// fitness
	// Required: true
	Fitness *string `json:"fitness"`

	// hash
	// Required: true
	Hash *string `json:"hash"`

	// level
	// Required: true
	Level *int32 `json:"level"`

	// meta cycle
	MetaCycle int32 `json:"metaCycle,omitempty"`

	// meta cycle position
	MetaCyclePosition int32 `json:"metaCyclePosition,omitempty"`

	// meta level
	MetaLevel int32 `json:"metaLevel,omitempty"`

	// meta level position
	MetaLevelPosition int32 `json:"metaLevelPosition,omitempty"`

	// meta voting period
	MetaVotingPeriod int32 `json:"metaVotingPeriod,omitempty"`

	// meta voting period position
	MetaVotingPeriodPosition int32 `json:"metaVotingPeriodPosition,omitempty"`

	// nonce hash
	NonceHash string `json:"nonceHash,omitempty"`

	// operations hash
	OperationsHash string `json:"operationsHash,omitempty"`

	// period kind
	PeriodKind string `json:"periodKind,omitempty"`

	// predecessor
	// Required: true
	Predecessor *string `json:"predecessor"`

	// proto
	// Required: true
	Proto *int32 `json:"proto"`

	// protocol
	// Required: true
	Protocol *string `json:"protocol"`

	// signature
	Signature string `json:"signature,omitempty"`

	// timestamp
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// validation pass
	// Required: true
	ValidationPass *int32 `json:"validationPass"`
}

// Validate validates this blocks row
func (m *BlocksRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFitness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePredecessor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationPass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlocksRow) validateFitness(formats strfmt.Registry) error {

	if err := validate.Required("fitness", "body", m.Fitness); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateHash(formats strfmt.Registry) error {

	if err := validate.Required("hash", "body", m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validatePredecessor(formats strfmt.Registry) error {

	if err := validate.Required("predecessor", "body", m.Predecessor); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateProto(formats strfmt.Registry) error {

	if err := validate.Required("proto", "body", m.Proto); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *BlocksRow) validateValidationPass(formats strfmt.Registry) error {

	if err := validate.Required("validationPass", "body", m.ValidationPass); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlocksRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlocksRow) UnmarshalBinary(b []byte) error {
	var res BlocksRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

package models

import "time"

type Block struct {
	Level                    uint                   `json:"level"`
	Proto                    uint                   `json:"proto"`
	Predecessor              string                 `json:"predecessor"`
	Timestamp                time.Time              `json:"timestamp"`
	ValidationPass           uint                   `json:"validation_pass"`
	Fitness                  string                 `json:"fitness"`
	Context                  string                 `json:"context"`
	Signature                string                 `json:"signature"`
	Protocol                 string                 `json:"protocol"`
	ChainID                  string                 `json:"chain_id"`
	Hash                     string                 `json:"hash"`
	OperationsHash           string                 `json:"operations_hash"`
	PeriodKind               string                 `json:"period_kind"`
	CurrentExpectedQuorum    uint                   `json:"current_expected_quorum"`
	ActiveProposal           string                 `json:"active_proposal"`
	Baker                    string                 `json:"baker"`
	NonceHash                string                 `json:"nonce_hash"`
	ConsumedGas              float64                `json:"consumed_gas"`
	MetaLevel                uint                   `json:"meta_level"`
	MetaLevelPosition        uint                   `json:"meta_level_position"`
	MetaCycle                uint                   `json:"meta_cycle"`
	MetaCyclePosition        uint                   `json:"meta_cycle_position"`
	MetaVotingPeriod         uint                   `json:"meta_voting_period"`
	MetaVotingPeriodPosition uint                   `json:"meta_voting_period_position"`
	ExpectedCommitment       bool                   `json:"expected_commitment"`
	Delegates                []*Delegate            `json:"delegates"`            // This line is infered from other tables.
	Proposals                []*Proposal            `json:"proposals"`            // This line is infered from other tables.
	Rolls                    []*Roll                `json:"rolls"`                // This line is infered from other tables.
	Ballots                  []*Ballot              `json:"ballots"`              // This line is infered from other tables.
	AccountsCheckpoint       []*AccountsCheckpoint  `json:"accounts_checkpoint"`  // This line is infered from other tables.
	OperationGroups          []*OperationGroup      `json:"operation_groups"`     // This line is infered from other tables.
	DelegatesCheckpoint      []*DelegatesCheckpoint `json:"delegates_checkpoint"` // This line is infered from other tables.
	Accounts                 []*Account             `json:"accounts"`             // This line is infered from other tables.

}

package models

import "time"

type Operation struct {
	OperationID         uint         `gorm:"primary_key;AUTO_INCREMENT" json:"operation_id"`
	Operation           *Operation   `json:"operation"` // This line is infered from column name "operation_id".
	OperationGroupHash  string       `json:"operation_group_hash"`
	Kind                string       `json:"kind"`
	Level               uint         `json:"level"`
	Delegate            string       `json:"delegate"`
	Slots               string       `json:"slots"`
	Nonce               string       `json:"nonce"`
	Pkh                 string       `json:"pkh"`
	Secret              string       `json:"secret"`
	Source              string       `json:"source"`
	Fee                 float64      `json:"fee"`
	Counter             float64      `json:"counter"`
	GasLimit            float64      `json:"gas_limit"`
	StorageLimit        float64      `json:"storage_limit"`
	PublicKey           string       `json:"public_key"`
	Amount              float64      `json:"amount"`
	Destination         string       `json:"destination"`
	Parameters          string       `json:"parameters"`
	ManagerPubkey       string       `json:"manager_pubkey"`
	Balance             float64      `json:"balance"`
	Spendable           bool         `json:"spendable"`
	Delegatable         bool         `json:"delegatable"`
	Script              string       `json:"script"`
	Storage             string       `json:"storage"`
	Status              string       `json:"status"`
	ConsumedGas         float64      `json:"consumed_gas"`
	StorageSize         float64      `json:"storage_size"`
	PaidStorageSizeDiff float64      `json:"paid_storage_size_diff"`
	OriginatedContracts string       `json:"originated_contracts"`
	BlockHash           string       `json:"block_hash"`
	BlockLevel          uint         `json:"block_level"`
	Timestamp           time.Time    `json:"timestamp"`
	Operations          []*Operation `json:"operations"` // This line is infered from other tables.

}

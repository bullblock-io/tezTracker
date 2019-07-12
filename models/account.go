package models

type Account struct {
	AccountID          string                `gorm:"primary_key;AUTO_INCREMENT" json:"account_id"`
	Account            *Account              `json:"account"` // This line is infered from column name "account_id".
	BlockID            string                `json:"block_id"`
	Block              *Block                `json:"block"` // This line is infered from column name "block_id".
	Manager            string                `json:"manager"`
	Spendable          bool                  `json:"spendable"`
	DelegateSetable    bool                  `json:"delegate_setable"`
	DelegateValue      string                `json:"delegate_value"`
	Counter            uint                  `json:"counter"`
	Script             string                `json:"script"`
	Storage            string                `json:"storage"`
	Balance            float64               `json:"balance"`
	BlockLevel         float64               `json:"block_level" sql:"DEFAULT:'-1'::integer"`
	AccountsCheckpoint []*AccountsCheckpoint `json:"accounts_checkpoint"` // This line is infered from other tables.
	DelegatedContracts []*DelegatedContract  `json:"delegated_contracts"` // This line is infered from other tables.
	Accounts           []*Account            `json:"accounts"`            // This line is infered from other tables.

}

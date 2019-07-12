package models

type Delegate struct {
	Pkh              string  `gorm:"primary_key;AUTO_INCREMENT" json:"pkh"`
	BlockID          string  `json:"block_id"`
	Block            *Block  `json:"block"` // This line is infered from column name "block_id".
	Balance          float64 `json:"balance"`
	FrozenBalance    float64 `json:"frozen_balance"`
	StakingBalance   float64 `json:"staking_balance"`
	DelegatedBalance float64 `json:"delegated_balance"`
	Deactivated      bool    `json:"deactivated"`
	GracePeriod      uint    `json:"grace_period"`
	BlockLevel       uint    `json:"block_level" sql:"DEFAULT:'-1'::integer"`
}

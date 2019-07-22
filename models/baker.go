package models

type Baker struct {
	AccountID      string `json:"pkh"`
	StakingBalance int64  `json:"staking_balance"`
	Blocks         int64  `json:"blocks"`
	Endorsements   int64  `json:"endorsements"`
	Fees           int64  `json:"fees"`
}

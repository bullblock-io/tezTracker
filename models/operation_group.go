package models

type OperationGroup struct {
	Protocol  string `json:"protocol"`
	ChainID   string `json:"chain_id"`
	Hash      string `gorm:"primary_key;AUTO_INCREMENT" json:"hash"`
	Branch    string `json:"branch"`
	Signature string `json:"signature"`
	BlockID   string `json:"block_id"`
	Block     *Block `json:"block"` // This line is infered from column name "block_id".

}

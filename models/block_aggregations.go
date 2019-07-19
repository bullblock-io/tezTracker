package models

type BlockAggregationView struct {
	Level  int64 `json:"level"`
	Volume int64 `json:"volume"`
	Fees   int64 `json:"fees"`
}

func (*BlockAggregationView) TableName() string {
	return "block_aggregation_view"
}

package cmc

// MarketData is a Price and Price Change with json deserialization.
type MarketData struct{
	Price  float64  `json:"usd"`
	Price24hChange float64`json:"usd_24h_change"`
}

// GetPrice returns the price in USD.
func (md *MarketData) GetPrice() float64{
	return md.Price
}
// GetPriceChange returns the price change during the last 24 hours in percents.
func (md *MarketData) GetPriceChange() float64{
	return md.Price24hChange
}
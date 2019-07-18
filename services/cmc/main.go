package cmc


import (
	"encoding/json"

	coingecko "github.com/superoo7/go-gecko/v3"
	"github.com/bullblock-io/tezTracker/models"
)

type tezosMarketData struct{
	Tezos MarketData `json:"tezos"`
}
// CoinGecko is market data provider.
type CoinGecko struct{}

// GetTezosMarketData gets the tezos price and price change from CoinGecko API.
// TODO: some caching layer should be implemented.
func (CoinGecko) GetTezosMarketData() (md models.MarketInfo, err error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=tezos&vs_currencies=usd&include_market_cap=false&include_24hr_vol=false&include_24hr_change=true&include_last_updated_at=true"
	
	cg := coingecko.NewClient(nil)
	b, err := cg.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var tmd tezosMarketData
	err = json.Unmarshal(b, &tmd)
	if err != nil{
		return nil, err
	}

	return &tmd.Tezos, nil
}
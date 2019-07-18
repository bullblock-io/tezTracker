package cmc

import (
	"encoding/json"
	"fmt"

	"github.com/bullblock-io/tezTracker/models"
	coingecko "github.com/superoo7/go-gecko/v3"
)

const baseURL = "https://api.coingecko.com/api/v3"

type tezosMarketData struct {
	Tezos USDMarketData `json:"tezos"`
}

// CoinGecko is market data provider.
type CoinGecko struct{}

// GetTezosMarketData gets the tezos price and price change from CoinGecko API.
// TODO: some caching layer should be implemented.
func (CoinGecko) GetTezosMarketData() (md models.MarketInfo, err error) {
	url := fmt.Sprintf("%s/simple/price?ids=tezos&vs_currencies=usd&include_24hr_change=true&include_last_updated_at=true", baseURL)

	cg := coingecko.NewClient(nil)
	b, err := cg.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var tmd tezosMarketData
	err = json.Unmarshal(b, &tmd)
	if err != nil {
		return nil, err
	}

	return &tmd.Tezos, nil
}

package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)
// Info renders price info into OpenAPI model.
func Info(mi models.MarketInfo) *genModels.Info {
	p := mi.GetPrice()
	p24 := mi.GetPriceChange()
	return &genModels.Info{Price: &p, Price24hChange: &p24}
}

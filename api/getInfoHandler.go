package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	info "github.com/bullblock-io/tezTracker/gen/restapi/operations/app_info"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// MarketDataProvider is an interface for getting actual price and price changes.
type MarketDataProvider interface {
	GetTezosMarketData() (md models.MarketInfo, err error)
}

type getInfoHandler struct {
	provider MarketDataProvider
}

// Handle serves the Get Info request.
func (h *getInfoHandler) Handle(params info.GetInfoParams) middleware.Responder {
	md, err := h.provider.GetTezosMarketData()
	if err != nil {
		logrus.Errorf("failed to get market data: %s", err.Error())
		return info.NewGetInfoInternalServerError()
	}

	return info.NewGetInfoOK().WithPayload(render.Info(md))
}

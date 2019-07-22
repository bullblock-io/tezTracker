package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getBakerListHandler struct {
	db *gorm.DB
}

// Handle serves the Get Baker List request.
func (h *getBakerListHandler) Handle(params accounts.GetBakersListParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	limit := uint64(0)
	if params.Limit != nil {
		limit = uint64(*params.Limit)
	}
	after := ""
	if params.AfterID != nil {
		after = *params.AfterID
	}
	accs, err := service.BakerList(limit, after)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetBakersListNotFound()
	}
	return accounts.NewGetBakersListOK().WithPayload(render.Bakers(accs))
}

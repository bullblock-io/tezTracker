package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getBakingRightsHandler struct {
	db *gorm.DB
}

// Handle serves the Get Block List request.
func (h *getBakingRightsHandler) Handle(params blocks.GetBakingRightsParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	limiter := NewLimiter(params.Limit, params.Offset)
	priorityTo := 0
	if params.PrioritiesTo != nil {
		priorityTo = int(*params.PrioritiesTo)
	}
	count, bs, err := service.BakingRightsList(params.BlockID, priorityTo, limiter)
	if err != nil {
		logrus.Errorf("failed to get baking rights: %s", err.Error())
		return blocks.NewGetBakingRightsNotFound()
	}

	return blocks.NewGetBakingRightsOK().WithPayload(render.BlocksBakingRights(bs)).WithXTotalCount(count)
}

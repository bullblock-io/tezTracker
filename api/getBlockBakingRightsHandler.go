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

type getBlockBakingRightsHandler struct {
	db *gorm.DB
}

// Handle serves the Get Block request.
func (h *getBlockBakingRightsHandler) Handle(params blocks.GetBlockBakingRightsParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	rights, count, err := service.GetBlockBakingRights(params.Hash)

	if err != nil {
		if err == services.ErrNotFound {
			return blocks.NewGetBlockBakingRightsNotFound()
		}
		logrus.Errorf("failed to get block: %s", err.Error())
		return blocks.NewGetBlockBakingRightsInternalServerError()
	}

	return blocks.NewGetBlockBakingRightsOK().WithPayload(render.BakingRights(rights)).WithXTotalCount(count)
}

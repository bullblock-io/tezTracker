package api

import (
	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// SetHandlers initializes the API handlers with underlying services.
func SetHandlers(serv *operations.TezTrackerAPI, db *gorm.DB) {
	serv.Logger = logrus.Infof
	serv.BlocksGetBlocksHeadHandler = &getHeadBlockHandler{db}
	serv.BlocksGetBlocksListHandler = &getBlockListHandler{db}
}

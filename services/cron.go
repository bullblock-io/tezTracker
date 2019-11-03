package services

import (
	"time"

	"github.com/bullblock-io/tezTracker/config"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services/counter"
	"github.com/jinzhu/gorm"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

func InitCron(cfg config.Config, db *gorm.DB) *gron.Cron {
	cron := gron.New()

	if cfg.CounterIntervalHours > 0 {
		dur := time.Duration(cfg.CounterIntervalHours) * time.Hour
		log.Tracef("Sheduling counter saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			unitOfWork := repos.New(db)

			err := counter.SaveCounters(unitOfWork.GetOperation(), unitOfWork.GetOperationCounter())
			if err != nil {
				log.Errorf("counter saver failed: %s", err.Error())
				return
			}

		})
	} else {
		log.Tracef(" no Sheduling counter saver every")
	}

	return cron
}

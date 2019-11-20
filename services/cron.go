package services

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/bullblock-io/tezTracker/config"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services/counter"
	"github.com/bullblock-io/tezTracker/services/future_rights"
	"github.com/bullblock-io/tezTracker/services/rpc_client"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client"
	"github.com/bullblock-io/tezTracker/services/snapshots"
	"github.com/jinzhu/gorm"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

func AddToCron(cron *gron.Cron, cfg config.Config, db *gorm.DB, rpcConfig client.TransportConfig, network models.Network) {

	if cfg.CounterIntervalHours > 0 {
		dur := time.Duration(cfg.CounterIntervalHours) * time.Hour
		log.Infof("Sheduling counter saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			unitOfWork := repos.New(db)

			err := counter.SaveCounters(unitOfWork.GetOperation(), unitOfWork.GetOperationCounter())
			if err != nil {
				log.Errorf("counter saver failed: %s", err.Error())
				return
			}

		})
	} else {
		log.Infof("no sheduling counter due to missing CounterIntervalHours in config")
	}
	if cfg.FutureRightsIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.FutureRightsIntervalMinutes) * time.Minute
		log.Infof("Sheduling future rights parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network))
				count, err := future_rights.SaveNewBakingRights(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("BakingRights saver failed: %s", err.Error())
					return
				}
				log.Tracef("BakingRights saved %d rights", count)
			} else {
				log.Tracef("skipping BakingRights saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling future rights parser due to missing FutureRightsIntervalMinutes in config")
	}
	if cfg.SnapshotCheckIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.SnapshotCheckIntervalMinutes) * time.Minute
		log.Infof("Sheduling snapshots parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network))
				count, err := snapshots.SaveNewSnapshots(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("Snapshots saver failed: %s", err.Error())
					return
				}
				log.Tracef("Snapshots saved %d rights", count)
			} else {
				log.Tracef("skipping Snapshots saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling snapshots parser due to missing FutureRightsIntervalMinutes in config")
	}

}

// cycle = 160
// level = 160*4096+1=655361

// snapshot=12

// snapshot_block = 153*4096+1+(12+1)*256-1=630016

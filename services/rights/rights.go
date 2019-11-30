package rights

import (
	"context"
	"fmt"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos/baking_rights"
	"github.com/bullblock-io/tezTracker/repos/block"
)

type RightsRepo interface {
	First() (found bool, right models.FutureBakingRight, err error)
	Create(right models.FutureBakingRight) error
}

type RightsProvider interface {
	RightsFor(ctx context.Context, blockFrom, blockTo int64) ([]models.BakingRight, error)
}

type UnitOfWork interface {
	Start(ctx context.Context)
	RollbackUnlessCommitted()
	Commit() error
	GetBlock() block.Repo
	GetBakingRight() baking_rights.Repo
}

const BlocksRangeSize = 256

func SaveOldBakingRights(ctx context.Context, unit UnitOfWork, provider RightsProvider) (count int, err error) {
	rightsRepo := unit.GetBakingRight()
	found, lastBlock, err := rightsRepo.First()
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, fmt.Errorf("shit")
	}
	if lastBlock.Level == 1 {
		return 0, nil
	}
	nextBlockToScan := lastBlock.Level - 1
	for nextBlockToScan >= 1 {
		startRange := nextBlockToScan - BlocksRangeSize + 1

		if startRange < 1 {
			startRange = 1
		}
		cnt, err := SaveRightsForBlockRange(ctx, startRange, nextBlockToScan, unit, provider)
		if err != nil {
			return 0, err
		}
		count += cnt
		nextBlockToScan = startRange - 1
	}
	return count, nil
}

func SaveRightsForBlockRange(ctx context.Context, blockFrom, blockTo int64, unit UnitOfWork, provider RightsProvider) (int, error) {
	rights, err := provider.RightsFor(ctx, blockFrom, blockTo)
	if err != nil {
		return 0, err
	}
	blockRepo := unit.GetBlock()
	blockHashes := make(map[int64]string)

	blockLevels := make([]int64, 0)
	for i := blockFrom; i <= blockTo; i++ {
		blockLevels = append(blockLevels, i)
	}
	blocks, err := blockRepo.Filter(models.BlockFilter{BlockLevels: blockLevels})
	if err != nil {
		return 0, err
	}
	for i := range blocks {
		blockHashes[blocks[i].Level.Int64] = blocks[i].Hash.String
	}
	for i := range rights {
		level := rights[i].Level
		if h, ok := blockHashes[level]; ok {
			rights[i].BlockHash = h

		} else {

			return 0, fmt.Errorf("shit")

		}
	}
	unit.Start(ctx)
	defer unit.RollbackUnlessCommitted()
	rightsRepo := unit.GetBakingRight()

	err = rightsRepo.CreateBulk(rights)
	if err != nil {
		return 0, err
	}

	err = unit.Commit()
	if err != nil {
		return 0, err
	}
	return len(rights), nil
}

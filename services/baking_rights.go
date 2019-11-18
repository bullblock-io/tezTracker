package services

import (
	"fmt"
	"strconv"

	"github.com/bullblock-io/tezTracker/models"
)

func (t *TezTracker) BakingRightsList(blockLevelOrHash []string, priorityTo int, limiter Limiter) (count int64, blocksWithRights []models.Block, err error) {
	filter := models.BakingRightFilter{
		PriorityTo: priorityTo,
	}
	count = int64(len(blockLevelOrHash))

	blockRepo := t.repoProvider.GetBlock()
	if count == 0 {
		last, err := blockRepo.Last()
		if err != nil {
			return 0, nil, err
		}
		lastLevel := last.Level.Int64
		rangeEnd := lastLevel - int64(limiter.Offset())
		if rangeEnd < 0 {
			return 0, nil, fmt.Errorf("out of range")
		}
		rangeStart := rangeEnd - int64(limiter.Limit())
		if rangeStart < 0 {
			rangeStart = 0
		}
		// we have a block with number 0
		count = lastLevel + 1
		for ; rangeStart <= rangeEnd; rangeStart++ {
			filter.BlockLevels = append(filter.BlockLevels, rangeStart)
		}
	} else {
		for i := range blockLevelOrHash {
			if level, e := strconv.ParseInt(blockLevelOrHash[i], 10, 64); e == nil {
				filter.BlockLevels = append(filter.BlockLevels, level)
			} else {
				filter.BlockHashes = append(filter.BlockHashes, blockLevelOrHash[i])
			}
		}
	}
	blocks, err := blockRepo.Filter(filter.BlockFilter)
	if err != nil {
		return count, nil, err
	}
	if len(filter.BlockHashes) > 0 {
		filter.BlockLevels = make([]int64, len(blocks))
		for i := range blocks {
			filter.BlockLevels[i] = blocks[i].Level.Int64
		}
	}
	r := t.repoProvider.GetBakingRight()
	rights, err := r.List(filter)
	if err != nil {
		return count, nil, err
	}
	blockMap := map[int64]*models.Block{}
	for i := range blocks {
		blockMap[blocks[i].Level.Int64] = &blocks[i]
	}
	for i := range rights {
		blockMap[rights[i].Level].BakingRights = append(blockMap[rights[i].Level].BakingRights, rights[i])
	}
	return count, blocks, nil
}

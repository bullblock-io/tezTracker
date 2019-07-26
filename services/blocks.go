package services

import (
	"fmt"
	"strconv"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/guregu/null"
)

// ErrNotFound is an error returned when the requested entity doesn't exist in the repository.
var ErrNotFound = fmt.Errorf("not found")

// // HeadBlock retrieves the last added block from the repository.
// func (t *TezTracker) HeadBlock() (models.Block, error) {
// 	r := t.repoProvider.GetBlock()
// 	return r.Last()
// }

func (t *TezTracker) HeadBlock() (models.Block, error) {
	r := t.repoProvider.GetNotification()
	for i := 0; i < 5; i++ {
		lvl, err := r.GetBlockUpdate()
		fmt.Printf("lvl %d, err: %s\n", lvl, err)
	}
	return t.repoProvider.GetBlock().Last()
}

// BlockList retrives up to limit of blocks before the specified level.
func (t *TezTracker) BlockList(beforeLevel uint64, limits Limiter) ([]models.Block, error) {
	r := t.repoProvider.GetBlock()
	return r.ListExtended(limits.Limit(), limits.Offset(), beforeLevel)
}

// GetBlockWithOperationGroups retrieves a block by hash or by level. It loads OperationGroups into the found block.
func (t *TezTracker) GetBlockWithOperationGroups(hashOrLevel string) (block models.Block, err error) {
	r := t.repoProvider.GetBlock()

	var filter models.Block
	if i, e := strconv.ParseInt(hashOrLevel, 10, 64); e == nil {
		filter.Level = null.IntFrom(i)
	} else {
		filter.Hash = null.StringFrom(hashOrLevel)
	}
	found, block, err := r.FindExtended(filter)
	if err != nil {
		return block, err
	}
	if !found {
		return block, ErrNotFound
	}
	ogr := t.repoProvider.GetOperationGroup()
	ogs, err := ogr.GetGroupsFor(block)
	if err != nil {
		return block, err
	}
	block.OperationGroups = ogs
	return block, nil
}

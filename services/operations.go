package services

import (
	"strconv"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/guregu/null"
)

// GetOperations gets the operations filtering by operation kinds and blocks wiht pagination.
func (t *TezTracker) GetOperations(kinds, inBlocks, accountIDs []string, limits Limiter, before int64) (operations []models.Operation, err error) {
	r := t.repoProvider.GetOperation()
	return r.List(kinds, inBlocks, accountIDs, limits.Limit(), limits.Offset(), before)
}

// GetBlockEndorsements finds a block and returns endorsements for it.
func (t *TezTracker) GetBlockEndorsements(hashOrLevel string, limits Limiter) (operations []models.Operation, err error) {
	r := t.repoProvider.GetBlock()
	var filter models.Block
	if i, e := strconv.ParseInt(hashOrLevel, 10, 64); e == nil {
		filter.Level = null.IntFrom(i)
	} else {
		filter.Hash = null.StringFrom(hashOrLevel)
	}
	found, block, err := r.Find(filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrNotFound
	}
	or := t.repoProvider.GetOperation()
	return or.EndorsementsFor(block.Level.Int64, limits.Limit(), limits.Offset())
}

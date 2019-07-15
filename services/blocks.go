package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

// HeadBlock retrieves the last added block from the repository.
func (t *TezTracker) HeadBlock() (models.Block, error) {
	r := t.repoProvider.GetBlock()
	return r.Last()
}

// BlockList retrives up to limit of blocks before the specified level.
func (t *TezTracker) BlockList(limit, beforeLevel uint64) ([]models.Block, error) {
	r := t.repoProvider.GetBlock()
	return r.List(limit, beforeLevel)
}

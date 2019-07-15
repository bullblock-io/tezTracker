package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

// HeadBlock retrieves the last added block from the repository.
func (t *TezTracker) HeadBlock() (models.Block, error) {
	r := t.repoProvider.GetBlock()
	return r.Last()
}

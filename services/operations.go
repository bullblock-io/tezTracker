package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

// GetOperations gets the operations filtering by operation kinds and blocks wiht pagination.
func (t *TezTracker) GetOperations(kinds, inBlocks, accountIDs []string, limit, before int64) (operations []models.Operation, err error) {
	r := t.repoProvider.GetOperation()
	return r.List(kinds, inBlocks, accountIDs, limit, before)
}

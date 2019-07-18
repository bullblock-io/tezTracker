package services

import (
	"github.com/bullblock-io/tezTracker/models"
)
// GetOperations gets the operations filtering by operation kinds and blocks wiht pagination.
func (t *TezTracker) GetOperations(kinds []string, inBlocks[]string, limit, before int64) (operations []models.Operation, err error) {
	r := t.repoProvider.GetOperation()
	return r.List(kinds, inBlocks, limit, before)
}

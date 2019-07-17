package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

func (t *TezTracker) GetOperations(kinds []string, limit, before int64) (operations []models.Operation, err error) {
	r := t.repoProvider.GetOperation()
	return r.List(kinds, limit, before)
}

package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) BakerList(limit uint64, after string) ([]models.Baker, error) {
	r := t.repoProvider.GetBaker()
	return r.List(limit, after)
}

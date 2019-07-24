package services

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/guregu/null"
)

// AccountList retrives up to limit of account before the specified id.
func (t *TezTracker) AccountList(before string, limits Limiter) ([]models.Account, error) {
	r := t.repoProvider.GetAccount()
	return r.List(limits.Limit(), limits.Offset(), before)
}

// GetAccount retrieves an account by its ID.
func (t *TezTracker) GetAccount(id string) (acc models.Account, err error) {
	r := t.repoProvider.GetAccount()

	filter := models.Account{AccountID: null.StringFrom(id)}

	found, acc, err := r.Find(filter)
	if err != nil {
		return acc, err
	}
	if !found {
		return acc, ErrNotFound
	}

	return acc, nil
}

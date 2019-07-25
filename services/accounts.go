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

// AccountDelegatorsList retrives up to limit of delegators accounts for the specified accountID.
func (t *TezTracker) AccountDelegatorsList(accountID string, limits Limiter) ([]models.Account, error) {
	r := t.repoProvider.GetAccount()
	filter := models.Account{DelegateValue: accountID}
	return r.Filter(filter, limits.Limit(), limits.Offset())
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

	bi, err := t.GetBakerInfo(id)
	if err != nil {
		return acc, err
	}
	acc.BakerInfo = bi
	return acc, nil
}

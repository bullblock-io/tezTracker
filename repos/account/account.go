package account

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./account.go -destination ./mock_account/main.go Repo
type (
	// Repository is the account repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(limit uint64, after string) ([]models.Account, error)
		Find(filter models.Account) (found bool, acc models.Account, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// List returns a list of accounts from the newest to oldest.
// limit defines the limit for the maximum number of accounts returned.
// before is used to paginate results based on the level. As the result is ordered descendingly the accounts with level < before will be returned.
func (r *Repository) List(limit uint64, after string) (accounts []models.Account, err error) {
	db := r.db.Model(&models.Account{})
	if after != "" {
		db = db.Where("account_id > ?", after)
	}
	err = db.Order("account_id asc").
		Limit(limit).
		Find(&accounts).Error
	return accounts, err
}

// Find looks up for an account with filter.
func (r *Repository) Find(filter models.Account) (found bool, acc models.Account, err error) {
	if res := r.db.Model(&filter).Where(&filter).Find(&acc); res.Error != nil {
		if res.RecordNotFound() {
			return false, acc, nil
		}
		return false, acc, err
	}
	return true, acc, nil
}

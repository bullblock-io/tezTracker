package operation

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the operation repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(kinds []string, inBlocks, accountIDs []string, limit, since int64) (operations []models.Operation, err error)
		EndorsementsFor(blockLevel int64) (operations []models.Operation, err error)
	}
)

const endorsementKind = "endorsement"

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// List returns a list of operations from the newest to oldest.
// limit defines the limit for the maximum number of operations returned.
// since is used to paginate results based on the operation id.
// As the result is ordered descendingly the operations with operation_id < since will be returned.
func (r *Repository) List(kinds []string, inBlocks, accountIDs []string, limit, since int64) (operations []models.Operation, err error) {
	db := r.db.Model(&models.Operation{})
	if since > 0 {
		db = db.Where("operation_id < ?", since)
	}
	if len(kinds) > 0 {
		db = db.Where("kind IN (?)", kinds)
	}
	if len(inBlocks) > 0 {
		db = db.Where("block_hash IN (?)", inBlocks)
	}
	if len(accountIDs) > 0 {
		db = db.Where("delegate IN (?) OR pkh IN (?) OR source IN (?) OR public_key IN (?) OR destination IN (?)", accountIDs, accountIDs, accountIDs, accountIDs, accountIDs)
	}
	err = db.Order("operation_id desc").
		Limit(limit).
		Find(&operations).Error
	return operations, err
}

// EndorsementsFor returns a list of endorsement operations for the provided block level.
func (r *Repository) EndorsementsFor(blockLevel int64) (operations []models.Operation, err error) {
	err = r.db.Model(&models.Operation{}).
		Where("kind = ?", endorsementKind).
		Where("level = ?", blockLevel).
		Find(&operations).Error
	return operations, err
}

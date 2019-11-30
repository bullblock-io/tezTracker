package baking_rights

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the baking rights repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(filter models.BakingRightFilter) (rights []models.BakingRight, err error)
		Find(filter models.BakingRightFilter) (found bool, right models.BakingRight, err error)
		First() (found bool, right models.BakingRight, err error)
		CreateBulk(rights []models.BakingRight) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.BakingRight{}),
	}
}

func (r *Repository) getDb(filter models.BakingRightFilter) *gorm.DB {
	db := r.db.Model(&models.BakingRight{})
	if len(filter.BlockLevels) != 0 {
		db = db.Where("level IN (?)", filter.BlockLevels)
	}
	if len(filter.Delegates) != 0 {
		db = db.Where("delegate IN (?)", filter.Delegates)
	}
	if filter.PriorityFrom != 0 {
		db = db.Where("priority >= ?", filter.PriorityFrom)
	}
	if filter.PriorityTo != 0 {
		db = db.Where("priority <= ?", filter.PriorityTo)
	}
	return db
}

// List returns a list of rights from the newest to oldest.
// limit defines the limit for the maximum number of rights returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the rights with level < since will be returned.
func (r *Repository) List(filter models.BakingRightFilter) (rights []models.BakingRight, err error) {
	db := r.getDb(filter)
	err = db.Order("level desc, priority asc").
		Find(&rights).Error
	return rights, err
}
func (r *Repository) First() (found bool, right models.BakingRight, err error) {
	if res := r.getDb(models.BakingRightFilter{}).
		Order("level asc, priority asc").
		Limit(1).Find(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, err
}

// Find looks up for rights with filter.
func (r *Repository) Find(filter models.BakingRightFilter) (found bool, right models.BakingRight, err error) {
	if res := r.getDb(filter).Find(&right); res.Error != nil {
		if res.RecordNotFound() {
			return false, right, nil
		}
		return false, right, res.Error
	}
	return true, right, nil
}
func (r *Repository) CreateBulk(rights []models.BakingRight) error {
	insertRecords := make([]interface{}, len(rights))
	for i := range rights {
		insertRecords[i] = rights[i]
	}
	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}

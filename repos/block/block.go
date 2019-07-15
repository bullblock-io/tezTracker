package block

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the block repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Last() (block models.Block, err error)
		List(limit, since uint64) (blocks []models.Block, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Last returns the last added block.
func (r *Repository) Last() (block models.Block, err error) {
	err = r.db.Model(&block).Order("level desc").First(&block).Error
	return block, err
}

// List returns a list of blocks from the newest to oldest.
// limit defines the limit for the maximum number of blocks returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the blocks with level < since will be returned.
func (r *Repository) List(limit, since uint64) (blocks []models.Block, err error) {
	db := r.db.Model(&models.Block{})
	if since > 0 {
		db = db.Where("level < ?", since)
	}
	err = db.Order("level desc").
		Limit(limit).
		Find(&blocks).Error
	return blocks, err
}

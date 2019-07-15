package repos

import (
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/jinzhu/gorm"
)
// Provider is the repository provider. 
type Provider struct {
	db *gorm.DB
}

// New creates a new instance of Provider with the underlying DB instance. 
func New(db *gorm.DB) *Provider {
	return &Provider{
		db: db,
	}
}

// GetBlock returns a new block repository.
func (u *Provider) GetBlock() block.Repo {
	return block.New(u.db)
}

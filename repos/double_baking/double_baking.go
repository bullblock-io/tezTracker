package double_baking

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

type (
	// Repository is the baking evidences repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(filter models.DoubleBakingEvidenceQueryOptions) (count int64, evidences []models.DoubleBakingEvidence, err error)
		Last() (found bool, evidence models.DoubleBakingEvidence, err error)
		Find(filter models.BakingRightFilter) (found bool, evidence models.DoubleBakingEvidence, err error)
		Create(evidence models.DoubleBakingEvidence) error
		CreateBulk(evidences []models.DoubleBakingEvidence) error
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.DoubleBakingEvidence{}),
	}
}

func (r *Repository) getDb(filter models.BakingRightFilter) *gorm.DB {
	db := r.db.Model(&models.DoubleBakingEvidence{})
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

// List returns a list of evidences from the newest to oldest.
// limit defines the limit for the maximum number of evidences returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the evidences with level < since will be returned.
func (r *Repository) List(filter models.DoubleBakingEvidenceQueryOptions) (evidences []models.DoubleBakingEvidence, err error) {
	db := r.getDb(filter)
	err = db.Order("level asc, priority asc").
		Find(&evidences).Error
	return evidences, err
}

func (r *Repository) Last() (found bool, evidence models.DoubleBakingEvidence, err error) {
	db := r.db.Model(&evidence)
	if res := db.Order("operation_id asc").Last(&evidence); res.Error != nil {
		if res.RecordNotFound() {
			return false, evidence, nil
		}
		return false, evidence, res.Error
	}
	return true, evidence, nil
}

// Find looks up for evidences with filter.
func (r *Repository) Find(filter models.BakingRightFilter) (found bool, evidence models.DoubleBakingEvidence, err error) {
	if res := r.getDb(filter).Find(&evidence); res.Error != nil {
		if res.RecordNotFound() {
			return false, evidence, nil
		}
		return false, evidence, res.Error
	}
	return true, evidence, nil
}

// Create creates a DoubleBakingEvidence.
func (r *Repository) Create(evidence models.DoubleBakingEvidence) error {
	return r.db.Model(&evidence).Create(&evidence).Error
}

func (r *Repository) CreateBulk(evidences []models.DoubleBakingEvidence) error {
	insertRecords := make([]interface{}, len(evidences))
	for i := range evidences {
		insertRecords[i] = evidences[i]
	}
	return gormbulk.BulkInsert(r.db, insertRecords, 2000)
}

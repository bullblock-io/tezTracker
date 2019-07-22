package baker

import (
	"github.com/bullblock-io/tezTracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./account.go -destination ./mock_account/main.go Repo
type (
	// Repository is the baker repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List(limit uint64, after string) ([]models.Baker, error)
	}

	BakerCounter struct {
		Baker string
		Count int64
	}
)

const endorsementKind = "endorsement"

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// List returns a list of bakers ordered by their staking balance.
// limit defines the limit for the maximum number of bakers returned.
// before is used to paginate results.
// TODO:The pagination is broken for now.
func (r *Repository) List(limit uint64, after string) (bakers []models.Baker, err error) {
	var delegates []models.Delegate
	db := r.db.Model(&models.Delegate{})
	if after != "" {
		db = db.Where("pkh > ?", after)
	}

	err = db.Order("staking_balance desc").
		Limit(limit).
		Find(&delegates).Error
	if err != nil {
		return nil, err
	}

	bakers = ConvertToBakers(delegates)
	bakers, err = r.ExtendBakers(bakers)
	if err != nil {
		return nil, err
	}

	return bakers, err
}

func ConvertToBaker(delegate *models.Delegate) models.Baker {
	return models.Baker{AccountID: delegate.Pkh, StakingBalance: delegate.StakingBalance}
}

func ConvertToBakers(delegates []models.Delegate) []models.Baker {
	count := len(delegates)
	bakers := make([]models.Baker, count)
	for i := range delegates {
		bakers[i] = ConvertToBaker(&delegates[i])
	}
	return bakers
}

func (r *Repository) ExtendBakers(bakers []models.Baker) (extended []models.Baker, err error) {
	count := len(bakers)
	ids := make([]string, count)
	m := make(map[string]*models.Baker, count)
	for i := range bakers {
		pkh := bakers[i].AccountID
		ids[i] = pkh
		m[pkh] = &bakers[i]
	}
	aggInfo, err := r.ListBlocksCount(ids)
	if err != nil {
		return bakers, err
	}
	for i := range aggInfo {
		baker := aggInfo[i].Baker
		if b, ok := m[baker]; ok {
			b.Blocks = aggInfo[i].Count
		}
	}
	aggInfo, err = r.ListEndorsementCount(ids)
	if err != nil {
		return bakers, err
	}
	for i := range aggInfo {
		baker := aggInfo[i].Baker
		if b, ok := m[baker]; ok {
			b.Endorsements = aggInfo[i].Count
		}
	}
	return bakers, nil
}

// ListBlocksCount returns a slice of block counters with the number of blocks baked by each baker among ids.
func (r *Repository) ListBlocksCount(ids []string) (counter []BakerCounter, err error) {
	err = r.db.Model(&models.Block{}).
		Where("baker IN (?)", ids).
		Select("baker, count(1) count").
		Group("baker").Scan(&counter).Error
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// ListBlocksCount returns a slice of block counters with the number of endorsements made by each baker among ids.
func (r *Repository) ListEndorsementCount(ids []string) (counter []BakerCounter, err error) {
	err = r.db.Model(&models.Operation{}).
		Where("delegate IN (?)", ids).
		Where("kind = ?", endorsementKind).
		Select("count(1) as count, delegate as baker").
		Group("delegate").Scan(&counter).Error
	if err != nil {
		return nil, err
	}

	return counter, nil
}

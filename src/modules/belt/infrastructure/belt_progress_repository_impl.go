package infrastructure_belt

import (
	domain_belt "bjj-tracker/src/modules/belt/domain"

	"gorm.io/gorm"
)

type BeltProgressRepositoryImpl struct {
	db *gorm.DB
}

func (r *BeltProgressRepositoryImpl) CreateBeltProgress(beltProgress *domain_belt.BeltProgress) (*domain_belt.BeltProgress, error) {
	beltProgress.ID = 0 // Reset ID to 0 to let GORM auto-generate it

	if err := r.db.Create(beltProgress).Error; err != nil {
		return nil, err
	}
	return beltProgress, nil
}

func NewBeltProgressRepository(db *gorm.DB) *BeltProgressRepositoryImpl {
	return &BeltProgressRepositoryImpl{
		db: db,
	}
}

package repository

import (
	"github.com/masnann/plant_care/features/guide"
	"github.com/masnann/plant_care/features/guide/domain"
	"gorm.io/gorm"
)

type GuideRepository struct {
	db *gorm.DB
}

func NewGuideRepository(db *gorm.DB) guide.RepoGuideInterface {
	return &GuideRepository{
		db: db,
	}
}

func (r *GuideRepository) CountGuides() (int64, error) {
	var count int64
	if err := r.db.Model(&domain.GuideModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *GuideRepository) GetGuidesWithPagination(page int, pageSize int) ([]domain.GuideModel, error) {
	var guides []domain.GuideModel

	offset := (page - 1) * pageSize
	if err := r.db.Offset(offset).Limit(pageSize).Find(&guides).Error; err != nil {
		return nil, err
	}

	return guides, nil
}

func (r *GuideRepository) GetByIdGuides(id uint64) (*domain.GuideModel, error) {
	var guides *domain.GuideModel
	if err := r.db.Where("id = ?", id).Find(&guides).Error; err != nil {
		return nil, err
	}
	return guides, nil
}

func (r *GuideRepository) SearchGuideByName(name string) ([]*domain.GuideModel, error) {
	var guides []*domain.GuideModel
	if err := r.db.Where("name LIKE ? AND deleted_at IS NULL", "%"+name+"%").Find(&guides).Error; err != nil {
		return nil, err
	}
	return guides, nil
}

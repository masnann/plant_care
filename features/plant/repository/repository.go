package repository

import (
	"github.com/masnann/plant_care/features/plant"
	"github.com/masnann/plant_care/features/plant/domain"
	"github.com/masnann/plant_care/utils"
	"gorm.io/gorm"
	"time"
)

type PlantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) plant.RepoPlantInterface {
	return &PlantRepository{db: db}
}

func (r *PlantRepository) GetPlantsWithPagination(userID uint64, offset, pageSize int, plants *[]*domain.PlantModel) error {
	if err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Offset(offset).Limit(pageSize).Find(&plants).Error; err != nil {
		return err
	}
	return nil
}

func (r *PlantRepository) InsertPlants(plant *domain.PlantModel) (*domain.PlantModel, error) {
	newData := &domain.PlantModel{
		UserID:    plant.UserID,
		Name:      plant.Name,
		Type:      plant.Type,
		Date:      utils.GetNowTime(),
		Photo:     plant.Photo,
		CreatedAt: utils.GetNowTime(),
		UpdatedAt: utils.GetNowTime(),
		DeletedAt: nil,
	}
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *PlantRepository) UpdatePlants(plant *domain.PlantModel) (*domain.PlantModel, error) {
	exitingPlant := domain.PlantModel{}
	newData := &domain.PlantModel{
		Name:      plant.Name,
		Type:      plant.Type,
		Date:      plant.Date,
		Photo:     plant.Photo,
		UpdatedAt: utils.GetNowTime(),
	}
	if err := r.db.Model(&exitingPlant).Where("id = ?", plant.ID).Updates(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil

}

func (r *PlantRepository) DeletePlants(plantID uint64) error {
	result := r.db.Model(&domain.PlantModel{}).Where("id = ?", plantID).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PlantRepository) GetPlantsByID(plantID uint) (*domain.PlantModel, error) {
	var plantData domain.PlantModel
	if err := r.db.Where("id = ? AND deleted_at IS NULL", plantID).First(&plantData).Error; err != nil {
		return nil, err
	}

	return &plantData, nil
}

func (r *PlantRepository) GetPlantsByUserID(userID uint64) ([]*domain.PlantModel, error) {
	var plants []*domain.PlantModel
	if err := r.db.Where("user_id = ?", userID).Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

func (r *PlantRepository) SearchPlantsByName(name string) ([]*domain.PlantModel, error) {
	var plants []*domain.PlantModel
	if err := r.db.Where("name LIKE ? AND deleted_at IS NULL", "%"+name+"%").Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

func (r *PlantRepository) SearchPlantsByType(types string) ([]*domain.PlantModel, error) {
	var plants []*domain.PlantModel
	if err := r.db.Where("type LIKE ?", "%"+types+"%").Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

func (r *PlantRepository) CountPlants(userID uint64) (uint64, error) {
	var count int64
	if err := r.db.Model(&domain.PlantModel{}).Where("user_id = ? AND deleted_at IS NULL", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint64(count), nil
}

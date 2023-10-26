package database

import (
	plant "github.com/masnann/plant_care/features/plant/domain"
	"github.com/masnann/plant_care/features/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(domain.UserModel{}, plant.PlantModel{})
}

package database

import (
	guide "github.com/masnann/plant_care/features/guide/domain"
	note "github.com/masnann/plant_care/features/note/domain"
	plant "github.com/masnann/plant_care/features/plant/domain"
	"github.com/masnann/plant_care/features/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(domain.UserModel{}, plant.PlantModel{}, guide.GuideModel{}, note.NoteModel{}, note.PhotoModel{})
}

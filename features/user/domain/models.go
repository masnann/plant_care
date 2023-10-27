package domain

import (
	note "github.com/masnann/plant_care/features/note/domain"
	plant "github.com/masnann/plant_care/features/plant/domain"
	"time"
)

type UserModel struct {
	ID        uint64             `gorm:"column:id;primary_key" json:"id"`
	Username  string             `gorm:"column:username" json:"username"`
	Email     string             `gorm:"column:email" json:"email"`
	Password  string             `gorm:"column:password" json:"password"`
	CreatedAt time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time         `gorm:"index;column:deleted_at" json:"deleted_at"`
	Plants    []plant.PlantModel `gorm:"foreignKey:UserID"`
	Notes     []note.NoteModel   `gorm:"foreignKey:UserID"`
}

func (UserModel) TableName() string {
	return "users"
}

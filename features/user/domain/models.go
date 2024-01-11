package domain

import (
	note "github.com/masnann/plant_care/features/note/domain"
	notification "github.com/masnann/plant_care/features/notification/domain"
	plant "github.com/masnann/plant_care/features/plant/domain"
	"time"
)

type UserModel struct {
	ID           uint64                           `gorm:"column:id;primary_key" json:"id"`
	Username     string                           `gorm:"column:username" json:"username"`
	Email        string                           `gorm:"column:email" json:"email"`
	Password     string                           `gorm:"column:password" json:"password"`
	CreatedAt    time.Time                        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time                        `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    *time.Time                       `gorm:"index;column:deleted_at" json:"deleted_at"`
	IsVerified   bool                             `gorm:"column:is_verified" json:"is_verified"`
	Otp          OtpModels                        `gorm:"foreignKey:UserID"`
	Plants       []plant.PlantModel               `gorm:"foreignKey:UserID"`
	Notes        []note.NoteModel                 `gorm:"foreignKey:UserID"`
	Notification []notification.NotificationModel `gorm:"foreignKey:UserID"`
}

type OtpModels struct {
	ID     uint64 `gorm:"column:id;primary_key" json:"id"`
	UserID uint64 `gorm:"column:user_id" json:"user_id"`
	Email  string `gorm:"column:email" json:"email"`
	Code   string `gorm:"column:code" json:"code"`
}

func (UserModel) TableName() string {
	return "users"
}

func (OtpModels) TableName() string {
	return "otp"
}

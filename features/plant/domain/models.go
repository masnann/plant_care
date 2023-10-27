package domain

import "time"

type PlantModel struct {
	ID        uint64     `gorm:"column:id;primary_key" json:"id"`
	UserID    uint64     `gorm:"column:user_id" json:"user_id"`
	Name      string     `gorm:"column:name" json:"name"`
	Type      string     `gorm:"column:type" json:"type"`
	Date      time.Time  `gorm:"column:date" json:"date"`
	Photo     string     `gorm:"column:photo" json:"photo"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (PlantModel) TableName() string {
	return "plants"
}

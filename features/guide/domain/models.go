package domain

import "time"

type GuideModel struct {
	ID          uint64     `gorm:"column:id;primary_key" json:"id"`
	Title       string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Date        time.Time  `gorm:"column:date" json:"date"`
	Photo       string     `gorm:"column:photo" json:"photo"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (GuideModel) TableName() string {
	return "guides"
}

package domain

import "time"

type NoteModel struct {
	ID          uint64       `gorm:"column:id;primary_key" json:"id"`
	UserID      uint64       `gorm:"column:user_id" json:"user_id"`
	PlantID     uint64       `gorm:"column:plant_id" json:"plant_id"`
	Date        time.Time    `gorm:"column:date" json:"date"`
	Title       string       `gorm:"column:title" json:"title"`
	Description string       `gorm:"column:description" json:"description"`
	CreatedAt   time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time   `gorm:"index;column:deleted_at" json:"deleted_at"`
	Photos      []PhotoModel `gorm:"foreignKey:NoteID" json:"photos"`
}

type PhotoModel struct {
	PhotoID     uint64     `gorm:"column:photo_id;primary_key" json:"photo_id"`
	NoteID      uint64     `gorm:"column:note_id" json:"note_id"`
	Photo       string     `gorm:"column:photo" json:"photo"`
	Description string     `gorm:"column:description" json:"description"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (NoteModel) TableName() string {
	return "notes"
}

func (PhotoModel) TableName() string {
	return "photo_models"
}

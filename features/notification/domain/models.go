package domain

import "time"

type NotificationModel struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	UserID    uint64     `gorm:"column:user_id" json:"user_id"`
	Message   string     `gorm:"column:message" json:"message"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index;column:deleted_at" json:"deleted_at"`
}

func (NotificationModel) TableName() string {
	return "notification"
}

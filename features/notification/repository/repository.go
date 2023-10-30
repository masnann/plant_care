package repository

import (
	"github.com/masnann/plant_care/features/notification"
	"github.com/masnann/plant_care/features/notification/domain"
	"github.com/masnann/plant_care/utils"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) notification.RepoNotificationInterface {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) InsertNotifications(notifyModel *domain.NotificationModel) (*domain.NotificationModel, error) {
	newData := &domain.NotificationModel{
		ID:        notifyModel.ID,
		UserID:    notifyModel.UserID,
		Message:   notifyModel.Message,
		CreatedAt: utils.GetNowTime(),
		UpdatedAt: utils.GetNowTime(),
		DeletedAt: nil,
	}
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *NotificationRepository) GetPaginationNotifications(userID uint64, offset int, pageSize int, notify *[]*domain.NotificationModel) error {
	if err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&notify).Error; err != nil {
		return err
	}
	return nil
}

func (r *NotificationRepository) CountNotifications(userID uint64) (uint64, error) {
	var count int64
	if err := r.db.Model(&domain.NotificationModel{}).Where("user_id = ? AND deleted_at IS NULL", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint64(count), nil
}

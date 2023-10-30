package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/notification"
	"github.com/masnann/plant_care/features/notification/domain"
	user "github.com/masnann/plant_care/features/user/domain"
)

type NotificationService struct {
	repo notification.RepoNotificationInterface
}

func NewNotificationService(repo notification.RepoNotificationInterface) notification.ServiceNotificationInterface {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) InsertNotifications(c echo.Context, plantName string) (*domain.NotificationModel, error) {
	currentUser := c.Get("CurrentUser").(*user.UserModel)
	message := fmt.Sprintf("%s berhasil ditambahkan. Jangan lupa untuk merawatnya dengan menyiram dan menambahkannya dalam daftar catatan!", plantName)
	newData := &domain.NotificationModel{
		UserID:  currentUser.ID,
		Message: message,
	}

	result, err := s.repo.InsertNotifications(newData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NotificationService) GetPaginationNotifications(userID uint64, page, pageSize int) ([]*domain.NotificationModel, error) {
	var notify []*domain.NotificationModel
	if err := s.repo.GetPaginationNotifications(userID, page, pageSize, &notify); err != nil {
		return nil, err
	}
	return notify, nil
}

func (s *NotificationService) CountNotifications(userID uint64) (uint64, error) {
	count, err := s.repo.CountNotifications(userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

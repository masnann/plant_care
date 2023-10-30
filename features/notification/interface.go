package notification

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/notification/domain"
)

type RepoNotificationInterface interface {
	InsertNotifications(notifyModel *domain.NotificationModel) (*domain.NotificationModel, error)
	GetPaginationNotifications(userID uint64, offset int, pageSize int, notify *[]*domain.NotificationModel) error
	CountNotifications(userID uint64) (uint64, error)
}

type ServiceNotificationInterface interface {
	InsertNotifications(c echo.Context, plantName string) (*domain.NotificationModel, error)
	GetPaginationNotifications(userID uint64, page, pageSize int) ([]*domain.NotificationModel, error)
	CountNotifications(userID uint64) (uint64, error)
}

type HandlerNotificationInterface interface {
	GetPaginationNotifications() echo.HandlerFunc
}

package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/notification"
	"github.com/masnann/plant_care/features/notification/domain"
	user "github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils/response"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type NotificationHandler struct {
	service notification.ServiceNotificationInterface
}

func (h *NotificationHandler) GetPaginationNotifications() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser := c.Get("CurrentUser").(*user.UserModel)

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			logrus.Error("Invalid page parameter", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Invalid page parameter ")
		}

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil {
			logrus.Error("Invalid pageSize parameter", err.Error())
			return response.SendErrorResponse(c, http.StatusBadRequest, "Invalid pageSize parameter ")
		}
		offset := (page - 1) * pageSize

		notify, err := h.service.GetPaginationNotifications(currentUser.ID, offset, pageSize)
		if err != nil {
			logrus.Error("Internal server error", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
		var notifyResponse []domain.GetNotificationResponse
		for _, value := range notify {
			notifyResponse = append(notifyResponse, domain.GetNotificationResponse{
				ID:      value.ID,
				UserID:  value.UserID,
				Message: value.Message,
			})
		}
		totalItems, err := h.service.CountNotifications(currentUser.ID)
		if err != nil {
			logrus.Error("Failed to count plants", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		}
		return response.PaginationResponse(c, notifyResponse, int(totalItems), page, pageSize, "Success")
	}
}

func NewNotificationHandler(service notification.ServiceNotificationInterface) notification.HandlerNotificationInterface {
	return &NotificationHandler{
		service: service,
	}
}

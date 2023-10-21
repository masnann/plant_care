package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/utils/response"
	"net/http"
)

type UserHandler struct {
	service user.ServiceUserInterface
}

func NewUserHandler(service user.ServiceUserInterface) user.HandlerUserInterface {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := h.service.GetAllUsers()
		if err != nil {
			c.Logger().Error("handler: failed to fetch all users:", err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		}
		if len(result) == 0 {
			return response.SendSuccessResponse(c, "Success Get All Users", nil)
		} else {
			return response.SendSuccessResponse(c, "Success Get All Users", result)
		}
	}
}

func (h *UserHandler) GetUserByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.QueryParam("email")
		if email == "" {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Email parameter is missing")
		}
		user, err := h.service.GetUserByEmail(email)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusNotFound, "User not found")
		}

		return response.SendSuccessResponse(c, "Success", user)
	}
}

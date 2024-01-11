package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"net/http"
)

type UserHandler struct {
	service user.ServiceUserInterface
	jwt     utils.JWTInterface
}

func NewUserHandler(service user.ServiceUserInterface, jwt utils.JWTInterface) user.HandlerUserInterface {
	return &UserHandler{
		service: service,
		jwt:     jwt,
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

func (h *UserHandler) UpdatePassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateRequest domain.UpdatePasswordRequest
		if err := c.Bind(&updateRequest); err != nil {
			return c.JSON(http.StatusBadRequest, "Failed to bind data: Binding data to the struct failed")
		}
		if err := utils.ValidateStruct(updateRequest); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation failed: "+err.Error())
		}
		currentUser := c.Get("CurrentUser").(*domain.UserModel)
		updateRequest.UserID = currentUser.ID

		err := h.service.ValidatePassword(currentUser.ID, updateRequest.OldPassword, updateRequest.OldPassword)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Gagal memperbarui kata sandi: "+err.Error())
		}

		newPasswordHash, err := utils.GenerateHash(updateRequest.NewPassword)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to generate hash for the new password")
		}

		err = h.service.UpdatePassword(currentUser.ID, newPasswordHash)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to update the password: "+err.Error())
		}
		return c.JSON(http.StatusOK, "Password updated successfully")
	}
}

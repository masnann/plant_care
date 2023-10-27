package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/config"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/auth/domain"
	"github.com/masnann/plant_care/features/user"
	users "github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthHandler struct {
	service     auth.ServiceAuthInterface
	userService user.ServiceUserInterface
	jwt         utils.JWTInterface
	cfg         config.Config
}

func NewAuthHandler(service auth.ServiceAuthInterface, userService user.ServiceUserInterface, jwt utils.JWTInterface, cfg config.Config) auth.HandlerAuthInterface {
	return &AuthHandler{
		service:     service,
		userService: userService,
		jwt:         jwt,
		cfg:         cfg,
	}
}

func (h *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		registerRequest := new(domain.RegisterRequest)
		if err := c.Bind(registerRequest); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}

		if err := utils.ValidateStruct(registerRequest); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Validation failed : "+err.Error())
		}

		_, err := h.userService.GetUserByEmail(registerRequest.Email)
		if err == nil {
			return response.SendErrorResponse(c, http.StatusConflict, "Email already registered")
		}

		newUser := &users.UserModel{
			Username: registerRequest.Username,
			Email:    registerRequest.Email,
			Password: registerRequest.Password,
		}

		createdUser, err := h.service.Register(newUser)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		}
		result := domain.RegisterResponse{
			Username: createdUser.Username,
			Email:    createdUser.Email,
		}

		return response.SendSuccessResponse(c, "Success", result)
	}
}

func (h *AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginRequest domain.LoginRequest
		if err := c.Bind(&loginRequest); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}

		if err := utils.ValidateStruct(loginRequest); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Validation failed : "+err.Error())
		}

		login, accessToken, err := h.service.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			logrus.Error("User not found" + err.Error())
			return response.SendErrorResponse(c, http.StatusInternalServerError, "User not found")
		}

		userLogin := &domain.LoginResponse{
			Email:       login.Email,
			AccessToken: accessToken,
		}

		return response.SendSuccessResponse(c, "Login successful", userLogin)
	}
}

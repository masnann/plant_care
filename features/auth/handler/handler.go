package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/auth/domain"
	"github.com/masnann/plant_care/features/user"
	users "github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/response"
	"net/http"
)

type AuthHandler struct {
	service     auth.ServiceAuthInterface
	userService user.ServiceUserInterface
	jwt         utils.JWTInterface
}

func NewAuthHandler(service auth.ServiceAuthInterface, userService user.ServiceUserInterface, jwt utils.JWTInterface) auth.HandlerAuthInterface {
	return &AuthHandler{
		service:     service,
		userService: userService,
		jwt:         jwt,
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

		user, accessToken, refreshToken, err := h.service.Login(loginRequest.Email, loginRequest.Password)
		if err != nil {
			return response.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		}

		userLogin := &domain.LoginResponse{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return response.SendSuccessResponse(c, "Login successful", userLogin)
	}
}

func (h *AuthHandler) RefreshJWT() echo.HandlerFunc {
	return func(c echo.Context) error {
		type RefreshInput struct {
			Token string `json:"access_token" form:"access_token"`
		}
		var input = RefreshInput{}
		if err := c.Bind(&input); err != nil {
			return response.SendErrorResponse(c, http.StatusBadRequest, "Error Binding Data")
		}

		currentToken := c.Get("user").(*jwt.Token)
		result := h.jwt.GenerateRefreshJWT(currentToken, currentToken)

		if result == nil {
			return response.SendErrorResponse(c, http.StatusUnauthorized, "Failed to refresh tokens")
		}

		return response.SendSuccessResponse(c, "Tokens refreshed successfully", result)
	}
}

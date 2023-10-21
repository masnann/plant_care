package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user/domain"
)

type RepoAuthInterface interface {
	Register(newData *domain.UserModel) (*domain.UserModel, error)
	Login(email, password string) (*domain.UserModel, error)
}

type ServiceAuthInterface interface {
	Register(newData *domain.UserModel) (*domain.UserModel, error)
	Login(email, password string) (*domain.UserModel, string, string, error)
}

type HandlerAuthInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	RefreshJWT() echo.HandlerFunc
}

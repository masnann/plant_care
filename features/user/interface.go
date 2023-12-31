package user

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user/domain"
)

type RepoUserInterface interface {
	GetAllUsers() ([]*domain.UserModel, error)
	GetUserByEmail(email string) (*domain.UserModel, error)
	GetUserById(userId uint64) (*domain.UserModel, error)
}

type ServiceUserInterface interface {
	GetAllUsers() ([]*domain.UserModel, error)
	GetUserByEmail(email string) (*domain.UserModel, error)
	GetUserById(userId uint64) (*domain.UserModel, error)
}

type HandlerUserInterface interface {
	GetAllUsers() echo.HandlerFunc
	GetUserByEmail() echo.HandlerFunc
}

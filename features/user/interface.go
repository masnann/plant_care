package user

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user/domain"
)

type RepoUserInterface interface {
	GetAllUsers() ([]*domain.UserModel, error)
	GetUserByEmail(email string) (*domain.UserModel, error)
	GetUserById(userId uint64) (*domain.UserModel, error)
	GetUsersPassword(oldPass string) (*domain.UserModel, error)
	UpdatePassword(userID uint64, newPasswordHash string) error
}

type ServiceUserInterface interface {
	GetAllUsers() ([]*domain.UserModel, error)
	GetUserByEmail(email string) (*domain.UserModel, error)
	GetUserById(userId uint64) (*domain.UserModel, error)
	GetUsersPassword(oldPass string) (*domain.UserModel, error)
	ValidatePassword(userID uint64, oldPassword, newPassword string) error
	UpdatePassword(userID uint64, newPasswordHash string) error
}

type HandlerUserInterface interface {
	GetAllUsers() echo.HandlerFunc
	GetUserByEmail() echo.HandlerFunc
	UpdatePassword() echo.HandlerFunc
}

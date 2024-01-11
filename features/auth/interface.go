package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/user/domain"
)

type RepoAuthInterface interface {
	Register(newData *domain.UserModel) (*domain.UserModel, error)
	Login(email, password string) (*domain.UserModel, error)
	VerifyEmail(email string) error
	SaveOtp(userID uint64, email, otpCode string) error
	CheckOtp(email, code string) (*domain.OtpModels, error)
}

type ServiceAuthInterface interface {
	Register(newData *domain.UserModel) (*domain.UserModel, error)
	Login(email, password string) (*domain.UserModel, string, error)
	VerifyEmail(email, code string) error
}

type HandlerAuthInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	VerifyEmail() echo.HandlerFunc
}

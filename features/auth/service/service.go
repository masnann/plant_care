package service

import (
	"errors"
	"fmt"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"github.com/masnann/plant_care/utils/email"
)

type AuthService struct {
	repo  auth.RepoAuthInterface
	utils utils.JWTInterface
	email email.EmailServiceInterface
	otp   email.VerificationCodeGenerator
}

func NewAuthService(repo auth.RepoAuthInterface, utils utils.JWTInterface, email email.EmailServiceInterface, otp email.VerificationCodeGenerator) auth.ServiceAuthInterface {
	return &AuthService{
		repo:  repo,
		utils: utils,
		email: email,
		otp:   otp,
	}
}

func (s *AuthService) Register(newData *domain.UserModel) (*domain.UserModel, error) {
	result, err := s.repo.Register(newData)
	if err != nil {
		return nil, err
	}

	//expirationTime := time.Now().Add(2 * time.Minute)
	otp := s.otp.GenerateCode(6)
	fmt.Println(result.ID)
	err = s.repo.SaveOtp(result.ID, result.Email, otp)
	if err != nil {
		return nil, err
	}

	err = s.email.SendVerificationEmail([]string{result.Email}, otp, result.Email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *AuthService) Login(email, password string) (*domain.UserModel, string, error) {
	user, err := s.repo.Login(email, password)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := s.utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, accessToken, nil
}

func (s *AuthService) VerifyEmail(email, code string) error {
	otp, err := s.repo.CheckOtp(email, code)
	if err != nil {
		return errors.New("otp tidak ada")
	}

	err = s.repo.VerifyEmail(otp.Email)
	if err != nil {
		return errors.New("gagal")
	}

	return nil
}

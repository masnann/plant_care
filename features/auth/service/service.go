package service

import (
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
)

type AuthService struct {
	repo  auth.RepoAuthInterface
	utils utils.JWTInterface
}

func NewAuthService(repo auth.RepoAuthInterface, utils utils.JWTInterface) auth.ServiceAuthInterface {
	return &AuthService{
		repo:  repo,
		utils: utils,
	}
}

func (s *AuthService) Register(newData *domain.UserModel) (*domain.UserModel, error) {
	result, err := s.repo.Register(newData)
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

package service

import (
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
)

type UserService struct {
	repo user.RepoUserInterface
	jwt  utils.JWTInterface
}

func NewUserService(repo user.RepoUserInterface, jwt utils.JWTInterface) user.ServiceUserInterface {
	return &UserService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *UserService) GetAllUsers() ([]*domain.UserModel, error) {
	result, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) GetUserByEmail(email string) (*domain.UserModel, error) {
	result, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return result, nil
}

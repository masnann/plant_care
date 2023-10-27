package service

import (
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/features/user/domain"
)

type UserService struct {
	repo user.RepoUserInterface
}

func NewUserService(repo user.RepoUserInterface) user.ServiceUserInterface {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]*domain.UserModel, error) {
	result, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) GetUserById(userId uint64) (*domain.UserModel, error) {
	result, err := s.repo.GetUserById(userId)
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

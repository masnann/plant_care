package service

import (
	"errors"
	"fmt"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
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

func (s *UserService) GetUsersPassword(oldPass string) (*domain.UserModel, error) {
	result, err := s.repo.GetUsersPassword(oldPass)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) ValidatePassword(userID uint64, oldPassword, newPassword string) error {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return err
	}

	isValidPassword, err := utils.ComparePassword(user.Password, oldPassword)
	if err != nil || !isValidPassword {
		return errors.New("Password lama salah")
	}

	if oldPassword == newPassword {
		return errors.New("Password baru tidak boleh sama dengan password lama")
	}

	fmt.Println(user.Password)
	fmt.Println(newPassword)
	if err != nil || !isValidPassword {
		return errors.New("Password salah")
	}

	return nil
}

func (s *UserService) UpdatePassword(userID uint64, newPasswordHash string) error {
	user, err := s.repo.GetUserById(userID)
	if err != nil {
		return err
	}
	err = s.repo.UpdatePassword(user.ID, newPasswordHash)
	if err != nil {
		return err
	}

	return nil
}

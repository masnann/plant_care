package repository

import (
	"errors"
	"github.com/masnann/plant_care/features/user"
	"github.com/masnann/plant_care/features/user/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.RepoUserInterface {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) GetUserById(userId uint64) (*domain.UserModel, error) {
	var user domain.UserModel
	if err := r.db.Where("id", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("id not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]*domain.UserModel, error) {
	var users []*domain.UserModel
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*domain.UserModel, error) {
	var user domain.UserModel
	if err := r.db.Where("email", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, errors.New("email not found")
		}
		return nil, err
	}
	return &user, nil
}

package repository

import (
	"errors"
	"github.com/masnann/plant_care/features/auth"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/utils"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.RepoAuthInterface {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(newData *domain.UserModel) (*domain.UserModel, error) {
	hashPassword, err := utils.GenerateHash(newData.Password)
	if err != nil {
		return nil, err
	}
	dbData := &domain.UserModel{
		Username: newData.Username,
		Email:    newData.Email,
		Password: hashPassword,
	}
	if err := r.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *AuthRepository) Login(email, password string) (*domain.UserModel, error) {
	var user domain.UserModel
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	isValidPassword, err := utils.ComparePassword(user.Password, password)
	if err != nil || !isValidPassword {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

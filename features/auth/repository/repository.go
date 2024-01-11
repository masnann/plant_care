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

	newData.ID = dbData.ID

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

func (r *AuthRepository) CheckOtp(email, code string) (*domain.OtpModels, error) {
	var otp domain.OtpModels

	if err := r.db.Where("email = ? AND code = ?", code, email).First(&otp).Error; err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *AuthRepository) VerifyEmail(email string) error {
	var user domain.UserModel

	if err := r.db.Model(&user).Where("email = ?", email).Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) SaveOtp(userID uint64, email, otpCode string) error {
	otp := domain.OtpModels{
		UserID: userID,
		Email:  email,
		Code:   otpCode,
	}

	if err := r.db.Create(&otp).Error; err != nil {
		return err
	}
	return nil
}

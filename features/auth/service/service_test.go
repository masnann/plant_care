package service

import (
	"errors"
	"github.com/masnann/plant_care/features/auth/mocks"
	"github.com/masnann/plant_care/features/user/domain"
	mocks2 "github.com/masnann/plant_care/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {
	jwt := mocks2.NewJWTInterface(t)
	repo := mocks.NewRepoAuthInterface(t)
	service := NewAuthService(repo, jwt)
	newUser := domain.UserModel{
		Username: "user1",
		Email:    "user@mail.com",
		Password: "a",
	}

	t.Run("Success insert", func(t *testing.T) {
		repo.On("Register", &newUser).Return(&newUser, nil).Once()
		result, err := service.Register(&newUser)
		assert.Nil(t, err)
		assert.Equal(t, newUser.Username, result.Username)
		assert.Equal(t, newUser.Email, result.Email)
		assert.Equal(t, newUser.Password, result.Password)
		repo.AssertExpectations(t)
	})
	t.Run("Failed insert", func(t *testing.T) {
		expectedErr := errors.New("registration failed")
		repo.On("Register", &newUser).Return(nil, expectedErr).Once()
		result, err := service.Register(&newUser)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	jwt := mocks2.NewJWTInterface(t)
	repo := mocks.NewRepoAuthInterface(t)
	service := NewAuthService(repo, jwt)
	email := "user@example.com"
	password := "a"
	user := &domain.UserModel{
		ID:       1,
		Username: "user",
		Email:    email,
	}

	t.Run("Successful login", func(t *testing.T) {

		expectedAccessToken := "mocked-access-token"
		jwt.On("GenerateJWT", user.ID).Return(expectedAccessToken, nil).Once()
		repo.On("Login", email, password).Return(user, nil).Once()

		resultUser, accessToken, err := service.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, user, resultUser)
		assert.Equal(t, expectedAccessToken, accessToken)

		jwt.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong Password", func(t *testing.T) {
		email := "user@example.com"
		password := "invalid-password"

		repo.On("Login", email, password).Return(nil, errors.New("invalid credentials")).Once()

		resultUser, accessToken, err := service.Login(email, password)
		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, accessToken)

		repo.AssertExpectations(t)
	})

	t.Run("Token error", func(t *testing.T) {

		jwt.On("GenerateJWT", user.ID).Return("", errors.New("JWT generation error")).Once()
		repo.On("Login", email, password).Return(user, nil).Once()

		resultUser, accessToken, err := service.Login(email, password)
		assert.Error(t, err)
		assert.Nil(t, resultUser)
		assert.Empty(t, accessToken)

		jwt.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}

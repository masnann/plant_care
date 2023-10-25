package service

import (
	"errors"
	"github.com/masnann/plant_care/features/user/domain"
	"github.com/masnann/plant_care/features/user/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserById(t *testing.T) {
	repo := mocks.NewRepoUserInterface(t)
	service := NewUserService(repo)
	userID := uint64(1)

	newUser := &domain.UserModel{
		ID:       userID,
		Username: "user1",
		Email:    "user@mail.com",
		Password: "a",
	}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("GetUserById", userID).Return(newUser, nil).Once()

		result, err := service.GetUserById(userID)

		assert.Nil(t, err)
		assert.Equal(t, newUser.ID, result.ID)
		assert.Equal(t, newUser.Username, result.Username)
		assert.Equal(t, newUser.Email, result.Email)
		assert.Equal(t, newUser.Password, result.Password)

		repo.AssertExpectations(t)
	})
	t.Run("Failed Case", func(t *testing.T) {
		expectedErr := errors.New("GetUserById")
		repo.On("GetUserById", userID).Return(nil, expectedErr).Once()
		result, err := service.GetUserById(userID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

}

func TestGetUserByEmail(t *testing.T) {
	repo := mocks.NewRepoUserInterface(t)
	service := NewUserService(repo)
	email := "user1@mail.com"

	t.Run("Success Case", func(t *testing.T) {

		expectedUser := &domain.UserModel{
			ID:       1,
			Username: "user1",
			Email:    email,
			Password: "a",
		}

		repo.On("GetUserByEmail", email).Return(expectedUser, nil).Once()

		result, err := service.GetUserByEmail(email)

		assert.Nil(t, err)
		assert.Equal(t, expectedUser.ID, result.ID)
		assert.Equal(t, expectedUser.Username, result.Username)
		assert.Equal(t, expectedUser.Email, result.Email)
		assert.Equal(t, expectedUser.Password, result.Password)

		repo.AssertExpectations(t)
	})
	t.Run("Failed Case", func(t *testing.T) {
		expectedErr := errors.New("GetUserByEmail")
		repo.On("GetUserByEmail", email).Return(nil, expectedErr).Once()
		result, err := service.GetUserByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllUsers(t *testing.T) {
	repo := mocks.NewRepoUserInterface(t)
	service := NewUserService(repo)

	t.Run("Success Case", func(t *testing.T) {
		expectedUsers := []*domain.UserModel{
			{
				ID:       1,
				Username: "user1",
				Email:    "user1@mail.com",
				Password: "a",
			},
			{
				ID:       2,
				Username: "user2",
				Email:    "user2@mail.com",
				Password: "b",
			},
		}

		repo.On("GetAllUsers").Return(expectedUsers, nil).Once()

		result, err := service.GetAllUsers()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedUsers))
		for i, user := range expectedUsers {
			assert.Equal(t, user.ID, result[i].ID)
			assert.Equal(t, user.Username, result[i].Username)
			assert.Equal(t, user.Email, result[i].Email)
			assert.Equal(t, user.Password, result[i].Password)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		expectedErr := errors.New("GetAllUsers")
		repo.On("GetAllUsers").Return(nil, expectedErr).Once()

		result, err := service.GetAllUsers()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

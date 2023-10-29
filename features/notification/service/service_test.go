package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/notification/domain"
	"github.com/masnann/plant_care/features/notification/mocks"
	user "github.com/masnann/plant_care/features/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotificationService_InsertNotifications(t *testing.T) {
	repo := mocks.NewRepoNotificationInterface(t)
	service := NewNotificationService(repo)
	currentUser := &user.UserModel{
		ID: 1,
	}
	plantName := "Pohon Mangga"

	notify := &domain.NotificationModel{
		UserID:  currentUser.ID,
		Message: "Pohon Mangga berhasil ditambahkan. Jangan lupa untuk merawatnya dengan menyiram dan menambahkannya dalam daftar catatan!",
	}

	t.Run("Success Case", func(t *testing.T) {

		repo.On("InsertNotifications", notify).Return(notify, nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("CurrentUser", currentUser)
		result, err := service.InsertNotifications(c, plantName)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, notify.UserID, result.UserID)
		assert.Equal(t, notify.Message, result.Message)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {

		expectedErr := errors.New("InsertNotifications failed")
		repo.On("InsertNotifications", mock.Anything).Return(nil, expectedErr).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("CurrentUser", currentUser)

		result, err := service.InsertNotifications(c, plantName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNotificationService_GetPaginationNotifications(t *testing.T) {
	repo := mocks.NewRepoNotificationInterface(t)
	service := NewNotificationService(repo)
	userID := uint64(1)
	offset := 0
	pageSize := 10

	expectedPlants := []*domain.NotificationModel{
		{
			ID:      1,
			UserID:  userID,
			Message: "Pohon Mangga berhasil ditambahkan. Jangan lupa untuk merawatnya dengan menyiram dan menambahkannya dalam daftar catatan!",
		},
		{
			ID:      2,
			UserID:  userID,
			Message: "Pohon Mangga berhasil ditambahkan. Jangan lupa untuk merawatnya dengan menyiram dan menambahkannya dalam daftar catatan!",
		},
	}

	t.Run("Success Case", func(t *testing.T) {
		
		repo.On("GetPaginationNotifications", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.NotificationModel")).Run(func(args mock.Arguments) {
			plants := args.Get(3).(*[]*domain.NotificationModel)
			*plants = expectedPlants
		}).Return(nil).Once()

		result, err := service.GetPaginationNotifications(userID, offset, pageSize)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedPlants))
		for i, plant := range expectedPlants {
			assert.Equal(t, plant.ID, result[i].ID)
			assert.Equal(t, plant.UserID, result[i].UserID)
			assert.Equal(t, plant.Message, result[i].Message)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		expectedErr := errors.New("GetPaginationNotifications")
		repo.On("GetPaginationNotifications", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.NotificationModel")).Return(expectedErr).Once()

		result, err := service.GetPaginationNotifications(userID, offset, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNotificationService_CountNotifications(t *testing.T) {
	repo := mocks.NewRepoNotificationInterface(t)
	service := NewNotificationService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)

		expectedCount := uint64(5)

		repo.On("CountNotifications", userID).Return(expectedCount, nil).Once()

		result, err := service.CountNotifications(userID)

		assert.Nil(t, err)
		assert.Equal(t, expectedCount, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)

		expectedErr := errors.New("CountNotifications")
		repo.On("CountNotifications", userID).Return(uint64(0), expectedErr).Once()

		result, err := service.CountNotifications(userID)

		assert.Error(t, err)
		assert.Equal(t, uint64(0), result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

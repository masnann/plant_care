package service

import (
	"errors"
	"github.com/masnann/plant_care/features/plant/domain"
	"github.com/masnann/plant_care/features/plant/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPlantService_SearchPlantsByType(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userId := uint64(1)
		plantType := "Flower"

		expectedPlants := []*domain.PlantModel{
			{
				ID:   1,
				Name: "Rose",
				Type: "Flower",
			},
			{
				ID:   2,
				Name: "Lily",
				Type: "Flower",
			},
		}
		repo.On("SearchPlantsByType", userId, plantType).Return(expectedPlants, nil).Once()

		result, err := service.SearchPlantsByType(userId, plantType)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedPlants))
		for i, plant := range expectedPlants {
			assert.Equal(t, plant.ID, result[i].ID)
			assert.Equal(t, plant.Name, result[i].Name)
			assert.Equal(t, plant.Type, result[i].Type)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userId := uint64(1)
		plantType := "Tree"

		expectedErr := errors.New("SearchPlantsByType")
		repo.On("SearchPlantsByType", userId, plantType).Return(nil, expectedErr).Once()

		result, err := service.SearchPlantsByType(userId, plantType)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

}

func TestPlantService_SearchPlantsByName(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userId := uint64(1)
		plantName := "Rose"

		expectedPlants := []*domain.PlantModel{
			{
				ID:   1,
				Name: "Rose",
				Type: "Flower",
			},
			{
				ID:   2,
				Name: "Lily",
				Type: "Flower",
			},
		}
		repo.On("SearchPlantsByName", userId, plantName).Return(expectedPlants, nil).Once()

		result, err := service.SearchPlantsByName(userId, plantName)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedPlants))
		for i, plant := range expectedPlants {
			assert.Equal(t, plant.ID, result[i].ID)
			assert.Equal(t, plant.Name, result[i].Name)
			assert.Equal(t, plant.Type, result[i].Type)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userId := uint64(1)
		plantName := "Tree"

		expectedErr := errors.New("SearchPlantsByName")
		repo.On("SearchPlantsByName", userId, plantName).Return(nil, expectedErr).Once()

		result, err := service.SearchPlantsByName(userId, plantName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

}

func TestPlantService_GetPlantsWithPagination(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)
		offset := 0
		pageSize := 10

		expectedPlants := []*domain.PlantModel{
			{
				ID:   1,
				Name: "Rose",
				Type: "Flower",
			},
			{
				ID:   2,
				Name: "Lily",
				Type: "Flower",
			},
		}

		repo.On("GetPlantsWithPagination", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.PlantModel")).Run(func(args mock.Arguments) {
			plants := args.Get(3).(*[]*domain.PlantModel)
			*plants = expectedPlants
		}).Return(nil).Once()

		result, err := service.GetPlantsWithPagination(userID, offset, pageSize)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedPlants))
		for i, plant := range expectedPlants {
			assert.Equal(t, plant.ID, result[i].ID)
			assert.Equal(t, plant.Name, result[i].Name)
			assert.Equal(t, plant.Type, result[i].Type)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)
		offset := 10
		pageSize := 5

		expectedErr := errors.New("GetPlantsWithPagination")
		repo.On("GetPlantsWithPagination", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.PlantModel")).Return(expectedErr).Once()

		result, err := service.GetPlantsWithPagination(userID, offset, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_InsertPlants(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		plant := &domain.PlantModel{
			UserID: 1,
			Name:   "Rose",
			Type:   "Flower",
			Photo:  "rose.jpg",
		}

		expectedPlant := &domain.PlantModel{
			ID:     1,
			UserID: plant.UserID,
			Name:   plant.Name,
			Type:   plant.Type,
			Photo:  plant.Photo,
		}

		repo.On("InsertPlants", plant).Return(expectedPlant, nil).Once()

		result, err := service.InsertPlants(plant)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedPlant.ID, result.ID)
		assert.Equal(t, expectedPlant.UserID, result.UserID)
		assert.Equal(t, expectedPlant.Name, result.Name)
		assert.Equal(t, expectedPlant.Type, result.Type)
		assert.Equal(t, expectedPlant.Photo, result.Photo)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		plant := &domain.PlantModel{
			UserID: 1,
			Name:   "Tulip",
			Type:   "Flower",
			Photo:  "tulip.jpg",
		}

		expectedErr := errors.New("InsertPlants")
		repo.On("InsertPlants", plant).Return(nil, expectedErr).Once()

		result, err := service.InsertPlants(plant)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_UpdatePlants(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		plant := &domain.PlantModel{
			ID:     1,
			UserID: 1,
			Name:   "Rose",
			Type:   "Flower",
			Photo:  "rose.jpg",
		}

		updatedPlant := &domain.PlantModel{
			ID:     plant.ID,
			UserID: plant.UserID,
			Name:   "Updated Rose",
			Type:   "Updated Flower",
			Photo:  "updated_rose.jpg",
		}

		repo.On("UpdatePlants", updatedPlant).Return(updatedPlant, nil).Once()

		result, err := service.UpdatePlants(updatedPlant)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedPlant.ID, result.ID)
		assert.Equal(t, updatedPlant.UserID, result.UserID)
		assert.Equal(t, updatedPlant.Name, result.Name)
		assert.Equal(t, updatedPlant.Type, result.Type)
		assert.Equal(t, updatedPlant.Photo, result.Photo)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		plant := &domain.PlantModel{
			ID:     1,
			UserID: 1,
			Name:   "Tulip",
			Type:   "Flower",
			Photo:  "tulip.jpg",
		}

		expectedErr := errors.New("UpdatePlants")
		repo.On("UpdatePlants", plant).Return(nil, expectedErr).Once()

		result, err := service.UpdatePlants(plant)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_DeletePlants(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		plantID := uint64(1)

		repo.On("DeletePlants", plantID).Return(nil).Once()

		err := service.DeletePlants(plantID)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		plantID := uint64(2)

		expectedErr := errors.New("DeletePlants")
		repo.On("DeletePlants", plantID).Return(expectedErr).Once()

		err := service.DeletePlants(plantID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_GetPlantsByID(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		plantID := uint(1)

		expectedPlant := &domain.PlantModel{
			ID:     1,
			UserID: 1,
			Name:   "Rose",
			Type:   "Flower",
			Photo:  "rose.jpg",
		}

		repo.On("GetPlantsByID", plantID).Return(expectedPlant, nil).Once()

		result, err := service.GetPlantsByID(plantID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedPlant.ID, result.ID)
		assert.Equal(t, expectedPlant.UserID, result.UserID)
		assert.Equal(t, expectedPlant.Name, result.Name)
		assert.Equal(t, expectedPlant.Type, result.Type)
		assert.Equal(t, expectedPlant.Photo, result.Photo)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		plantID := uint(2)

		expectedErr := errors.New("GetPlantsByID")
		repo.On("GetPlantsByID", plantID).Return(nil, expectedErr).Once()

		result, err := service.GetPlantsByID(plantID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_GetPlantsByUserID(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)

		expectedPlants := []*domain.PlantModel{
			{
				ID:     1,
				UserID: userID,
				Name:   "Rose",
				Type:   "Flower",
				Photo:  "rose.jpg",
			},
			{
				ID:     2,
				UserID: userID,
				Name:   "Tulip",
				Type:   "Flower",
				Photo:  "tulip.jpg",
			},
		}

		repo.On("GetPlantsByUserID", userID).Return(expectedPlants, nil).Once()

		result, err := service.GetPlantsByUserID(userID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(expectedPlants), len(result))

		for i, expectedPlant := range expectedPlants {
			assert.Equal(t, expectedPlant.ID, result[i].ID)
			assert.Equal(t, expectedPlant.UserID, result[i].UserID)
			assert.Equal(t, expectedPlant.Name, result[i].Name)
			assert.Equal(t, expectedPlant.Type, result[i].Type)
			assert.Equal(t, expectedPlant.Photo, result[i].Photo)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)

		expectedErr := errors.New("GetPlantsByUserID")
		repo.On("GetPlantsByUserID", userID).Return(nil, expectedErr).Once()

		result, err := service.GetPlantsByUserID(userID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_CountPlants(t *testing.T) {
	repo := mocks.NewRepoPlantInterface(t)
	service := NewPlantService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)

		expectedCount := uint64(5)

		repo.On("CountPlants", userID).Return(expectedCount, nil).Once()

		result, err := service.CountPlants(userID)

		assert.Nil(t, err)
		assert.Equal(t, expectedCount, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)

		expectedErr := errors.New("CountPlants")
		repo.On("CountPlants", userID).Return(uint64(0), expectedErr).Once()

		result, err := service.CountPlants(userID)

		assert.Error(t, err)
		assert.Equal(t, uint64(0), result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

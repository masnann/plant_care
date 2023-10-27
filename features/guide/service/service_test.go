package service

import (
	"errors"
	"github.com/masnann/plant_care/features/guide/domain"
	"github.com/masnann/plant_care/features/guide/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGuideService_GetPaginatedGuides(t *testing.T) {
	repo := mocks.NewRepoGuideInterface(t)
	service := NewGuideService(repo)
	page := 1
	pageSize := 10

	expectedGuides := []domain.GuideModel{
		{
			ID:          1,
			Title:       "Guide 1",
			Description: "Description 1",
		},
		{
			ID:          2,
			Title:       "Guide 2",
			Description: "Description 2",
		},
	}

	t.Run("Success Case", func(t *testing.T) {
		repo.On("GetGuidesWithPagination", page, pageSize).Return(expectedGuides, nil).Once()

		result, err := service.GetGuidesWithPagination(page, pageSize)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedGuides))
		for i, guide := range expectedGuides {
			assert.Equal(t, guide.ID, result[i].ID)
			assert.Equal(t, guide.Title, result[i].Title)
			assert.Equal(t, guide.Description, result[i].Description)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		expectedErr := errors.New("GetGuidesWithPagination")
		repo.On("GetGuidesWithPagination", page, pageSize).Return(nil, expectedErr).Once()

		result, err := service.GetGuidesWithPagination(page, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestPlantService_CountPlants(t *testing.T) {
	repo := mocks.NewRepoGuideInterface(t)
	service := NewGuideService(repo)

	t.Run("Success Case", func(t *testing.T) {

		expectedCount := int64(5)

		repo.On("CountGuides").Return(expectedCount, nil).Once()

		result, err := service.CountGuides()

		assert.Nil(t, err)
		assert.Equal(t, expectedCount, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {

		expectedErr := errors.New("CountGuides")
		repo.On("CountGuides").Return(int64(0), expectedErr).Once()

		result, err := service.CountGuides()

		assert.Error(t, err)
		assert.Equal(t, int64(0), result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestGuideService_SearchGuideByName(t *testing.T) {
	repo := mocks.NewRepoGuideInterface(t)
	service := NewGuideService(repo)
	expectedGuides := []*domain.GuideModel{
		{
			ID:          1,
			Title:       "Guide to Planting Roses",
			Description: "This guide will help you plant and care for roses in your garden.",
			Date:        time.Date(2023, 10, 26, 22, 25, 55, 48000000, time.FixedZone("WIB", 7*3600)),
			Photo:       "https://res.cloudinary.com/dufa4bel6/image/upload/v1698333954/plantcare/nf44sqhzhd3hfjyjhbph.jpg",
		},
		{
			ID:          2,
			Title:       "Guide to Caring for Orange Trees",
			Description: "Learn how to care for orange trees so they can grow lush and bear abundant fruit.",
			Date:        time.Date(2023, 10, 26, 22, 25, 55, 48000000, time.FixedZone("WIB", 7*3600)),
			Photo:       "https://res.cloudinary.com/dufa4bel6/image/upload/v1698333954/plantcare/nf44sqhzhd3hfjyjhbph.jpg",
		},
	}
	t.Run("Success Case", func(t *testing.T) {
		guideName := "Rose"

		repo.On("SearchGuideByName", guideName).Return(expectedGuides, nil).Once()

		result, err := service.SearchGuideByName(guideName)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedGuides))
		for i, guide := range expectedGuides {
			assert.Equal(t, guide.ID, result[i].ID)
			assert.Equal(t, guide.Title, result[i].Title)
			assert.Equal(t, guide.Description, result[i].Description)
			assert.Equal(t, guide.Date, result[i].Date)
			assert.Equal(t, guide.Photo, result[i].Photo)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		guideName := "Rose"

		expectedErr := errors.New("SearchGuideByName")
		repo.On("SearchGuideByName", guideName).Return(nil, expectedErr).Once()

		result, err := service.SearchGuideByName(guideName)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestGuideService_GetByIdGuides(t *testing.T) {
	repo := mocks.NewRepoGuideInterface(t)
	service := NewGuideService(repo)

	expectedGuide := &domain.GuideModel{
		ID:          1,
		Title:       "Guide to Planting Roses",
		Description: "This guide will help you plant and care for roses in your garden.",
		Date:        time.Date(2023, 10, 26, 22, 25, 55, 48000000, time.FixedZone("WIB", 7*3600)),
		Photo:       "https://res.cloudinary.com/dufa4bel6/image/upload/v1698333954/plantcare/nf44sqhzhd3hfjyjhbph.jpg",
	}

	t.Run("Success Case", func(t *testing.T) {
		guideId := uint64(1)

		repo.On("GetByIdGuides", guideId).Return(expectedGuide, nil).Once()

		result, err := service.GetByIdGuides(guideId)
		assert.Nil(t, err)
		assert.Equal(t, expectedGuide.ID, result.ID)
		assert.Equal(t, expectedGuide.Title, result.Title)
		assert.Equal(t, expectedGuide.Description, result.Description)
		assert.Equal(t, expectedGuide.Date, result.Date)
		assert.Equal(t, expectedGuide.Photo, result.Photo)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		guideId := uint64(1)

		expectedErr := errors.New("GetByIdGuides")
		repo.On("GetByIdGuides", guideId).Return(nil, expectedErr).Once()

		result, err := service.GetByIdGuides(guideId)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

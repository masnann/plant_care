package guide

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/guide/domain"
)

type RepoGuideInterface interface {
	GetGuidesWithPagination(page int, pageSize int) ([]domain.GuideModel, error)
	CountGuides() (int64, error)
	GetByIdGuides(id uint64) (*domain.GuideModel, error)
	SearchGuideByName(name string) ([]*domain.GuideModel, error)
}

type ServiceGuideInterface interface {
	GetGuidesWithPagination(page int, pageSize int) ([]domain.GuideModel, error)
	CountGuides() (int64, error)
	GetByIdGuides(id uint64) (*domain.GuideModel, error)
	SearchGuideByName(name string) ([]*domain.GuideModel, error)
}

type HandlerGuideInterface interface {
	GetGuidesWithPagination() echo.HandlerFunc
	GetGuidesById() echo.HandlerFunc
	SearchGuideByName() echo.HandlerFunc
}

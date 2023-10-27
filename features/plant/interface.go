package plant

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/plant/domain"
)

type RepoPlantInterface interface {
	InsertPlants(plant *domain.PlantModel) (*domain.PlantModel, error)
	GetPlantsByUserID(userID uint64) ([]*domain.PlantModel, error)
	GetPlantsWithPagination(userID uint64, offset, pageSize int, plants *[]*domain.PlantModel) error
	UpdatePlants(plant *domain.PlantModel) (*domain.PlantModel, error)
	GetPlantsByID(plantID uint) (*domain.PlantModel, error)
	DeletePlants(plantID uint64) error
	SearchPlantsByName(userID uint64, name string) ([]*domain.PlantModel, error)
	SearchPlantsByType(userIO uint64, types string) ([]*domain.PlantModel, error)
	CountPlants(userID uint64) (uint64, error)
}

type ServicePlantInterface interface {
	InsertPlants(plant *domain.PlantModel) (*domain.PlantModel, error)
	GetPlantsByUserID(userID uint64) ([]*domain.PlantModel, error)
	GetPlantsWithPagination(userID uint64, offset, pageSize int) ([]*domain.PlantModel, error)
	UpdatePlants(plant *domain.PlantModel) (*domain.PlantModel, error)
	GetPlantsByID(plantID uint) (*domain.PlantModel, error)
	DeletePlants(plantID uint64) error
	SearchPlantsByName(userID uint64, name string) ([]*domain.PlantModel, error)
	SearchPlantsByType(userIO uint64, types string) ([]*domain.PlantModel, error)
	CountPlants(userID uint64) (uint64, error)
}

type HandlerPlantInterface interface {
	InsertPlants() echo.HandlerFunc
	GetPaginationPlants() echo.HandlerFunc
	UpdatePlants() echo.HandlerFunc
	DeletePlants() echo.HandlerFunc
	SearchPlantsByName() echo.HandlerFunc
	SearchPlantsByType() echo.HandlerFunc
}

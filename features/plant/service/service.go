package service

import (
	"github.com/masnann/plant_care/features/plant"
	"github.com/masnann/plant_care/features/plant/domain"
)

type PlantService struct {
	repo plant.RepoPlantInterface
}

func NewPlantService(repo plant.RepoPlantInterface) plant.ServicePlantInterface {
	return &PlantService{
		repo: repo,
	}
}
func (s *PlantService) SearchPlantsByType(types string) ([]*domain.PlantModel, error) {
	result, err := s.repo.SearchPlantsByType(types)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PlantService) SearchPlantsByName(name string) ([]*domain.PlantModel, error) {
	result, err := s.repo.SearchPlantsByName(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PlantService) GetPlantsWithPagination(userID uint64, offset, pageSize int) ([]*domain.PlantModel, error) {
	var plants []*domain.PlantModel
	if err := s.repo.GetPlantsWithPagination(userID, offset, pageSize, &plants); err != nil {
		return nil, err
	}
	return plants, nil
}

func (s *PlantService) InsertPlants(plant *domain.PlantModel) (*domain.PlantModel, error) {
	result, err := s.repo.InsertPlants(plant)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PlantService) UpdatePlants(plant *domain.PlantModel) (*domain.PlantModel, error) {
	result, err := s.repo.UpdatePlants(plant)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *PlantService) DeletePlants(plantID uint64) error {
	err := s.repo.DeletePlants(plantID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PlantService) GetPlantsByID(plantID uint) (*domain.PlantModel, error) {
	result, err := s.repo.GetPlantsByID(plantID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PlantService) GetPlantsByUserID(userID uint64) ([]*domain.PlantModel, error) {
	result, err := s.repo.GetPlantsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *PlantService) CountPlants(userID uint64) (uint64, error) {
	count, err := s.repo.CountPlants(userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

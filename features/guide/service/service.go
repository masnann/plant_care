package service

import (
	"github.com/masnann/plant_care/features/guide"
	"github.com/masnann/plant_care/features/guide/domain"
)

type GuideService struct {
	repo guide.RepoGuideInterface
}

func NewGuideService(repo guide.RepoGuideInterface) guide.ServiceGuideInterface {
	return &GuideService{
		repo: repo,
	}
}
func (s *GuideService) SearchGuideByName(name string) ([]*domain.GuideModel, error) {
	result, err := s.repo.SearchGuideByName(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GuideService) GetByIdGuides(id uint64) (*domain.GuideModel, error) {
	result, err := s.repo.GetByIdGuides(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GuideService) GetGuidesWithPagination(page int, pageSize int) ([]domain.GuideModel, error) {
	result, err := s.repo.GetGuidesWithPagination(page, pageSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GuideService) CountGuides() (int64, error) {
	result, err := s.repo.CountGuides()
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/masnann/plant_care/features/plant/domain"
	mock "github.com/stretchr/testify/mock"
)

// RepoPlantInterface is an autogenerated mock type for the RepoPlantInterface type
type RepoPlantInterface struct {
	mock.Mock
}

// CountPlants provides a mock function with given fields: userID
func (_m *RepoPlantInterface) CountPlants(userID uint64) (uint64, error) {
	ret := _m.Called(userID)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (uint64, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePlants provides a mock function with given fields: plantID
func (_m *RepoPlantInterface) DeletePlants(plantID uint64) error {
	ret := _m.Called(plantID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(plantID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPlantsByID provides a mock function with given fields: plantID
func (_m *RepoPlantInterface) GetPlantsByID(plantID uint) (*domain.PlantModel, error) {
	ret := _m.Called(plantID)

	var r0 *domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.PlantModel, error)); ok {
		return rf(plantID)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.PlantModel); ok {
		r0 = rf(plantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(plantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlantsByUserID provides a mock function with given fields: userID
func (_m *RepoPlantInterface) GetPlantsByUserID(userID uint64) ([]*domain.PlantModel, error) {
	ret := _m.Called(userID)

	var r0 []*domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*domain.PlantModel, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*domain.PlantModel); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlantsWithPagination provides a mock function with given fields: userID, offset, pageSize, plants
func (_m *RepoPlantInterface) GetPlantsWithPagination(userID uint64, offset int, pageSize int, plants *[]*domain.PlantModel) error {
	ret := _m.Called(userID, offset, pageSize, plants)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, int, int, *[]*domain.PlantModel) error); ok {
		r0 = rf(userID, offset, pageSize, plants)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertPlants provides a mock function with given fields: _a0
func (_m *RepoPlantInterface) InsertPlants(_a0 *domain.PlantModel) (*domain.PlantModel, error) {
	ret := _m.Called(_a0)

	var r0 *domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PlantModel) (*domain.PlantModel, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*domain.PlantModel) *domain.PlantModel); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PlantModel) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPlantsByName provides a mock function with given fields: userID, name
func (_m *RepoPlantInterface) SearchPlantsByName(userID uint64, name string) ([]*domain.PlantModel, error) {
	ret := _m.Called(userID, name)

	var r0 []*domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, string) ([]*domain.PlantModel, error)); ok {
		return rf(userID, name)
	}
	if rf, ok := ret.Get(0).(func(uint64, string) []*domain.PlantModel); ok {
		r0 = rf(userID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(userID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchPlantsByType provides a mock function with given fields: userIO, types
func (_m *RepoPlantInterface) SearchPlantsByType(userIO uint64, types string) ([]*domain.PlantModel, error) {
	ret := _m.Called(userIO, types)

	var r0 []*domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, string) ([]*domain.PlantModel, error)); ok {
		return rf(userIO, types)
	}
	if rf, ok := ret.Get(0).(func(uint64, string) []*domain.PlantModel); ok {
		r0 = rf(userIO, types)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(userIO, types)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePlants provides a mock function with given fields: _a0
func (_m *RepoPlantInterface) UpdatePlants(_a0 *domain.PlantModel) (*domain.PlantModel, error) {
	ret := _m.Called(_a0)

	var r0 *domain.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PlantModel) (*domain.PlantModel, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*domain.PlantModel) *domain.PlantModel); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PlantModel) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepoPlantInterface creates a new instance of RepoPlantInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepoPlantInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepoPlantInterface {
	mock := &RepoPlantInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

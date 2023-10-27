// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/masnann/plant_care/features/note/domain"
	mock "github.com/stretchr/testify/mock"
)

// ServiceNoteInterface is an autogenerated mock type for the ServiceNoteInterface type
type ServiceNoteInterface struct {
	mock.Mock
}

// CountNotes provides a mock function with given fields: userID
func (_m *ServiceNoteInterface) CountNotes(userID uint64) (uint64, error) {
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

// DeleteNotes provides a mock function with given fields: noteID
func (_m *ServiceNoteInterface) DeleteNotes(noteID uint64) error {
	ret := _m.Called(noteID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(noteID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNotesPhotos provides a mock function with given fields: photoID
func (_m *ServiceNoteInterface) DeleteNotesPhotos(photoID uint64) error {
	ret := _m.Called(photoID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(photoID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNoteByID provides a mock function with given fields: noteID
func (_m *ServiceNoteInterface) GetNoteByID(noteID uint64) (*domain.NoteModel, error) {
	ret := _m.Called(noteID)

	var r0 *domain.NoteModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*domain.NoteModel, error)); ok {
		return rf(noteID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *domain.NoteModel); ok {
		r0 = rf(noteID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.NoteModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(noteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotePhotoByID provides a mock function with given fields: noteID
func (_m *ServiceNoteInterface) GetNotePhotoByID(noteID uint64) (*domain.PhotoModel, error) {
	ret := _m.Called(noteID)

	var r0 *domain.PhotoModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*domain.PhotoModel, error)); ok {
		return rf(noteID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *domain.PhotoModel); ok {
		r0 = rf(noteID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PhotoModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(noteID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotesWithPagination provides a mock function with given fields: userID, offset, pageSize
func (_m *ServiceNoteInterface) GetNotesWithPagination(userID uint64, offset int, pageSize int) ([]*domain.NoteModel, error) {
	ret := _m.Called(userID, offset, pageSize)

	var r0 []*domain.NoteModel
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, int, int) ([]*domain.NoteModel, error)); ok {
		return rf(userID, offset, pageSize)
	}
	if rf, ok := ret.Get(0).(func(uint64, int, int) []*domain.NoteModel); ok {
		r0 = rf(userID, offset, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.NoteModel)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, int, int) error); ok {
		r1 = rf(userID, offset, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNote provides a mock function with given fields: noteModel
func (_m *ServiceNoteInterface) InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error) {
	ret := _m.Called(noteModel)

	var r0 *domain.NoteModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.NoteModel) (*domain.NoteModel, error)); ok {
		return rf(noteModel)
	}
	if rf, ok := ret.Get(0).(func(*domain.NoteModel) *domain.NoteModel); ok {
		r0 = rf(noteModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.NoteModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.NoteModel) error); ok {
		r1 = rf(noteModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertNotePhoto provides a mock function with given fields: photoModel
func (_m *ServiceNoteInterface) InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error) {
	ret := _m.Called(photoModel)

	var r0 *domain.PhotoModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PhotoModel) (*domain.PhotoModel, error)); ok {
		return rf(photoModel)
	}
	if rf, ok := ret.Get(0).(func(*domain.PhotoModel) *domain.PhotoModel); ok {
		r0 = rf(photoModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PhotoModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PhotoModel) error); ok {
		r1 = rf(photoModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNotes provides a mock function with given fields: notes
func (_m *ServiceNoteInterface) UpdateNotes(notes *domain.NoteModel) (*domain.NoteModel, error) {
	ret := _m.Called(notes)

	var r0 *domain.NoteModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.NoteModel) (*domain.NoteModel, error)); ok {
		return rf(notes)
	}
	if rf, ok := ret.Get(0).(func(*domain.NoteModel) *domain.NoteModel); ok {
		r0 = rf(notes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.NoteModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.NoteModel) error); ok {
		r1 = rf(notes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNotesPhotos provides a mock function with given fields: photos
func (_m *ServiceNoteInterface) UpdateNotesPhotos(photos *domain.PhotoModel) (*domain.PhotoModel, error) {
	ret := _m.Called(photos)

	var r0 *domain.PhotoModel
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.PhotoModel) (*domain.PhotoModel, error)); ok {
		return rf(photos)
	}
	if rf, ok := ret.Get(0).(func(*domain.PhotoModel) *domain.PhotoModel); ok {
		r0 = rf(photos)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.PhotoModel)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.PhotoModel) error); ok {
		r1 = rf(photos)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServiceNoteInterface creates a new instance of ServiceNoteInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceNoteInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceNoteInterface {
	mock := &ServiceNoteInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
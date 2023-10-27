package service

import (
	"errors"
	"github.com/masnann/plant_care/features/note/domain"
	"github.com/masnann/plant_care/features/note/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNoteService_GetNotesWithPagination(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)
		offset := 0
		pageSize := 10

		expectedNotes := []*domain.NoteModel{
			{
				ID:          1,
				UserID:      userID,
				PlantID:     1,
				Date:        time.Now(),
				Title:       "Note 1",
				Description: "Description 1",
			},
			{
				ID:          2,
				UserID:      userID,
				PlantID:     2,
				Date:        time.Now(),
				Title:       "Note 2",
				Description: "Description 2",
			},
		}

		repo.On("GetNotesWithPagination", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.NoteModel")).Run(func(args mock.Arguments) {
			plants := args.Get(3).(*[]*domain.NoteModel)
			*plants = expectedNotes
		}).Return(nil).Once()

		result, err := service.GetNotesWithPagination(userID, offset, pageSize)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, len(expectedNotes))
		for i, note := range expectedNotes {
			assert.Equal(t, note.ID, result[i].ID)
			assert.Equal(t, note.UserID, result[i].UserID)
			assert.Equal(t, note.PlantID, result[i].PlantID)
			assert.Equal(t, note.Date, result[i].Date)
			assert.Equal(t, note.Title, result[i].Title)
			assert.Equal(t, note.Description, result[i].Description)
		}
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)
		offset := 10
		pageSize := 5
		expectedErr := errors.New("GetNotesWithPagination")
		repo.On("GetNotesWithPagination", userID, offset, pageSize, mock.AnythingOfType("*[]*domain.NoteModel")).Return(expectedErr).Once()

		result, err := service.GetNotesWithPagination(userID, offset, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)
		repo.AssertExpectations(t)
	})

}

func TestNoteService_InsertNotes(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		note := &domain.NoteModel{
			UserID:      1,
			PlantID:     1,
			Date:        time.Now(),
			Title:       "Note 1",
			Description: "Description 1",
		}

		expectedNote := &domain.NoteModel{
			ID:          1,
			UserID:      note.UserID,
			PlantID:     note.PlantID,
			Date:        note.Date,
			Title:       note.Title,
			Description: note.Description,
		}

		repo.On("InsertNote", note).Return(expectedNote, nil).Once()

		result, err := service.InsertNote(note)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedNote.ID, result.ID)
		assert.Equal(t, expectedNote.UserID, result.UserID)
		assert.Equal(t, expectedNote.PlantID, result.PlantID)
		assert.Equal(t, expectedNote.Date, result.Date)
		assert.Equal(t, expectedNote.Title, result.Title)
		assert.Equal(t, expectedNote.Description, result.Description)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		note := &domain.NoteModel{
			UserID:      1,
			PlantID:     2,
			Date:        time.Now(),
			Title:       "Note 2",
			Description: "Description 2",
		}

		expectedErr := errors.New("InsertNote")
		repo.On("InsertNote", note).Return(nil, expectedErr).Once()

		result, err := service.InsertNote(note)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_UpdateNotes(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		note := &domain.NoteModel{
			ID:          1,
			UserID:      1,
			PlantID:     1,
			Date:        time.Now(),
			Title:       "Note 1",
			Description: "Description 1",
		}

		updatedNote := &domain.NoteModel{
			ID:          note.ID,
			UserID:      note.UserID,
			PlantID:     note.PlantID,
			Date:        note.Date,
			Title:       "Updated Note 1",
			Description: "Updated Description 1",
		}

		repo.On("UpdateNotes", updatedNote).Return(updatedNote, nil).Once()

		result, err := service.UpdateNotes(updatedNote)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, updatedNote.ID, result.ID)
		assert.Equal(t, updatedNote.UserID, result.UserID)
		assert.Equal(t, updatedNote.PlantID, result.PlantID)
		assert.Equal(t, updatedNote.Date, result.Date)
		assert.Equal(t, updatedNote.Title, result.Title)
		assert.Equal(t, updatedNote.Description, result.Description)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		note := &domain.NoteModel{
			ID:          2,
			UserID:      1,
			PlantID:     2,
			Date:        time.Now(),
			Title:       "Note 2",
			Description: "Description 2",
		}

		expectedErr := errors.New("UpdateNotes")
		repo.On("UpdateNotes", note).Return(nil, expectedErr).Once()

		result, err := service.UpdateNotes(note)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_DeleteNotes(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		noteID := uint64(1)

		repo.On("DeleteNotes", noteID).Return(nil).Once()

		err := service.DeleteNotes(noteID)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		noteID := uint64(2)

		expectedErr := errors.New("DeleteNotes")
		repo.On("DeleteNotes", noteID).Return(expectedErr).Once()

		err := service.DeleteNotes(noteID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_GetNotesByID(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		noteID := uint64(1)

		expectedNote := &domain.NoteModel{
			ID:          1,
			UserID:      1,
			PlantID:     1,
			Date:        time.Now(),
			Title:       "Note 1",
			Description: "Description 1",
		}

		repo.On("GetNoteByID", noteID).Return(expectedNote, nil).Once()

		result, err := service.GetNoteByID(noteID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedNote.ID, result.ID)
		assert.Equal(t, expectedNote.UserID, result.UserID)
		assert.Equal(t, expectedNote.PlantID, result.PlantID)
		assert.Equal(t, expectedNote.Date, result.Date)
		assert.Equal(t, expectedNote.Title, result.Title)
		assert.Equal(t, expectedNote.Description, result.Description)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		noteID := uint64(2)

		expectedErr := errors.New("GetNotesByID")
		repo.On("GetNoteByID", noteID).Return(nil, expectedErr).Once()

		result, err := service.GetNoteByID(noteID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_CountNotes(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		userID := uint64(1)

		expectedCount := uint64(5)

		repo.On("CountNotes", userID).Return(expectedCount, nil).Once()

		result, err := service.CountNotes(userID)

		assert.Nil(t, err)
		assert.Equal(t, expectedCount, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		userID := uint64(2)

		expectedErr := errors.New("CountNotes")
		repo.On("CountNotes", userID).Return(uint64(0), expectedErr).Once()

		result, err := service.CountNotes(userID)

		assert.Error(t, err)
		assert.Equal(t, uint64(0), result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_InsertNotePhoto(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		photo := &domain.PhotoModel{
			NoteID:      1,
			Photo:       "photo.jpg",
			Description: "Photo 1",
		}

		expectedPhoto := &domain.PhotoModel{
			PhotoID:     1,
			NoteID:      photo.NoteID,
			Photo:       photo.Photo,
			Description: photo.Description,
		}

		repo.On("InsertNotePhoto", photo).Return(expectedPhoto, nil).Once()

		result, err := service.InsertNotePhoto(photo)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedPhoto.PhotoID, result.PhotoID)
		assert.Equal(t, expectedPhoto.NoteID, result.NoteID)
		assert.Equal(t, expectedPhoto.Photo, result.Photo)
		assert.Equal(t, expectedPhoto.Description, result.Description)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		photo := &domain.PhotoModel{
			NoteID:      2,
			Photo:       "photo2.jpg",
			Description: "Photo 2",
		}

		expectedErr := errors.New("InsertNotePhoto")
		repo.On("InsertNotePhoto", photo).Return(nil, expectedErr).Once()

		result, err := service.InsertNotePhoto(photo)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestNoteService_GetNotePhotoByID(t *testing.T) {
	repo := mocks.NewRepoNoteInterface(t)
	service := NewNoteService(repo)

	t.Run("Success Case", func(t *testing.T) {
		photoID := uint64(1)

		expectedPhoto := &domain.PhotoModel{
			PhotoID:     photoID,
			NoteID:      1,
			Photo:       "sample.jpg",
			Description: "Sample photo",
		}

		repo.On("GetNotePhotoByID", photoID).Return(expectedPhoto, nil).Once()

		result, err := service.GetNotePhotoByID(photoID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedPhoto.PhotoID, result.PhotoID)
		assert.Equal(t, expectedPhoto.NoteID, result.NoteID)
		assert.Equal(t, expectedPhoto.Photo, result.Photo)
		assert.Equal(t, expectedPhoto.Description, result.Description)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case", func(t *testing.T) {
		photoID := uint64(2)

		expectedErr := errors.New("GetNotePhotoByID")
		repo.On("GetNotePhotoByID", photoID).Return(nil, expectedErr).Once()

		result, err := service.GetNotePhotoByID(photoID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

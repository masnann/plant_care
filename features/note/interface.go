package note

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/note/domain"
)

type RepoNoteInterface interface {
	InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error)
	GetNoteByID(noteID uint64) (*domain.NoteModel, error)
	GetNotesWithPagination(userID uint64, offset, pageSize int, notes *[]*domain.NoteModel) error
	CountNotes(userID uint64) (uint64, error)
	UpdateNotes(note *domain.NoteModel) (*domain.NoteModel, error)
	DeleteNotes(noteID uint64) error
	InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error)
	UpdateNotesPhoto(photo *domain.PhotoModel) (*domain.PhotoModel, error)
	GetNotePhotoByID(noteID uint64) (*domain.PhotoModel, error)
	DeleteNotesPhoto(photoID uint64) error
}

type ServiceNoteInterface interface {
	InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error)
	GetNoteByID(noteID uint64) (*domain.NoteModel, error)
	GetNotesWithPagination(userID uint64, offset, pageSize int) ([]*domain.NoteModel, error)
	CountNotes(userID uint64) (uint64, error)
	UpdateNotes(notes *domain.NoteModel) (*domain.NoteModel, error)
	DeleteNotes(noteID uint64) error
	InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error)
	UpdateNotesPhotos(photos *domain.PhotoModel) (*domain.PhotoModel, error)
	GetNotePhotoByID(noteID uint64) (*domain.PhotoModel, error)
	DeleteNotesPhotos(photoID uint64) error
}

type HandlerNoteInterface interface {
	InsertNotes() echo.HandlerFunc
	GetNotesWithPagination() echo.HandlerFunc
	UpdateNotes() echo.HandlerFunc
	DeleteNotes() echo.HandlerFunc
	InsertNotePhoto() echo.HandlerFunc
	UpdateNotesPhotos() echo.HandlerFunc
	DeleteNotesPhotos() echo.HandlerFunc
}

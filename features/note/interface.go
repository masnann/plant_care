package note

import (
	"github.com/labstack/echo/v4"
	"github.com/masnann/plant_care/features/note/domain"
)

type RepoNoteInterface interface {
	InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error)
	InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error)
	GetNoteByID(noteID uint64) (*domain.NoteModel, error)
}

type ServiceNoteInterface interface {
	InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error)
	InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error)
	GetNoteByID(noteID uint64) (*domain.NoteModel, error)
}

type HandlerNoteInterface interface {
	InsertNotes() echo.HandlerFunc
	InsertNotePhoto() echo.HandlerFunc
}

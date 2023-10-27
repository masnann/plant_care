package service

import (
	"github.com/masnann/plant_care/features/note"
	"github.com/masnann/plant_care/features/note/domain"
)

type NoteService struct {
	repo note.RepoNoteInterface
}

func NewNoteService(repo note.RepoNoteInterface) note.ServiceNoteInterface {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) GetNoteByID(noteID uint64) (*domain.NoteModel, error) {
	result, err := s.repo.GetNoteByID(noteID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *NoteService) InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error) {
	result, err := s.repo.InsertNote(noteModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NoteService) InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error) {
	result, err := s.repo.InsertNotePhoto(photoModel)
	if err != nil {
		return nil, err
	}
	return result, nil
}

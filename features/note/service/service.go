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

func (s *NoteService) UpdateNotes(notes *domain.NoteModel) (*domain.NoteModel, error) {
	result, err := s.repo.UpdateNotes(notes)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NoteService) GetNotesWithPagination(userID uint64, offset, pageSize int) ([]*domain.NoteModel, error) {
	var notes []*domain.NoteModel
	if err := s.repo.GetNotesWithPagination(userID, offset, pageSize, &notes); err != nil {
		return nil, err
	}
	return notes, nil
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

func (s *NoteService) CountNotes(userID uint64) (uint64, error) {
	count, err := s.repo.CountNotes(userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *NoteService) DeleteNotes(noteID uint64) error {
	err := s.repo.DeleteNotes(noteID)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) UpdateNotesPhotos(photos *domain.PhotoModel) (*domain.PhotoModel, error) {
	result, err := s.repo.UpdateNotesPhoto(photos)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NoteService) GetNotePhotoByID(noteID uint64) (*domain.PhotoModel, error) {
	result, err := s.repo.GetNotePhotoByID(noteID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *NoteService) DeleteNotesPhotos(photoID uint64) error {
	err := s.repo.DeleteNotesPhoto(photoID)
	if err != nil {
		return err
	}
	return nil
}

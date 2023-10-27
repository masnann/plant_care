package repository

import (
	"github.com/masnann/plant_care/features/note"
	"github.com/masnann/plant_care/features/note/domain"
	"github.com/masnann/plant_care/utils"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) note.RepoNoteInterface {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) InsertNote(noteModel *domain.NoteModel) (*domain.NoteModel, error) {
	newData := &domain.NoteModel{
		ID:          noteModel.ID,
		UserID:      noteModel.UserID,
		PlantID:     noteModel.PlantID,
		Title:       noteModel.Title,
		Description: noteModel.Description,
		Date:        utils.GetNowTime(),
		CreatedAt:   utils.GetNowTime(),
		UpdatedAt:   utils.GetNowTime(),
		DeletedAt:   nil,
	}
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *NoteRepository) InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error) {
	newData := &domain.PhotoModel{
		NoteID:      photoModel.NoteID,
		URL:         photoModel.URL,
		Description: photoModel.Description,
	}

	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *NoteRepository) GetNoteByID(noteID uint64) (*domain.NoteModel, error) {
	var notes domain.NoteModel
	if err := r.db.Where("id = ?", noteID).First(&notes).Error; err != nil {
		return nil, err
	}
	return &notes, nil
}

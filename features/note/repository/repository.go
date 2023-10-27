package repository

import (
	"github.com/masnann/plant_care/features/note"
	"github.com/masnann/plant_care/features/note/domain"
	"github.com/masnann/plant_care/utils"
	"gorm.io/gorm"
	"time"
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

func (r *NoteRepository) GetNoteByID(noteID uint64) (*domain.NoteModel, error) {
	var notes domain.NoteModel
	if err := r.db.Where("id = ?", noteID).First(&notes).Error; err != nil {
		return nil, err
	}
	return &notes, nil
}

func (r *NoteRepository) GetNotesWithPagination(userID uint64, offset, pageSize int, notes *[]*domain.NoteModel) error {
	if err := r.db.Preload("Photos", "deleted_at IS NULL").Where("user_id = ? AND deleted_at IS NULL", userID).Offset(offset).Limit(pageSize).Find(&notes).Error; err != nil {
		return err
	}
	return nil
}

func (r *NoteRepository) CountNotes(userID uint64) (uint64, error) {
	var count int64
	if err := r.db.Model(&domain.NoteModel{}).Where("user_id = ? AND deleted_at IS NULL", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint64(count), nil
}

func (r *NoteRepository) UpdateNotes(note *domain.NoteModel) (*domain.NoteModel, error) {
	exitingNotes := domain.NoteModel{}
	newData := &domain.NoteModel{
		UserID:      note.UserID,
		PlantID:     note.PlantID,
		Title:       note.Title,
		Description: note.Description,
		UpdatedAt:   utils.GetNowTime(),
	}
	if err := r.db.Model(&exitingNotes).Where("id = ?", note.ID).Updates(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *NoteRepository) DeleteNotes(noteID uint64) error {
	result := r.db.Model(&domain.NoteModel{}).Where("id = ?", noteID).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *NoteRepository) InsertNotePhoto(photoModel *domain.PhotoModel) (*domain.PhotoModel, error) {
	newData := &domain.PhotoModel{
		NoteID:      photoModel.NoteID,
		Photo:       photoModel.Photo,
		Description: photoModel.Description,
		CreatedAt:   utils.GetNowTime(),
		UpdatedAt:   utils.GetNowTime(),
		DeletedAt:   nil,
	}

	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}

func (r *NoteRepository) UpdateNotesPhoto(photo *domain.PhotoModel) (*domain.PhotoModel, error) {
	existingPhoto := domain.PhotoModel{}
	newData := &domain.PhotoModel{
		Photo:       photo.Photo,
		Description: photo.Description,
		UpdatedAt:   utils.GetNowTime(),
	}
	if err := r.db.Model(&existingPhoto).Where("photo_id = ?", photo.PhotoID).Updates(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *NoteRepository) GetNotePhotoByID(noteID uint64) (*domain.PhotoModel, error) {
	var photos domain.PhotoModel
	if err := r.db.Where("photo_id = ?", noteID).First(&photos).Error; err != nil {
		return nil, err
	}
	return &photos, nil
}

func (r *NoteRepository) DeleteNotesPhoto(photoID uint64) error {
	result := r.db.Model(&domain.PhotoModel{}).Where("photo_id = ?", photoID).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

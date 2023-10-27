package domain

type InsertNoteRequest struct {
	UserId      uint64 `json:"user_id"`
	PlantID     uint64 `form:"plant_id" json:"plant_id" validate:"required"`
	Title       string `form:"title" json:"title" validate:"required"`
	Description string `form:"description" json:"description"`
}

type InsertNotePhotoRequest struct {
	UserId      uint64 `json:"user_id"`
	NoteId      uint64 `form:"note_id" json:"note_id"`
	Description string `form:"description" json:"description"`
	URL         string `json:"url" form:"url"`
}

type NoteHealthRecord struct {
	Description string `json:"description" form:"description"`
}

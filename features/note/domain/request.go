package domain

type InsertNoteRequest struct {
	UserId      uint64 `json:"user_id"`
	PlantID     uint64 `form:"plant_id" json:"plant_id" validate:"required"`
	Title       string `form:"title" json:"title" validate:"required"`
	Description string `form:"description" json:"description"`
}

type UpdateNoteRequest struct {
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}

type InsertNotePhotoRequest struct {
	UserId      uint64 `json:"user_id"`
	NoteId      uint64 `form:"note_id" json:"note_id"`
	Description string `form:"description" json:"description" validate:"required"`
	Photo       string `json:"photo" form:"photo"`
}

type UpdateNotePhotoRequest struct {
	UserId      uint64 `json:"user_id"`
	NoteId      uint64 `form:"note_id" json:"note_id"`
	Description string `form:"description" json:"description" validate:"required"`
	Photo       string `json:"photo" form:"photo"`
}

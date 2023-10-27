package domain

import "time"

type GetResponse struct {
	ID          uint64       `json:"id"`
	UserID      uint64       `json:"user_id"`
	PlantID     uint64       `json:"plant_id"`
	Date        time.Time    `json:"date"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Photos      []PhotoModel `json:"photos"`
}

type GetNotesResponse struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	PlantID     uint64    `json:"plant_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type NotesPhotoResponse struct {
	PhotoID     uint64 `json:"photo_id"`
	NoteID      uint64 `json:"note_id"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

package domain

import "time"

type GetGuideResponse struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Photo       string    `json:"photo"`
}

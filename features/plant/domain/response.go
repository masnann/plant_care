package domain

import "time"

type InsertResponse struct {
	Name  string    `form:"name" json:"name"`
	Type  string    `form:"type" json:"type"`
	Date  time.Time `form:"date" json:"date"`
	Photo string    `form:"date" json:"photo"`
}

type UpdateResponse struct {
	Name  string    `form:"name" json:"name"`
	Type  string    `form:"type" json:"type" `
	Date  time.Time `form:"date" json:"date"`
	Photo string    `form:"date" json:"photo"`
}

type GetResponse struct {
	ID    uint64    `json:"id"`
	Name  string    `json:"name"`
	Type  string    `json:"type" `
	Date  time.Time `json:"date"`
	Photo string    `json:"photo"`
}

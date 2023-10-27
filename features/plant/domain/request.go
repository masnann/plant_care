package domain

type InsertRequest struct {
	UserID uint64 `json:"user_id"`
	Name   string `form:"name" json:"name" validate:"required"`
	Type   string `form:"type" json:"type" validate:"required"`
	Photo  string `form:"photo" json:"photo"`
}

type UpdateRequest struct {
	UserID uint64 `json:"user_id"`
	Name   string `form:"name" json:"name"`
	Type   string `form:"type" json:"type"`
	Photo  string `form:"photo" json:"photo"`
}

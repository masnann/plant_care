package domain

type MessageRequest struct {
	Message string `form:"message" json:"message" validate:"required"`
}

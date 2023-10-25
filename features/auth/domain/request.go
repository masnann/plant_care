package domain

type RegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required"`
	Email    string `form:"email"json:"email" validate:"required"`
	Password string `form:"password"json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

package domain

type RegisterRequest struct {
	Username string `form:"username" json:"username" validate:"required,min=4,max=20"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

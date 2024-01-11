package domain

type UpdatePasswordRequest struct {
	UserID          uint64 `json:"user_id"`
	OldPassword     string `json:"old_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

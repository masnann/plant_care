package domain

type GetNotificationResponse struct {
	ID      uint64 `json:"id"`
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}

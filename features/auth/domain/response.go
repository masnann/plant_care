package domain

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	ID           uint64 `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

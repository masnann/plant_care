package domain

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

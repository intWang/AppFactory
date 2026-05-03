package domain

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type Provider struct {
	Key       string `json:"key"`
	Enabled   bool   `json:"enabled"`
	Available bool   `json:"available"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type RegisterResponse struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

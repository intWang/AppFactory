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

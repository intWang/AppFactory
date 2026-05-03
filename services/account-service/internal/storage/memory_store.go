package storage

import "appfactory/account-service/internal/domain"

type MemoryStore struct {
	CurrentUser domain.User
	Providers   []domain.Provider
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		CurrentUser: domain.User{
			ID:       "user-local-1",
			Email:    "local@example.com",
			Nickname: "Local User",
		},
		Providers: []domain.Provider{
			{Key: "google", Enabled: false, Available: true},
			{Key: "apple", Enabled: false, Available: true},
			{Key: "wechat", Enabled: false, Available: true},
		},
	}
}

package storage

import (
	"context"
	"time"

	"appfactory/account-service/internal/domain"
)

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

func (s *MemoryStore) Register(_ context.Context, req domain.RegisterRequest) (domain.User, error) {
	user := domain.User{
		ID:       "user-local-created",
		Email:    req.Email,
		Nickname: req.Nickname,
	}
	s.CurrentUser = user
	return user, nil
}

func (s *MemoryStore) GetCurrentUser(_ context.Context) (domain.User, error) {
	if s.CurrentUser.ID == "" {
		s.CurrentUser = domain.User{
			ID:       "user-local-seeded",
			Email:    "local@example.com",
			Nickname: "Local User",
		}
	}
	if s.CurrentUser.ID == "" {
		s.CurrentUser.ID = time.Now().Format(time.RFC3339Nano)
	}
	return s.CurrentUser, nil
}

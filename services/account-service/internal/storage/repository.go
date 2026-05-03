package storage

import (
	"context"

	"appfactory/account-service/internal/domain"
)

type Repository interface {
	Register(context.Context, domain.RegisterRequest) (domain.User, error)
	GetCurrentUser(context.Context) (domain.User, error)
}

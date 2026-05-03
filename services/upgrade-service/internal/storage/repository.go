package storage

import (
	"context"

	"appfactory/upgrade-service/internal/domain"
)

type Repository interface {
	GetTarget(context.Context, string) (domain.VersionTarget, error)
}

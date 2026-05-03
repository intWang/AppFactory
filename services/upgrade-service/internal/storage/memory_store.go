package storage

import (
	"context"
	"fmt"

	"appfactory/upgrade-service/internal/domain"
)

type MemoryStore struct {
	ClientTarget  domain.VersionTarget
	ServiceTarget domain.VersionTarget
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		ClientTarget: domain.VersionTarget{
			ProductSlug:       "shared-client",
			TargetType:        "client",
			CurrentVersion:    "26.2.20.01",
			CurrentBuild:      1,
			LatestVersion:     "26.2.20.04",
			LatestBuild:       4,
			ForceUpgradeAfter: 3,
		},
		ServiceTarget: domain.VersionTarget{
			ProductSlug:       "shared-service",
			TargetType:        "service",
			CurrentVersion:    "26.2.20.01",
			CurrentBuild:      1,
			LatestVersion:     "26.2.20.03",
			LatestBuild:       3,
			ForceUpgradeAfter: 3,
		},
	}
}

func (s *MemoryStore) GetTarget(_ context.Context, targetType string) (domain.VersionTarget, error) {
	switch targetType {
	case "client":
		return s.ClientTarget, nil
	case "service":
		return s.ServiceTarget, nil
	default:
		return domain.VersionTarget{}, fmt.Errorf("unknown target type: %s", targetType)
	}
}

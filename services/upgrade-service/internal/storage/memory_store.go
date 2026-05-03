package storage

import "appfactory/upgrade-service/internal/domain"

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

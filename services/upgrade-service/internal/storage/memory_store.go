package storage

import (
	"context"
	"fmt"

	"appfactory/upgrade-service/internal/domain"
)

type MemoryStore struct {
	ClientTarget  domain.VersionTarget
	ServiceTarget domain.VersionTarget
	Releases      map[string]domain.ReleaseVersion
	Deployments   map[string]domain.DeploymentRecord
	ActiveIDs     map[string]string
}

func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{
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
		Releases:    map[string]domain.ReleaseVersion{},
		Deployments: map[string]domain.DeploymentRecord{},
		ActiveIDs:   map[string]string{},
	}

	clientRelease := domain.ReleaseVersion{
		ID:           "release-client-4",
		ProductSlug:  store.ClientTarget.ProductSlug,
		TargetType:   "client",
		VersionLabel: store.ClientTarget.LatestVersion,
		BuildNumber:  store.ClientTarget.LatestBuild,
		UpgradeURL:   "https://example.com/client/26.2.20.04",
	}
	serviceRelease := domain.ReleaseVersion{
		ID:           "release-service-3",
		ProductSlug:  store.ServiceTarget.ProductSlug,
		TargetType:   "service",
		VersionLabel: store.ServiceTarget.LatestVersion,
		BuildNumber:  store.ServiceTarget.LatestBuild,
		UpgradeURL:   "https://example.com/service/26.2.20.03",
	}

	store.Releases[clientRelease.ID] = clientRelease
	store.Releases[serviceRelease.ID] = serviceRelease
	store.ActiveIDs["client"] = clientRelease.ID
	store.ActiveIDs["service"] = serviceRelease.ID

	return store
}

func (s *MemoryStore) GetActiveTargets(_ context.Context) (domain.TargetBundle, error) {
	return domain.TargetBundle{
		Client:  s.ClientTarget,
		Service: s.ServiceTarget,
	}, nil
}

func (s *MemoryStore) CreateRelease(_ context.Context, req domain.CreateReleaseRequest) (domain.ReleaseVersion, error) {
	release := domain.ReleaseVersion{
		ID:           fmt.Sprintf("release-%s-%d", req.TargetType, req.BuildNumber),
		ProductSlug:  req.ProductSlug,
		TargetType:   req.TargetType,
		VersionLabel: req.VersionLabel,
		BuildNumber:  req.BuildNumber,
		UpgradeURL:   req.UpgradeURL,
	}
	s.Releases[release.ID] = release
	return release, nil
}

func (s *MemoryStore) CreateDeployment(_ context.Context, req domain.CreateDeploymentRequest) (domain.DeploymentRecord, error) {
	if _, ok := s.Releases[req.TargetVersionID]; !ok {
		return domain.DeploymentRecord{}, fmt.Errorf("unknown release version: %s", req.TargetVersionID)
	}
	deployment := domain.DeploymentRecord{
		ID:              fmt.Sprintf("deployment-%d", len(s.Deployments)+1),
		TargetVersionID: req.TargetVersionID,
		Environment:     req.Environment,
		Status:          req.Status,
	}
	s.Deployments[deployment.ID] = deployment
	return deployment, nil
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

func (s *MemoryStore) SwitchTarget(_ context.Context, req domain.SwitchTargetRequest) (domain.VersionTarget, error) {
	release, err := s.releaseFor(req.TargetType, req.ToVersionID)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	s.ActiveIDs[req.TargetType] = release.ID
	target := s.applyRelease(req.TargetType, release)
	return target, nil
}

func (s *MemoryStore) RollbackTarget(_ context.Context, req domain.RollbackTargetRequest) (domain.VersionTarget, error) {
	release, err := s.releaseFor(req.TargetType, req.RolledBackToVersionID)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	s.ActiveIDs[req.TargetType] = release.ID
	target := s.applyRelease(req.TargetType, release)
	return target, nil
}

func (s *MemoryStore) releaseFor(targetType, releaseID string) (domain.ReleaseVersion, error) {
	release, ok := s.Releases[releaseID]
	if !ok {
		return domain.ReleaseVersion{}, fmt.Errorf("unknown release version: %s", releaseID)
	}
	if release.TargetType != targetType {
		return domain.ReleaseVersion{}, fmt.Errorf("release %s does not match target type %s", releaseID, targetType)
	}
	return release, nil
}

func (s *MemoryStore) applyRelease(targetType string, release domain.ReleaseVersion) domain.VersionTarget {
	switch targetType {
	case "client":
		s.ClientTarget.ProductSlug = release.ProductSlug
		s.ClientTarget.LatestVersion = release.VersionLabel
		s.ClientTarget.LatestBuild = release.BuildNumber
		s.ClientTarget.CurrentVersion = release.VersionLabel
		s.ClientTarget.CurrentBuild = release.BuildNumber
		return s.ClientTarget
	case "service":
		s.ServiceTarget.ProductSlug = release.ProductSlug
		s.ServiceTarget.LatestVersion = release.VersionLabel
		s.ServiceTarget.LatestBuild = release.BuildNumber
		s.ServiceTarget.CurrentVersion = release.VersionLabel
		s.ServiceTarget.CurrentBuild = release.BuildNumber
		return s.ServiceTarget
	default:
		return domain.VersionTarget{}
	}
}

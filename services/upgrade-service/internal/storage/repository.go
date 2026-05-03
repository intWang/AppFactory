package storage

import (
	"context"

	"appfactory/upgrade-service/internal/domain"
)

type Repository interface {
	GetTarget(context.Context, string) (domain.VersionTarget, error)
	GetActiveTargets(context.Context) (domain.TargetBundle, error)
	CreateRelease(context.Context, domain.CreateReleaseRequest) (domain.ReleaseVersion, error)
	CreateDeployment(context.Context, domain.CreateDeploymentRequest) (domain.DeploymentRecord, error)
	SwitchTarget(context.Context, domain.SwitchTargetRequest) (domain.VersionTarget, error)
	RollbackTarget(context.Context, domain.RollbackTargetRequest) (domain.VersionTarget, error)
}

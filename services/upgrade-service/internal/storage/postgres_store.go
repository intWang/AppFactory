package storage

import (
	"context"
	"fmt"
	"time"

	"appfactory/upgrade-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool             *pgxpool.Pool
	forcedUpgradeGap int
}

func NewPostgresStore(ctx context.Context, dsn string, forcedUpgradeGap int) (*PostgresStore, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{pool: pool, forcedUpgradeGap: forcedUpgradeGap}, nil
}

func (s *PostgresStore) GetTarget(ctx context.Context, targetType string) (domain.VersionTarget, error) {
	var target domain.VersionTarget
	err := s.pool.QueryRow(ctx, `
SELECT
  at.product_slug,
  at.target_type,
  rv.version_label,
  rv.build_number
FROM active_targets at
JOIN release_versions rv ON rv.id = at.active_version_id
WHERE at.target_type = $1
LIMIT 1
`, targetType).Scan(
		&target.ProductSlug,
		&target.TargetType,
		&target.LatestVersion,
		&target.LatestBuild,
	)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	target.ForceUpgradeAfter = s.forcedUpgradeGap
	return target, nil
}

func (s *PostgresStore) GetActiveTargets(ctx context.Context) (domain.TargetBundle, error) {
	client, err := s.GetTarget(ctx, "client")
	if err != nil {
		return domain.TargetBundle{}, err
	}
	service, err := s.GetTarget(ctx, "service")
	if err != nil {
		return domain.TargetBundle{}, err
	}
	return domain.TargetBundle{Client: client, Service: service}, nil
}

func (s *PostgresStore) ListReleases(ctx context.Context) ([]domain.ReleaseVersion, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, product_slug, target_type, version_label, build_number, COALESCE(upgrade_url, '')
FROM release_versions
ORDER BY created_at DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var releases []domain.ReleaseVersion
	for rows.Next() {
		var release domain.ReleaseVersion
		if err := rows.Scan(&release.ID, &release.ProductSlug, &release.TargetType, &release.VersionLabel, &release.BuildNumber, &release.UpgradeURL); err != nil {
			return nil, err
		}
		releases = append(releases, release)
	}
	return releases, rows.Err()
}

func (s *PostgresStore) ListDeployments(ctx context.Context) ([]domain.DeploymentRecord, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, target_version_id, environment, status
FROM deployment_records
ORDER BY deployed_at DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deployments []domain.DeploymentRecord
	for rows.Next() {
		var deployment domain.DeploymentRecord
		if err := rows.Scan(&deployment.ID, &deployment.TargetVersionID, &deployment.Environment, &deployment.Status); err != nil {
			return nil, err
		}
		deployments = append(deployments, deployment)
	}
	return deployments, rows.Err()
}

func (s *PostgresStore) ListSwitchEvents(ctx context.Context) ([]domain.SwitchEvent, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, product_slug, target_type, COALESCE(from_version_id, ''), to_version_id, operator
FROM switch_events
ORDER BY switched_at DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.SwitchEvent
	for rows.Next() {
		var event domain.SwitchEvent
		if err := rows.Scan(&event.ID, &event.ProductSlug, &event.TargetType, &event.FromVersionID, &event.ToVersionID, &event.Operator); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (s *PostgresStore) ListRollbackEvents(ctx context.Context) ([]domain.RollbackEvent, error) {
	rows, err := s.pool.Query(ctx, `
SELECT id, product_slug, target_type, rolled_back_from_version_id, rolled_back_to_version_id, operator
FROM rollback_events
ORDER BY rolled_back_at DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.RollbackEvent
	for rows.Next() {
		var event domain.RollbackEvent
		if err := rows.Scan(&event.ID, &event.ProductSlug, &event.TargetType, &event.RolledBackFromVersion, &event.RolledBackToVersion, &event.Operator); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (s *PostgresStore) CreateRelease(ctx context.Context, req domain.CreateReleaseRequest) (domain.ReleaseVersion, error) {
	release := domain.ReleaseVersion{
		ID:           fmt.Sprintf("release-%d", time.Now().UnixNano()),
		ProductSlug:  req.ProductSlug,
		TargetType:   req.TargetType,
		VersionLabel: req.VersionLabel,
		BuildNumber:  req.BuildNumber,
		UpgradeURL:   req.UpgradeURL,
	}
	_, err := s.pool.Exec(ctx, `
INSERT INTO release_versions (id, product_slug, target_type, version_label, build_number, upgrade_url)
VALUES ($1, $2, $3, $4, $5, $6)
`, release.ID, release.ProductSlug, release.TargetType, release.VersionLabel, release.BuildNumber, release.UpgradeURL)
	if err != nil {
		return domain.ReleaseVersion{}, err
	}
	return release, nil
}

func (s *PostgresStore) CreateDeployment(ctx context.Context, req domain.CreateDeploymentRequest) (domain.DeploymentRecord, error) {
	deployment := domain.DeploymentRecord{
		ID:              fmt.Sprintf("deployment-%d", time.Now().UnixNano()),
		TargetVersionID: req.TargetVersionID,
		Environment:     req.Environment,
		Status:          req.Status,
	}
	_, err := s.pool.Exec(ctx, `
INSERT INTO deployment_records (id, target_version_id, environment, status)
VALUES ($1, $2, $3, $4)
`, deployment.ID, deployment.TargetVersionID, deployment.Environment, deployment.Status)
	if err != nil {
		return domain.DeploymentRecord{}, err
	}
	return deployment, nil
}

func (s *PostgresStore) SwitchTarget(ctx context.Context, req domain.SwitchTargetRequest) (domain.VersionTarget, error) {
	release, err := s.releaseByID(ctx, req.ToVersionID)
	if err != nil {
		return domain.VersionTarget{}, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	defer tx.Rollback(ctx)

	var fromVersionID *string
	_ = tx.QueryRow(ctx, `
SELECT active_version_id
FROM active_targets
WHERE product_slug = $1 AND target_type = $2
`, req.ProductSlug, req.TargetType).Scan(&fromVersionID)

	_, err = tx.Exec(ctx, `
INSERT INTO active_targets (id, product_slug, target_type, active_version_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT (product_slug, target_type)
DO UPDATE SET active_version_id = EXCLUDED.active_version_id, updated_at = NOW()
`, fmt.Sprintf("active-%s-%s", req.ProductSlug, req.TargetType), req.ProductSlug, req.TargetType, req.ToVersionID)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	_, err = tx.Exec(ctx, `
INSERT INTO switch_events (id, product_slug, target_type, from_version_id, to_version_id, operator)
VALUES ($1, $2, $3, $4, $5, $6)
`, fmt.Sprintf("switch-%d", time.Now().UnixNano()), req.ProductSlug, req.TargetType, fromVersionID, req.ToVersionID, req.Operator)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.VersionTarget{}, err
	}
	return s.versionTargetFromRelease(release), nil
}

func (s *PostgresStore) RollbackTarget(ctx context.Context, req domain.RollbackTargetRequest) (domain.VersionTarget, error) {
	release, err := s.releaseByID(ctx, req.RolledBackToVersionID)
	if err != nil {
		return domain.VersionTarget{}, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	defer tx.Rollback(ctx)

	var currentVersionID string
	if err := tx.QueryRow(ctx, `
SELECT active_version_id
FROM active_targets
WHERE product_slug = $1 AND target_type = $2
`, req.ProductSlug, req.TargetType).Scan(&currentVersionID); err != nil {
		return domain.VersionTarget{}, err
	}
	_, err = tx.Exec(ctx, `
UPDATE active_targets
SET active_version_id = $1, updated_at = NOW()
WHERE product_slug = $2 AND target_type = $3
`, req.RolledBackToVersionID, req.ProductSlug, req.TargetType)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	_, err = tx.Exec(ctx, `
INSERT INTO rollback_events (id, product_slug, target_type, rolled_back_from_version_id, rolled_back_to_version_id, operator)
VALUES ($1, $2, $3, $4, $5, $6)
`, fmt.Sprintf("rollback-%d", time.Now().UnixNano()), req.ProductSlug, req.TargetType, currentVersionID, req.RolledBackToVersionID, req.Operator)
	if err != nil {
		return domain.VersionTarget{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.VersionTarget{}, err
	}
	return s.versionTargetFromRelease(release), nil
}

func (s *PostgresStore) releaseByID(ctx context.Context, id string) (domain.ReleaseVersion, error) {
	var release domain.ReleaseVersion
	err := s.pool.QueryRow(ctx, `
SELECT id, product_slug, target_type, version_label, build_number, COALESCE(upgrade_url, '')
FROM release_versions
WHERE id = $1
`, id).Scan(
		&release.ID,
		&release.ProductSlug,
		&release.TargetType,
		&release.VersionLabel,
		&release.BuildNumber,
		&release.UpgradeURL,
	)
	if err != nil {
		return domain.ReleaseVersion{}, err
	}
	return release, nil
}

func (s *PostgresStore) versionTargetFromRelease(release domain.ReleaseVersion) domain.VersionTarget {
	return domain.VersionTarget{
		ProductSlug:       release.ProductSlug,
		TargetType:        release.TargetType,
		CurrentVersion:    release.VersionLabel,
		CurrentBuild:      release.BuildNumber,
		LatestVersion:     release.VersionLabel,
		LatestBuild:       release.BuildNumber,
		ForceUpgradeAfter: s.forcedUpgradeGap,
	}
}

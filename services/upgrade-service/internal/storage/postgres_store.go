package storage

import (
	"context"

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

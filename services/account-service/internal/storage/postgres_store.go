package storage

import (
	"context"
	"fmt"
	"time"

	"appfactory/account-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, dsn string) (*PostgresStore, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{pool: pool}, nil
}

func (s *PostgresStore) Register(ctx context.Context, req domain.RegisterRequest) (domain.User, error) {
	user := domain.User{
		ID:       fmt.Sprintf("user-%d", time.Now().UnixNano()),
		Email:    req.Email,
		Nickname: req.Nickname,
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return domain.User{}, err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`INSERT INTO users (id, email, nickname) VALUES ($1, $2, $3)`,
		user.ID, user.Email, user.Nickname,
	); err != nil {
		return domain.User{}, err
	}
	if _, err := tx.Exec(ctx,
		`INSERT INTO local_credentials (user_id, password_hash) VALUES ($1, $2)`,
		user.ID, req.Password,
	); err != nil {
		return domain.User{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *PostgresStore) GetCurrentUser(ctx context.Context) (domain.User, error) {
	var user domain.User
	err := s.pool.QueryRow(ctx,
		`SELECT id, email, nickname FROM users ORDER BY created_at DESC LIMIT 1`,
	).Scan(&user.ID, &user.Email, &user.Nickname)
	return user, err
}

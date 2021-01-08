package store

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/snocorp/goteam/api"
)

type UserStore interface {
	Create(ctx context.Context, name, email, password string) (api.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (api.User, error)
}

type dbUserStore struct {
	db *sql.DB
}

func (s *dbUserStore) GetByID(ctx context.Context, id uuid.UUID) (u api.User, err error) {
	row := s.db.QueryRowContext(ctx, "SELECT name, email, passwordHash FROM users WHERE id = ?", id)

	err = row.Scan(&u)

	return
}

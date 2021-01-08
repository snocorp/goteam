package store

import "github.com/google/uuid"

import "github.com/snocorp/goteam/api"

import "context"

import "database/sql"

type SessionStore interface {
	Create(ctx context.Context, id, userID uuid.UUID) (api.Session, error)
}

type dbSessionStore struct {
	db *sql.DB
}

func (svc *dbSessionStore) Create(ctx context.Context, id, userID uuid.UUID) (s api.Session, err error) {
	_, err = svc.db.ExecContext(ctx, "INSERT INTO sessions (id, user_id) VALUES (?, ?)", id, userID)

	s = api.Session{
		ID:     id,
		UserID: userID,
	}
	return
}

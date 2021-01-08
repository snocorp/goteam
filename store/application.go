package store

import "github.com/snocorp/goteam/api"

import "github.com/google/uuid"

import "database/sql"

import "context"

type ApplicationStore interface {
	Create(ctx context.Context, id uuid.UUID, name string) (api.Application, error)
}

type dbApplicationStore struct {
	db *sql.DB
}

func (s *dbApplicationStore) Create(ctx context.Context, id uuid.UUID, name string) (app api.Application, err error) {
	_, err = s.db.ExecContext(ctx, "INSERT INTO applications (id, name) VALUES (?, ?)", id, name)

	app = api.Application{
		ID:   id,
		Name: name,
	}
	return
}

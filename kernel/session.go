package kernel

import "github.com/google/uuid"

import "github.com/snocorp/goteam/api"

import "github.com/snocorp/goteam/store"

import "context"

type sessionService struct {
	ctx   context.Context
	store store.SessionStore
}

func (svc *sessionService) Create(userID uuid.UUID, password string) (s api.Session, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	s, err = svc.store.Create(svc.ctx, id, userID)
	return
}

package kernel

import (
	"github.com/snocorp/goteam/api"
	"github.com/snocorp/goteam/store"

	"context"

	"github.com/google/uuid"
)

type applicationService struct {
	ctx   context.Context
	store store.ApplicationStore
}

func (svc *applicationService) Create(name string) (app api.Application, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	app, err = svc.store.Create(svc.ctx, id, name)
	return
}

package kernel

import (
	"context"
	"github.com/google/uuid"
	"github.com/snocorp/goteam/api"
	"github.com/snocorp/goteam/store"
)

type userService struct {
	ctx       context.Context
	userStore store.UserStore
}

func (svc *userService) Create(name, email, password string) (u api.User, err error) {
	u, err = svc.userStore.Create(svc.ctx, name, email, password)
	return
}

func (svc *userService) get(id uuid.UUID) (u api.User, err error) {
	u, err = svc.userStore.GetByID(svc.ctx, id)
	return
}

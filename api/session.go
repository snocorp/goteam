package api

import "github.com/google/uuid"

type Session struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type SessionService interface {
	Create(userID uuid.UUID, password string) (s Session, err error)
}

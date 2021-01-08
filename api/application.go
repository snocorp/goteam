package api

import "github.com/google/uuid"

type Application struct {
	ID   uuid.UUID
	Name string
}

type ApplicationService interface {
	Create(name string) (Application, error)
}

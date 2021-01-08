package api

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}

type UserService interface {
	Create(name, email, password string) (u User, err error)
	Get(id uuid.UUID) (u User, err error)
}

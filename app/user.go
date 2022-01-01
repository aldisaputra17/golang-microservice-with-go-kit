package app

import "context"

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Password string `json:"pasword"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

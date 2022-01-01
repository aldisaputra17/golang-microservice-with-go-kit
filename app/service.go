package app

import (
	"context"
	"errors"
)

type Service interface {
	CreateUser(ctx context.Context, name string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

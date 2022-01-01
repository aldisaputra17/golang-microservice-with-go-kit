package app

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/log"
)

var RepoErr = errors.New("Unable the handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (id, name, password)
		VALUES ($1, $2, $3)`

	if user.Name == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var name string
	err := repo.db.QueryRow("SELECT name FROM users WHERE id=$1", id).Scan(&name)
	if err != nil {
		return "", RepoErr
	}

	return name, nil
}

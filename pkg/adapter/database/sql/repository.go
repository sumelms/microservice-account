package sql

import (
	"context"
	"database/sql"
	"github.com/sumelms/microservice-user/pkg/domain"

	"github.com/go-kit/kit/log"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) domain.Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "Repository", "sql"),
	}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return domain.RepoErr
	}

	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := r.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", domain.RepoErr
	}

	return email, nil
}

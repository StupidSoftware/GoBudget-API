package repository

import (
	"context"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, user *model.User) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO users (id, username, password, created_at, balance) VALUES ($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.Password, user.CreatedAt, user.Balance)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return nil
}

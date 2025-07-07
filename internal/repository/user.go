package repository

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx *gin.Context, user *model.User) error
	GetByUsername(ctx *gin.Context, username string) (*model.User, error)
	Delete(ctx *gin.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx *gin.Context, user *model.User) error {
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

func (r *userRepository) GetByUsername(ctx *gin.Context, username string) (*model.User, error) {
	var user model.User

	err := r.db.QueryRow(ctx, "SELECT id, username, password, created_at, balance FROM users WHERE username = $1", username).Scan(
		&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.Balance)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(ctx *gin.Context, id string) error {
	return nil
}

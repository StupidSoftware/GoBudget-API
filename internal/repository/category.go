package repository

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository interface {
	Create(ctx *gin.Context, category *model.Category) error
	GetAll(ctx *gin.Context) ([]*model.Category, error)
	GetByID(ctx *gin.Context, id int) (*model.Category, error)
	Delete(ctx *gin.Context, id int) error
}

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx *gin.Context, category *model.Category) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO categories (id, name, user_id) VALUES ($1, $2, $3)",
		category.ID, category.Name, category.UserID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

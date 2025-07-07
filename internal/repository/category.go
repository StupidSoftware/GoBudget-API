package repository

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository interface {
	Create(ctx *gin.Context, category *model.Category) error
	GetByUserID(ctx *gin.Context, userID string) ([]*model.Category, error)
	CategoryAlreadyExists(ctx *gin.Context, category *model.Category) (bool, error)
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

func (r *categoryRepository) GetByUserID(ctx *gin.Context, userID string) ([]*model.Category, error) {
	var categories []*model.Category

	rows, err := r.db.Query(ctx, "SELECT id, name, user_id FROM categories WHERE user_id = $1 OR user_id IS NULL", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Name, &category.UserID)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func (r *categoryRepository) CategoryAlreadyExists(ctx *gin.Context, category *model.Category) (bool, error) {
	var count int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM categories WHERE name = $1 AND user_id = $2 OR name = $1 AND user_id IS NULL", category.Name, category.UserID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

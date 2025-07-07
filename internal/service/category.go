package service

import (
	"errors"

	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type CategoryService interface {
	Create(ctx *gin.Context, category *model.Category) *utils.CustomError
	GetByUserID(ctx *gin.Context) ([]*model.Category, *utils.CustomError)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *categoryService {
	return &categoryService{
		repo: repo,
	}
}

func (c *categoryService) Create(ctx *gin.Context, category *model.Category) *utils.CustomError {
	category.ID = uuid.New()

	if err := c.repo.Create(ctx, category); err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == "23503" {
				return &utils.CustomError{
					Message: "User not found",
					Code:    404,
					Err:     errors.New("user not found"),
				}
			}
		}

		return &utils.CustomError{
			Message: err.Error(),
			Code:    500,
			Err:     err,
		}
	}

	return nil
}

func (c *categoryService) GetByUserID(ctx *gin.Context) ([]*model.Category, *utils.CustomError) {
	userID, ok := ctx.Get("user_id")

	if !ok {
		return nil, &utils.CustomError{
			Message: "user id is required",
			Code:    400,
			Err:     errors.New("user id is required"),
		}
	}

	categories, err := c.repo.GetByUserID(ctx, userID.(string))
	if err != nil {
		return nil, &utils.CustomError{
			Message: err.Error(),
			Code:    500,
			Err:     err,
		}
	}

	return categories, nil
}

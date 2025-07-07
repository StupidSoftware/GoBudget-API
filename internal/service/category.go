package service

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryService interface {
	Create(ctx *gin.Context, user *model.Category) *utils.CustomError
	Login(ctx *gin.Context, user *model.Category) (string, *utils.CustomError)
	Delete(ctx *gin.Context, id uuid.UUID) error
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
		return &utils.CustomError{
			Message: err.Error(),
			Code:    500,
			Err:     err,
		}
	}

	return nil
}

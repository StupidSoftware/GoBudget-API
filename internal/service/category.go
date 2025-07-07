package service

import (
	"errors"

	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var logger = config.GetLogger("service")

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

	exists, err := c.repo.CategoryAlreadyExists(ctx, category)

	if err != nil {
		return &utils.CustomError{
			Message: err.Error(),
			Code:    500,
			Err:     err,
		}
	}

	if exists {
		return &utils.CustomError{
			Message: "Category already exists",
			Code:    409,
			Err:     errors.New("category already exists"),
		}
	}

	if err := c.repo.Create(ctx, category); err != nil {
		return utils.NewCustomPGError("user not found", 404, err)
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

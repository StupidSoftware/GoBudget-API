package controller

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
}

type categoryController struct {
	svc service.CategoryService
}

func NewCategoryController(svc service.CategoryService) *categoryController {
	return &categoryController{
		svc: svc,
	}
}

// @Summary Create a new category
// @Tags Category
// @Description Create a new category
// @Security BearerAuth
// @Accept json
// @Param category body model.Category true "Category"
// @Success 201 {object} model.Category
// @Failure 400 {object} utils.CustomError
// @Failure 409 {object} utils.CustomError
// @Failure 500 {object} utils.CustomError
// @Router /categories [post]
func (c *categoryController) Create(ctx *gin.Context) {
	var category model.Category

	if err := ctx.BindJSON(&category); err != nil {
		logger.Errorf("Error binding JSON: %v", err)
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := c.svc.Create(ctx, &category); err != nil {
		logger.Errorf("Error creating category: %v", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Category created",
	})
}

// @Summary Get all categories
// @Tags Category
// @Description Get all categories
// @Accept json
// @Security BearerAuth
// @Produce json
// @Success 200 {array} model.Category
// @Failure 400 {object} utils.CustomError
// @Failure 500 {object} utils.CustomError
// @Router /categories [get]
func (c *categoryController) GetByUserID(ctx *gin.Context) {
	categories, err := c.svc.GetByUserID(ctx)
	if err != nil {
		logger.Errorf("Error getting categories: %v", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(200, categories)
}

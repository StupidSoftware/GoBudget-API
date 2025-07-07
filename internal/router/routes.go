package router

import (
	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/controller"
	"github.com/breno5g/GoBudget/internal/middleware"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		userRoutes(v1)
		categoryRoutes(v1)
	}
}

func userRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	db := config.GetDB()
	repo := repository.NewRepository(db)
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	users := r.Group("/users")
	{
		users.POST("/", ctrl.Create)
		users.GET("/login", ctrl.Login)
	}
	return users
}

func categoryRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	categories := r.Group("/categories")
	{
		categories.Use(middleware.AuthRequired())
		{
			categories.POST("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "Hello World",
				})
			})
		}
	}

	return categories
}

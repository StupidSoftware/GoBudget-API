package router

import (
	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/controller"
	"github.com/breno5g/GoBudget/internal/middleware"
	"github.com/breno5g/GoBudget/internal/repository"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		userRoutes(v1)
		categoryRoutes(v1)
		transactionRoutes(v1)

		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})

		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	}
}

func userRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	db := config.GetDB()
	repo := repository.NewUserRepository(db)
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
	db := config.GetDB()
	repo := repository.NewCategoryRepository(db)
	svc := service.NewCategoryService(repo)
	ctrl := controller.NewCategoryController(svc)

	categories := r.Group("/categories")
	{
		categories.Use(middleware.AuthRequired())
		{
			categories.POST("/", ctrl.Create)
			categories.GET("/", ctrl.GetByUserID)
		}
	}

	return categories
}

func transactionRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	db := config.GetDB()
	repo := repository.NewTransactionRepository(db)
	svc := service.NewTransactionService(repo)
	ctrl := controller.NewTransactionController(svc)

	transactions := r.Group("/transactions")
	{
		transactions.Use(middleware.AuthRequired())
		{
			transactions.POST("/", ctrl.Create)
			transactions.GET("/", ctrl.GetByUserID)
		}
	}

	return transactions
}

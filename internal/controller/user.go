package controller

import (
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(ctx *gin.Context)
	GetByUsername(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	svc service.UserService
}

func NewUserController(svc service.UserService) *controller {
	return &controller{
		svc: svc,
	}
}

func (c *controller) Create(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err = c.svc.Create(ctx, &user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "User created",
	})
}

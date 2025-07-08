package controller

import (
	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/model"
	"github.com/breno5g/GoBudget/internal/service"
	"github.com/gin-gonic/gin"
)

var logger = config.GetLogger("controller")

type UserController interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
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

// @Summary Create a new user
// @Tags User
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} utils.CustomError
// @Failure 409 {object} utils.CustomError
// @Failure 500 {object} utils.CustomError
// @Router /users [post]
func (c *controller) Create(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := c.svc.Create(ctx, &user); err != nil {
		logger.Errorf("Error creating user: %v", err)
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "User created",
	})
}

func (c *controller) Login(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	token, err := c.svc.Login(ctx, &user)

	if err != nil {
		logger.Errorf("Error logging in: %v", err)
		ctx.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
